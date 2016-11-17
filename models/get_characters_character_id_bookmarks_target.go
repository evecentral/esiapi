package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// GetCharactersCharacterIDBookmarksTarget target object
// swagger:model get_characters_character_id_bookmarks_target
type GetCharactersCharacterIDBookmarksTarget struct {

	// coordinates
	Coordinates *GetCharactersCharacterIDBookmarksCoordinates `json:"coordinates,omitempty"`

	// item
	Item *GetCharactersCharacterIDBookmarksItem `json:"item,omitempty"`

	// get_characters_character_id_bookmarks_location_id
	//
	// location_id integer
	// Required: true
	LocationID *int64 `json:"location_id"`
}

// Validate validates this get characters character id bookmarks target
func (m *GetCharactersCharacterIDBookmarksTarget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCoordinates(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateItem(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLocationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCharactersCharacterIDBookmarksTarget) validateCoordinates(formats strfmt.Registry) error {

	if swag.IsZero(m.Coordinates) { // not required
		return nil
	}

	if m.Coordinates != nil {

		if err := m.Coordinates.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *GetCharactersCharacterIDBookmarksTarget) validateItem(formats strfmt.Registry) error {

	if swag.IsZero(m.Item) { // not required
		return nil
	}

	if m.Item != nil {

		if err := m.Item.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *GetCharactersCharacterIDBookmarksTarget) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("location_id", "body", m.LocationID); err != nil {
		return err
	}

	return nil
}
