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

// GetFwLeaderboardsOKBodyKills get_fw_leaderboards_kills
//
// Top 4 rankings of factions by number of kills from yesterday, last week and in total.
// swagger:model getFwLeaderboardsOKBodyKills
type GetFwLeaderboardsOKBodyKills struct {

	// get_fw_leaderboards_active_total
	//
	// Top 4 ranking of factions active in faction warfare by total kills. A faction is considered "active" if they have participated in faction warfare in the past 14 days.
	// Required: true
	// Max Items: 4
	ActiveTotal []*GetFwLeaderboardsOKBodyKillsActiveTotalItems `json:"active_total"`

	// get_fw_leaderboards_last_week
	//
	// Top 4 ranking of factions by kills in the past week
	// Required: true
	// Max Items: 4
	LastWeek []*GetFwLeaderboardsOKBodyKillsLastWeekItems `json:"last_week"`

	// get_fw_leaderboards_yesterday
	//
	// Top 4 ranking of factions by kills in the past day
	// Required: true
	// Max Items: 4
	Yesterday []*GetFwLeaderboardsOKBodyKillsYesterdayItems `json:"yesterday"`
}

// Validate validates this get fw leaderboards o k body kills
func (m *GetFwLeaderboardsOKBodyKills) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActiveTotal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastWeek(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateYesterday(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetFwLeaderboardsOKBodyKills) validateActiveTotal(formats strfmt.Registry) error {

	if err := validate.Required("active_total", "body", m.ActiveTotal); err != nil {
		return err
	}

	iActiveTotalSize := int64(len(m.ActiveTotal))

	if err := validate.MaxItems("active_total", "body", iActiveTotalSize, 4); err != nil {
		return err
	}

	for i := 0; i < len(m.ActiveTotal); i++ {
		if swag.IsZero(m.ActiveTotal[i]) { // not required
			continue
		}

		if m.ActiveTotal[i] != nil {
			if err := m.ActiveTotal[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("active_total" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetFwLeaderboardsOKBodyKills) validateLastWeek(formats strfmt.Registry) error {

	if err := validate.Required("last_week", "body", m.LastWeek); err != nil {
		return err
	}

	iLastWeekSize := int64(len(m.LastWeek))

	if err := validate.MaxItems("last_week", "body", iLastWeekSize, 4); err != nil {
		return err
	}

	for i := 0; i < len(m.LastWeek); i++ {
		if swag.IsZero(m.LastWeek[i]) { // not required
			continue
		}

		if m.LastWeek[i] != nil {
			if err := m.LastWeek[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("last_week" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetFwLeaderboardsOKBodyKills) validateYesterday(formats strfmt.Registry) error {

	if err := validate.Required("yesterday", "body", m.Yesterday); err != nil {
		return err
	}

	iYesterdaySize := int64(len(m.Yesterday))

	if err := validate.MaxItems("yesterday", "body", iYesterdaySize, 4); err != nil {
		return err
	}

	for i := 0; i < len(m.Yesterday); i++ {
		if swag.IsZero(m.Yesterday[i]) { // not required
			continue
		}

		if m.Yesterday[i] != nil {
			if err := m.Yesterday[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("yesterday" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetFwLeaderboardsOKBodyKills) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetFwLeaderboardsOKBodyKills) UnmarshalBinary(b []byte) error {
	var res GetFwLeaderboardsOKBodyKills
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
