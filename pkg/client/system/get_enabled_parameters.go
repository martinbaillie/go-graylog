// Code generated by go-swagger; DO NOT EDIT.

package system

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

// NewGetEnabledParams creates a new GetEnabledParams object
// with the default values initialized.
func NewGetEnabledParams() *GetEnabledParams {

	return &GetEnabledParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEnabledParamsWithTimeout creates a new GetEnabledParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEnabledParamsWithTimeout(timeout time.Duration) *GetEnabledParams {

	return &GetEnabledParams{

		timeout: timeout,
	}
}

// NewGetEnabledParamsWithContext creates a new GetEnabledParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEnabledParamsWithContext(ctx context.Context) *GetEnabledParams {

	return &GetEnabledParams{

		Context: ctx,
	}
}

// NewGetEnabledParamsWithHTTPClient creates a new GetEnabledParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEnabledParamsWithHTTPClient(client *http.Client) *GetEnabledParams {

	return &GetEnabledParams{
		HTTPClient: client,
	}
}

/*GetEnabledParams contains all the parameters to send to the API endpoint
for the get enabled operation typically these are written to a http.Request
*/
type GetEnabledParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get enabled params
func (o *GetEnabledParams) WithTimeout(timeout time.Duration) *GetEnabledParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get enabled params
func (o *GetEnabledParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get enabled params
func (o *GetEnabledParams) WithContext(ctx context.Context) *GetEnabledParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get enabled params
func (o *GetEnabledParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get enabled params
func (o *GetEnabledParams) WithHTTPClient(client *http.Client) *GetEnabledParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get enabled params
func (o *GetEnabledParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetEnabledParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}