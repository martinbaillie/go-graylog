// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewDeletePermissionsParams creates a new DeletePermissionsParams object
// with the default values initialized.
func NewDeletePermissionsParams() *DeletePermissionsParams {
	var ()
	return &DeletePermissionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePermissionsParamsWithTimeout creates a new DeletePermissionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePermissionsParamsWithTimeout(timeout time.Duration) *DeletePermissionsParams {
	var ()
	return &DeletePermissionsParams{

		timeout: timeout,
	}
}

// NewDeletePermissionsParamsWithContext creates a new DeletePermissionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePermissionsParamsWithContext(ctx context.Context) *DeletePermissionsParams {
	var ()
	return &DeletePermissionsParams{

		Context: ctx,
	}
}

// NewDeletePermissionsParamsWithHTTPClient creates a new DeletePermissionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePermissionsParamsWithHTTPClient(client *http.Client) *DeletePermissionsParams {
	var ()
	return &DeletePermissionsParams{
		HTTPClient: client,
	}
}

/*DeletePermissionsParams contains all the parameters to send to the API endpoint
for the delete permissions operation typically these are written to a http.Request
*/
type DeletePermissionsParams struct {

	/*Username
	  The name of the user to modify.

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete permissions params
func (o *DeletePermissionsParams) WithTimeout(timeout time.Duration) *DeletePermissionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete permissions params
func (o *DeletePermissionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete permissions params
func (o *DeletePermissionsParams) WithContext(ctx context.Context) *DeletePermissionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete permissions params
func (o *DeletePermissionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete permissions params
func (o *DeletePermissionsParams) WithHTTPClient(client *http.Client) *DeletePermissionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete permissions params
func (o *DeletePermissionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUsername adds the username to the delete permissions params
func (o *DeletePermissionsParams) WithUsername(username string) *DeletePermissionsParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the delete permissions params
func (o *DeletePermissionsParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePermissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}