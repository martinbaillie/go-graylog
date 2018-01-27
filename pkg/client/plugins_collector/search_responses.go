// Code generated by go-swagger; DO NOT EDIT.

package plugins_collector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SearchReader is a Reader for the Search structure.
type SearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSearchOK creates a SearchOK with default headers values
func NewSearchOK() *SearchOK {
	return &SearchOK{}
}

/*SearchOK handles this case with default header values.

No response was specified
*/
type SearchOK struct {
	Payload SearchOKBody
}

func (o *SearchOK) Error() string {
	return fmt.Sprintf("[GET /messages/{index}/{messageId}][%d] searchOK  %+v", 200, o.Payload)
}

func (o *SearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchNotFound creates a SearchNotFound with default headers values
func NewSearchNotFound() *SearchNotFound {
	return &SearchNotFound{}
}

/*SearchNotFound handles this case with default header values.

Message does not exist.
*/
type SearchNotFound struct {
}

func (o *SearchNotFound) Error() string {
	return fmt.Sprintf("[GET /messages/{index}/{messageId}][%d] searchNotFound ", 404)
}

func (o *SearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*SearchOKBody search o k body
swagger:model SearchOKBody
*/
type SearchOKBody interface{}