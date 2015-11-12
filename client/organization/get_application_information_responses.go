package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"
)

type GetApplicationInformationReader struct {
	formats strfmt.Registry
}

func (o *GetApplicationInformationReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		var result GetApplicationInformationOK
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 400:
		var result GetApplicationInformationBadRequest
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getApplicationInformationBadRequest", &result, response.Code())

	case 401:
		var result GetApplicationInformationUnauthorized
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getApplicationInformationUnauthorized", &result, response.Code())

	case 403:
		var result GetApplicationInformationForbidden
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getApplicationInformationForbidden", &result, response.Code())

	case 404:
		var result GetApplicationInformationNotFound
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("getApplicationInformationNotFound", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", nil, 0)
	}
}

/*
Successful invocation
*/
type GetApplicationInformationOK struct {
}

func (o *GetApplicationInformationOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Bad Request
*/
type GetApplicationInformationBadRequest struct {
}

func (o *GetApplicationInformationBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Session required
*/
type GetApplicationInformationUnauthorized struct {
}

func (o *GetApplicationInformationUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Unauthorized access
*/
type GetApplicationInformationForbidden struct {
}

func (o *GetApplicationInformationForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Not found
*/
type GetApplicationInformationNotFound struct {
}

func (o *GetApplicationInformationNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
