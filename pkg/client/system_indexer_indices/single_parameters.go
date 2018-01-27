// Code generated by go-swagger; DO NOT EDIT.

package system_indexer_indices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSingleParams creates a new SingleParams object
// with the default values initialized.
func NewSingleParams() *SingleParams {
	var ()
	return &SingleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSingleParamsWithTimeout creates a new SingleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSingleParamsWithTimeout(timeout time.Duration) *SingleParams {
	var ()
	return &SingleParams{

		timeout: timeout,
	}
}

// NewSingleParamsWithContext creates a new SingleParams object
// with the default values initialized, and the ability to set a context for a request
func NewSingleParamsWithContext(ctx context.Context) *SingleParams {
	var ()
	return &SingleParams{

		Context: ctx,
	}
}

// NewSingleParamsWithHTTPClient creates a new SingleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSingleParamsWithHTTPClient(client *http.Client) *SingleParams {
	var ()
	return &SingleParams{
		HTTPClient: client,
	}
}

/*SingleParams contains all the parameters to send to the API endpoint
for the single operation typically these are written to a http.Request
*/
type SingleParams struct {

	/*Limit
	  Limit

	*/
	Limit int64
	/*Offset
	  Offset

	*/
	Offset int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the single params
func (o *SingleParams) WithTimeout(timeout time.Duration) *SingleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the single params
func (o *SingleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the single params
func (o *SingleParams) WithContext(ctx context.Context) *SingleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the single params
func (o *SingleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the single params
func (o *SingleParams) WithHTTPClient(client *http.Client) *SingleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the single params
func (o *SingleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the single params
func (o *SingleParams) WithLimit(limit int64) *SingleParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the single params
func (o *SingleParams) SetLimit(limit int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the single params
func (o *SingleParams) WithOffset(offset int64) *SingleParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the single params
func (o *SingleParams) SetOffset(offset int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *SingleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param limit
	qrLimit := o.Limit
	qLimit := swag.FormatInt64(qrLimit)
	if qLimit != "" {
		if err := r.SetQueryParam("limit", qLimit); err != nil {
			return err
		}
	}

	// query param offset
	qrOffset := o.Offset
	qOffset := swag.FormatInt64(qrOffset)
	if qOffset != "" {
		if err := r.SetQueryParam("offset", qOffset); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}