// Code generated by go-swagger; DO NOT EDIT.

package search_decorators

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

// NewSearchAbsoluteParams creates a new SearchAbsoluteParams object
// with the default values initialized.
func NewSearchAbsoluteParams() *SearchAbsoluteParams {
	var (
		decorateDefault = bool(true)
	)
	return &SearchAbsoluteParams{
		Decorate: &decorateDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewSearchAbsoluteParamsWithTimeout creates a new SearchAbsoluteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSearchAbsoluteParamsWithTimeout(timeout time.Duration) *SearchAbsoluteParams {
	var (
		decorateDefault = bool(true)
	)
	return &SearchAbsoluteParams{
		Decorate: &decorateDefault,

		timeout: timeout,
	}
}

// NewSearchAbsoluteParamsWithContext creates a new SearchAbsoluteParams object
// with the default values initialized, and the ability to set a context for a request
func NewSearchAbsoluteParamsWithContext(ctx context.Context) *SearchAbsoluteParams {
	var (
		decorateDefault = bool(true)
	)
	return &SearchAbsoluteParams{
		Decorate: &decorateDefault,

		Context: ctx,
	}
}

// NewSearchAbsoluteParamsWithHTTPClient creates a new SearchAbsoluteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSearchAbsoluteParamsWithHTTPClient(client *http.Client) *SearchAbsoluteParams {
	var (
		decorateDefault = bool(true)
	)
	return &SearchAbsoluteParams{
		Decorate:   &decorateDefault,
		HTTPClient: client,
	}
}

/*SearchAbsoluteParams contains all the parameters to send to the API endpoint
for the search absolute operation typically these are written to a http.Request
*/
type SearchAbsoluteParams struct {

	/*Decorate
	  Run decorators on search result

	*/
	Decorate *bool
	/*Fields
	  Comma separated list of fields to return

	*/
	Fields *string
	/*Filter
	  Filter

	*/
	Filter *string
	/*From
	  Timerange start. See description for date format

	*/
	From string
	/*Limit
	  Maximum number of messages to return.

	*/
	Limit *int64
	/*Offset
	  Offset

	*/
	Offset *int64
	/*Query
	  Query (Lucene syntax)

	*/
	Query string
	/*Sort
	  Sorting (field:asc / field:desc)

	*/
	Sort *string
	/*To
	  Timerange end. See description for date format

	*/
	To string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the search absolute params
func (o *SearchAbsoluteParams) WithTimeout(timeout time.Duration) *SearchAbsoluteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search absolute params
func (o *SearchAbsoluteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search absolute params
func (o *SearchAbsoluteParams) WithContext(ctx context.Context) *SearchAbsoluteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search absolute params
func (o *SearchAbsoluteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search absolute params
func (o *SearchAbsoluteParams) WithHTTPClient(client *http.Client) *SearchAbsoluteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search absolute params
func (o *SearchAbsoluteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDecorate adds the decorate to the search absolute params
func (o *SearchAbsoluteParams) WithDecorate(decorate *bool) *SearchAbsoluteParams {
	o.SetDecorate(decorate)
	return o
}

// SetDecorate adds the decorate to the search absolute params
func (o *SearchAbsoluteParams) SetDecorate(decorate *bool) {
	o.Decorate = decorate
}

// WithFields adds the fields to the search absolute params
func (o *SearchAbsoluteParams) WithFields(fields *string) *SearchAbsoluteParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the search absolute params
func (o *SearchAbsoluteParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the search absolute params
func (o *SearchAbsoluteParams) WithFilter(filter *string) *SearchAbsoluteParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the search absolute params
func (o *SearchAbsoluteParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithFrom adds the from to the search absolute params
func (o *SearchAbsoluteParams) WithFrom(from string) *SearchAbsoluteParams {
	o.SetFrom(from)
	return o
}

// SetFrom adds the from to the search absolute params
func (o *SearchAbsoluteParams) SetFrom(from string) {
	o.From = from
}

// WithLimit adds the limit to the search absolute params
func (o *SearchAbsoluteParams) WithLimit(limit *int64) *SearchAbsoluteParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search absolute params
func (o *SearchAbsoluteParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the search absolute params
func (o *SearchAbsoluteParams) WithOffset(offset *int64) *SearchAbsoluteParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the search absolute params
func (o *SearchAbsoluteParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQuery adds the query to the search absolute params
func (o *SearchAbsoluteParams) WithQuery(query string) *SearchAbsoluteParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the search absolute params
func (o *SearchAbsoluteParams) SetQuery(query string) {
	o.Query = query
}

// WithSort adds the sort to the search absolute params
func (o *SearchAbsoluteParams) WithSort(sort *string) *SearchAbsoluteParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the search absolute params
func (o *SearchAbsoluteParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithTo adds the to to the search absolute params
func (o *SearchAbsoluteParams) WithTo(to string) *SearchAbsoluteParams {
	o.SetTo(to)
	return o
}

// SetTo adds the to to the search absolute params
func (o *SearchAbsoluteParams) SetTo(to string) {
	o.To = to
}

// WriteToRequest writes these params to a swagger request
func (o *SearchAbsoluteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Decorate != nil {

		// query param decorate
		var qrDecorate bool
		if o.Decorate != nil {
			qrDecorate = *o.Decorate
		}
		qDecorate := swag.FormatBool(qrDecorate)
		if qDecorate != "" {
			if err := r.SetQueryParam("decorate", qDecorate); err != nil {
				return err
			}
		}

	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	if o.Filter != nil {

		// query param filter
		var qrFilter string
		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {
			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}

	}

	// query param from
	qrFrom := o.From
	qFrom := qrFrom
	if qFrom != "" {
		if err := r.SetQueryParam("from", qFrom); err != nil {
			return err
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	// query param query
	qrQuery := o.Query
	qQuery := qrQuery
	if qQuery != "" {
		if err := r.SetQueryParam("query", qQuery); err != nil {
			return err
		}
	}

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

	}

	// query param to
	qrTo := o.To
	qTo := qrTo
	if qTo != "" {
		if err := r.SetQueryParam("to", qTo); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}