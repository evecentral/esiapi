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

// GetCorporationsCorporationIDOrdersHistoryOKBodyItems get_corporations_corporation_id_orders_history_200_ok
//
// 200 ok object
// swagger:model getCorporationsCorporationIdOrdersHistoryOKBodyItems
type GetCorporationsCorporationIDOrdersHistoryOKBodyItems struct {

	// get_corporations_corporation_id_orders_history_duration
	//
	// Number of days the order was valid for (starting from the issued date). An order expires at time issued + duration
	// Required: true
	Duration *int32 `json:"duration"`

	// get_corporations_corporation_id_orders_history_escrow
	//
	// For buy orders, the amount of ISK in escrow
	Escrow float64 `json:"escrow,omitempty"`

	// get_corporations_corporation_id_orders_history_is_buy_order
	//
	// True if the order is a bid (buy) order
	IsBuyOrder bool `json:"is_buy_order,omitempty"`

	// get_corporations_corporation_id_orders_history_issued
	//
	// Date and time when this order was issued
	// Required: true
	// Format: date-time
	Issued *strfmt.DateTime `json:"issued"`

	// get_corporations_corporation_id_orders_history_location_id
	//
	// ID of the location where order was placed
	// Required: true
	LocationID *int64 `json:"location_id"`

	// get_corporations_corporation_id_orders_history_min_volume
	//
	// For buy orders, the minimum quantity that will be accepted in a matching sell order
	MinVolume int32 `json:"min_volume,omitempty"`

	// get_corporations_corporation_id_orders_history_order_id
	//
	// Unique order ID
	// Required: true
	OrderID *int64 `json:"order_id"`

	// get_corporations_corporation_id_orders_history_price
	//
	// Cost per unit for this order
	// Required: true
	Price *float64 `json:"price"`

	// get_corporations_corporation_id_orders_history_range
	//
	// Valid order range, numbers are ranges in jumps
	// Required: true
	// Enum: [1 10 2 20 3 30 4 40 5 region solarsystem station]
	Range *string `json:"range"`

	// get_corporations_corporation_id_orders_history_region_id
	//
	// ID of the region where order was placed
	// Required: true
	RegionID *int32 `json:"region_id"`

	// get_corporations_corporation_id_orders_history_state
	//
	// Current order state
	// Required: true
	// Enum: [cancelled expired]
	State *string `json:"state"`

	// get_corporations_corporation_id_orders_history_type_id
	//
	// The type ID of the item transacted in this order
	// Required: true
	TypeID *int32 `json:"type_id"`

	// get_corporations_corporation_id_orders_history_volume_remain
	//
	// Quantity of items still required or offered
	// Required: true
	VolumeRemain *int32 `json:"volume_remain"`

	// get_corporations_corporation_id_orders_history_volume_total
	//
	// Quantity of items required or offered at time order was placed
	// Required: true
	VolumeTotal *int32 `json:"volume_total"`

	// get_corporations_corporation_id_orders_history_wallet_division
	//
	// The corporation wallet division used for this order
	// Required: true
	// Maximum: 7
	// Minimum: 1
	WalletDivision *int32 `json:"wallet_division"`
}

// Validate validates this get corporations corporation Id orders history o k body items
func (m *GetCorporationsCorporationIDOrdersHistoryOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssued(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeRemain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeTotal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWalletDivision(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCorporationsCorporationIDOrdersHistoryOKBodyItems) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDOrdersHistoryOKBodyItems) validateIssued(formats strfmt.Registry) error {

	if err := validate.Required("issued", "body", m.Issued); err != nil {
		return err
	}

	if err := validate.FormatOf("issued", "body", "date-time", m.Issued.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDOrdersHistoryOKBodyItems) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("location_id", "body", m.LocationID); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDOrdersHistoryOKBodyItems) validateOrderID(formats strfmt.Registry) error {

	if err := validate.Required("order_id", "body", m.OrderID); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDOrdersHistoryOKBodyItems) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", m.Price); err != nil {
		return err
	}

	return nil
}

var getCorporationsCorporationIdOrdersHistoryOKBodyItemsTypeRangePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["1","10","2","20","3","30","4","40","5","region","solarsystem","station"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdOrdersHistoryOKBodyItemsTypeRangePropEnum = append(getCorporationsCorporationIdOrdersHistoryOKBodyItemsTypeRangePropEnum, v)
	}
}

const (

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItemsRangeNr1 captures enum value "1"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItemsRangeNr1 string = "1"

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItemsRangeNr10 captures enum value "10"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItemsRangeNr10 string = "10"

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItemsRangeNr2 captures enum value "2"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItemsRangeNr2 string = "2"

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItemsRangeNr20 captures enum value "20"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItemsRangeNr20 string = "20"

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItemsRangeNr3 captures enum value "3"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItemsRangeNr3 string = "3"

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItemsRangeNr30 captures enum value "30"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItemsRangeNr30 string = "30"

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItemsRangeNr4 captures enum value "4"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItemsRangeNr4 string = "4"

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItemsRangeNr40 captures enum value "40"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItemsRangeNr40 string = "40"

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItemsRangeNr5 captures enum value "5"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItemsRangeNr5 string = "5"

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItemsRangeRegion captures enum value "region"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItemsRangeRegion string = "region"

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItemsRangeSolarsystem captures enum value "solarsystem"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItemsRangeSolarsystem string = "solarsystem"

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItemsRangeStation captures enum value "station"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItemsRangeStation string = "station"
)

// prop value enum
func (m *GetCorporationsCorporationIDOrdersHistoryOKBodyItems) validateRangeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCorporationsCorporationIdOrdersHistoryOKBodyItemsTypeRangePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCorporationsCorporationIDOrdersHistoryOKBodyItems) validateRange(formats strfmt.Registry) error {

	if err := validate.Required("range", "body", m.Range); err != nil {
		return err
	}

	// value enum
	if err := m.validateRangeEnum("range", "body", *m.Range); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDOrdersHistoryOKBodyItems) validateRegionID(formats strfmt.Registry) error {

	if err := validate.Required("region_id", "body", m.RegionID); err != nil {
		return err
	}

	return nil
}

var getCorporationsCorporationIdOrdersHistoryOKBodyItemsTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cancelled","expired"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdOrdersHistoryOKBodyItemsTypeStatePropEnum = append(getCorporationsCorporationIdOrdersHistoryOKBodyItemsTypeStatePropEnum, v)
	}
}

const (

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItemsStateCancelled captures enum value "cancelled"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItemsStateCancelled string = "cancelled"

	// GetCorporationsCorporationIDOrdersHistoryOKBodyItemsStateExpired captures enum value "expired"
	GetCorporationsCorporationIDOrdersHistoryOKBodyItemsStateExpired string = "expired"
)

// prop value enum
func (m *GetCorporationsCorporationIDOrdersHistoryOKBodyItems) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCorporationsCorporationIdOrdersHistoryOKBodyItemsTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCorporationsCorporationIDOrdersHistoryOKBodyItems) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDOrdersHistoryOKBodyItems) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", m.TypeID); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDOrdersHistoryOKBodyItems) validateVolumeRemain(formats strfmt.Registry) error {

	if err := validate.Required("volume_remain", "body", m.VolumeRemain); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDOrdersHistoryOKBodyItems) validateVolumeTotal(formats strfmt.Registry) error {

	if err := validate.Required("volume_total", "body", m.VolumeTotal); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDOrdersHistoryOKBodyItems) validateWalletDivision(formats strfmt.Registry) error {

	if err := validate.Required("wallet_division", "body", m.WalletDivision); err != nil {
		return err
	}

	if err := validate.MinimumInt("wallet_division", "body", int64(*m.WalletDivision), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("wallet_division", "body", int64(*m.WalletDivision), 7, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCorporationsCorporationIDOrdersHistoryOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCorporationsCorporationIDOrdersHistoryOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDOrdersHistoryOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
