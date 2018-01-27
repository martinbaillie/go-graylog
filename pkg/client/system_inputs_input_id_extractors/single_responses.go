// Code generated by go-swagger; DO NOT EDIT.

package system_inputs_input_id_extractors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// SingleReader is a Reader for the Single structure.
type SingleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SingleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSingleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewSingleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSingleOK creates a SingleOK with default headers values
func NewSingleOK() *SingleOK {
	return &SingleOK{}
}

/*SingleOK handles this case with default header values.

No response was specified
*/
type SingleOK struct {
	Payload *models.ExtractorSummary
}

func (o *SingleOK) Error() string {
	return fmt.Sprintf("[GET /system/inputs/{inputId}/extractors/{extractorId}][%d] singleOK  %+v", 200, o.Payload)
}

func (o *SingleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExtractorSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSingleNotFound creates a SingleNotFound with default headers values
func NewSingleNotFound() *SingleNotFound {
	return &SingleNotFound{}
}

/*SingleNotFound handles this case with default header values.

No such extractor on this input.
*/
type SingleNotFound struct {
}

func (o *SingleNotFound) Error() string {
	return fmt.Sprintf("[GET /system/inputs/{inputId}/extractors/{extractorId}][%d] singleNotFound ", 404)
}

func (o *SingleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}