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

// GetCharactersCharacterIDPlanetsPlanetIDOKBodyLinksItems get_characters_character_id_planets_planet_id_link
//
// link object
// swagger:model getCharactersCharacterIdPlanetsPlanetIdOKBodyLinksItems
type GetCharactersCharacterIDPlanetsPlanetIDOKBodyLinksItems struct {

	// get_characters_character_id_planets_planet_id_destination_pin_id
	//
	// destination_pin_id integer
	// Required: true
	DestinationPinID *int64 `json:"destination_pin_id"`

	// get_characters_character_id_planets_planet_id_link_level
	//
	// link_level integer
	// Required: true
	// Maximum: 10
	// Minimum: 0
	LinkLevel *int32 `json:"link_level"`

	// get_characters_character_id_planets_planet_id_source_pin_id
	//
	// source_pin_id integer
	// Required: true
	SourcePinID *int64 `json:"source_pin_id"`
}

// Validate validates this get characters character Id planets planet Id o k body links items
func (m *GetCharactersCharacterIDPlanetsPlanetIDOKBodyLinksItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinationPinID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinkLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourcePinID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCharactersCharacterIDPlanetsPlanetIDOKBodyLinksItems) validateDestinationPinID(formats strfmt.Registry) error {

	if err := validate.Required("destination_pin_id", "body", m.DestinationPinID); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDPlanetsPlanetIDOKBodyLinksItems) validateLinkLevel(formats strfmt.Registry) error {

	if err := validate.Required("link_level", "body", m.LinkLevel); err != nil {
		return err
	}

	if err := validate.MinimumInt("link_level", "body", int64(*m.LinkLevel), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("link_level", "body", int64(*m.LinkLevel), 10, false); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDPlanetsPlanetIDOKBodyLinksItems) validateSourcePinID(formats strfmt.Registry) error {

	if err := validate.Required("source_pin_id", "body", m.SourcePinID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDPlanetsPlanetIDOKBodyLinksItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDPlanetsPlanetIDOKBodyLinksItems) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDPlanetsPlanetIDOKBodyLinksItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
