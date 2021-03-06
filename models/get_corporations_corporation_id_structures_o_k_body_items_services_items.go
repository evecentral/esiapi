// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetCorporationsCorporationIDStructuresOKBodyItemsServicesItems get_corporations_corporation_id_structures_service
//
// service object
// swagger:model getCorporationsCorporationIdStructuresOKBodyItemsServicesItems
type GetCorporationsCorporationIDStructuresOKBodyItemsServicesItems struct {

	// get_corporations_corporation_id_structures_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// get_corporations_corporation_id_structures_service_state
	//
	// state string
	// Required: true
	// Enum: [online offline cleanup]
	State *string `json:"state"`
}

// Validate validates this get corporations corporation Id structures o k body items services items
func (m *GetCorporationsCorporationIDStructuresOKBodyItemsServicesItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCorporationsCorporationIDStructuresOKBodyItemsServicesItems) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var getCorporationsCorporationIdStructuresOKBodyItemsServicesItemsTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["online","offline","cleanup"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdStructuresOKBodyItemsServicesItemsTypeStatePropEnum = append(getCorporationsCorporationIdStructuresOKBodyItemsServicesItemsTypeStatePropEnum, v)
	}
}

const (

	// GetCorporationsCorporationIDStructuresOKBodyItemsServicesItemsStateOnline captures enum value "online"
	GetCorporationsCorporationIDStructuresOKBodyItemsServicesItemsStateOnline string = "online"

	// GetCorporationsCorporationIDStructuresOKBodyItemsServicesItemsStateOffline captures enum value "offline"
	GetCorporationsCorporationIDStructuresOKBodyItemsServicesItemsStateOffline string = "offline"

	// GetCorporationsCorporationIDStructuresOKBodyItemsServicesItemsStateCleanup captures enum value "cleanup"
	GetCorporationsCorporationIDStructuresOKBodyItemsServicesItemsStateCleanup string = "cleanup"
)

// prop value enum
func (m *GetCorporationsCorporationIDStructuresOKBodyItemsServicesItems) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCorporationsCorporationIdStructuresOKBodyItemsServicesItemsTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCorporationsCorporationIDStructuresOKBodyItemsServicesItems) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCorporationsCorporationIDStructuresOKBodyItemsServicesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCorporationsCorporationIDStructuresOKBodyItemsServicesItems) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDStructuresOKBodyItemsServicesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
