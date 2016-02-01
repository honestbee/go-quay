package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*RevertTag Reverts a tag to a specific image

swagger:model RevertTag
*/
type RevertTag struct {

	/* Image identifier to which the tag should point

	Required: true
	*/
	Image string `json:"image,omitempty"`
}

// Validate validates this revert tag
func (m *RevertTag) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RevertTag) validateImage(formats strfmt.Registry) error {

	if err := validate.RequiredString("image", "body", string(m.Image)); err != nil {
		return err
	}

	return nil
}
