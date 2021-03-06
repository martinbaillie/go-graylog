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

// UpdateAdapterReader is a Reader for the UpdateAdapter structure.
type UpdateAdapterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAdapterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateAdapterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateAdapterOK creates a UpdateAdapterOK with default headers values
func NewUpdateAdapterOK() *UpdateAdapterOK {
	return &UpdateAdapterOK{}
}

/*UpdateAdapterOK handles this case with default header values.

No response was specified
*/
type UpdateAdapterOK struct {
	Payload *models.DataAdapterAPI
}

func (o *UpdateAdapterOK) Error() string {
	return fmt.Sprintf("[PUT /system/lookup/adapters/{idOrName}][%d] updateAdapterOK  %+v", 200, o.Payload)
}

func (o *UpdateAdapterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataAdapterAPI)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
