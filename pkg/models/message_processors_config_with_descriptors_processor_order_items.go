// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MessageProcessorsConfigWithDescriptorsProcessorOrderItems message processors config with descriptors processor order items
// swagger:model messageProcessorsConfigWithDescriptorsProcessorOrderItems
type MessageProcessorsConfigWithDescriptorsProcessorOrderItems struct {

	// class name
	ClassName string `json:"class_name,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this message processors config with descriptors processor order items
func (m *MessageProcessorsConfigWithDescriptorsProcessorOrderItems) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *MessageProcessorsConfigWithDescriptorsProcessorOrderItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MessageProcessorsConfigWithDescriptorsProcessorOrderItems) UnmarshalBinary(b []byte) error {
	var res MessageProcessorsConfigWithDescriptorsProcessorOrderItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
