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

// GetCharactersCharacterIDContactsLabelsOKBodyItems get_characters_character_id_contacts_labels_200_ok
//
// 200 ok object
// swagger:model getCharactersCharacterIdContactsLabelsOKBodyItems
type GetCharactersCharacterIDContactsLabelsOKBodyItems struct {

	// get_characters_character_id_contacts_labels_label_id
	//
	// label_id integer
	// Required: true
	LabelID *int64 `json:"label_id"`

	// get_characters_character_id_contacts_labels_label_name
	//
	// label_name string
	// Required: true
	LabelName *string `json:"label_name"`
}

// Validate validates this get characters character Id contacts labels o k body items
func (m *GetCharactersCharacterIDContactsLabelsOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabelID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCharactersCharacterIDContactsLabelsOKBodyItems) validateLabelID(formats strfmt.Registry) error {

	if err := validate.Required("label_id", "body", m.LabelID); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDContactsLabelsOKBodyItems) validateLabelName(formats strfmt.Registry) error {

	if err := validate.Required("label_name", "body", m.LabelName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDContactsLabelsOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDContactsLabelsOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDContactsLabelsOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
