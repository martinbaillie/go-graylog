// Code generated by go-swagger; DO NOT EDIT.

package system_cluster_traffic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// MongoStatsReader is a Reader for the MongoStats structure.
type MongoStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MongoStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewMongoStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMongoStatsOK creates a MongoStatsOK with default headers values
func NewMongoStatsOK() *MongoStatsOK {
	return &MongoStatsOK{}
}

/*MongoStatsOK handles this case with default header values.

No response was specified
*/
type MongoStatsOK struct {
	Payload *models.MongoStats
}

func (o *MongoStatsOK) Error() string {
	return fmt.Sprintf("[GET /system/cluster/stats/mongo][%d] mongoStatsOK  %+v", 200, o.Payload)
}

func (o *MongoStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MongoStats)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}