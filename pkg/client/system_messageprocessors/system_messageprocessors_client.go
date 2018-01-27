// Code generated by go-swagger; DO NOT EDIT.

package system_messageprocessors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new system messageprocessors API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for system messageprocessors API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Adapters lists available data adapters
*/
func (a *Client) Adapters(params *AdaptersParams, authInfo runtime.ClientAuthInfoWriter) (*AdaptersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAdaptersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "adapters",
		Method:             "GET",
		PathPattern:        "/system/lookup/adapters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AdaptersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AdaptersOK), nil

}

/*
AvailableAdapterTypes lists available data adapter types
*/
func (a *Client) AvailableAdapterTypes(params *AvailableAdapterTypesParams, authInfo runtime.ClientAuthInfoWriter) (*AvailableAdapterTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAvailableAdapterTypesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "availableAdapterTypes",
		Method:             "GET",
		PathPattern:        "/system/lookup/types/adapters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AvailableAdapterTypesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AvailableAdapterTypesOK), nil

}

/*
AvailableCacheTypes lists available caches types
*/
func (a *Client) AvailableCacheTypes(params *AvailableCacheTypesParams, authInfo runtime.ClientAuthInfoWriter) (*AvailableCacheTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAvailableCacheTypesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "availableCacheTypes",
		Method:             "GET",
		PathPattern:        "/system/lookup/types/caches",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AvailableCacheTypesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AvailableCacheTypesOK), nil

}

/*
Caches lists available caches
*/
func (a *Client) Caches(params *CachesParams, authInfo runtime.ClientAuthInfoWriter) (*CachesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCachesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "caches",
		Method:             "GET",
		PathPattern:        "/system/lookup/caches",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CachesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CachesOK), nil

}

/*
DeleteCache deletes the given cache

The cache cannot be in use by any lookup table, otherwise the request will fail.
*/
func (a *Client) DeleteCache(params *DeleteCacheParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteCacheOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCacheParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCache",
		Method:             "DELETE",
		PathPattern:        "/system/lookup/caches/{idOrName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteCacheReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteCacheOK), nil

}

/*
ErrorStates retrieves the runtime error states of the given lookup tables caches and adapters
*/
func (a *Client) ErrorStates(params *ErrorStatesParams, authInfo runtime.ClientAuthInfoWriter) (*ErrorStatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewErrorStatesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "errorStates",
		Method:             "POST",
		PathPattern:        "/system/lookup/errorstates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ErrorStatesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ErrorStatesOK), nil

}

/*
GetAdapter lists the given data adapter
*/
func (a *Client) GetAdapter(params *GetAdapterParams, authInfo runtime.ClientAuthInfoWriter) (*GetAdapterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAdapterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAdapter",
		Method:             "GET",
		PathPattern:        "/system/lookup/adapters/{idOrName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAdapterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAdapterOK), nil

}

/*
PerformAdapterLookup queries a lookup table
*/
func (a *Client) PerformAdapterLookup(params *PerformAdapterLookupParams, authInfo runtime.ClientAuthInfoWriter) (*PerformAdapterLookupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPerformAdapterLookupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "performAdapterLookup",
		Method:             "GET",
		PathPattern:        "/system/lookup/adapters/{name}/query",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PerformAdapterLookupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PerformAdapterLookupOK), nil

}

/*
PerformLookup queries a lookup table
*/
func (a *Client) PerformLookup(params *PerformLookupParams, authInfo runtime.ClientAuthInfoWriter) (*PerformLookupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPerformLookupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "performLookup",
		Method:             "GET",
		PathPattern:        "/system/lookup/tables/{name}/query",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PerformLookupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PerformLookupOK), nil

}

/*
PerformPurge purges lookup table cache
*/
func (a *Client) PerformPurge(params *PerformPurgeParams, authInfo runtime.ClientAuthInfoWriter) (*PerformPurgeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPerformPurgeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "performPurge",
		Method:             "POST",
		PathPattern:        "/system/lookup/tables/{idOrName}/purge",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PerformPurgeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PerformPurgeOK), nil

}

/*
Tables lists configured lookup tables
*/
func (a *Client) Tables(params *TablesParams, authInfo runtime.ClientAuthInfoWriter) (*TablesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTablesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tables",
		Method:             "GET",
		PathPattern:        "/system/lookup/tables",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TablesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TablesOK), nil

}

/*
UpdateTable updates the given lookup table
*/
func (a *Client) UpdateTable(params *UpdateTableParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateTableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateTableParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateTable",
		Method:             "PUT",
		PathPattern:        "/system/lookup/tables/{idOrName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateTableReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateTableOK), nil

}

/*
ValidateAdapter validates the data adapter config
*/
func (a *Client) ValidateAdapter(params *ValidateAdapterParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateAdapterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateAdapterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "validateAdapter",
		Method:             "POST",
		PathPattern:        "/system/lookup/adapters/validate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ValidateAdapterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ValidateAdapterOK), nil

}

/*
ValidateCache validates the cache config
*/
func (a *Client) ValidateCache(params *ValidateCacheParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateCacheOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateCacheParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "validateCache",
		Method:             "POST",
		PathPattern:        "/system/lookup/caches/validate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ValidateCacheReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ValidateCacheOK), nil

}

/*
ValidateTable validates the lookup table config
*/
func (a *Client) ValidateTable(params *ValidateTableParams, authInfo runtime.ClientAuthInfoWriter) (*ValidateTableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateTableParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "validateTable",
		Method:             "POST",
		PathPattern:        "/system/lookup/tables/validate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ValidateTableReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ValidateTableOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}