// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GetUniverseConstellationsConstellationIDNotFoundBody get_universe_constellations_constellation_id_not_found
//
// Not found
// swagger:model getUniverseConstellationsConstellationIdNotFoundBody
type GetUniverseConstellationsConstellationIDNotFoundBody struct {

	// get_universe_constellations_constellation_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get universe constellations constellation Id not found body
func (m *GetUniverseConstellationsConstellationIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetUniverseConstellationsConstellationIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUniverseConstellationsConstellationIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseConstellationsConstellationIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
