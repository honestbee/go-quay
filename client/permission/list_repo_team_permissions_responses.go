package permission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"
)

type ListRepoTeamPermissionsReader struct {
	formats strfmt.Registry
}

func (o *ListRepoTeamPermissionsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		var result ListRepoTeamPermissionsOK
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 400:
		var result ListRepoTeamPermissionsBadRequest
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("listRepoTeamPermissionsBadRequest", &result, response.Code())

	case 401:
		var result ListRepoTeamPermissionsUnauthorized
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("listRepoTeamPermissionsUnauthorized", &result, response.Code())

	case 403:
		var result ListRepoTeamPermissionsForbidden
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("listRepoTeamPermissionsForbidden", &result, response.Code())

	case 404:
		var result ListRepoTeamPermissionsNotFound
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("listRepoTeamPermissionsNotFound", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", nil, 0)
	}
}

/*
Successful invocation
*/
type ListRepoTeamPermissionsOK struct {
}

func (o *ListRepoTeamPermissionsOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Bad Request
*/
type ListRepoTeamPermissionsBadRequest struct {
}

func (o *ListRepoTeamPermissionsBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Session required
*/
type ListRepoTeamPermissionsUnauthorized struct {
}

func (o *ListRepoTeamPermissionsUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Unauthorized access
*/
type ListRepoTeamPermissionsForbidden struct {
}

func (o *ListRepoTeamPermissionsForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Not found
*/
type ListRepoTeamPermissionsNotFound struct {
}

func (o *ListRepoTeamPermissionsNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
