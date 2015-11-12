package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*
Description of a the new prototype role

swagger:model PrototypeUpdate
*/
type PrototypeUpdate struct {

	/* Role that should be applied to the permission

	Required: true
	*/
	Role string `json:"role"`
}

// Validate validates this prototype update
func (m *PrototypeUpdate) Validate(formats strfmt.Registry) error {

	var res []error

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}

	return nil
}

var prototypeUpdateRoleEnum []interface{}

func (m *PrototypeUpdate) validateRoleEnum(path, location string, value string) error {
	if prototypeUpdateRoleEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["read","write","admin"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			prototypeUpdateRoleEnum = append(prototypeUpdateRoleEnum, v)
		}
	}
	return validate.Enum(path, location, value, prototypeUpdateRoleEnum)
}

func (m *PrototypeUpdate) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", string(m.Role)); err != nil {
		return err
	}

	if err := m.validateRoleEnum("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}
