// Code generated by go-swagger; DO NOT EDIT.

package plugins_map_mapdata

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

// NewListConfigurationsParams creates a new ListConfigurationsParams object
// with the default values initialized.
func NewListConfigurationsParams() *ListConfigurationsParams {

	return &ListConfigurationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListConfigurationsParamsWithTimeout creates a new ListConfigurationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListConfigurationsParamsWithTimeout(timeout time.Duration) *ListConfigurationsParams {

	return &ListConfigurationsParams{

		timeout: timeout,
	}
}

// NewListConfigurationsParamsWithContext creates a new ListConfigurationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListConfigurationsParamsWithContext(ctx context.Context) *ListConfigurationsParams {

	return &ListConfigurationsParams{

		Context: ctx,
	}
}

// NewListConfigurationsParamsWithHTTPClient creates a new ListConfigurationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListConfigurationsParamsWithHTTPClient(client *http.Client) *ListConfigurationsParams {

	return &ListConfigurationsParams{
		HTTPClient: client,
	}
}

/*ListConfigurationsParams contains all the parameters to send to the API endpoint
for the list configurations operation typically these are written to a http.Request
*/
type ListConfigurationsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list configurations params
func (o *ListConfigurationsParams) WithTimeout(timeout time.Duration) *ListConfigurationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list configurations params
func (o *ListConfigurationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list configurations params
func (o *ListConfigurationsParams) WithContext(ctx context.Context) *ListConfigurationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list configurations params
func (o *ListConfigurationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list configurations params
func (o *ListConfigurationsParams) WithHTTPClient(client *http.Client) *ListConfigurationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list configurations params
func (o *ListConfigurationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListConfigurationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
