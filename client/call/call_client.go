// Code generated by go-swagger; DO NOT EDIT.

package call

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new call API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for call API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetAppsAppCalls gets app bound calls

Get app-bound calls can filter to route-bound calls.
*/
func (a *Client) GetAppsAppCalls(params *GetAppsAppCallsParams) (*GetAppsAppCallsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAppsAppCallsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAppsAppCalls",
		Method:             "GET",
		PathPattern:        "/apps/{app}/calls",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAppsAppCallsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAppsAppCallsOK), nil

}

/*
GetAppsAppCallsCall gets call information

Get call information
*/
func (a *Client) GetAppsAppCallsCall(params *GetAppsAppCallsCallParams) (*GetAppsAppCallsCallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAppsAppCallsCallParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAppsAppCallsCall",
		Method:             "GET",
		PathPattern:        "/apps/{app}/calls/{call}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAppsAppCallsCallReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAppsAppCallsCallOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
