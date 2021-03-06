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

// GetCorporationCorporationIDMiningExtractionsOKBodyItems get_corporation_corporation_id_mining_extractions_200_ok
//
// 200 ok object
// swagger:model getCorporationCorporationIdMiningExtractionsOKBodyItems
type GetCorporationCorporationIDMiningExtractionsOKBodyItems struct {

	// get_corporation_corporation_id_mining_extractions_chunk_arrival_time
	//
	// The time at which the chunk being extracted will arrive and can be fractured by the moon mining drill.
	//
	// Required: true
	// Format: date-time
	ChunkArrivalTime *strfmt.DateTime `json:"chunk_arrival_time"`

	// get_corporation_corporation_id_mining_extractions_extraction_start_time
	//
	// The time at which the current extraction was initiated.
	//
	// Required: true
	// Format: date-time
	ExtractionStartTime *strfmt.DateTime `json:"extraction_start_time"`

	// get_corporation_corporation_id_mining_extractions_moon_id
	//
	// moon_id integer
	// Required: true
	MoonID *int32 `json:"moon_id"`

	// get_corporation_corporation_id_mining_extractions_natural_decay_time
	//
	// The time at which the chunk being extracted will naturally fracture if it is not first fractured by the moon mining drill.
	//
	// Required: true
	// Format: date-time
	NaturalDecayTime *strfmt.DateTime `json:"natural_decay_time"`

	// get_corporation_corporation_id_mining_extractions_structure_id
	//
	// structure_id integer
	// Required: true
	StructureID *int64 `json:"structure_id"`
}

// Validate validates this get corporation corporation Id mining extractions o k body items
func (m *GetCorporationCorporationIDMiningExtractionsOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChunkArrivalTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtractionStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoonID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNaturalDecayTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStructureID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCorporationCorporationIDMiningExtractionsOKBodyItems) validateChunkArrivalTime(formats strfmt.Registry) error {

	if err := validate.Required("chunk_arrival_time", "body", m.ChunkArrivalTime); err != nil {
		return err
	}

	if err := validate.FormatOf("chunk_arrival_time", "body", "date-time", m.ChunkArrivalTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationCorporationIDMiningExtractionsOKBodyItems) validateExtractionStartTime(formats strfmt.Registry) error {

	if err := validate.Required("extraction_start_time", "body", m.ExtractionStartTime); err != nil {
		return err
	}

	if err := validate.FormatOf("extraction_start_time", "body", "date-time", m.ExtractionStartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationCorporationIDMiningExtractionsOKBodyItems) validateMoonID(formats strfmt.Registry) error {

	if err := validate.Required("moon_id", "body", m.MoonID); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationCorporationIDMiningExtractionsOKBodyItems) validateNaturalDecayTime(formats strfmt.Registry) error {

	if err := validate.Required("natural_decay_time", "body", m.NaturalDecayTime); err != nil {
		return err
	}

	if err := validate.FormatOf("natural_decay_time", "body", "date-time", m.NaturalDecayTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationCorporationIDMiningExtractionsOKBodyItems) validateStructureID(formats strfmt.Registry) error {

	if err := validate.Required("structure_id", "body", m.StructureID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCorporationCorporationIDMiningExtractionsOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCorporationCorporationIDMiningExtractionsOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetCorporationCorporationIDMiningExtractionsOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
