package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// GetUniverseTypesTypeIDNotFound Type not found
// swagger:model get_universe_types_type_id_not_found
type GetUniverseTypesTypeIDNotFound struct {

	// get_universe_types_type_id_error
	//
	// error message
	Error string `json:"error,omitempty"`
}

// Validate validates this get universe types type id not found
func (m *GetUniverseTypesTypeIDNotFound) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
