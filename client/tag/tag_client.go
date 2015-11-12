package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new tag API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tag API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*Change which image a tag points to or create a new tag.
 */
func (a *Client) ChangeTagImage(params ChangeTagImageParams, authInfo client.AuthInfoWriter) (*ChangeTagImageOK, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:       "changeTagImage",
		Params:   &params,
		Reader:   &ChangeTagImageReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChangeTagImageOK), nil
}

/*Delete the specified repository tag.
 */
func (a *Client) DeleteFullTag(params DeleteFullTagParams, authInfo client.AuthInfoWriter) (*DeleteFullTagNoContent, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:       "deleteFullTag",
		Params:   &params,
		Reader:   &DeleteFullTagReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteFullTagNoContent), nil
}

/*ListRepoTags list repo tags API
 */
func (a *Client) ListRepoTags(params ListRepoTagsParams, authInfo client.AuthInfoWriter) (*ListRepoTagsOK, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:       "listRepoTags",
		Params:   &params,
		Reader:   &ListRepoTagsReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListRepoTagsOK), nil
}

/*List the images for the specified repository tag.
 */
func (a *Client) ListTagImages(params ListTagImagesParams, authInfo client.AuthInfoWriter) (*ListTagImagesOK, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:       "listTagImages",
		Params:   &params,
		Reader:   &ListTagImagesReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListTagImagesOK), nil
}

/*Reverts a repository tag back to a previous image in the repository.
 */
func (a *Client) RevertTag(params RevertTagParams, authInfo client.AuthInfoWriter) (*RevertTagOK, error) {
	// TODO: Validate the params before sending

	result, err := a.transport.Submit(&client.Operation{
		ID:       "revertTag",
		Params:   &params,
		Reader:   &RevertTagReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RevertTagOK), nil
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
