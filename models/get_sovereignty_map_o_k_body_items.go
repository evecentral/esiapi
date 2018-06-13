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

// GetSovereigntyMapOKBodyItems get_sovereignty_map_200_ok
//
// 200 ok object
// swagger:model getSovereigntyMapOKBodyItems
type GetSovereigntyMapOKBodyItems struct {

	// get_sovereignty_map_alliance_id
	//
	// alliance_id integer
	AllianceID int32 `json:"alliance_id,omitempty"`

	// get_sovereignty_map_corporation_id
	//
	// corporation_id integer
	CorporationID int32 `json:"corporation_id,omitempty"`

	// get_sovereignty_map_faction_id
	//
	// faction_id integer
	FactionID int32 `json:"faction_id,omitempty"`

	// get_sovereignty_map_system_id
	//
	// system_id integer
	// Required: true
	SystemID *int32 `json:"system_id"`
}

// Validate validates this get sovereignty map o k body items
func (m *GetSovereigntyMapOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSystemID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSovereigntyMapOKBodyItems) validateSystemID(formats strfmt.Registry) error {

	if err := validate.Required("system_id", "body", m.SystemID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetSovereigntyMapOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSovereigntyMapOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetSovereigntyMapOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
