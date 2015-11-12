package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*
GetOrganizationApplicationParams contains all the parameters to send to the API endpoint
for the get organization application operation typically these are written to a http.Request
*/
type GetOrganizationApplicationParams struct {
	/*
	  The OAuth client ID
	*/
	ClientID string
	/*
	  The name of the organization
	*/
	Orgname string
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationApplicationParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param client_id
	if err := r.SetPathParam("client_id", o.ClientID); err != nil {
		return err
	}

	// path param orgname
	if err := r.SetPathParam("orgname", o.Orgname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
