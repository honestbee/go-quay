package organization

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

// ChangeOrganizationDetailsReader is a Reader for the ChangeOrganizationDetails structure.
type ChangeOrganizationDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *ChangeOrganizationDetailsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChangeOrganizationDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewChangeOrganizationDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewChangeOrganizationDetailsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewChangeOrganizationDetailsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewChangeOrganizationDetailsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeOrganizationDetailsOK creates a ChangeOrganizationDetailsOK with default headers values
func NewChangeOrganizationDetailsOK() *ChangeOrganizationDetailsOK {
	return &ChangeOrganizationDetailsOK{}
}

/*ChangeOrganizationDetailsOK handles this case with default header values.

Successful invocation
*/
type ChangeOrganizationDetailsOK struct {
}

func (o *ChangeOrganizationDetailsOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}][%d] changeOrganizationDetailsOK ", 200)
}

func (o *ChangeOrganizationDetailsOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeOrganizationDetailsBadRequest creates a ChangeOrganizationDetailsBadRequest with default headers values
func NewChangeOrganizationDetailsBadRequest() *ChangeOrganizationDetailsBadRequest {
	return &ChangeOrganizationDetailsBadRequest{}
}

/*ChangeOrganizationDetailsBadRequest handles this case with default header values.

Bad Request
*/
type ChangeOrganizationDetailsBadRequest struct {
	Payload *models.GeneralError
}

func (o *ChangeOrganizationDetailsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}][%d] changeOrganizationDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *ChangeOrganizationDetailsBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeOrganizationDetailsUnauthorized creates a ChangeOrganizationDetailsUnauthorized with default headers values
func NewChangeOrganizationDetailsUnauthorized() *ChangeOrganizationDetailsUnauthorized {
	return &ChangeOrganizationDetailsUnauthorized{}
}

/*ChangeOrganizationDetailsUnauthorized handles this case with default header values.

Session required
*/
type ChangeOrganizationDetailsUnauthorized struct {
}

func (o *ChangeOrganizationDetailsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}][%d] changeOrganizationDetailsUnauthorized ", 401)
}

func (o *ChangeOrganizationDetailsUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeOrganizationDetailsForbidden creates a ChangeOrganizationDetailsForbidden with default headers values
func NewChangeOrganizationDetailsForbidden() *ChangeOrganizationDetailsForbidden {
	return &ChangeOrganizationDetailsForbidden{}
}

/*ChangeOrganizationDetailsForbidden handles this case with default header values.

Unauthorized access
*/
type ChangeOrganizationDetailsForbidden struct {
}

func (o *ChangeOrganizationDetailsForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}][%d] changeOrganizationDetailsForbidden ", 403)
}

func (o *ChangeOrganizationDetailsForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChangeOrganizationDetailsNotFound creates a ChangeOrganizationDetailsNotFound with default headers values
func NewChangeOrganizationDetailsNotFound() *ChangeOrganizationDetailsNotFound {
	return &ChangeOrganizationDetailsNotFound{}
}

/*ChangeOrganizationDetailsNotFound handles this case with default header values.

Not found
*/
type ChangeOrganizationDetailsNotFound struct {
}

func (o *ChangeOrganizationDetailsNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}][%d] changeOrganizationDetailsNotFound ", 404)
}

func (o *ChangeOrganizationDetailsNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}
