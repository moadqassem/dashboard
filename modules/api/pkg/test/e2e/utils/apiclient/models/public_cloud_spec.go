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

// PublicCloudSpec PublicCloudSpec is a public counterpart of apiv1.CloudSpec.
//
// swagger:model PublicCloudSpec
type PublicCloudSpec struct {

	// datacenter name
	DatacenterName string `json:"dc,omitempty"`

	// alibaba
	Alibaba PublicAlibabaCloudSpec `json:"alibaba,omitempty"`

	// anexia
	Anexia PublicAnexiaCloudSpec `json:"anexia,omitempty"`

	// aws
	Aws *PublicAWSCloudSpec `json:"aws,omitempty"`

	// azure
	Azure *PublicAzureCloudSpec `json:"azure,omitempty"`

	// bringyourown
	Bringyourown PublicBringYourOwnCloudSpec `json:"bringyourown,omitempty"`

	// digitalocean
	Digitalocean PublicDigitaloceanCloudSpec `json:"digitalocean,omitempty"`

	// edge
	Edge PublicEdgeCloudSpec `json:"edge,omitempty"`

	// fake
	Fake PublicFakeCloudSpec `json:"fake,omitempty"`

	// gcp
	Gcp *PublicGCPCloudSpec `json:"gcp,omitempty"`

	// hetzner
	Hetzner PublicHetznerCloudSpec `json:"hetzner,omitempty"`

	// kubevirt
	Kubevirt *PublicKubevirtCloudSpec `json:"kubevirt,omitempty"`

	// nutanix
	Nutanix PublicNutanixCloudSpec `json:"nutanix,omitempty"`

	// openstack
	Openstack *PublicOpenstackCloudSpec `json:"openstack,omitempty"`

	// packet
	Packet PublicPacketCloudSpec `json:"packet,omitempty"`

	// vmwareclouddirector
	Vmwareclouddirector PublicVMwareCloudDirectorCloudSpec `json:"vmwareclouddirector,omitempty"`

	// vsphere
	Vsphere PublicVSphereCloudSpec `json:"vsphere,omitempty"`
}

// Validate validates this public cloud spec
func (m *PublicCloudSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubevirt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenstack(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicCloudSpec) validateAws(formats strfmt.Registry) error {
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

func (m *PublicCloudSpec) validateAzure(formats strfmt.Registry) error {
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

func (m *PublicCloudSpec) validateGcp(formats strfmt.Registry) error {
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

func (m *PublicCloudSpec) validateKubevirt(formats strfmt.Registry) error {
	if swag.IsZero(m.Kubevirt) { // not required
		return nil
	}

	if m.Kubevirt != nil {
		if err := m.Kubevirt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubevirt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kubevirt")
			}
			return err
		}
	}

	return nil
}

func (m *PublicCloudSpec) validateOpenstack(formats strfmt.Registry) error {
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

// ContextValidate validate this public cloud spec based on the context it is used
func (m *PublicCloudSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAws(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzure(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGcp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKubevirt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOpenstack(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicCloudSpec) contextValidateAws(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PublicCloudSpec) contextValidateAzure(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PublicCloudSpec) contextValidateGcp(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PublicCloudSpec) contextValidateKubevirt(ctx context.Context, formats strfmt.Registry) error {

	if m.Kubevirt != nil {
		if err := m.Kubevirt.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubevirt")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kubevirt")
			}
			return err
		}
	}

	return nil
}

func (m *PublicCloudSpec) contextValidateOpenstack(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *PublicCloudSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicCloudSpec) UnmarshalBinary(b []byte) error {
	var res PublicCloudSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
