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

// GetWarsWarIDOKBody get_wars_war_id_ok
//
// 200 ok object
// swagger:model getWarsWarIdOKBody
type GetWarsWarIDOKBody struct {

	// aggressor
	// Required: true
	Aggressor *GetWarsWarIDOKBodyAggressor `json:"aggressor"`

	// get_wars_war_id_allies
	//
	// allied corporations or alliances, each object contains either corporation_id or alliance_id
	// Max Items: 10000
	Allies []*GetWarsWarIDOKBodyAlliesItems `json:"allies"`

	// get_wars_war_id_declared
	//
	// Time that the war was declared
	// Required: true
	// Format: date-time
	Declared *strfmt.DateTime `json:"declared"`

	// defender
	// Required: true
	Defender *GetWarsWarIDOKBodyDefender `json:"defender"`

	// get_wars_war_id_finished
	//
	// Time the war ended and shooting was no longer allowed
	// Format: date-time
	Finished strfmt.DateTime `json:"finished,omitempty"`

	// get_wars_war_id_id
	//
	// ID of the specified war
	// Required: true
	ID *int32 `json:"id"`

	// get_wars_war_id_mutual
	//
	// Was the war declared mutual by both parties
	// Required: true
	Mutual *bool `json:"mutual"`

	// get_wars_war_id_open_for_allies
	//
	// Is the war currently open for allies or not
	// Required: true
	OpenForAllies *bool `json:"open_for_allies"`

	// get_wars_war_id_retracted
	//
	// Time the war was retracted but both sides could still shoot each other
	// Format: date-time
	Retracted strfmt.DateTime `json:"retracted,omitempty"`

	// get_wars_war_id_started
	//
	// Time when the war started and both sides could shoot each other
	// Format: date-time
	Started strfmt.DateTime `json:"started,omitempty"`
}

// Validate validates this get wars war Id o k body
func (m *GetWarsWarIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggressor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAllies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeclared(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefender(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFinished(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMutual(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenForAllies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRetracted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStarted(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetWarsWarIDOKBody) validateAggressor(formats strfmt.Registry) error {

	if err := validate.Required("aggressor", "body", m.Aggressor); err != nil {
		return err
	}

	if m.Aggressor != nil {
		if err := m.Aggressor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggressor")
			}
			return err
		}
	}

	return nil
}

func (m *GetWarsWarIDOKBody) validateAllies(formats strfmt.Registry) error {

	if swag.IsZero(m.Allies) { // not required
		return nil
	}

	iAlliesSize := int64(len(m.Allies))

	if err := validate.MaxItems("allies", "body", iAlliesSize, 10000); err != nil {
		return err
	}

	for i := 0; i < len(m.Allies); i++ {
		if swag.IsZero(m.Allies[i]) { // not required
			continue
		}

		if m.Allies[i] != nil {
			if err := m.Allies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetWarsWarIDOKBody) validateDeclared(formats strfmt.Registry) error {

	if err := validate.Required("declared", "body", m.Declared); err != nil {
		return err
	}

	if err := validate.FormatOf("declared", "body", "date-time", m.Declared.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetWarsWarIDOKBody) validateDefender(formats strfmt.Registry) error {

	if err := validate.Required("defender", "body", m.Defender); err != nil {
		return err
	}

	if m.Defender != nil {
		if err := m.Defender.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("defender")
			}
			return err
		}
	}

	return nil
}

func (m *GetWarsWarIDOKBody) validateFinished(formats strfmt.Registry) error {

	if swag.IsZero(m.Finished) { // not required
		return nil
	}

	if err := validate.FormatOf("finished", "body", "date-time", m.Finished.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetWarsWarIDOKBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *GetWarsWarIDOKBody) validateMutual(formats strfmt.Registry) error {

	if err := validate.Required("mutual", "body", m.Mutual); err != nil {
		return err
	}

	return nil
}

func (m *GetWarsWarIDOKBody) validateOpenForAllies(formats strfmt.Registry) error {

	if err := validate.Required("open_for_allies", "body", m.OpenForAllies); err != nil {
		return err
	}

	return nil
}

func (m *GetWarsWarIDOKBody) validateRetracted(formats strfmt.Registry) error {

	if swag.IsZero(m.Retracted) { // not required
		return nil
	}

	if err := validate.FormatOf("retracted", "body", "date-time", m.Retracted.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetWarsWarIDOKBody) validateStarted(formats strfmt.Registry) error {

	if swag.IsZero(m.Started) { // not required
		return nil
	}

	if err := validate.FormatOf("started", "body", "date-time", m.Started.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetWarsWarIDOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetWarsWarIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetWarsWarIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
