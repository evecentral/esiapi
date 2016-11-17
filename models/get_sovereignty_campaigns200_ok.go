package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// GetSovereigntyCampaigns200Ok 200 ok object
// swagger:model get_sovereignty_campaigns_200_ok
type GetSovereigntyCampaigns200Ok struct {

	// get_sovereignty_campaigns_attackers_score
	//
	// Score for all attacking parties, only present in Defense Events.
	//
	AttackersScore float32 `json:"attackers_score,omitempty"`

	// get_sovereignty_campaigns_campaign_id
	//
	// Unique ID for this campaign.
	// Required: true
	CampaignID *int64 `json:"campaign_id"`

	// get_sovereignty_campaigns_constellation_id
	//
	// The constellation in which the campaign will take place.
	//
	// Required: true
	ConstellationID *int64 `json:"constellation_id"`

	// get_sovereignty_campaigns_defender_id
	//
	// Defending alliance, only present in Defense Events
	//
	DefenderID int64 `json:"defender_id,omitempty"`

	// get_sovereignty_campaigns_defender_score
	//
	// Score for the defending alliance, only present in Defense Events.
	//
	DefenderScore float32 `json:"defender_score,omitempty"`

	// get_sovereignty_campaigns_event_type
	//
	// Type of event this campaign is for. tcu_defense, ihub_defense and station_defense are referred to as "Defense Events", station_freeport as "Freeport Events".
	//
	// Required: true
	EventType *string `json:"event_type"`

	// get_sovereignty_campaigns_participants
	//
	// Alliance participating and their respective scores, only present in Freeport Events.
	//
	Participants []*GetSovereigntyCampaignsParticipant `json:"participants"`

	// get_sovereignty_campaigns_solar_system_id
	//
	// The solar system the structure is located in.
	//
	// Required: true
	SolarSystemID *int64 `json:"solar_system_id"`

	// get_sovereignty_campaigns_start_time
	//
	// Time the event is scheduled to start.
	//
	// Required: true
	StartTime *strfmt.DateTime `json:"start_time"`

	// get_sovereignty_campaigns_structure_id
	//
	// The structure item ID that is related to this campaign.
	//
	// Required: true
	StructureID *int64 `json:"structure_id"`
}

// Validate validates this get sovereignty campaigns 200 ok
func (m *GetSovereigntyCampaigns200Ok) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCampaignID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConstellationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEventType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateParticipants(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSolarSystemID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStructureID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSovereigntyCampaigns200Ok) validateCampaignID(formats strfmt.Registry) error {

	if err := validate.Required("campaign_id", "body", m.CampaignID); err != nil {
		return err
	}

	return nil
}

func (m *GetSovereigntyCampaigns200Ok) validateConstellationID(formats strfmt.Registry) error {

	if err := validate.Required("constellation_id", "body", m.ConstellationID); err != nil {
		return err
	}

	return nil
}

var getSovereigntyCampaigns200OkTypeEventTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tcu_defense","ihub_defense","station_defense","station_freeport"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getSovereigntyCampaigns200OkTypeEventTypePropEnum = append(getSovereigntyCampaigns200OkTypeEventTypePropEnum, v)
	}
}

const (
	getSovereigntyCampaigns200OkEventTypeTcuDefense      string = "tcu_defense"
	getSovereigntyCampaigns200OkEventTypeIhubDefense     string = "ihub_defense"
	getSovereigntyCampaigns200OkEventTypeStationDefense  string = "station_defense"
	getSovereigntyCampaigns200OkEventTypeStationFreeport string = "station_freeport"
)

// prop value enum
func (m *GetSovereigntyCampaigns200Ok) validateEventTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getSovereigntyCampaigns200OkTypeEventTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetSovereigntyCampaigns200Ok) validateEventType(formats strfmt.Registry) error {

	if err := validate.Required("event_type", "body", m.EventType); err != nil {
		return err
	}

	// value enum
	if err := m.validateEventTypeEnum("event_type", "body", *m.EventType); err != nil {
		return err
	}

	return nil
}

func (m *GetSovereigntyCampaigns200Ok) validateParticipants(formats strfmt.Registry) error {

	if swag.IsZero(m.Participants) { // not required
		return nil
	}

	for i := 0; i < len(m.Participants); i++ {

		if swag.IsZero(m.Participants[i]) { // not required
			continue
		}

		if m.Participants[i] != nil {

			if err := m.Participants[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *GetSovereigntyCampaigns200Ok) validateSolarSystemID(formats strfmt.Registry) error {

	if err := validate.Required("solar_system_id", "body", m.SolarSystemID); err != nil {
		return err
	}

	return nil
}

func (m *GetSovereigntyCampaigns200Ok) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("start_time", "body", m.StartTime); err != nil {
		return err
	}

	return nil
}

func (m *GetSovereigntyCampaigns200Ok) validateStructureID(formats strfmt.Registry) error {

	if err := validate.Required("structure_id", "body", m.StructureID); err != nil {
		return err
	}

	return nil
}
