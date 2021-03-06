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

// GetKillmailsKillmailIDKillmailHashOKBody get_killmails_killmail_id_killmail_hash_ok
//
// 200 ok object
// swagger:model getKillmailsKillmailIdKillmailHashOKBody
type GetKillmailsKillmailIDKillmailHashOKBody struct {

	// get_killmails_killmail_id_killmail_hash_attackers
	//
	// attackers array
	// Required: true
	// Max Items: 10000
	Attackers []*GetKillmailsKillmailIDKillmailHashOKBodyAttackersItems `json:"attackers"`

	// get_killmails_killmail_id_killmail_hash_killmail_id
	//
	// ID of the killmail
	// Required: true
	KillmailID *int32 `json:"killmail_id"`

	// get_killmails_killmail_id_killmail_hash_killmail_time
	//
	// Time that the victim was killed and the killmail generated
	//
	// Required: true
	// Format: date-time
	KillmailTime *strfmt.DateTime `json:"killmail_time"`

	// get_killmails_killmail_id_killmail_hash_moon_id
	//
	// Moon if the kill took place at one
	MoonID int32 `json:"moon_id,omitempty"`

	// get_killmails_killmail_id_killmail_hash_solar_system_id
	//
	// Solar system that the kill took place in
	//
	// Required: true
	SolarSystemID *int32 `json:"solar_system_id"`

	// victim
	// Required: true
	Victim *GetKillmailsKillmailIDKillmailHashOKBodyVictim `json:"victim"`

	// get_killmails_killmail_id_killmail_hash_war_id
	//
	// War if the killmail is generated in relation to an official war
	//
	WarID int32 `json:"war_id,omitempty"`
}

// Validate validates this get killmails killmail Id killmail hash o k body
func (m *GetKillmailsKillmailIDKillmailHashOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttackers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKillmailID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKillmailTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSolarSystemID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVictim(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetKillmailsKillmailIDKillmailHashOKBody) validateAttackers(formats strfmt.Registry) error {

	if err := validate.Required("attackers", "body", m.Attackers); err != nil {
		return err
	}

	iAttackersSize := int64(len(m.Attackers))

	if err := validate.MaxItems("attackers", "body", iAttackersSize, 10000); err != nil {
		return err
	}

	for i := 0; i < len(m.Attackers); i++ {
		if swag.IsZero(m.Attackers[i]) { // not required
			continue
		}

		if m.Attackers[i] != nil {
			if err := m.Attackers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attackers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetKillmailsKillmailIDKillmailHashOKBody) validateKillmailID(formats strfmt.Registry) error {

	if err := validate.Required("killmail_id", "body", m.KillmailID); err != nil {
		return err
	}

	return nil
}

func (m *GetKillmailsKillmailIDKillmailHashOKBody) validateKillmailTime(formats strfmt.Registry) error {

	if err := validate.Required("killmail_time", "body", m.KillmailTime); err != nil {
		return err
	}

	if err := validate.FormatOf("killmail_time", "body", "date-time", m.KillmailTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetKillmailsKillmailIDKillmailHashOKBody) validateSolarSystemID(formats strfmt.Registry) error {

	if err := validate.Required("solar_system_id", "body", m.SolarSystemID); err != nil {
		return err
	}

	return nil
}

func (m *GetKillmailsKillmailIDKillmailHashOKBody) validateVictim(formats strfmt.Registry) error {

	if err := validate.Required("victim", "body", m.Victim); err != nil {
		return err
	}

	if m.Victim != nil {
		if err := m.Victim.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("victim")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetKillmailsKillmailIDKillmailHashOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetKillmailsKillmailIDKillmailHashOKBody) UnmarshalBinary(b []byte) error {
	var res GetKillmailsKillmailIDKillmailHashOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
