package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// PostCharactersCharacterIDMailForbidden Forbidden
// swagger:model post_characters_character_id_mail_forbidden
type PostCharactersCharacterIDMailForbidden struct {

	// post_characters_character_id_mail_403_forbidden
	//
	// Forbidden message
	Error string `json:"error,omitempty"`
}

// Validate validates this post characters character id mail forbidden
func (m *PostCharactersCharacterIDMailForbidden) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
