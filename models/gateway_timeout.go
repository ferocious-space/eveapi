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

// GatewayTimeout Gateway timeout
//
// # Gateway timeout model
//
// swagger:model gateway_timeout
type GatewayTimeout struct {

	// Gateway timeout message
	// Required: true
	Error *string `json:"error"`

	// number of seconds the request was given
	Timeout int64 `json:"timeout,omitempty"`
}

// Validate validates this gateway timeout
func (m *GatewayTimeout) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GatewayTimeout) validateError(formats strfmt.Registry) error {

	if err := validate.Required("error", "body", m.Error); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this gateway timeout based on context it is used
func (m *GatewayTimeout) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GatewayTimeout) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GatewayTimeout) UnmarshalBinary(b []byte) error {
	var res GatewayTimeout
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
