// Code generated by go-swagger; DO NOT EDIT.

package system_indexer_overview

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// IndexSetOpenReader is a Reader for the IndexSetOpen structure.
type IndexSetOpenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexSetOpenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIndexSetOpenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIndexSetOpenOK creates a IndexSetOpenOK with default headers values
func NewIndexSetOpenOK() *IndexSetOpenOK {
	return &IndexSetOpenOK{}
}

/*IndexSetOpenOK handles this case with default header values.

No response was specified
*/
type IndexSetOpenOK struct {
	Payload *models.OpenIndicesInfo
}

func (o *IndexSetOpenOK) Error() string {
	return fmt.Sprintf("[GET /system/indexer/indices/{indexSetId}/open][%d] indexSetOpenOK  %+v", 200, o.Payload)
}

func (o *IndexSetOpenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenIndicesInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}