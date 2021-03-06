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

// GetLoyaltyStoresCorporationIDOffersOKBodyItems get_loyalty_stores_corporation_id_offers_200_ok
//
// 200 ok object
// swagger:model getLoyaltyStoresCorporationIdOffersOKBodyItems
type GetLoyaltyStoresCorporationIDOffersOKBodyItems struct {

	// get_loyalty_stores_corporation_id_offers_ak_cost
	//
	// Analysis kredit cost
	AkCost int32 `json:"ak_cost,omitempty"`

	// get_loyalty_stores_corporation_id_offers_isk_cost
	//
	// isk_cost integer
	// Required: true
	IskCost *int64 `json:"isk_cost"`

	// get_loyalty_stores_corporation_id_offers_lp_cost
	//
	// lp_cost integer
	// Required: true
	LpCost *int32 `json:"lp_cost"`

	// get_loyalty_stores_corporation_id_offers_offer_id
	//
	// offer_id integer
	// Required: true
	OfferID *int32 `json:"offer_id"`

	// get_loyalty_stores_corporation_id_offers_quantity
	//
	// quantity integer
	// Required: true
	Quantity *int32 `json:"quantity"`

	// get_loyalty_stores_corporation_id_offers_required_items
	//
	// required_items array
	// Required: true
	// Max Items: 100
	RequiredItems []*GetLoyaltyStoresCorporationIDOffersOKBodyItemsRequiredItemsItems `json:"required_items"`

	// get_loyalty_stores_corporation_id_offers_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this get loyalty stores corporation Id offers o k body items
func (m *GetLoyaltyStoresCorporationIDOffersOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIskCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLpCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOfferID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequiredItems(formats); err != nil {
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

func (m *GetLoyaltyStoresCorporationIDOffersOKBodyItems) validateIskCost(formats strfmt.Registry) error {

	if err := validate.Required("isk_cost", "body", m.IskCost); err != nil {
		return err
	}

	return nil
}

func (m *GetLoyaltyStoresCorporationIDOffersOKBodyItems) validateLpCost(formats strfmt.Registry) error {

	if err := validate.Required("lp_cost", "body", m.LpCost); err != nil {
		return err
	}

	return nil
}

func (m *GetLoyaltyStoresCorporationIDOffersOKBodyItems) validateOfferID(formats strfmt.Registry) error {

	if err := validate.Required("offer_id", "body", m.OfferID); err != nil {
		return err
	}

	return nil
}

func (m *GetLoyaltyStoresCorporationIDOffersOKBodyItems) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", m.Quantity); err != nil {
		return err
	}

	return nil
}

func (m *GetLoyaltyStoresCorporationIDOffersOKBodyItems) validateRequiredItems(formats strfmt.Registry) error {

	if err := validate.Required("required_items", "body", m.RequiredItems); err != nil {
		return err
	}

	iRequiredItemsSize := int64(len(m.RequiredItems))

	if err := validate.MaxItems("required_items", "body", iRequiredItemsSize, 100); err != nil {
		return err
	}

	for i := 0; i < len(m.RequiredItems); i++ {
		if swag.IsZero(m.RequiredItems[i]) { // not required
			continue
		}

		if m.RequiredItems[i] != nil {
			if err := m.RequiredItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("required_items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetLoyaltyStoresCorporationIDOffersOKBodyItems) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", m.TypeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetLoyaltyStoresCorporationIDOffersOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetLoyaltyStoresCorporationIDOffersOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetLoyaltyStoresCorporationIDOffersOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
