// Code generated by go-swagger; DO NOT EDIT.

package system_fields

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new system fields API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for system fields API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Cycle cycles deflector to new next index in index set
*/
func (a *Client) Cycle(params *CycleParams, authInfo runtime.ClientAuthInfoWriter) (*CycleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCycleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "cycle",
		Method:             "POST",
		PathPattern:        "/system/deflector/{indexSetId}/cycle",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CycleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CycleOK), nil

}

/*
Deflector gets current deflector status in index set
*/
func (a *Client) Deflector(params *DeflectorParams, authInfo runtime.ClientAuthInfoWriter) (*DeflectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeflectorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deflector",
		Method:             "GET",
		PathPattern:        "/system/deflector/{indexSetId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeflectorReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeflectorOK), nil

}

/*
DeprecatedCycle cycles deflector to new next index
*/
func (a *Client) DeprecatedCycle(params *DeprecatedCycleParams, authInfo runtime.ClientAuthInfoWriter) (*DeprecatedCycleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeprecatedCycleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deprecatedCycle",
		Method:             "POST",
		PathPattern:        "/system/deflector/cycle",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeprecatedCycleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeprecatedCycleOK), nil

}

/*
DeprecatedDeflector gets current deflector status
*/
func (a *Client) DeprecatedDeflector(params *DeprecatedDeflectorParams, authInfo runtime.ClientAuthInfoWriter) (*DeprecatedDeflectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeprecatedDeflectorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deprecatedDeflector",
		Method:             "GET",
		PathPattern:        "/system/deflector",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeprecatedDeflectorReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeprecatedDeflectorOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
