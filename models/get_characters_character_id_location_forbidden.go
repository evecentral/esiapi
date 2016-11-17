package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// GetCharactersCharacterIDLocationForbidden Forbidden
// swagger:model get_characters_character_id_location_forbidden
type GetCharactersCharacterIDLocationForbidden struct {

	// get_characters_character_id_location_403_forbidden
	//
	// Forbidden message
	Error string `json:"error,omitempty"`
}

// Validate validates this get characters character id location forbidden
func (m *GetCharactersCharacterIDLocationForbidden) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}