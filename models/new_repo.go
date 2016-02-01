package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*NewRepo Description of a new repository

swagger:model NewRepo
*/
type NewRepo struct {

	/* Markdown encoded description for the repository

	Required: true
	*/
	Description string `json:"description,omitempty"`

	/* Namespace in which the repository should be created. If omitted, the username of the caller is used
	 */
	Namespace *string `json:"namespace,omitempty"`

	/* Repository name

	Required: true
	*/
	Repository string `json:"repository,omitempty"`

	/* Visibility which the repository will start with

	Required: true
	*/
	Visibility string `json:"visibility,omitempty"`
}

// Validate validates this new repo
func (m *NewRepo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRepository(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVisibility(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewRepo) validateDescription(formats strfmt.Registry) error {

	if err := validate.RequiredString("description", "body", string(m.Description)); err != nil {
		return err
	}

	return nil
}

func (m *NewRepo) validateRepository(formats strfmt.Registry) error {

	if err := validate.RequiredString("repository", "body", string(m.Repository)); err != nil {
		return err
	}

	return nil
}

var newRepoVisibilityEnum []interface{}

func (m *NewRepo) validateVisibilityEnum(path, location string, value string) error {
	if newRepoVisibilityEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["public","private"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			newRepoVisibilityEnum = append(newRepoVisibilityEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, newRepoVisibilityEnum); err != nil {
		return err
	}
	return nil
}

func (m *NewRepo) validateVisibility(formats strfmt.Registry) error {

	if err := validate.RequiredString("visibility", "body", string(m.Visibility)); err != nil {
		return err
	}

	if err := m.validateVisibilityEnum("visibility", "body", m.Visibility); err != nil {
		return err
	}

	return nil
}
