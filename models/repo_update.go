package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*
Fields which can be updated in a repository.

swagger:model RepoUpdate
*/
type RepoUpdate struct {

	/* Markdown encoded description for the repository

	Required: true
	*/
	Description string `json:"description"`
}

// Validate validates this repo update
func (m *RepoUpdate) Validate(formats strfmt.Registry) error {

	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}

func (m *RepoUpdate) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", string(m.Description)); err != nil {
		return err
	}

	return nil
}
