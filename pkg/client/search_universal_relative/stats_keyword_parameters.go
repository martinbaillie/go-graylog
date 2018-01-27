// Code generated by go-swagger; DO NOT EDIT.

package search_universal_relative

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

// NewStatsKeywordParams creates a new StatsKeywordParams object
// with the default values initialized.
func NewStatsKeywordParams() *StatsKeywordParams {
	var ()
	return &StatsKeywordParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStatsKeywordParamsWithTimeout creates a new StatsKeywordParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStatsKeywordParamsWithTimeout(timeout time.Duration) *StatsKeywordParams {
	var ()
	return &StatsKeywordParams{

		timeout: timeout,
	}
}

// NewStatsKeywordParamsWithContext creates a new StatsKeywordParams object
// with the default values initialized, and the ability to set a context for a request
func NewStatsKeywordParamsWithContext(ctx context.Context) *StatsKeywordParams {
	var ()
	return &StatsKeywordParams{

		Context: ctx,
	}
}

// NewStatsKeywordParamsWithHTTPClient creates a new StatsKeywordParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStatsKeywordParamsWithHTTPClient(client *http.Client) *StatsKeywordParams {
	var ()
	return &StatsKeywordParams{
		HTTPClient: client,
	}
}

/*StatsKeywordParams contains all the parameters to send to the API endpoint
for the stats keyword operation typically these are written to a http.Request
*/
type StatsKeywordParams struct {

	/*Field
	  Message field of numeric type to return statistics for

	*/
	Field string
	/*Filter
	  Filter

	*/
	Filter *string
	/*Keyword
	  Range keyword

	*/
	Keyword string
	/*Query
	  Query (Lucene syntax)

	*/
	Query string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stats keyword params
func (o *StatsKeywordParams) WithTimeout(timeout time.Duration) *StatsKeywordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stats keyword params
func (o *StatsKeywordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stats keyword params
func (o *StatsKeywordParams) WithContext(ctx context.Context) *StatsKeywordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stats keyword params
func (o *StatsKeywordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stats keyword params
func (o *StatsKeywordParams) WithHTTPClient(client *http.Client) *StatsKeywordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stats keyword params
func (o *StatsKeywordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithField adds the field to the stats keyword params
func (o *StatsKeywordParams) WithField(field string) *StatsKeywordParams {
	o.SetField(field)
	return o
}

// SetField adds the field to the stats keyword params
func (o *StatsKeywordParams) SetField(field string) {
	o.Field = field
}

// WithFilter adds the filter to the stats keyword params
func (o *StatsKeywordParams) WithFilter(filter *string) *StatsKeywordParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the stats keyword params
func (o *StatsKeywordParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithKeyword adds the keyword to the stats keyword params
func (o *StatsKeywordParams) WithKeyword(keyword string) *StatsKeywordParams {
	o.SetKeyword(keyword)
	return o
}

// SetKeyword adds the keyword to the stats keyword params
func (o *StatsKeywordParams) SetKeyword(keyword string) {
	o.Keyword = keyword
}

// WithQuery adds the query to the stats keyword params
func (o *StatsKeywordParams) WithQuery(query string) *StatsKeywordParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the stats keyword params
func (o *StatsKeywordParams) SetQuery(query string) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *StatsKeywordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param field
	qrField := o.Field
	qField := qrField
	if qField != "" {
		if err := r.SetQueryParam("field", qField); err != nil {
			return err
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

	// query param keyword
	qrKeyword := o.Keyword
	qKeyword := qrKeyword
	if qKeyword != "" {
		if err := r.SetQueryParam("keyword", qKeyword); err != nil {
			return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}