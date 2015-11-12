package permission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"
)

type GetUserTransitivePermissionReader struct {
	formats strfmt.Registry
}

func (o *GetUserTransitivePermissionReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		var result GetUserTransitivePermissionOK
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 400:
		var result GetUserTransitivePermissionBadRequest
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getUserTransitivePermissionBadRequest", &result, response.Code())

	case 401:
		var result GetUserTransitivePermissionUnauthorized
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getUserTransitivePermissionUnauthorized", &result, response.Code())

	case 403:
		var result GetUserTransitivePermissionForbidden
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getUserTransitivePermissionForbidden", &result, response.Code())

	case 404:
		var result GetUserTransitivePermissionNotFound
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getUserTransitivePermissionNotFound", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", nil, 0)
	}
}

/*
Successful invocation
*/
type GetUserTransitivePermissionOK struct {
}

func (o *GetUserTransitivePermissionOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Bad Request
*/
type GetUserTransitivePermissionBadRequest struct {
}

func (o *GetUserTransitivePermissionBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Session required
*/
type GetUserTransitivePermissionUnauthorized struct {
}

func (o *GetUserTransitivePermissionUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Unauthorized access
*/
type GetUserTransitivePermissionForbidden struct {
}

func (o *GetUserTransitivePermissionForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Not found
*/
type GetUserTransitivePermissionNotFound struct {
}

func (o *GetUserTransitivePermissionNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
