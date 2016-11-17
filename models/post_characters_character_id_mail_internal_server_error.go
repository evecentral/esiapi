package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// PostCharactersCharacterIDMailInternalServerError Internal server error
// swagger:model post_characters_character_id_mail_internal_server_error
type PostCharactersCharacterIDMailInternalServerError struct {

	// post_characters_character_id_mail_500_internal_server_error
	//
	// Internal server error message
	Error string `json:"error,omitempty"`
}

// Validate validates this post characters character id mail internal server error
func (m *PostCharactersCharacterIDMailInternalServerError) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
