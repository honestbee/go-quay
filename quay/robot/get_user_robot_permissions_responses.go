package robot

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

// GetUserRobotPermissionsReader is a Reader for the GetUserRobotPermissions structure.
type GetUserRobotPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetUserRobotPermissionsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUserRobotPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetUserRobotPermissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetUserRobotPermissionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetUserRobotPermissionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetUserRobotPermissionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUserRobotPermissionsOK creates a GetUserRobotPermissionsOK with default headers values
func NewGetUserRobotPermissionsOK() *GetUserRobotPermissionsOK {
	return &GetUserRobotPermissionsOK{}
}

/*GetUserRobotPermissionsOK handles this case with default header values.

Successful invocation
*/
type GetUserRobotPermissionsOK struct {
}

func (o *GetUserRobotPermissionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/robots/{robot_shortname}/permissions][%d] getUserRobotPermissionsOK ", 200)
}

func (o *GetUserRobotPermissionsOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserRobotPermissionsBadRequest creates a GetUserRobotPermissionsBadRequest with default headers values
func NewGetUserRobotPermissionsBadRequest() *GetUserRobotPermissionsBadRequest {
	return &GetUserRobotPermissionsBadRequest{}
}

/*GetUserRobotPermissionsBadRequest handles this case with default header values.

Bad Request
*/
type GetUserRobotPermissionsBadRequest struct {
	Payload *models.GeneralError
}

func (o *GetUserRobotPermissionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/robots/{robot_shortname}/permissions][%d] getUserRobotPermissionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserRobotPermissionsBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserRobotPermissionsUnauthorized creates a GetUserRobotPermissionsUnauthorized with default headers values
func NewGetUserRobotPermissionsUnauthorized() *GetUserRobotPermissionsUnauthorized {
	return &GetUserRobotPermissionsUnauthorized{}
}

/*GetUserRobotPermissionsUnauthorized handles this case with default header values.

Session required
*/
type GetUserRobotPermissionsUnauthorized struct {
}

func (o *GetUserRobotPermissionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/robots/{robot_shortname}/permissions][%d] getUserRobotPermissionsUnauthorized ", 401)
}

func (o *GetUserRobotPermissionsUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserRobotPermissionsForbidden creates a GetUserRobotPermissionsForbidden with default headers values
func NewGetUserRobotPermissionsForbidden() *GetUserRobotPermissionsForbidden {
	return &GetUserRobotPermissionsForbidden{}
}

/*GetUserRobotPermissionsForbidden handles this case with default header values.

Unauthorized access
*/
type GetUserRobotPermissionsForbidden struct {
}

func (o *GetUserRobotPermissionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/robots/{robot_shortname}/permissions][%d] getUserRobotPermissionsForbidden ", 403)
}

func (o *GetUserRobotPermissionsForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUserRobotPermissionsNotFound creates a GetUserRobotPermissionsNotFound with default headers values
func NewGetUserRobotPermissionsNotFound() *GetUserRobotPermissionsNotFound {
	return &GetUserRobotPermissionsNotFound{}
}

/*GetUserRobotPermissionsNotFound handles this case with default header values.

Not found
*/
type GetUserRobotPermissionsNotFound struct {
}

func (o *GetUserRobotPermissionsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/robots/{robot_shortname}/permissions][%d] getUserRobotPermissionsNotFound ", 404)
}

func (o *GetUserRobotPermissionsNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
