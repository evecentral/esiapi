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

// GetUniverseSystemsSystemIDOKBody get_universe_systems_system_id_ok
//
// 200 ok object
// swagger:model getUniverseSystemsSystemIdOKBody
type GetUniverseSystemsSystemIDOKBody struct {

	// get_universe_systems_system_id_constellation_id
	//
	// The constellation this solar system is in
	// Required: true
	ConstellationID *int32 `json:"constellation_id"`

	// get_universe_systems_system_id_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// get_universe_systems_system_id_planets
	//
	// planets array
	// Required: true
	// Max Items: 1000
	Planets []*GetUniverseSystemsSystemIDOKBodyPlanetsItems `json:"planets"`

	// position
	// Required: true
	Position *GetUniverseSystemsSystemIDOKBodyPosition `json:"position"`

	// get_universe_systems_system_id_security_class
	//
	// security_class string
	SecurityClass string `json:"security_class,omitempty"`

	// get_universe_systems_system_id_security_status
	//
	// security_status number
	// Required: true
	SecurityStatus *float32 `json:"security_status"`

	// get_universe_systems_system_id_star_id
	//
	// star_id integer
	// Required: true
	StarID *int32 `json:"star_id"`

	// get_universe_systems_system_id_stargates
	//
	// stargates array
	// Max Items: 25
	Stargates []int32 `json:"stargates"`

	// get_universe_systems_system_id_stations
	//
	// stations array
	// Max Items: 25
	Stations []int32 `json:"stations"`

	// get_universe_systems_system_id_system_id
	//
	// system_id integer
	// Required: true
	SystemID *int32 `json:"system_id"`
}

// Validate validates this get universe systems system Id o k body
func (m *GetUniverseSystemsSystemIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConstellationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStarID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStargates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetUniverseSystemsSystemIDOKBody) validateConstellationID(formats strfmt.Registry) error {

	if err := validate.Required("constellation_id", "body", m.ConstellationID); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseSystemsSystemIDOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseSystemsSystemIDOKBody) validatePlanets(formats strfmt.Registry) error {

	if err := validate.Required("planets", "body", m.Planets); err != nil {
		return err
	}

	iPlanetsSize := int64(len(m.Planets))

	if err := validate.MaxItems("planets", "body", iPlanetsSize, 1000); err != nil {
		return err
	}

	for i := 0; i < len(m.Planets); i++ {
		if swag.IsZero(m.Planets[i]) { // not required
			continue
		}

		if m.Planets[i] != nil {
			if err := m.Planets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("planets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetUniverseSystemsSystemIDOKBody) validatePosition(formats strfmt.Registry) error {

	if err := validate.Required("position", "body", m.Position); err != nil {
		return err
	}

	if m.Position != nil {
		if err := m.Position.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("position")
			}
			return err
		}
	}

	return nil
}

func (m *GetUniverseSystemsSystemIDOKBody) validateSecurityStatus(formats strfmt.Registry) error {

	if err := validate.Required("security_status", "body", m.SecurityStatus); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseSystemsSystemIDOKBody) validateStarID(formats strfmt.Registry) error {

	if err := validate.Required("star_id", "body", m.StarID); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseSystemsSystemIDOKBody) validateStargates(formats strfmt.Registry) error {

	if swag.IsZero(m.Stargates) { // not required
		return nil
	}

	iStargatesSize := int64(len(m.Stargates))

	if err := validate.MaxItems("stargates", "body", iStargatesSize, 25); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseSystemsSystemIDOKBody) validateStations(formats strfmt.Registry) error {

	if swag.IsZero(m.Stations) { // not required
		return nil
	}

	iStationsSize := int64(len(m.Stations))

	if err := validate.MaxItems("stations", "body", iStationsSize, 25); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseSystemsSystemIDOKBody) validateSystemID(formats strfmt.Registry) error {

	if err := validate.Required("system_id", "body", m.SystemID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetUniverseSystemsSystemIDOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUniverseSystemsSystemIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseSystemsSystemIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
