// Code generated by go-swagger; DO NOT EDIT.

package format

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new format API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for format API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	FormatSeriesServiceDelete(params *FormatSeriesServiceDeleteParams) (*FormatSeriesServiceDeleteOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  FormatSeriesServiceDelete deletes a series format

  Delete a series format when given an id
*/
func (a *Client) FormatSeriesServiceDelete(params *FormatSeriesServiceDeleteParams) (*FormatSeriesServiceDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFormatSeriesServiceDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "FormatSeriesService_Delete",
		Method:             "DELETE",
		PathPattern:        "/api/v1/format/series/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FormatSeriesServiceDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FormatSeriesServiceDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FormatSeriesServiceDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
