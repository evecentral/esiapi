// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetCharactersCharacterIDClonesOKBody get_characters_character_id_clones_ok
//
// 200 ok object
// swagger:model getCharactersCharacterIdClonesOKBody
type GetCharactersCharacterIDClonesOKBody struct {

	// home location
	HomeLocation *GetCharactersCharacterIDClonesOKBodyHomeLocation `json:"home_location,omitempty"`

	// get_characters_character_id_clones_jump_clones
	//
	// jump_clones array
	// Required: true
	// Max Items: 10
	JumpClones []*GetCharactersCharacterIDClonesOKBodyJumpClonesItems `json:"jump_clones"`

	// get_characters_character_id_clones_last_clone_jump_date
	//
	// last_clone_jump_date string
	// Format: date-time
	LastCloneJumpDate strfmt.DateTime `json:"last_clone_jump_date,omitempty"`

	// get_characters_character_id_clones_last_station_change_date
	//
	// last_station_change_date string
	// Format: date-time
	LastStationChangeDate strfmt.DateTime `json:"last_station_change_date,omitempty"`
}

// Validate validates this get characters character Id clones o k body
func (m *GetCharactersCharacterIDClonesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHomeLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJumpClones(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastCloneJumpDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastStationChangeDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCharactersCharacterIDClonesOKBody) validateHomeLocation(formats strfmt.Registry) error {

	if swag.IsZero(m.HomeLocation) { // not required
		return nil
	}

	if m.HomeLocation != nil {
		if err := m.HomeLocation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("home_location")
			}
			return err
		}
	}

	return nil
}

func (m *GetCharactersCharacterIDClonesOKBody) validateJumpClones(formats strfmt.Registry) error {

	if err := validate.Required("jump_clones", "body", m.JumpClones); err != nil {
		return err
	}

	iJumpClonesSize := int64(len(m.JumpClones))

	if err := validate.MaxItems("jump_clones", "body", iJumpClonesSize, 10); err != nil {
		return err
	}

	for i := 0; i < len(m.JumpClones); i++ {
		if swag.IsZero(m.JumpClones[i]) { // not required
			continue
		}

		if m.JumpClones[i] != nil {
			if err := m.JumpClones[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("jump_clones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetCharactersCharacterIDClonesOKBody) validateLastCloneJumpDate(formats strfmt.Registry) error {

	if swag.IsZero(m.LastCloneJumpDate) { // not required
		return nil
	}

	if err := validate.FormatOf("last_clone_jump_date", "body", "date-time", m.LastCloneJumpDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDClonesOKBody) validateLastStationChangeDate(formats strfmt.Registry) error {

	if swag.IsZero(m.LastStationChangeDate) { // not required
		return nil
	}

	if err := validate.FormatOf("last_station_change_date", "body", "date-time", m.LastStationChangeDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDClonesOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDClonesOKBody) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDClonesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
