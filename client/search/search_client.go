package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new search API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for search API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*Get a list of entities and resources that match the specified query.
 */
func (a *Client) ConductSearch(params ConductSearchParams, authInfo client.AuthInfoWriter) (*ConductSearchOK, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:       "conductSearch",
		Params:   &params,
		Reader:   &ConductSearchReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ConductSearchOK), nil
}

/*Get a list of entities that match the specified prefix.
 */
func (a *Client) GetMatchingEntities(params GetMatchingEntitiesParams) (*GetMatchingEntitiesOK, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:     "getMatchingEntities",
		Params: &params,
		Reader: &GetMatchingEntitiesReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMatchingEntitiesOK), nil
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
