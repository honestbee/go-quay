package discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new discovery API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for discovery API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*
List all of the API endpoints available in the swagger API format.
*/
func (a *Client) Discovery(params *DiscoveryParams) (*DiscoveryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDiscoveryParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "discovery",
		Method:             "GET",
		PathPattern:        "/api/v1/discovery",
		ProducesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DiscoveryReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*DiscoveryOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport client.Transport) {
	a.transport = transport
}

// NewAPIError creates a new API error
func NewAPIError(opName string, response interface{}, code int) APIError {
	return APIError{
		OperationName: opName,
		Response:      response,
		Code:          code,
	}
}

// APIError wraps an error model and captures the status code
type APIError struct {
	OperationName string
	Response      interface{}
	Code          int
}

func (a APIError) Error() string {
	return fmt.Sprintf("%s (status %d): %+v ", a.OperationName, a.Code, a.Response)
}
