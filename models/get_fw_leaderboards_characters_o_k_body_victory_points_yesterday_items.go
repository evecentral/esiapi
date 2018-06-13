// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GetFwLeaderboardsCharactersOKBodyVictoryPointsYesterdayItems get_fw_leaderboards_characters_yesterday_yesterday_1
//
// yesterday object
// swagger:model getFwLeaderboardsCharactersOKBodyVictoryPointsYesterdayItems
type GetFwLeaderboardsCharactersOKBodyVictoryPointsYesterdayItems struct {

	// get_fw_leaderboards_characters_yesterday_amount_1
	//
	// Amount of victory points
	Amount int32 `json:"amount,omitempty"`

	// get_fw_leaderboards_characters_yesterday_character_id_1
	//
	// character_id integer
	CharacterID int32 `json:"character_id,omitempty"`
}

// Validate validates this get fw leaderboards characters o k body victory points yesterday items
func (m *GetFwLeaderboardsCharactersOKBodyVictoryPointsYesterdayItems) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetFwLeaderboardsCharactersOKBodyVictoryPointsYesterdayItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetFwLeaderboardsCharactersOKBodyVictoryPointsYesterdayItems) UnmarshalBinary(b []byte) error {
	var res GetFwLeaderboardsCharactersOKBodyVictoryPointsYesterdayItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
