package discovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*
DiscoveryParams contains all the parameters to send to the API endpoint
for the discovery operation typically these are written to a http.Request
*/
type DiscoveryParams struct {
	/*
	  Whether to include internal APIs.
	*/
	Internal *bool
}

// WriteToRequest writes these params to a swagger request
func (o *DiscoveryParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Internal != nil {

		// query param internal
		if err := r.SetQueryParam("internal", swag.FormatBool(*o.Internal)); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
