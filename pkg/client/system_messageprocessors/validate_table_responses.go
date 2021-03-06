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

// ValidateTableReader is a Reader for the ValidateTable structure.
type ValidateTableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidateTableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewValidateTableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewValidateTableOK creates a ValidateTableOK with default headers values
func NewValidateTableOK() *ValidateTableOK {
	return &ValidateTableOK{}
}

/*ValidateTableOK handles this case with default header values.

No response was specified
*/
type ValidateTableOK struct {
	Payload *models.ValidationResult
}

func (o *ValidateTableOK) Error() string {
	return fmt.Sprintf("[POST /system/lookup/tables/validate][%d] validateTableOK  %+v", 200, o.Payload)
}

func (o *ValidateTableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
