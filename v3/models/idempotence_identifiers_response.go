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

// IdempotenceIdentifiersResponse Idempotence identifier status definition.
//
// Idempotence identifier status definition.
//
// swagger:model idempotence_identifiers_response
type IdempotenceIdentifiersResponse struct {

	// The client identifier string.
	ClientIdentifier string `json:"client_identifier,omitempty"`

	// The number of idempotence identifiers provided.
	// Required: true
	Count *int64 `json:"count"`

	// UTC date and time in RFC-3339 format of the expiration time (with reference to system time). Value is creation time + valid_duration
	// Format: date-time
	ExpirationTime strfmt.DateTime `json:"expiration_time,omitempty"`

	// uuid list
	// Required: true
	UUIDList []string `json:"uuid_list"`
}

// Validate validates this idempotence identifiers response
func (m *IdempotenceIdentifiersResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpirationTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUIDList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdempotenceIdentifiersResponse) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

func (m *IdempotenceIdentifiersResponse) validateExpirationTime(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpirationTime) { // not required
		return nil
	}

	if err := validate.FormatOf("expiration_time", "body", "date-time", m.ExpirationTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IdempotenceIdentifiersResponse) validateUUIDList(formats strfmt.Registry) error {

	if err := validate.Required("uuid_list", "body", m.UUIDList); err != nil {
		return err
	}

	for i := 0; i < len(m.UUIDList); i++ {

		if err := validate.Pattern("uuid_list"+"."+strconv.Itoa(i), "body", m.UUIDList[i], `^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this idempotence identifiers response based on context it is used
func (m *IdempotenceIdentifiersResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IdempotenceIdentifiersResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdempotenceIdentifiersResponse) UnmarshalBinary(b []byte) error {
	var res IdempotenceIdentifiersResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}