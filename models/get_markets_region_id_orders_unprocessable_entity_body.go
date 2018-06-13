// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GetMarketsRegionIDOrdersUnprocessableEntityBody get_markets_region_id_orders_unprocessable_entity
//
// Unprocessable entity
// swagger:model getMarketsRegionIdOrdersUnprocessableEntityBody
type GetMarketsRegionIDOrdersUnprocessableEntityBody struct {

	// get_markets_region_id_orders_422_unprocessable_entity
	//
	// Unprocessable entity message
	Error string `json:"error,omitempty"`
}

// Validate validates this get markets region Id orders unprocessable entity body
func (m *GetMarketsRegionIDOrdersUnprocessableEntityBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetMarketsRegionIDOrdersUnprocessableEntityBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetMarketsRegionIDOrdersUnprocessableEntityBody) UnmarshalBinary(b []byte) error {
	var res GetMarketsRegionIDOrdersUnprocessableEntityBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
