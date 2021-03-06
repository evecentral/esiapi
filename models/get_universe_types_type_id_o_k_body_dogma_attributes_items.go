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

// GetUniverseTypesTypeIDOKBodyDogmaAttributesItems get_universe_types_type_id_dogma_attribute
//
// dogma_attribute object
// swagger:model getUniverseTypesTypeIdOKBodyDogmaAttributesItems
type GetUniverseTypesTypeIDOKBodyDogmaAttributesItems struct {

	// get_universe_types_type_id_attribute_id
	//
	// attribute_id integer
	// Required: true
	AttributeID *int32 `json:"attribute_id"`

	// get_universe_types_type_id_value
	//
	// value number
	// Required: true
	Value *float32 `json:"value"`
}

// Validate validates this get universe types type Id o k body dogma attributes items
func (m *GetUniverseTypesTypeIDOKBodyDogmaAttributesItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetUniverseTypesTypeIDOKBodyDogmaAttributesItems) validateAttributeID(formats strfmt.Registry) error {

	if err := validate.Required("attribute_id", "body", m.AttributeID); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseTypesTypeIDOKBodyDogmaAttributesItems) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetUniverseTypesTypeIDOKBodyDogmaAttributesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUniverseTypesTypeIDOKBodyDogmaAttributesItems) UnmarshalBinary(b []byte) error {
	var res GetUniverseTypesTypeIDOKBodyDogmaAttributesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
