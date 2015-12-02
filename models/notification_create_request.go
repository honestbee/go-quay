package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*Information for creating a notification on a repository

swagger:model NotificationCreateRequest
*/
type NotificationCreateRequest struct {

	/* JSON config information for the specific method of notification

	Required: true
	*/
	Config interface{} `json:"config,omitempty"`

	/* The event on which the notification will respond

	Required: true
	*/
	Event string `json:"event,omitempty"`

	/* JSON config information for the specific event of notification
	 */
	EventConfig interface{} `json:"eventConfig,omitempty"`

	/* The method of notification (such as email or web callback)

	Required: true
	*/
	Method string `json:"method,omitempty"`

	/* The human-readable title of the notification
	 */
	Title string `json:"title,omitempty"`
}

// Validate validates this notification create request
func (m *NotificationCreateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEvent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMethod(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NotificationCreateRequest) validateConfig(formats strfmt.Registry) error {

	return nil
}

func (m *NotificationCreateRequest) validateEvent(formats strfmt.Registry) error {

	if err := validate.Required("event", "body", string(m.Event)); err != nil {
		return err
	}

	return nil
}

func (m *NotificationCreateRequest) validateMethod(formats strfmt.Registry) error {

	if err := validate.Required("method", "body", string(m.Method)); err != nil {
		return err
	}

	return nil
}
