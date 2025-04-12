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

// ErrorResponse ErrorResponse
//
// # Общая ошибка
//
// swagger:model ErrorResponse
type ErrorResponse struct {

	// error
	Error *ErrorResponseAO0Error `json:"error,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ErrorResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Error *ErrorResponseAO0Error `json:"error,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Error = dataAO0.Error

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ErrorResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	var dataAO0 struct {
		Error *ErrorResponseAO0Error `json:"error,omitempty"`
	}

	dataAO0.Error = m.Error

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this error response
func (m *ErrorResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ErrorResponse) validateError(formats strfmt.Registry) error {

	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this error response based on the context it is used
func (m *ErrorResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ErrorResponse) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if m.Error != nil {

		if swag.IsZero(m.Error) { // not required
			return nil
		}

		if err := m.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ErrorResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorResponse) UnmarshalBinary(b []byte) error {
	var res ErrorResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ErrorResponseAO0Error error response a o0 error
//
// swagger:model ErrorResponseAO0Error
type ErrorResponseAO0Error struct {

	// Message
	Message string `json:"message,omitempty"`
}

// Validate validates this error response a o0 error
func (m *ErrorResponseAO0Error) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this error response a o0 error based on context it is used
func (m *ErrorResponseAO0Error) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ErrorResponseAO0Error) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorResponseAO0Error) UnmarshalBinary(b []byte) error {
	var res ErrorResponseAO0Error
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
