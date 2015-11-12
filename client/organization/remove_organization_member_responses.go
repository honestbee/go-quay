package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"
)

type RemoveOrganizationMemberReader struct {
	formats strfmt.Registry
}

func (o *RemoveOrganizationMemberReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		var result RemoveOrganizationMemberNoContent
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 400:
		var result RemoveOrganizationMemberBadRequest
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("removeOrganizationMemberBadRequest", &result, response.Code())

	case 401:
		var result RemoveOrganizationMemberUnauthorized
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("removeOrganizationMemberUnauthorized", &result, response.Code())

	case 403:
		var result RemoveOrganizationMemberForbidden
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("removeOrganizationMemberForbidden", &result, response.Code())

	case 404:
		var result RemoveOrganizationMemberNotFound
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("removeOrganizationMemberNotFound", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", nil, 0)
	}
}

/*
Deleted
*/
type RemoveOrganizationMemberNoContent struct {
}

func (o *RemoveOrganizationMemberNoContent) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Bad Request
*/
type RemoveOrganizationMemberBadRequest struct {
}

func (o *RemoveOrganizationMemberBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Session required
*/
type RemoveOrganizationMemberUnauthorized struct {
}

func (o *RemoveOrganizationMemberUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Unauthorized access
*/
type RemoveOrganizationMemberForbidden struct {
}

func (o *RemoveOrganizationMemberForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Not found
*/
type RemoveOrganizationMemberNotFound struct {
}

func (o *RemoveOrganizationMemberNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
