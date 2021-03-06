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

// GetKillmailsKillmailIDKillmailHashOKBodyVictimItemsItemsItemsItems get_killmails_killmail_id_killmail_hash_items_item
//
// item object
// swagger:model getKillmailsKillmailIdKillmailHashOKBodyVictimItemsItemsItemsItems
type GetKillmailsKillmailIDKillmailHashOKBodyVictimItemsItemsItemsItems struct {

	// get_killmails_killmail_id_killmail_hash_item_flag
	//
	// flag integer
	// Required: true
	Flag *int32 `json:"flag"`

	// get_killmails_killmail_id_killmail_hash_item_item_type_id
	//
	// item_type_id integer
	// Required: true
	ItemTypeID *int32 `json:"item_type_id"`

	// get_killmails_killmail_id_killmail_hash_item_quantity_destroyed
	//
	// quantity_destroyed integer
	QuantityDestroyed int64 `json:"quantity_destroyed,omitempty"`

	// get_killmails_killmail_id_killmail_hash_item_quantity_dropped
	//
	// quantity_dropped integer
	QuantityDropped int64 `json:"quantity_dropped,omitempty"`

	// get_killmails_killmail_id_killmail_hash_item_singleton
	//
	// singleton integer
	// Required: true
	Singleton *int32 `json:"singleton"`
}

// Validate validates this get killmails killmail Id killmail hash o k body victim items items items items
func (m *GetKillmailsKillmailIDKillmailHashOKBodyVictimItemsItemsItemsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemTypeID(formats); err != nil {
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

func (m *GetKillmailsKillmailIDKillmailHashOKBodyVictimItemsItemsItemsItems) validateFlag(formats strfmt.Registry) error {

	if err := validate.Required("flag", "body", m.Flag); err != nil {
		return err
	}

	return nil
}

func (m *GetKillmailsKillmailIDKillmailHashOKBodyVictimItemsItemsItemsItems) validateItemTypeID(formats strfmt.Registry) error {

	if err := validate.Required("item_type_id", "body", m.ItemTypeID); err != nil {
		return err
	}

	return nil
}

func (m *GetKillmailsKillmailIDKillmailHashOKBodyVictimItemsItemsItemsItems) validateSingleton(formats strfmt.Registry) error {

	if err := validate.Required("singleton", "body", m.Singleton); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetKillmailsKillmailIDKillmailHashOKBodyVictimItemsItemsItemsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetKillmailsKillmailIDKillmailHashOKBodyVictimItemsItemsItemsItems) UnmarshalBinary(b []byte) error {
	var res GetKillmailsKillmailIDKillmailHashOKBodyVictimItemsItemsItemsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
