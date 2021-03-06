// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PauseReader is a Reader for the Pause structure.
type PauseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PauseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPauseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPauseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPauseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPauseOK creates a PauseOK with default headers values
func NewPauseOK() *PauseOK {
	return &PauseOK{}
}

/*PauseOK handles this case with default header values.

No response was specified
*/
type PauseOK struct {
}

func (o *PauseOK) Error() string {
	return fmt.Sprintf("[POST /streams/{streamId}/pause][%d] pauseOK ", 200)
}

func (o *PauseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPauseBadRequest creates a PauseBadRequest with default headers values
func NewPauseBadRequest() *PauseBadRequest {
	return &PauseBadRequest{}
}

/*PauseBadRequest handles this case with default header values.

Invalid or missing Stream id.
*/
type PauseBadRequest struct {
}

func (o *PauseBadRequest) Error() string {
	return fmt.Sprintf("[POST /streams/{streamId}/pause][%d] pauseBadRequest ", 400)
}

func (o *PauseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPauseNotFound creates a PauseNotFound with default headers values
func NewPauseNotFound() *PauseNotFound {
	return &PauseNotFound{}
}

/*PauseNotFound handles this case with default header values.

Stream not found.
*/
type PauseNotFound struct {
}

func (o *PauseNotFound) Error() string {
	return fmt.Sprintf("[POST /streams/{streamId}/pause][%d] pauseNotFound ", 404)
}

func (o *PauseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
