// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IndexSetResponseIndexSetsItemsRetentionStrategy index set response index sets items retention strategy
// swagger:model indexSetResponseIndexSetsItemsRetentionStrategy
type IndexSetResponseIndexSetsItemsRetentionStrategy struct {

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this index set response index sets items retention strategy
func (m *IndexSetResponseIndexSetsItemsRetentionStrategy) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *IndexSetResponseIndexSetsItemsRetentionStrategy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IndexSetResponseIndexSetsItemsRetentionStrategy) UnmarshalBinary(b []byte) error {
	var res IndexSetResponseIndexSetsItemsRetentionStrategy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}