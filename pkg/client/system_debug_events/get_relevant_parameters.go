// Code generated by go-swagger; DO NOT EDIT.

package system_debug_events

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

// NewGetRelevantParams creates a new GetRelevantParams object
// with the default values initialized.
func NewGetRelevantParams() *GetRelevantParams {

	return &GetRelevantParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRelevantParamsWithTimeout creates a new GetRelevantParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRelevantParamsWithTimeout(timeout time.Duration) *GetRelevantParams {

	return &GetRelevantParams{

		timeout: timeout,
	}
}

// NewGetRelevantParamsWithContext creates a new GetRelevantParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRelevantParamsWithContext(ctx context.Context) *GetRelevantParams {

	return &GetRelevantParams{

		Context: ctx,
	}
}

// NewGetRelevantParamsWithHTTPClient creates a new GetRelevantParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRelevantParamsWithHTTPClient(client *http.Client) *GetRelevantParams {

	return &GetRelevantParams{
		HTTPClient: client,
	}
}

/*GetRelevantParams contains all the parameters to send to the API endpoint
for the get relevant operation typically these are written to a http.Request
*/
type GetRelevantParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get relevant params
func (o *GetRelevantParams) WithTimeout(timeout time.Duration) *GetRelevantParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get relevant params
func (o *GetRelevantParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get relevant params
func (o *GetRelevantParams) WithContext(ctx context.Context) *GetRelevantParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get relevant params
func (o *GetRelevantParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get relevant params
func (o *GetRelevantParams) WithHTTPClient(client *http.Client) *GetRelevantParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get relevant params
func (o *GetRelevantParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetRelevantParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}