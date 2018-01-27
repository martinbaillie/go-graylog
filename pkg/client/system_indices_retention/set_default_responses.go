// Code generated by go-swagger; DO NOT EDIT.

package system_indices_retention

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// SetDefaultReader is a Reader for the SetDefault structure.
type SetDefaultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetDefaultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetDefaultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewSetDefaultForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetDefaultOK creates a SetDefaultOK with default headers values
func NewSetDefaultOK() *SetDefaultOK {
	return &SetDefaultOK{}
}

/*SetDefaultOK handles this case with default header values.

No response was specified
*/
type SetDefaultOK struct {
	Payload *models.IndexSetSummary
}

func (o *SetDefaultOK) Error() string {
	return fmt.Sprintf("[PUT /system/indices/index_sets/{id}/default][%d] setDefaultOK  %+v", 200, o.Payload)
}

func (o *SetDefaultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IndexSetSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetDefaultForbidden creates a SetDefaultForbidden with default headers values
func NewSetDefaultForbidden() *SetDefaultForbidden {
	return &SetDefaultForbidden{}
}

/*SetDefaultForbidden handles this case with default header values.

Unauthorized
*/
type SetDefaultForbidden struct {
}

func (o *SetDefaultForbidden) Error() string {
	return fmt.Sprintf("[PUT /system/indices/index_sets/{id}/default][%d] setDefaultForbidden ", 403)
}

func (o *SetDefaultForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}