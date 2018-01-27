// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SystemStatsProcess system stats process
// swagger:model systemStatsProcess
type SystemStatsProcess struct {

	// cpu
	CPU *SystemStatsProcessCPU `json:"cpu,omitempty"`

	// max file descriptors
	MaxFileDescriptors int64 `json:"max_file_descriptors,omitempty"`

	// memory
	Memory *SystemStatsProcessMemory `json:"memory,omitempty"`

	// open file descriptors
	OpenFileDescriptors int64 `json:"open_file_descriptors,omitempty"`

	// pid
	Pid int64 `json:"pid,omitempty"`
}

// Validate validates this system stats process
func (m *SystemStatsProcess) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPU(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemStatsProcess) validateCPU(formats strfmt.Registry) error {

	if swag.IsZero(m.CPU) { // not required
		return nil
	}

	if m.CPU != nil {

		if err := m.CPU.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cpu")
			}
			return err
		}
	}

	return nil
}

func (m *SystemStatsProcess) validateMemory(formats strfmt.Registry) error {

	if swag.IsZero(m.Memory) { // not required
		return nil
	}

	if m.Memory != nil {

		if err := m.Memory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("memory")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SystemStatsProcess) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemStatsProcess) UnmarshalBinary(b []byte) error {
	var res SystemStatsProcess
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}