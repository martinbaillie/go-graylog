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

// RemoveTableReader is a Reader for the RemoveTable structure.
type RemoveTableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveTableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRemoveTableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveTableOK creates a RemoveTableOK with default headers values
func NewRemoveTableOK() *RemoveTableOK {
	return &RemoveTableOK{}
}

/*RemoveTableOK handles this case with default header values.

No response was specified
*/
type RemoveTableOK struct {
	Payload *models.LookupTableAPI
}

func (o *RemoveTableOK) Error() string {
	return fmt.Sprintf("[DELETE /system/lookup/tables/{idOrName}][%d] removeTableOK  %+v", 200, o.Payload)
}

func (o *RemoveTableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LookupTableAPI)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
