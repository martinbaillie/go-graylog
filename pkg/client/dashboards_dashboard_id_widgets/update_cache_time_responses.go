// Code generated by go-swagger; DO NOT EDIT.

package dashboards_dashboard_id_widgets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateCacheTimeReader is a Reader for the UpdateCacheTime structure.
type UpdateCacheTimeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCacheTimeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateCacheTimeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewUpdateCacheTimeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateCacheTimeOK creates a UpdateCacheTimeOK with default headers values
func NewUpdateCacheTimeOK() *UpdateCacheTimeOK {
	return &UpdateCacheTimeOK{}
}

/*UpdateCacheTimeOK handles this case with default header values.

No response was specified
*/
type UpdateCacheTimeOK struct {
}

func (o *UpdateCacheTimeOK) Error() string {
	return fmt.Sprintf("[PUT /dashboards/{dashboardId}/widgets/{widgetId}/cachetime][%d] updateCacheTimeOK ", 200)
}

func (o *UpdateCacheTimeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateCacheTimeNotFound creates a UpdateCacheTimeNotFound with default headers values
func NewUpdateCacheTimeNotFound() *UpdateCacheTimeNotFound {
	return &UpdateCacheTimeNotFound{}
}

/*UpdateCacheTimeNotFound handles this case with default header values.

Widget not found.
*/
type UpdateCacheTimeNotFound struct {
}

func (o *UpdateCacheTimeNotFound) Error() string {
	return fmt.Sprintf("[PUT /dashboards/{dashboardId}/widgets/{widgetId}/cachetime][%d] updateCacheTimeNotFound ", 404)
}

func (o *UpdateCacheTimeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}