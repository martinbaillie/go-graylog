// Code generated by go-swagger; DO NOT EDIT.

package system_lookup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// SubsystemsReader is a Reader for the Subsystems structure.
type SubsystemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubsystemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSubsystemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSubsystemsOK creates a SubsystemsOK with default headers values
func NewSubsystemsOK() *SubsystemsOK {
	return &SubsystemsOK{}
}

/*SubsystemsOK handles this case with default header values.

No response was specified
*/
type SubsystemsOK struct {
	Payload *models.SubsystemSummary
}

func (o *SubsystemsOK) Error() string {
	return fmt.Sprintf("[GET /system/loggers/subsystems][%d] subsystemsOK  %+v", 200, o.Payload)
}

func (o *SubsystemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubsystemSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}