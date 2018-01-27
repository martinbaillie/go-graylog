// Code generated by go-swagger; DO NOT EDIT.

package system_indices_retention

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

// NewGlobalStatsParams creates a new GlobalStatsParams object
// with the default values initialized.
func NewGlobalStatsParams() *GlobalStatsParams {

	return &GlobalStatsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGlobalStatsParamsWithTimeout creates a new GlobalStatsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGlobalStatsParamsWithTimeout(timeout time.Duration) *GlobalStatsParams {

	return &GlobalStatsParams{

		timeout: timeout,
	}
}

// NewGlobalStatsParamsWithContext creates a new GlobalStatsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGlobalStatsParamsWithContext(ctx context.Context) *GlobalStatsParams {

	return &GlobalStatsParams{

		Context: ctx,
	}
}

// NewGlobalStatsParamsWithHTTPClient creates a new GlobalStatsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGlobalStatsParamsWithHTTPClient(client *http.Client) *GlobalStatsParams {

	return &GlobalStatsParams{
		HTTPClient: client,
	}
}

/*GlobalStatsParams contains all the parameters to send to the API endpoint
for the global stats operation typically these are written to a http.Request
*/
type GlobalStatsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the global stats params
func (o *GlobalStatsParams) WithTimeout(timeout time.Duration) *GlobalStatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the global stats params
func (o *GlobalStatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the global stats params
func (o *GlobalStatsParams) WithContext(ctx context.Context) *GlobalStatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the global stats params
func (o *GlobalStatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the global stats params
func (o *GlobalStatsParams) WithHTTPClient(client *http.Client) *GlobalStatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the global stats params
func (o *GlobalStatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GlobalStatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}