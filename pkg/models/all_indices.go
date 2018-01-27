// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AllIndices all indices
// swagger:model AllIndices
type AllIndices struct {

	// all
	All *AllIndicesAll `json:"all,omitempty"`

	// closed
	Closed *AllIndicesClosed `json:"closed,omitempty"`

	// reopened
	Reopened interface{} `json:"reopened,omitempty"`
}

// Validate validates this all indices
func (m *AllIndices) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAll(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateClosed(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AllIndices) validateAll(formats strfmt.Registry) error {

	if swag.IsZero(m.All) { // not required
		return nil
	}

	if m.All != nil {

		if err := m.All.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("all")
			}
			return err
		}
	}

	return nil
}

func (m *AllIndices) validateClosed(formats strfmt.Registry) error {

	if swag.IsZero(m.Closed) { // not required
		return nil
	}

	if m.Closed != nil {

		if err := m.Closed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("closed")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AllIndices) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AllIndices) UnmarshalBinary(b []byte) error {
	var res AllIndices
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}