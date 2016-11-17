package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// GetCharactersCharacterIDClonesHomeLocation home_location object
// swagger:model get_characters_character_id_clones_home_location
type GetCharactersCharacterIDClonesHomeLocation struct {

	// get_characters_character_id_clones_location_id
	//
	// location_id integer
	LocationID int64 `json:"location_id,omitempty"`

	// get_characters_character_id_clones_location_type
	//
	// location_type string
	LocationType string `json:"location_type,omitempty"`
}

// Validate validates this get characters character id clones home location
func (m *GetCharactersCharacterIDClonesHomeLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocationType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getCharactersCharacterIdClonesHomeLocationTypeLocationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["station","structure"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdClonesHomeLocationTypeLocationTypePropEnum = append(getCharactersCharacterIdClonesHomeLocationTypeLocationTypePropEnum, v)
	}
}

const (
	getCharactersCharacterIdClonesHomeLocationLocationTypeStation   string = "station"
	getCharactersCharacterIdClonesHomeLocationLocationTypeStructure string = "structure"
)

// prop value enum
func (m *GetCharactersCharacterIDClonesHomeLocation) validateLocationTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCharactersCharacterIdClonesHomeLocationTypeLocationTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCharactersCharacterIDClonesHomeLocation) validateLocationType(formats strfmt.Registry) error {

	if swag.IsZero(m.LocationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateLocationTypeEnum("location_type", "body", m.LocationType); err != nil {
		return err
	}

	return nil
}