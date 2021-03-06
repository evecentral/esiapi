// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GetFwLeaderboardsCorporationsOKBodyVictoryPointsYesterdayItems get_fw_leaderboards_corporations_yesterday_yesterday_1
//
// yesterday object
// swagger:model getFwLeaderboardsCorporationsOKBodyVictoryPointsYesterdayItems
type GetFwLeaderboardsCorporationsOKBodyVictoryPointsYesterdayItems struct {

	// get_fw_leaderboards_corporations_yesterday_amount_1
	//
	// Amount of victory points
	Amount int32 `json:"amount,omitempty"`

	// get_fw_leaderboards_corporations_yesterday_corporation_id_1
	//
	// corporation_id integer
	CorporationID int32 `json:"corporation_id,omitempty"`
}

// Validate validates this get fw leaderboards corporations o k body victory points yesterday items
func (m *GetFwLeaderboardsCorporationsOKBodyVictoryPointsYesterdayItems) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetFwLeaderboardsCorporationsOKBodyVictoryPointsYesterdayItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetFwLeaderboardsCorporationsOKBodyVictoryPointsYesterdayItems) UnmarshalBinary(b []byte) error {
	var res GetFwLeaderboardsCorporationsOKBodyVictoryPointsYesterdayItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
