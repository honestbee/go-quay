package trigger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"
)

type ManuallyStartBuildTriggerReader struct {
	formats strfmt.Registry
}

func (o *ManuallyStartBuildTriggerReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		var result ManuallyStartBuildTriggerOK
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return &result, nil

	case 400:
		var result ManuallyStartBuildTriggerBadRequest
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("manuallyStartBuildTriggerBadRequest", &result, response.Code())

	case 401:
		var result ManuallyStartBuildTriggerUnauthorized
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("manuallyStartBuildTriggerUnauthorized", &result, response.Code())

	case 403:
		var result ManuallyStartBuildTriggerForbidden
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("manuallyStartBuildTriggerForbidden", &result, response.Code())

	case 404:
		var result ManuallyStartBuildTriggerNotFound
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, NewAPIError("manuallyStartBuildTriggerNotFound", &result, response.Code())

	default:
		return nil, NewAPIError("unknown error", nil, 0)
	}
}

/*
Successful invocation
*/
type ManuallyStartBuildTriggerOK struct {
}

func (o *ManuallyStartBuildTriggerOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Bad Request
*/
type ManuallyStartBuildTriggerBadRequest struct {
}

func (o *ManuallyStartBuildTriggerBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Session required
*/
type ManuallyStartBuildTriggerUnauthorized struct {
}

func (o *ManuallyStartBuildTriggerUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Unauthorized access
*/
type ManuallyStartBuildTriggerForbidden struct {
}

func (o *ManuallyStartBuildTriggerForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
Not found
*/
type ManuallyStartBuildTriggerNotFound struct {
}

func (o *ManuallyStartBuildTriggerNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
