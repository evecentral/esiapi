package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// GetUniverseTypesTypeIDInternalServerError Internal server error
// swagger:model get_universe_types_type_id_internal_server_error
type GetUniverseTypesTypeIDInternalServerError struct {

	// get_universe_types_type_id_500_internal_server_error
	//
	// Internal server error message
	Error string `json:"error,omitempty"`
}

// Validate validates this get universe types type id internal server error
func (m *GetUniverseTypesTypeIDInternalServerError) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
