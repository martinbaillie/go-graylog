// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SimulationResponseMessagesItemsDecorationStats simulation response messages items decoration stats
// swagger:model simulationResponseMessagesItemsDecorationStats
type SimulationResponseMessagesItemsDecorationStats struct {

	// added fields
	AddedFields interface{} `json:"added_fields,omitempty"`

	// changed fields
	ChangedFields interface{} `json:"changed_fields,omitempty"`

	// removed fields
	RemovedFields interface{} `json:"removed_fields,omitempty"`
}

// Validate validates this simulation response messages items decoration stats
func (m *SimulationResponseMessagesItemsDecorationStats) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SimulationResponseMessagesItemsDecorationStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SimulationResponseMessagesItemsDecorationStats) UnmarshalBinary(b []byte) error {
	var res SimulationResponseMessagesItemsDecorationStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}