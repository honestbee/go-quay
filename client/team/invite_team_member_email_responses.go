package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"
)

type InviteTeamMemberEmailReader struct {
	formats strfmt.Registry
}

func (o *InviteTeamMemberEmailReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		var result InviteTeamMemberEmailOK
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 400:
		var result InviteTeamMemberEmailBadRequest
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("inviteTeamMemberEmailBadRequest", &result, response.Code())

	case 401:
		var result InviteTeamMemberEmailUnauthorized
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("inviteTeamMemberEmailUnauthorized", &result, response.Code())

	case 403:
		var result InviteTeamMemberEmailForbidden
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("inviteTeamMemberEmailForbidden", &result, response.Code())

	case 404:
		var result InviteTeamMemberEmailNotFound
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("inviteTeamMemberEmailNotFound", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", nil, 0)
	}
}

/*
Successful invocation
*/
type InviteTeamMemberEmailOK struct {
}

func (o *InviteTeamMemberEmailOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Bad Request
*/
type InviteTeamMemberEmailBadRequest struct {
}

func (o *InviteTeamMemberEmailBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Session required
*/
type InviteTeamMemberEmailUnauthorized struct {
}

func (o *InviteTeamMemberEmailUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Unauthorized access
*/
type InviteTeamMemberEmailForbidden struct {
}

func (o *InviteTeamMemberEmailForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Not found
*/
type InviteTeamMemberEmailNotFound struct {
}

func (o *InviteTeamMemberEmailNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
