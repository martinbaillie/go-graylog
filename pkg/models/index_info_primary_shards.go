// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IndexInfoPrimaryShards index info primary shards
// swagger:model indexInfoPrimaryShards
type IndexInfoPrimaryShards struct {

	// documents
	Documents *IndexInfoPrimaryShardsDocuments `json:"documents,omitempty"`

	// flush
	Flush *IndexInfoPrimaryShardsFlush `json:"flush,omitempty"`

	// get
	Get interface{} `json:"get,omitempty"`

	// index
	Index interface{} `json:"index,omitempty"`

	// merge
	Merge interface{} `json:"merge,omitempty"`

	// open search contexts
	OpenSearchContexts int64 `json:"open_search_contexts,omitempty"`

	// refresh
	Refresh interface{} `json:"refresh,omitempty"`

	// search fetch
	SearchFetch interface{} `json:"search_fetch,omitempty"`

	// search query
	SearchQuery interface{} `json:"search_query,omitempty"`

	// segments
	Segments int64 `json:"segments,omitempty"`

	// store size bytes
	StoreSizeBytes int64 `json:"store_size_bytes,omitempty"`
}

// Validate validates this index info primary shards
func (m *IndexInfoPrimaryShards) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocuments(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFlush(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IndexInfoPrimaryShards) validateDocuments(formats strfmt.Registry) error {

	if swag.IsZero(m.Documents) { // not required
		return nil
	}

	if m.Documents != nil {

		if err := m.Documents.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("documents")
			}
			return err
		}
	}

	return nil
}

func (m *IndexInfoPrimaryShards) validateFlush(formats strfmt.Registry) error {

	if swag.IsZero(m.Flush) { // not required
		return nil
	}

	if m.Flush != nil {

		if err := m.Flush.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flush")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IndexInfoPrimaryShards) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IndexInfoPrimaryShards) UnmarshalBinary(b []byte) error {
	var res IndexInfoPrimaryShards
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}