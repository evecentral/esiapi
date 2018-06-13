// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GetFwLeaderboardsCorporationsOKBodyVictoryPointsLastWeekItems get_fw_leaderboards_corporations_last_week_last_week_1
//
// last_week object
// swagger:model getFwLeaderboardsCorporationsOKBodyVictoryPointsLastWeekItems
type GetFwLeaderboardsCorporationsOKBodyVictoryPointsLastWeekItems struct {

	// get_fw_leaderboards_corporations_last_week_amount_1
	//
	// Amount of victory points
	Amount int32 `json:"amount,omitempty"`

	// get_fw_leaderboards_corporations_last_week_corporation_id_1
	//
	// corporation_id integer
	CorporationID int32 `json:"corporation_id,omitempty"`
}

// Validate validates this get fw leaderboards corporations o k body victory points last week items
func (m *GetFwLeaderboardsCorporationsOKBodyVictoryPointsLastWeekItems) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetFwLeaderboardsCorporationsOKBodyVictoryPointsLastWeekItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetFwLeaderboardsCorporationsOKBodyVictoryPointsLastWeekItems) UnmarshalBinary(b []byte) error {
	var res GetFwLeaderboardsCorporationsOKBodyVictoryPointsLastWeekItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}