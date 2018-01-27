// Code generated by go-swagger; DO NOT EDIT.

package cluster_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// GetJobReader is a Reader for the GetJob structure.
type GetJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetJobOK creates a GetJobOK with default headers values
func NewGetJobOK() *GetJobOK {
	return &GetJobOK{}
}

/*GetJobOK handles this case with default header values.

No response was specified
*/
type GetJobOK struct {
	Payload *models.SystemJobSummary
}

func (o *GetJobOK) Error() string {
	return fmt.Sprintf("[GET /cluster/jobs/{jobId}][%d] getJobOK  %+v", 200, o.Payload)
}

func (o *GetJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemJobSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}