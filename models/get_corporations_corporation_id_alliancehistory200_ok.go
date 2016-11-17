package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// GetCorporationsCorporationIDAlliancehistory200Ok 200 ok object
// swagger:model get_corporations_corporation_id_alliancehistory_200_ok
type GetCorporationsCorporationIDAlliancehistory200Ok struct {

	// alliance
	Alliance *GetCorporationsCorporationIDAlliancehistoryAlliance `json:"alliance,omitempty"`

	// get_corporations_corporation_id_alliancehistory_record_id
	//
	// An incrementing ID that can be used to canonically establish order of records in cases where dates may be ambiguous
	// Required: true
	RecordID *int32 `json:"record_id"`

	// get_corporations_corporation_id_alliancehistory_start_date
	//
	// start_date string
	// Required: true
	StartDate *strfmt.DateTime `json:"start_date"`
}

// Validate validates this get corporations corporation id alliancehistory 200 ok
func (m *GetCorporationsCorporationIDAlliancehistory200Ok) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlliance(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRecordID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCorporationsCorporationIDAlliancehistory200Ok) validateAlliance(formats strfmt.Registry) error {

	if swag.IsZero(m.Alliance) { // not required
		return nil
	}

	if m.Alliance != nil {

		if err := m.Alliance.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *GetCorporationsCorporationIDAlliancehistory200Ok) validateRecordID(formats strfmt.Registry) error {

	if err := validate.Required("record_id", "body", m.RecordID); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDAlliancehistory200Ok) validateStartDate(formats strfmt.Registry) error {

	if err := validate.Required("start_date", "body", m.StartDate); err != nil {
		return err
	}

	return nil
}
