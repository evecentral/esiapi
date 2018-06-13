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

// GetUniverseAsteroidBeltsAsteroidBeltIDOKBodyPosition get_universe_asteroid_belts_asteroid_belt_id_position
//
// position object
// swagger:model getUniverseAsteroidBeltsAsteroidBeltIdOKBodyPosition
type GetUniverseAsteroidBeltsAsteroidBeltIDOKBodyPosition struct {

	// get_universe_asteroid_belts_asteroid_belt_id_x
	//
	// x number
	// Required: true
	X *float64 `json:"x"`

	// get_universe_asteroid_belts_asteroid_belt_id_y
	//
	// y number
	// Required: true
	Y *float64 `json:"y"`

	// get_universe_asteroid_belts_asteroid_belt_id_z
	//
	// z number
	// Required: true
	Z *float64 `json:"z"`
}

// Validate validates this get universe asteroid belts asteroid belt Id o k body position
func (m *GetUniverseAsteroidBeltsAsteroidBeltIDOKBodyPosition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateX(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateY(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZ(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetUniverseAsteroidBeltsAsteroidBeltIDOKBodyPosition) validateX(formats strfmt.Registry) error {

	if err := validate.Required("x", "body", m.X); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseAsteroidBeltsAsteroidBeltIDOKBodyPosition) validateY(formats strfmt.Registry) error {

	if err := validate.Required("y", "body", m.Y); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseAsteroidBeltsAsteroidBeltIDOKBodyPosition) validateZ(formats strfmt.Registry) error {

	if err := validate.Required("z", "body", m.Z); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetUniverseAsteroidBeltsAsteroidBeltIDOKBodyPosition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUniverseAsteroidBeltsAsteroidBeltIDOKBodyPosition) UnmarshalBinary(b []byte) error {
	var res GetUniverseAsteroidBeltsAsteroidBeltIDOKBodyPosition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
