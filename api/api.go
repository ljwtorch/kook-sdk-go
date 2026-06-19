// Package api provides HTTP API functions for the KOOK platform.
//
// All functions accept a Doer interface as the client parameter,
// allowing them to work with any HTTP client implementation that
// satisfies the interface. The primary implementation is kook.Client.
package api

import (
	"context"
	"io"
)

// Doer defines the interface for executing HTTP requests against the KOOK API.
// The kook.Client type implements this interface.
type Doer interface {
	// Do executes an HTTP request. For GET requests, body is converted to query
	// parameters. For POST requests, body is serialized as JSON. The response
	// data is unmarshaled into result. Pass nil for result when no response
	// body is expected.
	Do(ctx context.Context, method, path string, body interface{}, result interface{}) error

	// DoMultipart executes a multipart/form-data HTTP request for file uploads.
	// fields contains additional form fields. fieldName and fileName describe
	// the file part. The response data is unmarshaled into result.
	DoMultipart(ctx context.Context, path string, fields map[string]string, fieldName string, fileName string, file io.Reader, result interface{}) error
}
