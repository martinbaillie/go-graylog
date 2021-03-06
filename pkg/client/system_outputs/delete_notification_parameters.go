// Code generated by go-swagger; DO NOT EDIT.

package system_outputs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteNotificationParams creates a new DeleteNotificationParams object
// with the default values initialized.
func NewDeleteNotificationParams() *DeleteNotificationParams {
	var ()
	return &DeleteNotificationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNotificationParamsWithTimeout creates a new DeleteNotificationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteNotificationParamsWithTimeout(timeout time.Duration) *DeleteNotificationParams {
	var ()
	return &DeleteNotificationParams{

		timeout: timeout,
	}
}

// NewDeleteNotificationParamsWithContext creates a new DeleteNotificationParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteNotificationParamsWithContext(ctx context.Context) *DeleteNotificationParams {
	var ()
	return &DeleteNotificationParams{

		Context: ctx,
	}
}

// NewDeleteNotificationParamsWithHTTPClient creates a new DeleteNotificationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteNotificationParamsWithHTTPClient(client *http.Client) *DeleteNotificationParams {
	var ()
	return &DeleteNotificationParams{
		HTTPClient: client,
	}
}

/*DeleteNotificationParams contains all the parameters to send to the API endpoint
for the delete notification operation typically these are written to a http.Request
*/
type DeleteNotificationParams struct {

	/*NotificationType*/
	NotificationType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete notification params
func (o *DeleteNotificationParams) WithTimeout(timeout time.Duration) *DeleteNotificationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete notification params
func (o *DeleteNotificationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete notification params
func (o *DeleteNotificationParams) WithContext(ctx context.Context) *DeleteNotificationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete notification params
func (o *DeleteNotificationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete notification params
func (o *DeleteNotificationParams) WithHTTPClient(client *http.Client) *DeleteNotificationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete notification params
func (o *DeleteNotificationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNotificationType adds the notificationType to the delete notification params
func (o *DeleteNotificationParams) WithNotificationType(notificationType string) *DeleteNotificationParams {
	o.SetNotificationType(notificationType)
	return o
}

// SetNotificationType adds the notificationType to the delete notification params
func (o *DeleteNotificationParams) SetNotificationType(notificationType string) {
	o.NotificationType = notificationType
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNotificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param notificationType
	if err := r.SetPathParam("notificationType", o.NotificationType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
