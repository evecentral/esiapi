package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// GetCharactersCharacterIDPortraitNotFound No image server for this datasource
// swagger:model get_characters_character_id_portrait_not_found
type GetCharactersCharacterIDPortraitNotFound struct {

	// get_characters_character_id_portrait_error
	//
	// error message
	Error string `json:"error,omitempty"`
}

// Validate validates this get characters character id portrait not found
func (m *GetCharactersCharacterIDPortraitNotFound) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}