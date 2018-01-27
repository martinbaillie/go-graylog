// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SourcesList sources list
// swagger:model SourcesList
type SourcesList struct {

	// range
	Range int64 `json:"range,omitempty"`

	// sources
	Sources interface{} `json:"sources,omitempty"`

	// took ms
	TookMs int64 `json:"took_ms,omitempty"`

	// total
	Total int64 `json:"total,omitempty"`
}

// Validate validates this sources list
func (m *SourcesList) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SourcesList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SourcesList) UnmarshalBinary(b []byte) error {
	var res SourcesList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}