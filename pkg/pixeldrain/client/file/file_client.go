// Code generated by go-swagger; DO NOT EDIT.

package file

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new file API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for file API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetFileIDThumbnail(params *GetFileIDThumbnailParams, opts ...ClientOption) (*GetFileIDThumbnailOK, error)

	DeleteFile(params *DeleteFileParams, opts ...ClientOption) (*DeleteFileOK, error)

	DownloadFile(params *DownloadFileParams, writer io.Writer, opts ...ClientOption) (*DownloadFileOK, error)

	GetFileInfo(params *GetFileInfoParams, opts ...ClientOption) (*GetFileInfoOK, error)

	UploadFile(params *UploadFileParams, opts ...ClientOption) (*UploadFileCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetFileIDThumbnail gets a thumbnail image representing the file

  Returns a PNG thumbnail image representing the file. The thumbnail is always 100*100 px. If the source file is parsable by imagemagick the thumbnail will be generated from the file, if not it will be a generic mime type icon.

*/
func (a *Client) GetFileIDThumbnail(params *GetFileIDThumbnailParams, opts ...ClientOption) (*GetFileIDThumbnailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFileIDThumbnailParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetFileIDThumbnail",
		Method:             "GET",
		PathPattern:        "/file/{id}/thumbnail",
		ProducesMediaTypes: []string{"application/json", "image/png"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFileIDThumbnailReader{formats: a.formats},
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
	success, ok := result.(*GetFileIDThumbnailOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetFileIDThumbnailDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteFile deletes a file

  Deletes a file. Only works when the users owns the file.
*/
func (a *Client) DeleteFile(params *DeleteFileParams, opts ...ClientOption) (*DeleteFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteFile",
		Method:             "DELETE",
		PathPattern:        "/file/{id}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteFileReader{formats: a.formats},
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
	success, ok := result.(*DeleteFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteFileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DownloadFile downloads a file

  Returns the full file associated with the ID. Supports byte range requests.

*/
func (a *Client) DownloadFile(params *DownloadFileParams, writer io.Writer, opts ...ClientOption) (*DownloadFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDownloadFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "downloadFile",
		Method:             "GET",
		PathPattern:        "/file/{id}",
		ProducesMediaTypes: []string{"application/json", "application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DownloadFileReader{formats: a.formats, writer: writer},
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
	success, ok := result.(*DownloadFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DownloadFileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetFileInfo retrieves information of a file

  Returns information about one or more files. You can also put a comma separated list of file IDs in the URL and it will return an array of file info, instead of a single object.

*/
func (a *Client) GetFileInfo(params *GetFileInfoParams, opts ...ClientOption) (*GetFileInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFileInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getFileInfo",
		Method:             "GET",
		PathPattern:        "/file/{id}/info",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetFileInfoReader{formats: a.formats},
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
	success, ok := result.(*GetFileInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetFileInfoDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UploadFile uploads a file

  Upload a file.
*/
func (a *Client) UploadFile(params *UploadFileParams, opts ...ClientOption) (*UploadFileCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadFileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "uploadFile",
		Method:             "POST",
		PathPattern:        "/file",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UploadFileReader{formats: a.formats},
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
	success, ok := result.(*UploadFileCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UploadFileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
