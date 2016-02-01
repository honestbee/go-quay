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

// CancelRepoBuildReader is a Reader for the CancelRepoBuild structure.
type CancelRepoBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *CancelRepoBuildReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewCancelRepoBuildNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCancelRepoBuildBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCancelRepoBuildUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCancelRepoBuildForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCancelRepoBuildNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewCancelRepoBuildNoContent creates a CancelRepoBuildNoContent with default headers values
func NewCancelRepoBuildNoContent() *CancelRepoBuildNoContent {
	return &CancelRepoBuildNoContent{}
}

/*CancelRepoBuildNoContent handles this case with default header values.

Deleted
*/
type CancelRepoBuildNoContent struct {
}

func (o *CancelRepoBuildNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/build/{build_uuid}][%d] cancelRepoBuildNoContent ", 204)
}

func (o *CancelRepoBuildNoContent) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCancelRepoBuildBadRequest creates a CancelRepoBuildBadRequest with default headers values
func NewCancelRepoBuildBadRequest() *CancelRepoBuildBadRequest {
	return &CancelRepoBuildBadRequest{}
}

/*CancelRepoBuildBadRequest handles this case with default header values.

Bad Request
*/
type CancelRepoBuildBadRequest struct {
	Payload *models.GeneralError
}

func (o *CancelRepoBuildBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/build/{build_uuid}][%d] cancelRepoBuildBadRequest  %+v", 400, o.Payload)
}

func (o *CancelRepoBuildBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelRepoBuildUnauthorized creates a CancelRepoBuildUnauthorized with default headers values
func NewCancelRepoBuildUnauthorized() *CancelRepoBuildUnauthorized {
	return &CancelRepoBuildUnauthorized{}
}

/*CancelRepoBuildUnauthorized handles this case with default header values.

Session required
*/
type CancelRepoBuildUnauthorized struct {
}

func (o *CancelRepoBuildUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/build/{build_uuid}][%d] cancelRepoBuildUnauthorized ", 401)
}

func (o *CancelRepoBuildUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCancelRepoBuildForbidden creates a CancelRepoBuildForbidden with default headers values
func NewCancelRepoBuildForbidden() *CancelRepoBuildForbidden {
	return &CancelRepoBuildForbidden{}
}

/*CancelRepoBuildForbidden handles this case with default header values.

Unauthorized access
*/
type CancelRepoBuildForbidden struct {
}

func (o *CancelRepoBuildForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/build/{build_uuid}][%d] cancelRepoBuildForbidden ", 403)
}

func (o *CancelRepoBuildForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCancelRepoBuildNotFound creates a CancelRepoBuildNotFound with default headers values
func NewCancelRepoBuildNotFound() *CancelRepoBuildNotFound {
	return &CancelRepoBuildNotFound{}
}

/*CancelRepoBuildNotFound handles this case with default header values.

Not found
*/
type CancelRepoBuildNotFound struct {
}

func (o *CancelRepoBuildNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/build/{build_uuid}][%d] cancelRepoBuildNotFound ", 404)
}

func (o *CancelRepoBuildNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
