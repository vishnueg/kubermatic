/*
Copyright 2022 The Kubermatic Kubernetes Platform contributors.

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

package openapi

import (
	"embed"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"strings"

	"github.com/gobuffalo/flect"

	ext "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	extv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	extv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apiextensions-apiserver/pkg/apiserver/validation"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/kube-openapi/pkg/validation/validate"
)

// NewValidatorFromCRD creates a new validator based on the supplied crd.
// The CRD must be in yaml or json representation.
// As a v1 CRD can contain multiple versions, a desired one needs to be specified
// When using beta1v1 CRDs, version can be left empty.
func NewValidatorFromCRD(crd io.Reader, version string) (*validate.SchemaValidator, error) {
	u := &unstructured.Unstructured{}
	dec := yaml.NewYAMLOrJSONDecoder(crd, 1024)
	if err := dec.Decode(u); err != nil {
		return nil, err
	}

	// Currently there is an issue in the k8s yaml package when working with documents that start with a blank line followed by the yaml document separator
	// (https://github.com/kubernetes/apimachinery/issues/133). So for now we work around the empty buffer by re-reading from it
	// Note: alternative yaml packages (e.g. gopkg.in/yaml) cannot be used, as they do not handle kubernetes-specifc struct tags
	if u.GetAPIVersion() == "" {
		err := dec.Decode(u)
		if err != nil {
			return nil, err
		}
	}

	var sv *validate.SchemaValidator
	var err error
	v := u.GetObjectKind().GroupVersionKind().Version
	switch v {
	case "v1beta1":
		sv, err = v1beta1Validator(u)
	case "v1":
		if version == "" {
			return nil, fmt.Errorf("when using v1 CRDs, desVer cannot be empty")
		}
		sv, err = v1Validator(u, version)
	default:
		return nil, fmt.Errorf("APIVersion %q of CRD is not supported", v)
	}
	if err != nil {
		return nil, err
	}

	return sv, nil
}

func v1beta1Validator(u *unstructured.Unstructured) (*validate.SchemaValidator, error) {
	crdv1beta1 := &extv1beta1.CustomResourceDefinition{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), crdv1beta1); err != nil {
		return nil, err
	}
	crdr := &ext.CustomResourceDefinition{}
	if err := extv1beta1.Convert_v1beta1_CustomResourceDefinition_To_apiextensions_CustomResourceDefinition(crdv1beta1, crdr, nil); err != nil {
		return nil, err
	}
	sv, _, err := validation.NewSchemaValidator(crdr.Spec.Validation)
	if err != nil {
		return nil, err
	}

	return sv, nil
}

func v1Validator(u *unstructured.Unstructured, desVer string) (*validate.SchemaValidator, error) {
	crdv1 := &extv1.CustomResourceDefinition{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), crdv1); err != nil {
		return nil, err
	}
	crdr := &ext.CustomResourceDefinition{}
	if err := extv1.Convert_v1_CustomResourceDefinition_To_apiextensions_CustomResourceDefinition(crdv1, crdr, nil); err != nil {
		return nil, err
	}

	var err error
	var sv *validate.SchemaValidator
	for _, ver := range crdr.Spec.Versions {
		if ver.Name == desVer {
			// If there is only one version in the CRD, the crdv1 to crd converter will move the validation
			// into the global .Spec.Validation. Therefore, we need to manually check if per-version or
			// global validation is enabled
			if ver.Schema != nil {
				sv, _, err = validation.NewSchemaValidator(ver.Schema)
			} else {
				sv, _, err = validation.NewSchemaValidator(crdr.Spec.Validation)
			}
			if err != nil {
				return nil, err
			}
			break
		}
	}
	if sv == nil {
		return nil, fmt.Errorf("could not find SchemaValidator for desired version %q", desVer)
	}

	return sv, nil
}

// Efs is an embedded fs that contains kubermatic CRD manifest
//go:embed crd/k8c.io
var Efs embed.FS

// CRDRootFolder is the root folder where CRD manifest live in Efs.
const CRDRootFolder = "crd/k8c.io/"

// NewValidatorForType returns a suitable validator from its catalogue based on the supplied TypeMeta.
func NewValidatorForType(tm *metav1.TypeMeta) (*validate.SchemaValidator, error) {
	// as filenames are being generated by controller-gen, use the same pluralization mechanism
	// https://github.com/kubernetes-sigs/controller-tools/blob/8cb5ce83c3cca425a4de0af3d2578e31a3cd6a48/pkg/crd/spec.go#L58
	kindPlural := strings.ToLower(flect.Pluralize(tm.GroupVersionKind().Kind))
	crdfpath := CRDRootFolder + tm.GroupVersionKind().Group + "_" + kindPlural + ".yaml"

	f, err := Efs.Open(crdfpath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, fmt.Errorf("could not find validator for type \"%s_%s\"", tm.GroupVersionKind().Group, kindPlural)
		}
		return nil, err
	}

	v, err := NewValidatorFromCRD(f, tm.GroupVersionKind().Version)
	if err != nil {
		return nil, err
	}

	return v, nil
}
