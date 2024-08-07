// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PcVMNicConfiguration Virtual Machine NIC
//
// Virtual Machine NIC.
//
// swagger:model pc_vm_nic_configuration
type PcVMNicConfiguration struct {

	// Network IP address.
	IPList []string `json:"ip_list"`

	// network configuration
	NetworkConfiguration *NetworkConfig `json:"network_configuration,omitempty"`
}

// Validate validates this pc vm nic configuration
func (m *PcVMNicConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PcVMNicConfiguration) validateIPList(formats strfmt.Registry) error {
	if swag.IsZero(m.IPList) { // not required
		return nil
	}

	for i := 0; i < len(m.IPList); i++ {

		if err := validate.Pattern("ip_list"+"."+strconv.Itoa(i), "body", m.IPList[i], `^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`); err != nil {
			return err
		}

	}

	return nil
}

func (m *PcVMNicConfiguration) validateNetworkConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.NetworkConfiguration) { // not required
		return nil
	}

	if m.NetworkConfiguration != nil {
		if err := m.NetworkConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network_configuration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("network_configuration")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this pc vm nic configuration based on the context it is used
func (m *PcVMNicConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNetworkConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PcVMNicConfiguration) contextValidateNetworkConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.NetworkConfiguration != nil {

		if swag.IsZero(m.NetworkConfiguration) { // not required
			return nil
		}

		if err := m.NetworkConfiguration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network_configuration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("network_configuration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PcVMNicConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PcVMNicConfiguration) UnmarshalBinary(b []byte) error {
	var res PcVMNicConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
