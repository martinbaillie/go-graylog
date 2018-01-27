// Code generated by go-swagger; DO NOT EDIT.

package system_cluster_stats

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SchemaReader is a Reader for the Schema structure.
type SchemaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SchemaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSchemaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSchemaOK creates a SchemaOK with default headers values
func NewSchemaOK() *SchemaOK {
	return &SchemaOK{}
}

/*SchemaOK handles this case with default header values.

No response was specified
*/
type SchemaOK struct {
	Payload SchemaOKBody
}

func (o *SchemaOK) Error() string {
	return fmt.Sprintf("[GET /system/cluster_config/{configClass}][%d] schemaOK  %+v", 200, o.Payload)
}

func (o *SchemaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SchemaOKBody schema o k body
swagger:model SchemaOKBody
*/
type SchemaOKBody interface{}