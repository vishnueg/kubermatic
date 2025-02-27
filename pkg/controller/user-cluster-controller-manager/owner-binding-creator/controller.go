/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package ownerbindingcreator

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	userclustercontrollermanager "k8c.io/kubermatic/v2/pkg/controller/user-cluster-controller-manager"
	predicateutil "k8c.io/kubermatic/v2/pkg/controller/util/predicate"
	handlercommon "k8c.io/kubermatic/v2/pkg/handler/common"

	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/rand"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/record"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const (
	// This controller creates cluster role bindings for the API roles.
	controllerName = "kkp-owner-binding-creator"
)

type reconciler struct {
	log             *zap.SugaredLogger
	client          ctrlruntimeclient.Client
	recorder        record.EventRecorder
	ownerEmail      string
	clusterIsPaused userclustercontrollermanager.IsPausedChecker
}

func Add(ctx context.Context, log *zap.SugaredLogger, mgr manager.Manager, ownerEmail string, clusterIsPaused userclustercontrollermanager.IsPausedChecker) error {
	log = log.Named(controllerName)

	r := &reconciler{
		log:             log,
		client:          mgr.GetClient(),
		recorder:        mgr.GetEventRecorderFor(controllerName),
		ownerEmail:      ownerEmail,
		clusterIsPaused: clusterIsPaused,
	}
	c, err := controller.New(controllerName, mgr, controller.Options{
		Reconciler: r,
	})
	if err != nil {
		return fmt.Errorf("failed to create controller: %w", err)
	}

	// Watch for changes to ClusterRoles
	if err = c.Watch(&source.Kind{Type: &rbacv1.ClusterRole{}}, &handler.EnqueueRequestForObject{}, predicateutil.ByLabel(handlercommon.UserClusterComponentKey, handlercommon.UserClusterRoleComponentValue)); err != nil {
		return fmt.Errorf("failed to establish watch for the ClusterRoles: %w", err)
	}
	// Watch for changes to ClusterRoleBindings
	if err = c.Watch(&source.Kind{Type: &rbacv1.ClusterRoleBinding{}}, enqueueAPIBindings(mgr.GetClient()), predicateutil.ByLabel(handlercommon.UserClusterComponentKey, handlercommon.UserClusterBindingComponentValue)); err != nil {
		return fmt.Errorf("failed to establish watch for the ClusterRoleBindings: %w", err)
	}

	return nil
}

func (r *reconciler) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	paused, err := r.clusterIsPaused(ctx)
	if err != nil {
		return reconcile.Result{}, fmt.Errorf("failed to check cluster pause status: %w", err)
	}
	if paused {
		return reconcile.Result{}, nil
	}

	log := r.log.With("clusterrole", request.Name)
	log.Debug("Reconciling")

	err = r.reconcile(ctx, log, request.Name)
	if err != nil {
		log.Errorw("Reconciling failed", zap.Error(err))
		r.recorder.Event(&rbacv1.ClusterRole{ObjectMeta: metav1.ObjectMeta{Name: request.Name}}, corev1.EventTypeWarning, "AddBindingFailed", err.Error())
	}

	return reconcile.Result{}, err
}

func (r *reconciler) reconcile(ctx context.Context, log *zap.SugaredLogger, clusterRoleName string) error {
	labels := map[string]string{handlercommon.UserClusterComponentKey: handlercommon.UserClusterBindingComponentValue}

	clusterRoleBindingList := &rbacv1.ClusterRoleBindingList{}
	if err := r.client.List(ctx, clusterRoleBindingList, ctrlruntimeclient.MatchingLabels(labels)); err != nil {
		return fmt.Errorf("failed to list ClusterRoleBindings: %w", err)
	}

	var existingClusterRoleBinding *rbacv1.ClusterRoleBinding
	for _, clusterRoleBinding := range clusterRoleBindingList.Items {
		if clusterRoleBinding.RoleRef.Name == clusterRoleName {
			existingClusterRoleBinding = clusterRoleBinding.DeepCopy()
			break
		}
	}

	// Create Cluster Role Binding if doesn't exist
	if existingClusterRoleBinding == nil {
		log.Infow("Creating ClusterRoleBinding", "clusterrole", clusterRoleName)

		crb := &rbacv1.ClusterRoleBinding{
			ObjectMeta: metav1.ObjectMeta{
				Name:   fmt.Sprintf("%s:%s", rand.String(10), clusterRoleName),
				Labels: labels,
			},
			RoleRef: rbacv1.RoleRef{
				APIGroup: rbacv1.GroupName,
				Kind:     "ClusterRole",
				Name:     clusterRoleName,
			},
		}

		// Bind user who created the cluster to the `cluster-admin` ClusterRole.
		// Add cluster owner only once when binding doesn't exist yet.
		// Later the user can remove/add subjects from the binding using the API.
		if clusterRoleName == "cluster-admin" {
			crb.Subjects = []rbacv1.Subject{
				{
					Kind:     rbacv1.UserKind,
					APIGroup: rbacv1.GroupName,
					Name:     r.ownerEmail,
				},
			}
		}

		if err := r.client.Create(ctx, crb); err != nil {
			return fmt.Errorf("failed to create ClusterRoleBinding: %w", err)
		}
	}

	return nil
}

// enqueueAPIBindings enqueues the ClusterRoleBindings with a special label component=userClusterRole.
func enqueueAPIBindings(client ctrlruntimeclient.Client) handler.EventHandler {
	return handler.EnqueueRequestsFromMapFunc(func(_ ctrlruntimeclient.Object) []reconcile.Request {
		clusterRoleList := &rbacv1.ClusterRoleList{}
		if err := client.List(context.Background(), clusterRoleList, ctrlruntimeclient.MatchingLabels{handlercommon.UserClusterComponentKey: handlercommon.UserClusterRoleComponentValue}); err != nil {
			utilruntime.HandleError(fmt.Errorf("failed to list ClusterRoles: %w", err))
			return []reconcile.Request{}
		}

		var requests []reconcile.Request
		for _, clusterRole := range clusterRoleList.Items {
			requests = append(requests, reconcile.Request{NamespacedName: types.NamespacedName{Name: clusterRole.Name}})
		}
		return requests
	})
}
