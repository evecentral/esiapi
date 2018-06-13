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

// PostCharactersCharacterIDFittingsParamsBody post_characters_character_id_fittings_fitting
//
// fitting object
// swagger:model postCharactersCharacterIdFittingsParamsBody
type PostCharactersCharacterIDFittingsParamsBody struct {

	// post_characters_character_id_fittings_description
	//
	// description string
	// Required: true
	// Max Length: 500
	// Min Length: 0
	Description *string `json:"description"`

	// post_characters_character_id_fittings_items
	//
	// items array
	// Required: true
	// Max Items: 255
	// Min Items: 1
	Items []*PostCharactersCharacterIDFittingsParamsBodyItemsItems `json:"items"`

	// post_characters_character_id_fittings_name
	//
	// name string
	// Required: true
	// Max Length: 50
	// Min Length: 1
	Name *string `json:"name"`

	// post_characters_character_id_fittings_ship_type_id
	//
	// ship_type_id integer
	// Required: true
	ShipTypeID *int32 `json:"ship_type_id"`
}

// Validate validates this post characters character Id fittings params body
func (m *PostCharactersCharacterIDFittingsParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipTypeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostCharactersCharacterIDFittingsParamsBody) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 500); err != nil {
		return err
	}

	return nil
}

func (m *PostCharactersCharacterIDFittingsParamsBody) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("items", "body", m.Items); err != nil {
		return err
	}

	iItemsSize := int64(len(m.Items))

	if err := validate.MinItems("items", "body", iItemsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("items", "body", iItemsSize, 255); err != nil {
		return err
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostCharactersCharacterIDFittingsParamsBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 50); err != nil {
		return err
	}

	return nil
}

func (m *PostCharactersCharacterIDFittingsParamsBody) validateShipTypeID(formats strfmt.Registry) error {

	if err := validate.Required("ship_type_id", "body", m.ShipTypeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostCharactersCharacterIDFittingsParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostCharactersCharacterIDFittingsParamsBody) UnmarshalBinary(b []byte) error {
	var res PostCharactersCharacterIDFittingsParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}