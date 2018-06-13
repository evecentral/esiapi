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

// GetDogmaDynamicItemsTypeIDItemIDOKBodyDogmaAttributesItems get_dogma_dynamic_items_type_id_item_id_dogma_attribute
//
// dogma_attribute object
// swagger:model getDogmaDynamicItemsTypeIdItemIdOKBodyDogmaAttributesItems
type GetDogmaDynamicItemsTypeIDItemIDOKBodyDogmaAttributesItems struct {

	// get_dogma_dynamic_items_type_id_item_id_attribute_id
	//
	// attribute_id integer
	// Required: true
	AttributeID *int32 `json:"attribute_id"`

	// get_dogma_dynamic_items_type_id_item_id_value
	//
	// value number
	// Required: true
	Value *float32 `json:"value"`
}

// Validate validates this get dogma dynamic items type Id item Id o k body dogma attributes items
func (m *GetDogmaDynamicItemsTypeIDItemIDOKBodyDogmaAttributesItems) Validate(formats strfmt.Registry) error {
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

func (m *GetDogmaDynamicItemsTypeIDItemIDOKBodyDogmaAttributesItems) validateAttributeID(formats strfmt.Registry) error {

	if err := validate.Required("attribute_id", "body", m.AttributeID); err != nil {
		return err
	}

	return nil
}

func (m *GetDogmaDynamicItemsTypeIDItemIDOKBodyDogmaAttributesItems) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetDogmaDynamicItemsTypeIDItemIDOKBodyDogmaAttributesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetDogmaDynamicItemsTypeIDItemIDOKBodyDogmaAttributesItems) UnmarshalBinary(b []byte) error {
	var res GetDogmaDynamicItemsTypeIDItemIDOKBodyDogmaAttributesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
