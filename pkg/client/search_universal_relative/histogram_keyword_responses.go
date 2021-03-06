// Code generated by go-swagger; DO NOT EDIT.

package search_universal_relative

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// HistogramKeywordReader is a Reader for the HistogramKeyword structure.
type HistogramKeywordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HistogramKeywordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewHistogramKeywordOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewHistogramKeywordBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHistogramKeywordOK creates a HistogramKeywordOK with default headers values
func NewHistogramKeywordOK() *HistogramKeywordOK {
	return &HistogramKeywordOK{}
}

/*HistogramKeywordOK handles this case with default header values.

No response was specified
*/
type HistogramKeywordOK struct {
	Payload *models.HistogramResult
}

func (o *HistogramKeywordOK) Error() string {
	return fmt.Sprintf("[GET /search/universal/keyword/histogram][%d] histogramKeywordOK  %+v", 200, o.Payload)
}

func (o *HistogramKeywordOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HistogramResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHistogramKeywordBadRequest creates a HistogramKeywordBadRequest with default headers values
func NewHistogramKeywordBadRequest() *HistogramKeywordBadRequest {
	return &HistogramKeywordBadRequest{}
}

/*HistogramKeywordBadRequest handles this case with default header values.

Invalid interval provided.
*/
type HistogramKeywordBadRequest struct {
}

func (o *HistogramKeywordBadRequest) Error() string {
	return fmt.Sprintf("[GET /search/universal/keyword/histogram][%d] histogramKeywordBadRequest ", 400)
}

func (o *HistogramKeywordBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
