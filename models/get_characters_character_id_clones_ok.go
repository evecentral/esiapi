package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// GetCharactersCharacterIDClonesOk 200 ok object
// swagger:model get_characters_character_id_clones_ok
type GetCharactersCharacterIDClonesOk struct {

	// home location
	HomeLocation *GetCharactersCharacterIDClonesHomeLocation `json:"home_location,omitempty"`

	// get_characters_character_id_clones_jump_clones
	//
	// jump_clones array
	// Required: true
	JumpClones []*GetCharactersCharacterIDClonesJumpClone `json:"jump_clones"`

	// get_characters_character_id_clones_last_jump_date
	//
	// last_jump_date string
	LastJumpDate strfmt.DateTime `json:"last_jump_date,omitempty"`
}

// Validate validates this get characters character id clones ok
func (m *GetCharactersCharacterIDClonesOk) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHomeLocation(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateJumpClones(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCharactersCharacterIDClonesOk) validateHomeLocation(formats strfmt.Registry) error {

	if swag.IsZero(m.HomeLocation) { // not required
		return nil
	}

	if m.HomeLocation != nil {

		if err := m.HomeLocation.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *GetCharactersCharacterIDClonesOk) validateJumpClones(formats strfmt.Registry) error {

	if err := validate.Required("jump_clones", "body", m.JumpClones); err != nil {
		return err
	}

	for i := 0; i < len(m.JumpClones); i++ {

		if swag.IsZero(m.JumpClones[i]) { // not required
			continue
		}

		if m.JumpClones[i] != nil {

			if err := m.JumpClones[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
