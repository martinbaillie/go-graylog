// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IndexSetSummary index set summary
// swagger:model IndexSetSummary
type IndexSetSummary struct {

	// creation date
	CreationDate strfmt.DateTime `json:"creation_date,omitempty"`

	// default
	Default bool `json:"default,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// index analyzer
	IndexAnalyzer string `json:"index_analyzer,omitempty"`

	// index optimization disabled
	IndexOptimizationDisabled bool `json:"index_optimization_disabled,omitempty"`

	// index optimization max num segments
	IndexOptimizationMaxNumSegments int64 `json:"index_optimization_max_num_segments,omitempty"`

	// index prefix
	IndexPrefix string `json:"index_prefix,omitempty"`

	// replicas
	Replicas int64 `json:"replicas,omitempty"`

	// retention strategy
	RetentionStrategy *IndexSetSummaryRetentionStrategy `json:"retention_strategy,omitempty"`

	// retention strategy class
	RetentionStrategyClass string `json:"retention_strategy_class,omitempty"`

	// rotation strategy
	RotationStrategy *IndexSetSummaryRotationStrategy `json:"rotation_strategy,omitempty"`

	// rotation strategy class
	RotationStrategyClass string `json:"rotation_strategy_class,omitempty"`

	// shards
	Shards int64 `json:"shards,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// writable
	Writable bool `json:"writable,omitempty"`
}

// Validate validates this index set summary
func (m *IndexSetSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRetentionStrategy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRotationStrategy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IndexSetSummary) validateCreationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creation_date", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IndexSetSummary) validateRetentionStrategy(formats strfmt.Registry) error {

	if swag.IsZero(m.RetentionStrategy) { // not required
		return nil
	}

	if m.RetentionStrategy != nil {

		if err := m.RetentionStrategy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("retention_strategy")
			}
			return err
		}
	}

	return nil
}

func (m *IndexSetSummary) validateRotationStrategy(formats strfmt.Registry) error {

	if swag.IsZero(m.RotationStrategy) { // not required
		return nil
	}

	if m.RotationStrategy != nil {

		if err := m.RotationStrategy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rotation_strategy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IndexSetSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IndexSetSummary) UnmarshalBinary(b []byte) error {
	var res IndexSetSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}