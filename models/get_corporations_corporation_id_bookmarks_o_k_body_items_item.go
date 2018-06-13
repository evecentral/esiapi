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

// GetCorporationsCorporationIDBookmarksOKBodyItemsItem get_corporations_corporation_id_bookmarks_item
//
// Optional object that is returned if a bookmark was made on a particular item.
// swagger:model getCorporationsCorporationIdBookmarksOKBodyItemsItem
type GetCorporationsCorporationIDBookmarksOKBodyItemsItem struct {

	// get_corporations_corporation_id_bookmarks_item_id
	//
	// item_id integer
	// Required: true
	ItemID *int64 `json:"item_id"`

	// get_corporations_corporation_id_bookmarks_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this get corporations corporation Id bookmarks o k body items item
func (m *GetCorporationsCorporationIDBookmarksOKBodyItemsItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItemID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCorporationsCorporationIDBookmarksOKBodyItemsItem) validateItemID(formats strfmt.Registry) error {

	if err := validate.Required("item_id", "body", m.ItemID); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDBookmarksOKBodyItemsItem) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", m.TypeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCorporationsCorporationIDBookmarksOKBodyItemsItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCorporationsCorporationIDBookmarksOKBodyItemsItem) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDBookmarksOKBodyItemsItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
