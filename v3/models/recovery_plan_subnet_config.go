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

// RecoveryPlanSubnetConfig Subnet configuration for the Recovery Plan
//
// Subnet configuration for network mapping in the Recovery Plan.
//
// swagger:model recovery_plan_subnet_config
type RecoveryPlanSubnetConfig struct {

	// External connectivity state of the subnet. This is applicable only for the subnet to be created in public cloud Availability Zone.
	//
	ExternalConnectivityState string `json:"external_connectivity_state,omitempty"`

	// Gateway IP address for the subnet.
	//
	// Required: true
	// Pattern: ^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$
	GatewayIP *string `json:"gateway_ip"`

	// Configurations for creating Layer2 stretches for the subnet.
	//
	L2StretchConfig []*RecoveryPlanL2StretchConfig `json:"l2_stretch_config"`

	// Prefix length for the subnet.
	//
	// Required: true
	// Maximum: 32
	// Minimum: 0
	PrefixLength *int32 `json:"prefix_length"`

	// Network subnet range carved out of given CIDR for this network. For example, if given gateway_ip is 10.0.0.0 and prefix_length is 24 then subnet_range can have start_ip_address and end_ip_address anywhere in the range 10.0.0.0 to 10.0.0.256.
	//
	SubnetRange *RecoveryPlanSubnetRangeConfig `json:"subnet_range,omitempty"`
}

// Validate validates this recovery plan subnet config
func (m *RecoveryPlanSubnetConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGatewayIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateL2StretchConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrefixLength(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnetRange(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoveryPlanSubnetConfig) validateGatewayIP(formats strfmt.Registry) error {

	if err := validate.Required("gateway_ip", "body", m.GatewayIP); err != nil {
		return err
	}

	if err := validate.Pattern("gateway_ip", "body", *m.GatewayIP, `^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`); err != nil {
		return err
	}

	return nil
}

func (m *RecoveryPlanSubnetConfig) validateL2StretchConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.L2StretchConfig) { // not required
		return nil
	}

	for i := 0; i < len(m.L2StretchConfig); i++ {
		if swag.IsZero(m.L2StretchConfig[i]) { // not required
			continue
		}

		if m.L2StretchConfig[i] != nil {
			if err := m.L2StretchConfig[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("l2_stretch_config" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("l2_stretch_config" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RecoveryPlanSubnetConfig) validatePrefixLength(formats strfmt.Registry) error {

	if err := validate.Required("prefix_length", "body", m.PrefixLength); err != nil {
		return err
	}

	if err := validate.MinimumInt("prefix_length", "body", int64(*m.PrefixLength), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("prefix_length", "body", int64(*m.PrefixLength), 32, false); err != nil {
		return err
	}

	return nil
}

func (m *RecoveryPlanSubnetConfig) validateSubnetRange(formats strfmt.Registry) error {
	if swag.IsZero(m.SubnetRange) { // not required
		return nil
	}

	if m.SubnetRange != nil {
		if err := m.SubnetRange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subnet_range")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subnet_range")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this recovery plan subnet config based on the context it is used
func (m *RecoveryPlanSubnetConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateL2StretchConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubnetRange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoveryPlanSubnetConfig) contextValidateL2StretchConfig(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.L2StretchConfig); i++ {

		if m.L2StretchConfig[i] != nil {

			if swag.IsZero(m.L2StretchConfig[i]) { // not required
				return nil
			}

			if err := m.L2StretchConfig[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("l2_stretch_config" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("l2_stretch_config" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RecoveryPlanSubnetConfig) contextValidateSubnetRange(ctx context.Context, formats strfmt.Registry) error {

	if m.SubnetRange != nil {

		if swag.IsZero(m.SubnetRange) { // not required
			return nil
		}

		if err := m.SubnetRange.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subnet_range")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subnet_range")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecoveryPlanSubnetConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoveryPlanSubnetConfig) UnmarshalBinary(b []byte) error {
	var res RecoveryPlanSubnetConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}