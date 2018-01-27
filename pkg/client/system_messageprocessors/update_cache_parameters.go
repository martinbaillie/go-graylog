// Code generated by go-swagger; DO NOT EDIT.

package system_messageprocessors

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

	models "github.com/martinbaillie/go-graylog/pkg/models"
)

// NewUpdateCacheParams creates a new UpdateCacheParams object
// with the default values initialized.
func NewUpdateCacheParams() *UpdateCacheParams {
	var ()
	return &UpdateCacheParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCacheParamsWithTimeout creates a new UpdateCacheParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateCacheParamsWithTimeout(timeout time.Duration) *UpdateCacheParams {
	var ()
	return &UpdateCacheParams{

		timeout: timeout,
	}
}

// NewUpdateCacheParamsWithContext creates a new UpdateCacheParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateCacheParamsWithContext(ctx context.Context) *UpdateCacheParams {
	var ()
	return &UpdateCacheParams{

		Context: ctx,
	}
}

// NewUpdateCacheParamsWithHTTPClient creates a new UpdateCacheParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateCacheParamsWithHTTPClient(client *http.Client) *UpdateCacheParams {
	var ()
	return &UpdateCacheParams{
		HTTPClient: client,
	}
}

/*UpdateCacheParams contains all the parameters to send to the API endpoint
for the update cache operation typically these are written to a http.Request
*/
type UpdateCacheParams struct {

	/*Body*/
	Body *models.CacheAPI
	/*IDOrName*/
	IDOrName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update cache params
func (o *UpdateCacheParams) WithTimeout(timeout time.Duration) *UpdateCacheParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update cache params
func (o *UpdateCacheParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update cache params
func (o *UpdateCacheParams) WithContext(ctx context.Context) *UpdateCacheParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update cache params
func (o *UpdateCacheParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update cache params
func (o *UpdateCacheParams) WithHTTPClient(client *http.Client) *UpdateCacheParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update cache params
func (o *UpdateCacheParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update cache params
func (o *UpdateCacheParams) WithBody(body *models.CacheAPI) *UpdateCacheParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update cache params
func (o *UpdateCacheParams) SetBody(body *models.CacheAPI) {
	o.Body = body
}

// WithIDOrName adds the iDOrName to the update cache params
func (o *UpdateCacheParams) WithIDOrName(iDOrName string) *UpdateCacheParams {
	o.SetIDOrName(iDOrName)
	return o
}

// SetIDOrName adds the idOrName to the update cache params
func (o *UpdateCacheParams) SetIDOrName(iDOrName string) {
	o.IDOrName = iDOrName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCacheParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param idOrName
	if err := r.SetPathParam("idOrName", o.IDOrName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}