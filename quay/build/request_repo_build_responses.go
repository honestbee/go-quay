package build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/coreos/go-quay/models"
)

// RequestRepoBuildReader is a Reader for the RequestRepoBuild structure.
type RequestRepoBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *RequestRepoBuildReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRequestRepoBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRequestRepoBuildBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewRequestRepoBuildUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRequestRepoBuildForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRequestRepoBuildNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewRequestRepoBuildOK creates a RequestRepoBuildOK with default headers values
func NewRequestRepoBuildOK() *RequestRepoBuildOK {
	return &RequestRepoBuildOK{}
}

/*RequestRepoBuildOK handles this case with default header values.

Successful invocation
*/
type RequestRepoBuildOK struct {
}

func (o *RequestRepoBuildOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/build/][%d] requestRepoBuildOK ", 200)
}

func (o *RequestRepoBuildOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRequestRepoBuildBadRequest creates a RequestRepoBuildBadRequest with default headers values
func NewRequestRepoBuildBadRequest() *RequestRepoBuildBadRequest {
	return &RequestRepoBuildBadRequest{}
}

/*RequestRepoBuildBadRequest handles this case with default header values.

Bad Request
*/
type RequestRepoBuildBadRequest struct {
	Payload *models.GeneralError
}

func (o *RequestRepoBuildBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/build/][%d] requestRepoBuildBadRequest  %+v", 400, o.Payload)
}

func (o *RequestRepoBuildBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequestRepoBuildUnauthorized creates a RequestRepoBuildUnauthorized with default headers values
func NewRequestRepoBuildUnauthorized() *RequestRepoBuildUnauthorized {
	return &RequestRepoBuildUnauthorized{}
}

/*RequestRepoBuildUnauthorized handles this case with default header values.

Session required
*/
type RequestRepoBuildUnauthorized struct {
}

func (o *RequestRepoBuildUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/build/][%d] requestRepoBuildUnauthorized ", 401)
}

func (o *RequestRepoBuildUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRequestRepoBuildForbidden creates a RequestRepoBuildForbidden with default headers values
func NewRequestRepoBuildForbidden() *RequestRepoBuildForbidden {
	return &RequestRepoBuildForbidden{}
}

/*RequestRepoBuildForbidden handles this case with default header values.

Unauthorized access
*/
type RequestRepoBuildForbidden struct {
}

func (o *RequestRepoBuildForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/build/][%d] requestRepoBuildForbidden ", 403)
}

func (o *RequestRepoBuildForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRequestRepoBuildNotFound creates a RequestRepoBuildNotFound with default headers values
func NewRequestRepoBuildNotFound() *RequestRepoBuildNotFound {
	return &RequestRepoBuildNotFound{}
}

/*RequestRepoBuildNotFound handles this case with default header values.

Not found
*/
type RequestRepoBuildNotFound struct {
}

func (o *RequestRepoBuildNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/build/][%d] requestRepoBuildNotFound ", 404)
}

func (o *RequestRepoBuildNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
