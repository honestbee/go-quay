package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"
)

type GetOrganizationTeamMembersReader struct {
	formats strfmt.Registry
}

func (o *GetOrganizationTeamMembersReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		var result GetOrganizationTeamMembersOK
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 400:
		var result GetOrganizationTeamMembersBadRequest
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getOrganizationTeamMembersBadRequest", &result, response.Code())

	case 401:
		var result GetOrganizationTeamMembersUnauthorized
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getOrganizationTeamMembersUnauthorized", &result, response.Code())

	case 403:
		var result GetOrganizationTeamMembersForbidden
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getOrganizationTeamMembersForbidden", &result, response.Code())

	case 404:
		var result GetOrganizationTeamMembersNotFound
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getOrganizationTeamMembersNotFound", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", nil, 0)
	}
}

/*
Successful invocation
*/
type GetOrganizationTeamMembersOK struct {
}

func (o *GetOrganizationTeamMembersOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Bad Request
*/
type GetOrganizationTeamMembersBadRequest struct {
}

func (o *GetOrganizationTeamMembersBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Session required
*/
type GetOrganizationTeamMembersUnauthorized struct {
}

func (o *GetOrganizationTeamMembersUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Unauthorized access
*/
type GetOrganizationTeamMembersForbidden struct {
}

func (o *GetOrganizationTeamMembersForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Not found
*/
type GetOrganizationTeamMembersNotFound struct {
}

func (o *GetOrganizationTeamMembersNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
