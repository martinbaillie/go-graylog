// Code generated by go-swagger; DO NOT EDIT.

package streams_streamid_outputs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CheckConditionsReader is a Reader for the CheckConditions structure.
type CheckConditionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckConditionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCheckConditionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCheckConditionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCheckConditionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCheckConditionsOK creates a CheckConditionsOK with default headers values
func NewCheckConditionsOK() *CheckConditionsOK {
	return &CheckConditionsOK{}
}

/*CheckConditionsOK handles this case with default header values.

No response was specified
*/
type CheckConditionsOK struct {
	Payload CheckConditionsOKBody
}

func (o *CheckConditionsOK) Error() string {
	return fmt.Sprintf("[GET /streams/{streamId}/alerts/check][%d] checkConditionsOK  %+v", 200, o.Payload)
}

func (o *CheckConditionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckConditionsBadRequest creates a CheckConditionsBadRequest with default headers values
func NewCheckConditionsBadRequest() *CheckConditionsBadRequest {
	return &CheckConditionsBadRequest{}
}

/*CheckConditionsBadRequest handles this case with default header values.

Invalid ObjectId.
*/
type CheckConditionsBadRequest struct {
}

func (o *CheckConditionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /streams/{streamId}/alerts/check][%d] checkConditionsBadRequest ", 400)
}

func (o *CheckConditionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCheckConditionsNotFound creates a CheckConditionsNotFound with default headers values
func NewCheckConditionsNotFound() *CheckConditionsNotFound {
	return &CheckConditionsNotFound{}
}

/*CheckConditionsNotFound handles this case with default header values.

Stream not found.
*/
type CheckConditionsNotFound struct {
}

func (o *CheckConditionsNotFound) Error() string {
	return fmt.Sprintf("[GET /streams/{streamId}/alerts/check][%d] checkConditionsNotFound ", 404)
}

func (o *CheckConditionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*CheckConditionsOKBody check conditions o k body
swagger:model CheckConditionsOKBody
*/
type CheckConditionsOKBody interface{}
