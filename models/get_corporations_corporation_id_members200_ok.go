package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// GetCorporationsCorporationIDMembers200Ok 200 ok object
// swagger:model get_corporations_corporation_id_members_200_ok
type GetCorporationsCorporationIDMembers200Ok struct {

	// get_corporations_corporation_id_members_character_id
	//
	// character_id integer
	// Required: true
	CharacterID *int32 `json:"character_id"`
}

// Validate validates this get corporations corporation id members 200 ok
func (m *GetCorporationsCorporationIDMembers200Ok) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCharacterID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCorporationsCorporationIDMembers200Ok) validateCharacterID(formats strfmt.Registry) error {

	if err := validate.Required("character_id", "body", m.CharacterID); err != nil {
		return err
	}

	return nil
}