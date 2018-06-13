// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GetAlliancesAllianceIDIconsOKBody get_alliances_alliance_id_icons_ok
//
// 200 ok object
// swagger:model getAlliancesAllianceIdIconsOKBody
type GetAlliancesAllianceIDIconsOKBody struct {

	// get_alliances_alliance_id_icons_px128x128
	//
	// px128x128 string
	Px128x128 string `json:"px128x128,omitempty"`

	// get_alliances_alliance_id_icons_px64x64
	//
	// px64x64 string
	Px64x64 string `json:"px64x64,omitempty"`
}

// Validate validates this get alliances alliance Id icons o k body
func (m *GetAlliancesAllianceIDIconsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetAlliancesAllianceIDIconsOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetAlliancesAllianceIDIconsOKBody) UnmarshalBinary(b []byte) error {
	var res GetAlliancesAllianceIDIconsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}