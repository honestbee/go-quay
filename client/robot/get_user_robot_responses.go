package robot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"
)

type GetUserRobotReader struct {
	formats strfmt.Registry
}

func (o *GetUserRobotReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		var result GetUserRobotOK
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 400:
		var result GetUserRobotBadRequest
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getUserRobotBadRequest", &result, response.Code())

	case 401:
		var result GetUserRobotUnauthorized
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getUserRobotUnauthorized", &result, response.Code())

	case 403:
		var result GetUserRobotForbidden
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getUserRobotForbidden", &result, response.Code())

	case 404:
		var result GetUserRobotNotFound
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getUserRobotNotFound", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", nil, 0)
	}
}

/*
Successful invocation
*/
type GetUserRobotOK struct {
}

func (o *GetUserRobotOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Bad Request
*/
type GetUserRobotBadRequest struct {
}

func (o *GetUserRobotBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Session required
*/
type GetUserRobotUnauthorized struct {
}

func (o *GetUserRobotUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Unauthorized access
*/
type GetUserRobotForbidden struct {
}

func (o *GetUserRobotForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Not found
*/
type GetUserRobotNotFound struct {
}

func (o *GetUserRobotNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
