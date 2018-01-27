// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClusterStatsMongoHostInfo cluster stats mongo host info
// swagger:model clusterStatsMongoHostInfo
type ClusterStatsMongoHostInfo struct {

	// extra
	Extra *ClusterStatsMongoHostInfoExtra `json:"extra,omitempty"`

	// os
	Os *ClusterStatsMongoHostInfoOs `json:"os,omitempty"`

	// system
	System *ClusterStatsMongoHostInfoSystem `json:"system,omitempty"`
}

// Validate validates this cluster stats mongo host info
func (m *ClusterStatsMongoHostInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExtra(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSystem(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterStatsMongoHostInfo) validateExtra(formats strfmt.Registry) error {

	if swag.IsZero(m.Extra) { // not required
		return nil
	}

	if m.Extra != nil {

		if err := m.Extra.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extra")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterStatsMongoHostInfo) validateOs(formats strfmt.Registry) error {

	if swag.IsZero(m.Os) { // not required
		return nil
	}

	if m.Os != nil {

		if err := m.Os.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("os")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterStatsMongoHostInfo) validateSystem(formats strfmt.Registry) error {

	if swag.IsZero(m.System) { // not required
		return nil
	}

	if m.System != nil {

		if err := m.System.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("system")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterStatsMongoHostInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterStatsMongoHostInfo) UnmarshalBinary(b []byte) error {
	var res ClusterStatsMongoHostInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}