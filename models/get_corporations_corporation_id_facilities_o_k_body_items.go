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

// GetCorporationsCorporationIDFacilitiesOKBodyItems get_corporations_corporation_id_facilities_200_ok
//
// 200 ok object
// swagger:model getCorporationsCorporationIdFacilitiesOKBodyItems
type GetCorporationsCorporationIDFacilitiesOKBodyItems struct {

	// get_corporations_corporation_id_facilities_facility_id
	//
	// facility_id integer
	// Required: true
	FacilityID *int64 `json:"facility_id"`

	// get_corporations_corporation_id_facilities_system_id
	//
	// system_id integer
	// Required: true
	SystemID *int32 `json:"system_id"`

	// get_corporations_corporation_id_facilities_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this get corporations corporation Id facilities o k body items
func (m *GetCorporationsCorporationIDFacilitiesOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFacilityID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCorporationsCorporationIDFacilitiesOKBodyItems) validateFacilityID(formats strfmt.Registry) error {

	if err := validate.Required("facility_id", "body", m.FacilityID); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDFacilitiesOKBodyItems) validateSystemID(formats strfmt.Registry) error {

	if err := validate.Required("system_id", "body", m.SystemID); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDFacilitiesOKBodyItems) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", m.TypeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCorporationsCorporationIDFacilitiesOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCorporationsCorporationIDFacilitiesOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDFacilitiesOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
