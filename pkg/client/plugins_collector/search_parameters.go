// Code generated by go-swagger; DO NOT EDIT.

package plugins_collector

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

// NewSearchParams creates a new SearchParams object
// with the default values initialized.
func NewSearchParams() *SearchParams {
	var ()
	return &SearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSearchParamsWithTimeout creates a new SearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSearchParamsWithTimeout(timeout time.Duration) *SearchParams {
	var ()
	return &SearchParams{

		timeout: timeout,
	}
}

// NewSearchParamsWithContext creates a new SearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewSearchParamsWithContext(ctx context.Context) *SearchParams {
	var ()
	return &SearchParams{

		Context: ctx,
	}
}

// NewSearchParamsWithHTTPClient creates a new SearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSearchParamsWithHTTPClient(client *http.Client) *SearchParams {
	var ()
	return &SearchParams{
		HTTPClient: client,
	}
}

/*SearchParams contains all the parameters to send to the API endpoint
for the search operation typically these are written to a http.Request
*/
type SearchParams struct {

	/*Index
	  The index this message is stored in.

	*/
	Index string
	/*MessageID*/
	MessageID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the search params
func (o *SearchParams) WithTimeout(timeout time.Duration) *SearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search params
func (o *SearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search params
func (o *SearchParams) WithContext(ctx context.Context) *SearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search params
func (o *SearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search params
func (o *SearchParams) WithHTTPClient(client *http.Client) *SearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search params
func (o *SearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIndex adds the index to the search params
func (o *SearchParams) WithIndex(index string) *SearchParams {
	o.SetIndex(index)
	return o
}

// SetIndex adds the index to the search params
func (o *SearchParams) SetIndex(index string) {
	o.Index = index
}

// WithMessageID adds the messageID to the search params
func (o *SearchParams) WithMessageID(messageID string) *SearchParams {
	o.SetMessageID(messageID)
	return o
}

// SetMessageID adds the messageId to the search params
func (o *SearchParams) SetMessageID(messageID string) {
	o.MessageID = messageID
}

// WriteToRequest writes these params to a swagger request
func (o *SearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param index
	if err := r.SetPathParam("index", o.Index); err != nil {
		return err
	}

	// path param messageId
	if err := r.SetPathParam("messageId", o.MessageID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}