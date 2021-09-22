// Code generated by go-swagger; DO NOT EDIT.

package list

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new list API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for list API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateFileList(params *CreateFileListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateFileListCreated, error)

	GetFileList(params *GetFileListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFileListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateFileList creates a list of files that can be viewed together on the file viewer page

  Creates a list of files that can be viewed together on the file viewer page.
*/
func (a *Client) CreateFileList(params *CreateFileListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateFileListCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateFileListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createFileList",
		Method:             "POST",
		PathPattern:        "/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateFileListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateFileListCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateFileListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetFileList returns information about a file list and the files in it

  Returns information about a file list and the files in it.
*/
func (a *Client) GetFileList(params *GetFileListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetFileListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFileListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFileList",
		Method:             "GET",
		PathPattern:        "/list/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFileListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFileListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetFileListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
