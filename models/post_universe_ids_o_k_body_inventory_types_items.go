// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PostUniverseIdsOKBodyInventoryTypesItems post_universe_ids_inventory_type
//
// inventory_type object
// swagger:model postUniverseIdsOKBodyInventoryTypesItems
type PostUniverseIdsOKBodyInventoryTypesItems struct {

	// post_universe_ids_inventory_type_id
	//
	// id integer
	ID int32 `json:"id,omitempty"`

	// post_universe_ids_inventory_type_name
	//
	// name string
	Name string `json:"name,omitempty"`
}

// Validate validates this post universe ids o k body inventory types items
func (m *PostUniverseIdsOKBodyInventoryTypesItems) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostUniverseIdsOKBodyInventoryTypesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostUniverseIdsOKBodyInventoryTypesItems) UnmarshalBinary(b []byte) error {
	var res PostUniverseIdsOKBodyInventoryTypesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}