package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// GetCorporationsCorporationIDMembersInternalServerError Internal server error
// swagger:model get_corporations_corporation_id_members_internal_server_error
type GetCorporationsCorporationIDMembersInternalServerError struct {

	// get_corporations_corporation_id_members_500_internal_server_error
	//
	// Internal server error message
	Error string `json:"error,omitempty"`
}

// Validate validates this get corporations corporation id members internal server error
func (m *GetCorporationsCorporationIDMembersInternalServerError) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
