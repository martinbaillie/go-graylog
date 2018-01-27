// Code generated by go-swagger; DO NOT EDIT.

package search_universal_absolute

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

// NewAddMemberParams creates a new AddMemberParams object
// with the default values initialized.
func NewAddMemberParams() *AddMemberParams {
	var ()
	return &AddMemberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddMemberParamsWithTimeout creates a new AddMemberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddMemberParamsWithTimeout(timeout time.Duration) *AddMemberParams {
	var ()
	return &AddMemberParams{

		timeout: timeout,
	}
}

// NewAddMemberParamsWithContext creates a new AddMemberParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddMemberParamsWithContext(ctx context.Context) *AddMemberParams {
	var ()
	return &AddMemberParams{

		Context: ctx,
	}
}

// NewAddMemberParamsWithHTTPClient creates a new AddMemberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddMemberParamsWithHTTPClient(client *http.Client) *AddMemberParams {
	var ()
	return &AddMemberParams{
		HTTPClient: client,
	}
}

/*AddMemberParams contains all the parameters to send to the API endpoint
for the add member operation typically these are written to a http.Request
*/
type AddMemberParams struct {

	/*JSONBody
	  Placeholder because PUT requests should have a body. Set to '{}', the content will be ignored.

	*/
	JSONBody *string
	/*Rolename*/
	Rolename string
	/*Username*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add member params
func (o *AddMemberParams) WithTimeout(timeout time.Duration) *AddMemberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add member params
func (o *AddMemberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add member params
func (o *AddMemberParams) WithContext(ctx context.Context) *AddMemberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add member params
func (o *AddMemberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add member params
func (o *AddMemberParams) WithHTTPClient(client *http.Client) *AddMemberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add member params
func (o *AddMemberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJSONBody adds the jSONBody to the add member params
func (o *AddMemberParams) WithJSONBody(jSONBody *string) *AddMemberParams {
	o.SetJSONBody(jSONBody)
	return o
}

// SetJSONBody adds the jsonBody to the add member params
func (o *AddMemberParams) SetJSONBody(jSONBody *string) {
	o.JSONBody = jSONBody
}

// WithRolename adds the rolename to the add member params
func (o *AddMemberParams) WithRolename(rolename string) *AddMemberParams {
	o.SetRolename(rolename)
	return o
}

// SetRolename adds the rolename to the add member params
func (o *AddMemberParams) SetRolename(rolename string) {
	o.Rolename = rolename
}

// WithUsername adds the username to the add member params
func (o *AddMemberParams) WithUsername(username string) *AddMemberParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the add member params
func (o *AddMemberParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *AddMemberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.JSONBody != nil {
		if err := r.SetBodyParam(o.JSONBody); err != nil {
			return err
		}
	}

	// path param rolename
	if err := r.SetPathParam("rolename", o.Rolename); err != nil {
		return err
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}