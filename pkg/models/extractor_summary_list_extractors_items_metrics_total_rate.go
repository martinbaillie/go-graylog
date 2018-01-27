// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ExtractorSummaryListExtractorsItemsMetricsTotalRate extractor summary list extractors items metrics total rate
// swagger:model extractorSummaryListExtractorsItemsMetricsTotalRate
type ExtractorSummaryListExtractorsItemsMetricsTotalRate struct {

	// fifteen minute
	FifteenMinute float64 `json:"fifteen_minute,omitempty"`

	// five minute
	FiveMinute float64 `json:"five_minute,omitempty"`

	// mean
	Mean float64 `json:"mean,omitempty"`

	// one minute
	OneMinute float64 `json:"one_minute,omitempty"`

	// total
	Total float64 `json:"total,omitempty"`
}

// Validate validates this extractor summary list extractors items metrics total rate
func (m *ExtractorSummaryListExtractorsItemsMetricsTotalRate) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ExtractorSummaryListExtractorsItemsMetricsTotalRate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExtractorSummaryListExtractorsItemsMetricsTotalRate) UnmarshalBinary(b []byte) error {
	var res ExtractorSummaryListExtractorsItemsMetricsTotalRate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}