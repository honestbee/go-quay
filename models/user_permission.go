package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*Description of a user permission.

swagger:model UserPermission
*/
type UserPermission struct {

	/* Role to use for the user

	Required: true
	*/
	Role string `json:"role,omitempty"`
}

// Validate validates this user permission
func (m *UserPermission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRole(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var userPermissionRoleEnum []interface{}

func (m *UserPermission) validateRoleEnum(path, location string, value string) error {
	if userPermissionRoleEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["read","write","admin"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			userPermissionRoleEnum = append(userPermissionRoleEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, userPermissionRoleEnum); err != nil {
		return err
	}
	return nil
}

func (m *UserPermission) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", string(m.Role)); err != nil {
		return err
	}

	if err := m.validateRoleEnum("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}
