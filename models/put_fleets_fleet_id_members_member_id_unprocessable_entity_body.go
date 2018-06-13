// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PutFleetsFleetIDMembersMemberIDUnprocessableEntityBody put_fleets_fleet_id_members_member_id_unprocessable_entity
//
// 422 unprocessable entity object
// swagger:model putFleetsFleetIdMembersMemberIdUnprocessableEntityBody
type PutFleetsFleetIDMembersMemberIDUnprocessableEntityBody struct {

	// put_fleets_fleet_id_members_member_id_error
	//
	// error message
	Error string `json:"error,omitempty"`
}

// Validate validates this put fleets fleet Id members member Id unprocessable entity body
func (m *PutFleetsFleetIDMembersMemberIDUnprocessableEntityBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PutFleetsFleetIDMembersMemberIDUnprocessableEntityBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutFleetsFleetIDMembersMemberIDUnprocessableEntityBody) UnmarshalBinary(b []byte) error {
	var res PutFleetsFleetIDMembersMemberIDUnprocessableEntityBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
