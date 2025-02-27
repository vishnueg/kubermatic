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

// KubeOneCloudSpec kube one cloud spec
//
// swagger:model KubeOneCloudSpec
type KubeOneCloudSpec struct {

	// aws
	Aws *KubeOneAWSCloudSpec `json:"aws,omitempty"`

	// azure
	Azure *KubeOneAzureCloudSpec `json:"azure,omitempty"`

	// digitalocean
	Digitalocean *KubeOneDigitalOceanCloudSpec `json:"digitalocean,omitempty"`

	// equinix
	Equinix *KubeOneEquinixCloudSpec `json:"equinix,omitempty"`

	// gcp
	Gcp *KubeOneGCPCloudSpec `json:"gcp,omitempty"`

	// hetzner
	Hetzner *KubeOneHetznerCloudSpec `json:"hetzner,omitempty"`

	// nutanix
	Nutanix *KubeOneNutanixCloudSpec `json:"nutanix,omitempty"`

	// openstack
	Openstack *KubeOneOpenStackCloudSpec `json:"openstack,omitempty"`

	// vsphere
	Vsphere *KubeOneVSphereCloudSpec `json:"vsphere,omitempty"`
}

// Validate validates this kube one cloud spec
func (m *KubeOneCloudSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDigitalocean(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEquinix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHetzner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNutanix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenstack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVsphere(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubeOneCloudSpec) validateAws(formats strfmt.Registry) error {
	if swag.IsZero(m.Aws) { // not required
		return nil
	}

	if m.Aws != nil {
		if err := m.Aws.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *KubeOneCloudSpec) validateAzure(formats strfmt.Registry) error {
	if swag.IsZero(m.Azure) { // not required
		return nil
	}

	if m.Azure != nil {
		if err := m.Azure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

func (m *KubeOneCloudSpec) validateDigitalocean(formats strfmt.Registry) error {
	if swag.IsZero(m.Digitalocean) { // not required
		return nil
	}

	if m.Digitalocean != nil {
		if err := m.Digitalocean.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("digitalocean")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("digitalocean")
			}
			return err
		}
	}

	return nil
}

func (m *KubeOneCloudSpec) validateEquinix(formats strfmt.Registry) error {
	if swag.IsZero(m.Equinix) { // not required
		return nil
	}

	if m.Equinix != nil {
		if err := m.Equinix.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("equinix")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("equinix")
			}
			return err
		}
	}

	return nil
}

func (m *KubeOneCloudSpec) validateGcp(formats strfmt.Registry) error {
	if swag.IsZero(m.Gcp) { // not required
		return nil
	}

	if m.Gcp != nil {
		if err := m.Gcp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcp")
			}
			return err
		}
	}

	return nil
}

func (m *KubeOneCloudSpec) validateHetzner(formats strfmt.Registry) error {
	if swag.IsZero(m.Hetzner) { // not required
		return nil
	}

	if m.Hetzner != nil {
		if err := m.Hetzner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hetzner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hetzner")
			}
			return err
		}
	}

	return nil
}

func (m *KubeOneCloudSpec) validateNutanix(formats strfmt.Registry) error {
	if swag.IsZero(m.Nutanix) { // not required
		return nil
	}

	if m.Nutanix != nil {
		if err := m.Nutanix.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nutanix")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nutanix")
			}
			return err
		}
	}

	return nil
}

func (m *KubeOneCloudSpec) validateOpenstack(formats strfmt.Registry) error {
	if swag.IsZero(m.Openstack) { // not required
		return nil
	}

	if m.Openstack != nil {
		if err := m.Openstack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("openstack")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("openstack")
			}
			return err
		}
	}

	return nil
}

func (m *KubeOneCloudSpec) validateVsphere(formats strfmt.Registry) error {
	if swag.IsZero(m.Vsphere) { // not required
		return nil
	}

	if m.Vsphere != nil {
		if err := m.Vsphere.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vsphere")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vsphere")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this kube one cloud spec based on the context it is used
func (m *KubeOneCloudSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAws(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzure(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDigitalocean(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEquinix(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGcp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHetzner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNutanix(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOpenstack(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVsphere(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubeOneCloudSpec) contextValidateAws(ctx context.Context, formats strfmt.Registry) error {

	if m.Aws != nil {
		if err := m.Aws.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *KubeOneCloudSpec) contextValidateAzure(ctx context.Context, formats strfmt.Registry) error {

	if m.Azure != nil {
		if err := m.Azure.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

func (m *KubeOneCloudSpec) contextValidateDigitalocean(ctx context.Context, formats strfmt.Registry) error {

	if m.Digitalocean != nil {
		if err := m.Digitalocean.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("digitalocean")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("digitalocean")
			}
			return err
		}
	}

	return nil
}

func (m *KubeOneCloudSpec) contextValidateEquinix(ctx context.Context, formats strfmt.Registry) error {

	if m.Equinix != nil {
		if err := m.Equinix.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("equinix")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("equinix")
			}
			return err
		}
	}

	return nil
}

func (m *KubeOneCloudSpec) contextValidateGcp(ctx context.Context, formats strfmt.Registry) error {

	if m.Gcp != nil {
		if err := m.Gcp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcp")
			}
			return err
		}
	}

	return nil
}

func (m *KubeOneCloudSpec) contextValidateHetzner(ctx context.Context, formats strfmt.Registry) error {

	if m.Hetzner != nil {
		if err := m.Hetzner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hetzner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hetzner")
			}
			return err
		}
	}

	return nil
}

func (m *KubeOneCloudSpec) contextValidateNutanix(ctx context.Context, formats strfmt.Registry) error {

	if m.Nutanix != nil {
		if err := m.Nutanix.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nutanix")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nutanix")
			}
			return err
		}
	}

	return nil
}

func (m *KubeOneCloudSpec) contextValidateOpenstack(ctx context.Context, formats strfmt.Registry) error {

	if m.Openstack != nil {
		if err := m.Openstack.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("openstack")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("openstack")
			}
			return err
		}
	}

	return nil
}

func (m *KubeOneCloudSpec) contextValidateVsphere(ctx context.Context, formats strfmt.Registry) error {

	if m.Vsphere != nil {
		if err := m.Vsphere.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vsphere")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vsphere")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KubeOneCloudSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubeOneCloudSpec) UnmarshalBinary(b []byte) error {
	var res KubeOneCloudSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
