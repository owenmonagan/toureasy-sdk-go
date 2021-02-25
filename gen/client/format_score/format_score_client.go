// Code generated by go-swagger; DO NOT EDIT.

package format_score

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new format score API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for format score API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	FormatScoreServiceCreate(params *FormatScoreServiceCreateParams) (*FormatScoreServiceCreateOK, error)

	FormatScoreServiceDelete(params *FormatScoreServiceDeleteParams) (*FormatScoreServiceDeleteOK, error)

	FormatScoreServiceUpdate(params *FormatScoreServiceUpdateParams) (*FormatScoreServiceUpdateOK, error)

	FormatScoreServiceView(params *FormatScoreServiceViewParams) (*FormatScoreServiceViewOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  FormatScoreServiceCreate adds a score format

  Add a score format to the server.
*/
func (a *Client) FormatScoreServiceCreate(params *FormatScoreServiceCreateParams) (*FormatScoreServiceCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFormatScoreServiceCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "FormatScoreService_Create",
		Method:             "POST",
		PathPattern:        "/api/v1/format/score",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FormatScoreServiceCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FormatScoreServiceCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FormatScoreServiceCreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  FormatScoreServiceDelete deletes a score format

  Delete a score format when given an id
*/
func (a *Client) FormatScoreServiceDelete(params *FormatScoreServiceDeleteParams) (*FormatScoreServiceDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFormatScoreServiceDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "FormatScoreService_Delete",
		Method:             "DELETE",
		PathPattern:        "/api/v1/format/score/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FormatScoreServiceDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FormatScoreServiceDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FormatScoreServiceDeleteDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  FormatScoreServiceUpdate updates a score format

  Update a score format when given an id
*/
func (a *Client) FormatScoreServiceUpdate(params *FormatScoreServiceUpdateParams) (*FormatScoreServiceUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFormatScoreServiceUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "FormatScoreService_Update",
		Method:             "PATCH",
		PathPattern:        "/api/v1/format/score/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FormatScoreServiceUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FormatScoreServiceUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FormatScoreServiceUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  FormatScoreServiceView views a score format

  View a score format when given an id
*/
func (a *Client) FormatScoreServiceView(params *FormatScoreServiceViewParams) (*FormatScoreServiceViewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFormatScoreServiceViewParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "FormatScoreService_View",
		Method:             "GET",
		PathPattern:        "/api/v1/format/score/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &FormatScoreServiceViewReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FormatScoreServiceViewOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FormatScoreServiceViewDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}