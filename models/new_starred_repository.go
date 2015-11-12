package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*
NewStarredRepository new starred repository

swagger:model NewStarredRepository
*/
type NewStarredRepository struct {

	/* Namespace in which the repository belongs

	Required: true
	*/
	Namespace string `json:"namespace"`

	/* Repository name

	Required: true
	*/
	Repository string `json:"repository"`
}

// Validate validates this new starred repository
func (m *NewStarredRepository) Validate(formats strfmt.Registry) error {

	var res []error

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepository(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}

func (m *NewStarredRepository) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", string(m.Namespace)); err != nil {
		return err
	}

	return nil
}

func (m *NewStarredRepository) validateRepository(formats strfmt.Registry) error {

	if err := validate.Required("repository", "body", string(m.Repository)); err != nil {
		return err
	}

	return nil
}
