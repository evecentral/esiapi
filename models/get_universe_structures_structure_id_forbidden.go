package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// GetUniverseStructuresStructureIDForbidden Forbidden
// swagger:model get_universe_structures_structure_id_forbidden
type GetUniverseStructuresStructureIDForbidden struct {

	// get_universe_structures_structure_id_403_forbidden
	//
	// Forbidden message
	Error string `json:"error,omitempty"`
}

// Validate validates this get universe structures structure id forbidden
func (m *GetUniverseStructuresStructureIDForbidden) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
