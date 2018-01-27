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

// NewUpdateOutputParams creates a new UpdateOutputParams object
// with the default values initialized.
func NewUpdateOutputParams() *UpdateOutputParams {
	var ()
	return &UpdateOutputParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOutputParamsWithTimeout creates a new UpdateOutputParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateOutputParamsWithTimeout(timeout time.Duration) *UpdateOutputParams {
	var ()
	return &UpdateOutputParams{

		timeout: timeout,
	}
}

// NewUpdateOutputParamsWithContext creates a new UpdateOutputParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateOutputParamsWithContext(ctx context.Context) *UpdateOutputParams {
	var ()
	return &UpdateOutputParams{

		Context: ctx,
	}
}

// NewUpdateOutputParamsWithHTTPClient creates a new UpdateOutputParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateOutputParamsWithHTTPClient(client *http.Client) *UpdateOutputParams {
	var ()
	return &UpdateOutputParams{
		HTTPClient: client,
	}
}

/*UpdateOutputParams contains all the parameters to send to the API endpoint
for the update output operation typically these are written to a http.Request
*/
type UpdateOutputParams struct {

	/*JSONBody*/
	JSONBody interface{}
	/*ID*/
	ID string
	/*OutputID*/
	OutputID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update output params
func (o *UpdateOutputParams) WithTimeout(timeout time.Duration) *UpdateOutputParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update output params
func (o *UpdateOutputParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update output params
func (o *UpdateOutputParams) WithContext(ctx context.Context) *UpdateOutputParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update output params
func (o *UpdateOutputParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update output params
func (o *UpdateOutputParams) WithHTTPClient(client *http.Client) *UpdateOutputParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update output params
func (o *UpdateOutputParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJSONBody adds the jSONBody to the update output params
func (o *UpdateOutputParams) WithJSONBody(jSONBody interface{}) *UpdateOutputParams {
	o.SetJSONBody(jSONBody)
	return o
}

// SetJSONBody adds the jsonBody to the update output params
func (o *UpdateOutputParams) SetJSONBody(jSONBody interface{}) {
	o.JSONBody = jSONBody
}

// WithID adds the id to the update output params
func (o *UpdateOutputParams) WithID(id string) *UpdateOutputParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update output params
func (o *UpdateOutputParams) SetID(id string) {
	o.ID = id
}

// WithOutputID adds the outputID to the update output params
func (o *UpdateOutputParams) WithOutputID(outputID string) *UpdateOutputParams {
	o.SetOutputID(outputID)
	return o
}

// SetOutputID adds the outputId to the update output params
func (o *UpdateOutputParams) SetOutputID(outputID string) {
	o.OutputID = outputID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOutputParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.JSONBody != nil {
		if err := r.SetBodyParam(o.JSONBody); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param output_id
	if err := r.SetPathParam("output_id", o.OutputID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}