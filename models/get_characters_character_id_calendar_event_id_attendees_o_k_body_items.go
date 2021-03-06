// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetCharactersCharacterIDCalendarEventIDAttendeesOKBodyItems get_characters_character_id_calendar_event_id_attendees_200_ok
//
// character_id and response of an attendee
// swagger:model getCharactersCharacterIdCalendarEventIdAttendeesOKBodyItems
type GetCharactersCharacterIDCalendarEventIDAttendeesOKBodyItems struct {

	// get_characters_character_id_calendar_event_id_attendees_character_id
	//
	// character_id integer
	CharacterID int32 `json:"character_id,omitempty"`

	// get_characters_character_id_calendar_event_id_attendees_event_response
	//
	// event_response string
	// Enum: [declined not_responded accepted tentative]
	EventResponse string `json:"event_response,omitempty"`
}

// Validate validates this get characters character Id calendar event Id attendees o k body items
func (m *GetCharactersCharacterIDCalendarEventIDAttendeesOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getCharactersCharacterIdCalendarEventIdAttendeesOKBodyItemsTypeEventResponsePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["declined","not_responded","accepted","tentative"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdCalendarEventIdAttendeesOKBodyItemsTypeEventResponsePropEnum = append(getCharactersCharacterIdCalendarEventIdAttendeesOKBodyItemsTypeEventResponsePropEnum, v)
	}
}

const (

	// GetCharactersCharacterIDCalendarEventIDAttendeesOKBodyItemsEventResponseDeclined captures enum value "declined"
	GetCharactersCharacterIDCalendarEventIDAttendeesOKBodyItemsEventResponseDeclined string = "declined"

	// GetCharactersCharacterIDCalendarEventIDAttendeesOKBodyItemsEventResponseNotResponded captures enum value "not_responded"
	GetCharactersCharacterIDCalendarEventIDAttendeesOKBodyItemsEventResponseNotResponded string = "not_responded"

	// GetCharactersCharacterIDCalendarEventIDAttendeesOKBodyItemsEventResponseAccepted captures enum value "accepted"
	GetCharactersCharacterIDCalendarEventIDAttendeesOKBodyItemsEventResponseAccepted string = "accepted"

	// GetCharactersCharacterIDCalendarEventIDAttendeesOKBodyItemsEventResponseTentative captures enum value "tentative"
	GetCharactersCharacterIDCalendarEventIDAttendeesOKBodyItemsEventResponseTentative string = "tentative"
)

// prop value enum
func (m *GetCharactersCharacterIDCalendarEventIDAttendeesOKBodyItems) validateEventResponseEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCharactersCharacterIdCalendarEventIdAttendeesOKBodyItemsTypeEventResponsePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCharactersCharacterIDCalendarEventIDAttendeesOKBodyItems) validateEventResponse(formats strfmt.Registry) error {

	if swag.IsZero(m.EventResponse) { // not required
		return nil
	}

	// value enum
	if err := m.validateEventResponseEnum("event_response", "body", m.EventResponse); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDCalendarEventIDAttendeesOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDCalendarEventIDAttendeesOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDCalendarEventIDAttendeesOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
