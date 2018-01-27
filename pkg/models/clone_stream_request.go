// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CloneStreamRequest clone stream request
// swagger:model CloneStreamRequest
type CloneStreamRequest struct {

	// description
	Description string `json:"description,omitempty"`

	// index set id
	IndexSetID string `json:"index_set_id,omitempty"`

	// remove matches from default stream
	RemoveMatchesFromDefaultStream bool `json:"remove_matches_from_default_stream,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this clone stream request
func (m *CloneStreamRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CloneStreamRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloneStreamRequest) UnmarshalBinary(b []byte) error {
	var res CloneStreamRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}