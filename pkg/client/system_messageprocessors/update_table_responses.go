// Code generated by go-swagger; DO NOT EDIT.

package system_messageprocessors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// UpdateTableReader is a Reader for the UpdateTable structure.
type UpdateTableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateTableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateTableOK creates a UpdateTableOK with default headers values
func NewUpdateTableOK() *UpdateTableOK {
	return &UpdateTableOK{}
}

/*UpdateTableOK handles this case with default header values.

No response was specified
*/
type UpdateTableOK struct {
	Payload *models.LookupTableAPI
}

func (o *UpdateTableOK) Error() string {
	return fmt.Sprintf("[PUT /system/lookup/tables/{idOrName}][%d] updateTableOK  %+v", 200, o.Payload)
}

func (o *UpdateTableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LookupTableAPI)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}