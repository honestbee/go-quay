package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*NewPrototype Description of a new prototype

swagger:model NewPrototype
*/
type NewPrototype struct {

	/* ActivatingUser activating user
	 */
	ActivatingUser *NewPrototypeActivatingUser `json:"activating_user,omitempty"`

	/* Delegate delegate

	Required: true
	*/
	Delegate *NewPrototypeDelegate `json:"delegate,omitempty"`

	/* Role that should be applied to the delegate

	Required: true
	*/
	Role string `json:"role,omitempty"`
}

// Validate validates this new prototype
func (m *NewPrototype) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivatingUser(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDelegate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewPrototype) validateActivatingUser(formats strfmt.Registry) error {

	if swag.IsZero(m.ActivatingUser) { // not required
		return nil
	}

	if m.ActivatingUser != nil {

		if err := m.ActivatingUser.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *NewPrototype) validateDelegate(formats strfmt.Registry) error {

	if m.Delegate != nil {

		if err := m.Delegate.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

var newPrototypeRoleEnum []interface{}

func (m *NewPrototype) validateRoleEnum(path, location string, value string) error {
	if newPrototypeRoleEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["read","write","admin"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			newPrototypeRoleEnum = append(newPrototypeRoleEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, newPrototypeRoleEnum); err != nil {
		return err
	}
	return nil
}

func (m *NewPrototype) validateRole(formats strfmt.Registry) error {

	if err := validate.RequiredString("role", "body", string(m.Role)); err != nil {
		return err
	}

	if err := m.validateRoleEnum("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

/*NewPrototypeActivatingUser Repository creating user to whom the rule should apply

swagger:model NewPrototypeActivatingUser
*/
type NewPrototypeActivatingUser struct {

	/* The username for the activating_user

	Required: true
	*/
	Name string `json:"name,omitempty"`
}

// Validate validates this new prototype activating user
func (m *NewPrototypeActivatingUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewPrototypeActivatingUser) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("activating_user"+"."+"name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

/*NewPrototypeDelegate Information about the user or team to which the rule grants access

swagger:model NewPrototypeDelegate
*/
type NewPrototypeDelegate struct {

	/* Whether the delegate is a user or a team

	Required: true
	*/
	Kind string `json:"kind,omitempty"`

	/* The name for the delegate team or user

	Required: true
	*/
	Name string `json:"name,omitempty"`
}

// Validate validates this new prototype delegate
func (m *NewPrototypeDelegate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var newPrototypeDelegateKindEnum []interface{}

func (m *NewPrototypeDelegate) validateKindEnum(path, location string, value string) error {
	if newPrototypeDelegateKindEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["user","team"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			newPrototypeDelegateKindEnum = append(newPrototypeDelegateKindEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, newPrototypeDelegateKindEnum); err != nil {
		return err
	}
	return nil
}

func (m *NewPrototypeDelegate) validateKind(formats strfmt.Registry) error {

	if err := validate.RequiredString("delegate"+"."+"kind", "body", string(m.Kind)); err != nil {
		return err
	}

	if err := m.validateKindEnum("delegate"+"."+"kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *NewPrototypeDelegate) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("delegate"+"."+"name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}
