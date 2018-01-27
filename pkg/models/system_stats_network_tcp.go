// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SystemStatsNetworkTCP system stats network Tcp
// swagger:model systemStatsNetworkTcp
type SystemStatsNetworkTCP struct {

	// active opens
	ActiveOpens int64 `json:"active_opens,omitempty"`

	// attempt fails
	AttemptFails int64 `json:"attempt_fails,omitempty"`

	// curr estab
	CurrEstab int64 `json:"curr_estab,omitempty"`

	// estab resets
	EstabResets int64 `json:"estab_resets,omitempty"`

	// in errs
	InErrs int64 `json:"in_errs,omitempty"`

	// in segs
	InSegs int64 `json:"in_segs,omitempty"`

	// out rsts
	OutRsts int64 `json:"out_rsts,omitempty"`

	// out segs
	OutSegs int64 `json:"out_segs,omitempty"`

	// passive opens
	PassiveOpens int64 `json:"passive_opens,omitempty"`

	// retrans segs
	RetransSegs int64 `json:"retrans_segs,omitempty"`
}

// Validate validates this system stats network Tcp
func (m *SystemStatsNetworkTCP) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SystemStatsNetworkTCP) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemStatsNetworkTCP) UnmarshalBinary(b []byte) error {
	var res SystemStatsNetworkTCP
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}