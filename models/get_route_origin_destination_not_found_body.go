// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GetRouteOriginDestinationNotFoundBody get_route_origin_destination_not_found
//
// Not found
// swagger:model getRouteOriginDestinationNotFoundBody
type GetRouteOriginDestinationNotFoundBody struct {

	// get_route_origin_destination_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get route origin destination not found body
func (m *GetRouteOriginDestinationNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetRouteOriginDestinationNotFoundBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetRouteOriginDestinationNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetRouteOriginDestinationNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
