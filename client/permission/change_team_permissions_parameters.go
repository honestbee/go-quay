package permission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/coreos/go-quay/models"
)

/*
ChangeTeamPermissionsParams contains all the parameters to send to the API endpoint
for the change team permissions operation typically these are written to a http.Request
*/
type ChangeTeamPermissionsParams struct {
	/*
	  Request body contents.
	*/
	Body *models.TeamPermission
	/*
	  The full path of the repository. e.g. namespace/name
	*/
	Repository string
	/*
	  The name of the team to which the permission applies
	*/
	Teamname string
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeTeamPermissionsParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	// path param teamname
	if err := r.SetPathParam("teamname", o.Teamname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
