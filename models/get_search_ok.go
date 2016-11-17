package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// GetSearchOk 200 ok object
// swagger:model get_search_ok
type GetSearchOk struct {

	// get_search_agent
	//
	// agent array
	Agent []int64 `json:"agent"`

	// get_search_alliance
	//
	// alliance array
	Alliance []int64 `json:"alliance"`

	// get_search_character
	//
	// character array
	Character []int64 `json:"character"`

	// get_search_constellation
	//
	// constellation array
	Constellation []int64 `json:"constellation"`

	// get_search_corporation
	//
	// corporation array
	Corporation []int64 `json:"corporation"`

	// get_search_faction
	//
	// faction array
	Faction []int64 `json:"faction"`

	// get_search_inventorytype
	//
	// inventorytype array
	Inventorytype []int64 `json:"inventorytype"`

	// get_search_region
	//
	// region array
	Region []int64 `json:"region"`

	// get_search_solarsystem
	//
	// solarsystem array
	Solarsystem []int64 `json:"solarsystem"`

	// get_search_station
	//
	// station array
	Station []int64 `json:"station"`

	// get_search_wormhole
	//
	// wormhole array
	Wormhole []int64 `json:"wormhole"`
}

// Validate validates this get search ok
func (m *GetSearchOk) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAlliance(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCharacter(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConstellation(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCorporation(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFaction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInventorytype(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSolarsystem(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStation(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWormhole(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSearchOk) validateAgent(formats strfmt.Registry) error {

	if swag.IsZero(m.Agent) { // not required
		return nil
	}

	return nil
}

func (m *GetSearchOk) validateAlliance(formats strfmt.Registry) error {

	if swag.IsZero(m.Alliance) { // not required
		return nil
	}

	return nil
}

func (m *GetSearchOk) validateCharacter(formats strfmt.Registry) error {

	if swag.IsZero(m.Character) { // not required
		return nil
	}

	return nil
}

func (m *GetSearchOk) validateConstellation(formats strfmt.Registry) error {

	if swag.IsZero(m.Constellation) { // not required
		return nil
	}

	return nil
}

func (m *GetSearchOk) validateCorporation(formats strfmt.Registry) error {

	if swag.IsZero(m.Corporation) { // not required
		return nil
	}

	return nil
}

func (m *GetSearchOk) validateFaction(formats strfmt.Registry) error {

	if swag.IsZero(m.Faction) { // not required
		return nil
	}

	return nil
}

func (m *GetSearchOk) validateInventorytype(formats strfmt.Registry) error {

	if swag.IsZero(m.Inventorytype) { // not required
		return nil
	}

	return nil
}

func (m *GetSearchOk) validateRegion(formats strfmt.Registry) error {

	if swag.IsZero(m.Region) { // not required
		return nil
	}

	return nil
}

func (m *GetSearchOk) validateSolarsystem(formats strfmt.Registry) error {

	if swag.IsZero(m.Solarsystem) { // not required
		return nil
	}

	return nil
}

func (m *GetSearchOk) validateStation(formats strfmt.Registry) error {

	if swag.IsZero(m.Station) { // not required
		return nil
	}

	return nil
}

func (m *GetSearchOk) validateWormhole(formats strfmt.Registry) error {

	if swag.IsZero(m.Wormhole) { // not required
		return nil
	}

	return nil
}
