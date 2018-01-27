// Code generated by go-swagger; DO NOT EDIT.

package search_universal_keyword

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetAvailableReader is a Reader for the GetAvailable structure.
type GetAvailableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAvailableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAvailableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAvailableOK creates a GetAvailableOK with default headers values
func NewGetAvailableOK() *GetAvailableOK {
	return &GetAvailableOK{}
}

/*GetAvailableOK handles this case with default header values.

No response was specified
*/
type GetAvailableOK struct {
	Payload GetAvailableOKBody
}

func (o *GetAvailableOK) Error() string {
	return fmt.Sprintf("[GET /search/decorators/available][%d] getAvailableOK  %+v", 200, o.Payload)
}

func (o *GetAvailableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetAvailableOKBody get available o k body
swagger:model GetAvailableOKBody
*/
type GetAvailableOKBody interface{}