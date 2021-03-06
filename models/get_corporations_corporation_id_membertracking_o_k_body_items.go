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

// GetCorporationsCorporationIDMembertrackingOKBodyItems get_corporations_corporation_id_membertracking_200_ok
//
// 200 ok object
// swagger:model getCorporationsCorporationIdMembertrackingOKBodyItems
type GetCorporationsCorporationIDMembertrackingOKBodyItems struct {

	// get_corporations_corporation_id_membertracking_base_id
	//
	// base_id integer
	BaseID int32 `json:"base_id,omitempty"`

	// get_corporations_corporation_id_membertracking_character_id
	//
	// character_id integer
	// Required: true
	CharacterID *int32 `json:"character_id"`

	// get_corporations_corporation_id_membertracking_location_id
	//
	// location_id integer
	LocationID int64 `json:"location_id,omitempty"`

	// get_corporations_corporation_id_membertracking_logoff_date
	//
	// logoff_date string
	// Format: date-time
	LogoffDate strfmt.DateTime `json:"logoff_date,omitempty"`

	// get_corporations_corporation_id_membertracking_logon_date
	//
	// logon_date string
	// Format: date-time
	LogonDate strfmt.DateTime `json:"logon_date,omitempty"`

	// get_corporations_corporation_id_membertracking_ship_type_id
	//
	// ship_type_id integer
	ShipTypeID int32 `json:"ship_type_id,omitempty"`

	// get_corporations_corporation_id_membertracking_start_date
	//
	// start_date string
	// Format: date-time
	StartDate strfmt.DateTime `json:"start_date,omitempty"`
}

// Validate validates this get corporations corporation Id membertracking o k body items
func (m *GetCorporationsCorporationIDMembertrackingOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCharacterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogoffDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogonDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCorporationsCorporationIDMembertrackingOKBodyItems) validateCharacterID(formats strfmt.Registry) error {

	if err := validate.Required("character_id", "body", m.CharacterID); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDMembertrackingOKBodyItems) validateLogoffDate(formats strfmt.Registry) error {

	if swag.IsZero(m.LogoffDate) { // not required
		return nil
	}

	if err := validate.FormatOf("logoff_date", "body", "date-time", m.LogoffDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDMembertrackingOKBodyItems) validateLogonDate(formats strfmt.Registry) error {

	if swag.IsZero(m.LogonDate) { // not required
		return nil
	}

	if err := validate.FormatOf("logon_date", "body", "date-time", m.LogonDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDMembertrackingOKBodyItems) validateStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("start_date", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCorporationsCorporationIDMembertrackingOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCorporationsCorporationIDMembertrackingOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDMembertrackingOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
