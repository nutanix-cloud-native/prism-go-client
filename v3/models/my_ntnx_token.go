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

// MyNtnxToken MCM My Nutanix token.
//
// MCM My Nutanix token.
//
// swagger:model my_ntnx_token
type MyNtnxToken struct {

	// API Key issue by my.nutanix API Keys app.
	// Required: true
	APIKey *string `json:"api_key"`

	// Unique identifier for the API key.
	// Required: true
	KeyID *string `json:"key_id"`

	// my.nutanix scope Id assigned to the deployed component.
	// Required: true
	ScopeID *string `json:"scope_id"`

	// my.nutanix subscope Id assigned to the deployed component.
	// Required: true
	SubscopeID *string `json:"subscope_id"`
}

// Validate validates this my ntnx token
func (m *MyNtnxToken) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScopeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscopeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MyNtnxToken) validateAPIKey(formats strfmt.Registry) error {

	if err := validate.Required("api_key", "body", m.APIKey); err != nil {
		return err
	}

	return nil
}

func (m *MyNtnxToken) validateKeyID(formats strfmt.Registry) error {

	if err := validate.Required("key_id", "body", m.KeyID); err != nil {
		return err
	}

	return nil
}

func (m *MyNtnxToken) validateScopeID(formats strfmt.Registry) error {

	if err := validate.Required("scope_id", "body", m.ScopeID); err != nil {
		return err
	}

	return nil
}

func (m *MyNtnxToken) validateSubscopeID(formats strfmt.Registry) error {

	if err := validate.Required("subscope_id", "body", m.SubscopeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this my ntnx token based on context it is used
func (m *MyNtnxToken) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MyNtnxToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MyNtnxToken) UnmarshalBinary(b []byte) error {
	var res MyNtnxToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}