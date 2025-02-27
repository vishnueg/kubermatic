// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GKEClusterSpec GKEClusterSpec A Google Kubernetes Engine cluster.
//
// swagger:model GKEClusterSpec
type GKEClusterSpec struct {

	// Autopilot: Autopilot configuration for the cluster.
	Autopilot bool `json:"autopilot,omitempty"`

	// ClusterIpv4Cidr: The IP address range of the container pods in this
	// cluster, in CIDR
	// (http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing)
	// notation (e.g. `10.96.0.0/14`). Leave blank to have one automatically
	// chosen or specify a `/14` block in `10.0.0.0/8`.
	ClusterIPV4Cidr string `json:"clusterIpv4Cidr,omitempty"`

	// DefaultMaxPodsConstraint: The default constraint on the maximum
	// number of pods that can be run simultaneously on a node in the node
	// pool of this cluster. Only honored if cluster created with IP Alias
	// support.
	DefaultMaxPodsConstraint int64 `json:"defaultMaxPodsConstraint,omitempty"`

	// EnableKubernetesAlpha: Kubernetes alpha features are enabled on this
	// cluster. This includes alpha API groups (e.g. v1alpha1) and features
	// that may not be production ready in the kubernetes version of the
	// master and nodes. The cluster has no SLA for uptime and master/node
	// upgrades are disabled. Alpha enabled clusters are automatically
	// deleted thirty days after creation.
	EnableKubernetesAlpha bool `json:"enableKubernetesAlpha,omitempty"`

	// EnableTpu: Enable the ability to use Cloud TPUs in this cluster.
	EnableTpu bool `json:"enableTpu,omitempty"`

	// InitialClusterVersion: The initial Kubernetes version for this
	// cluster. Valid versions are those found in validMasterVersions
	// returned by getServerConfig. The version can be upgraded over time;
	// such upgrades are reflected in currentMasterVersion and
	// currentNodeVersion. Users may specify either explicit versions
	// offered by Kubernetes Engine or version aliases, which have the
	// following behavior: - "latest": picks the highest valid Kubernetes
	// version - "1.X": picks the highest valid patch+gke.N patch in the 1.X
	// version - "1.X.Y": picks the highest valid gke.N patch in the 1.X.Y
	// version - "1.X.Y-gke.N": picks an explicit Kubernetes version -
	// "","-": picks the default Kubernetes version
	InitialClusterVersion string `json:"initialClusterVersion,omitempty"`

	// InitialNodeCount: The number of nodes to create in this cluster. You
	// must ensure that your Compute Engine resource quota
	// (https://cloud.google.com/compute/quotas) is sufficient for this
	// number of instances. You must also have available firewall and routes
	// quota. For requests, this field should only be used in lieu of a
	// "node_pool" object, since this configuration (along with the
	// "node_config") will be used to create a "NodePool" object with an
	// auto-generated name. Do not use this and a node_pool at the same
	// time. This field is deprecated, use node_pool.initial_node_count
	// instead.
	InitialNodeCount int64 `json:"initialNodeCount,omitempty"`

	// Locations: The list of Google Compute Engine zones
	// (https://cloud.google.com/compute/docs/zones#available) in which the
	// cluster's nodes should be located. This field provides a default
	// value if NodePool.Locations
	// (https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters.nodePools#NodePool.FIELDS.locations)
	// are not specified during node pool creation. Warning: changing
	// cluster locations will update the NodePool.Locations
	// (https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.clusters.nodePools#NodePool.FIELDS.locations)
	// of all node pools and will result in nodes being added and/or
	// removed.
	Locations []string `json:"locations"`

	// Network: The name of the Google Compute Engine network
	// (https://cloud.google.com/compute/docs/networks-and-firewalls#networks)
	// to which the cluster is connected. If left unspecified, the `default`
	// network will be used.
	Network string `json:"network,omitempty"`

	// Subnetwork: The name of the Google Compute Engine subnetwork
	// (https://cloud.google.com/compute/docs/subnetworks) to which the
	// cluster is connected.
	Subnetwork string `json:"subnetwork,omitempty"`

	// TpuIpv4CidrBlock: [Output only] The IP address range of the Cloud
	// TPUs in this cluster, in CIDR
	// (http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing)
	// notation (e.g. `1.2.3.4/29`).
	TpuIPV4CidrBlock string `json:"tpuIpv4CidrBlock,omitempty"`

	// VerticalPodAutoscaling: Cluster-level Vertical Pod Autoscaling
	// configuration.
	VerticalPodAutoscaling bool `json:"verticalPodAutoscaling,omitempty"`

	// autoscaling
	Autoscaling *GKEClusterAutoscaling `json:"autoscaling,omitempty"`

	// node config
	NodeConfig *GKENodeConfig `json:"nodeConfig,omitempty"`
}

// Validate validates this g k e cluster spec
func (m *GKEClusterSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoscaling(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GKEClusterSpec) validateAutoscaling(formats strfmt.Registry) error {
	if swag.IsZero(m.Autoscaling) { // not required
		return nil
	}

	if m.Autoscaling != nil {
		if err := m.Autoscaling.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoscaling")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autoscaling")
			}
			return err
		}
	}

	return nil
}

func (m *GKEClusterSpec) validateNodeConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeConfig) { // not required
		return nil
	}

	if m.NodeConfig != nil {
		if err := m.NodeConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodeConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this g k e cluster spec based on the context it is used
func (m *GKEClusterSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAutoscaling(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNodeConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GKEClusterSpec) contextValidateAutoscaling(ctx context.Context, formats strfmt.Registry) error {

	if m.Autoscaling != nil {
		if err := m.Autoscaling.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("autoscaling")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("autoscaling")
			}
			return err
		}
	}

	return nil
}

func (m *GKEClusterSpec) contextValidateNodeConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.NodeConfig != nil {
		if err := m.NodeConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nodeConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GKEClusterSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GKEClusterSpec) UnmarshalBinary(b []byte) error {
	var res GKEClusterSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
