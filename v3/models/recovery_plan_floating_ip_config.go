// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RecoveryPlanFloatingIPConfig Floating IP configuration for a VM
//
// Configuration for assigning floating IP to a VM on the execution of the Recovery Plan.
//
// swagger:model recovery_plan_floating_ip_config
type RecoveryPlanFloatingIPConfig struct {

	// IP to be assigned to VM, in case of failover.
	//
	// Pattern: ^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$
	IP string `json:"ip,omitempty"`

	// Whether to allocate the floating IPs for the VMs dynamically.
	//
	ShouldAllocateDynamically bool `json:"should_allocate_dynamically,omitempty"`
}

// Validate validates this recovery plan floating ip config
func (m *RecoveryPlanFloatingIPConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoveryPlanFloatingIPConfig) validateIP(formats strfmt.Registry) error {
	if swag.IsZero(m.IP) { // not required
		return nil
	}

	if err := validate.Pattern("ip", "body", m.IP, `^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this recovery plan floating ip config based on context it is used
func (m *RecoveryPlanFloatingIPConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RecoveryPlanFloatingIPConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoveryPlanFloatingIPConfig) UnmarshalBinary(b []byte) error {
	var res RecoveryPlanFloatingIPConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}