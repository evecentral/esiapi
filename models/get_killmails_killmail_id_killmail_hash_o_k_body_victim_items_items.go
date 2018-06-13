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

// GetKillmailsKillmailIDKillmailHashOKBodyVictimItemsItems get_killmails_killmail_id_killmail_hash_item
//
// item object
// swagger:model getKillmailsKillmailIdKillmailHashOKBodyVictimItemsItems
type GetKillmailsKillmailIDKillmailHashOKBodyVictimItemsItems struct {

	// get_killmails_killmail_id_killmail_hash_flag
	//
	// Flag for the location of the item
	//
	// Required: true
	Flag *int32 `json:"flag"`

	// get_killmails_killmail_id_killmail_hash_item_type_id
	//
	// item_type_id integer
	// Required: true
	ItemTypeID *int32 `json:"item_type_id"`

	// get_killmails_killmail_id_killmail_hash_item_items
	//
	// items array
	// Max Items: 10000
	Items []*GetKillmailsKillmailIDKillmailHashOKBodyVictimItemsItemsItemsItems `json:"items"`

	// get_killmails_killmail_id_killmail_hash_quantity_destroyed
	//
	// How many of the item were destroyed if any
	//
	QuantityDestroyed int64 `json:"quantity_destroyed,omitempty"`

	// get_killmails_killmail_id_killmail_hash_quantity_dropped
	//
	// How many of the item were dropped if any
	//
	QuantityDropped int64 `json:"quantity_dropped,omitempty"`

	// get_killmails_killmail_id_killmail_hash_singleton
	//
	// singleton integer
	// Required: true
	Singleton *int32 `json:"singleton"`
}

// Validate validates this get killmails killmail Id killmail hash o k body victim items items
func (m *GetKillmailsKillmailIDKillmailHashOKBodyVictimItemsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemTypeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSingleton(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetKillmailsKillmailIDKillmailHashOKBodyVictimItemsItems) validateFlag(formats strfmt.Registry) error {

	if err := validate.Required("flag", "body", m.Flag); err != nil {
		return err
	}

	return nil
}

func (m *GetKillmailsKillmailIDKillmailHashOKBodyVictimItemsItems) validateItemTypeID(formats strfmt.Registry) error {

	if err := validate.Required("item_type_id", "body", m.ItemTypeID); err != nil {
		return err
	}

	return nil
}

func (m *GetKillmailsKillmailIDKillmailHashOKBodyVictimItemsItems) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(m.Items) { // not required
		return nil
	}

	iItemsSize := int64(len(m.Items))

	if err := validate.MaxItems("items", "body", iItemsSize, 10000); err != nil {
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

func (m *GetKillmailsKillmailIDKillmailHashOKBodyVictimItemsItems) validateSingleton(formats strfmt.Registry) error {

	if err := validate.Required("singleton", "body", m.Singleton); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetKillmailsKillmailIDKillmailHashOKBodyVictimItemsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetKillmailsKillmailIDKillmailHashOKBodyVictimItemsItems) UnmarshalBinary(b []byte) error {
	var res GetKillmailsKillmailIDKillmailHashOKBodyVictimItemsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
