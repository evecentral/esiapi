package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// GetUniverseTypesTypeIDOk 200 ok object
// swagger:model get_universe_types_type_id_ok
type GetUniverseTypesTypeIDOk struct {

	// get_universe_types_type_id_category_id
	//
	// category_id integer
	// Required: true
	CategoryID *int32 `json:"category_id"`

	// get_universe_types_type_id_graphic_id
	//
	// graphic_id integer
	GraphicID int32 `json:"graphic_id,omitempty"`

	// get_universe_types_type_id_group_id
	//
	// group_id integer
	// Required: true
	GroupID *int32 `json:"group_id"`

	// get_universe_types_type_id_icon_id
	//
	// icon_id integer
	IconID int32 `json:"icon_id,omitempty"`

	// get_universe_types_type_id_type_description
	//
	// type_description string
	// Required: true
	TypeDescription *string `json:"type_description"`

	// get_universe_types_type_id_type_name
	//
	// type_name string
	// Required: true
	TypeName *string `json:"type_name"`
}

// Validate validates this get universe types type id ok
func (m *GetUniverseTypesTypeIDOk) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategoryID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGroupID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTypeDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTypeName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetUniverseTypesTypeIDOk) validateCategoryID(formats strfmt.Registry) error {

	if err := validate.Required("category_id", "body", m.CategoryID); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseTypesTypeIDOk) validateGroupID(formats strfmt.Registry) error {

	if err := validate.Required("group_id", "body", m.GroupID); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseTypesTypeIDOk) validateTypeDescription(formats strfmt.Registry) error {

	if err := validate.Required("type_description", "body", m.TypeDescription); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseTypesTypeIDOk) validateTypeName(formats strfmt.Registry) error {

	if err := validate.Required("type_name", "body", m.TypeName); err != nil {
		return err
	}

	return nil
}
