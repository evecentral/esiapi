// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostCharactersAffiliationOKBodyItems post_characters_affiliation_200_ok
//
// 200 ok object
// swagger:model postCharactersAffiliationOKBodyItems
type PostCharactersAffiliationOKBodyItems struct {

	// post_characters_affiliation_alliance_id
	//
	// The character's alliance ID, if their corporation is in an alliance
	AllianceID int32 `json:"alliance_id,omitempty"`

	// post_characters_affiliation_character_id
	//
	// The character's ID
	// Required: true
	CharacterID *int32 `json:"character_id"`

	// post_characters_affiliation_corporation_id
	//
	// The character's corporation ID
	// Required: true
	CorporationID *int32 `json:"corporation_id"`

	// post_characters_affiliation_faction_id
	//
	// The character's faction ID, if their corporation is in a faction
	FactionID int32 `json:"faction_id,omitempty"`
}

// Validate validates this post characters affiliation o k body items
func (m *PostCharactersAffiliationOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCharacterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCorporationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostCharactersAffiliationOKBodyItems) validateCharacterID(formats strfmt.Registry) error {

	if err := validate.Required("character_id", "body", m.CharacterID); err != nil {
		return err
	}

	return nil
}

func (m *PostCharactersAffiliationOKBodyItems) validateCorporationID(formats strfmt.Registry) error {

	if err := validate.Required("corporation_id", "body", m.CorporationID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostCharactersAffiliationOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostCharactersAffiliationOKBodyItems) UnmarshalBinary(b []byte) error {
	var res PostCharactersAffiliationOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}