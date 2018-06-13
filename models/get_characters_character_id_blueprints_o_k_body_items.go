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

// GetCharactersCharacterIDBlueprintsOKBodyItems get_characters_character_id_blueprints_200_ok
//
// 200 ok object
// swagger:model getCharactersCharacterIdBlueprintsOKBodyItems
type GetCharactersCharacterIDBlueprintsOKBodyItems struct {

	// get_characters_character_id_blueprints_item_id
	//
	// Unique ID for this item.
	// Required: true
	ItemID *int64 `json:"item_id"`

	// get_characters_character_id_blueprints_location_flag
	//
	// Type of the location_id
	// Required: true
	// Enum: [AutoFit Cargo CorpseBay DroneBay FleetHangar Deliveries HiddenModifiers Hangar HangarAll LoSlot0 LoSlot1 LoSlot2 LoSlot3 LoSlot4 LoSlot5 LoSlot6 LoSlot7 MedSlot0 MedSlot1 MedSlot2 MedSlot3 MedSlot4 MedSlot5 MedSlot6 MedSlot7 HiSlot0 HiSlot1 HiSlot2 HiSlot3 HiSlot4 HiSlot5 HiSlot6 HiSlot7 AssetSafety Locked Unlocked Implant QuafeBay RigSlot0 RigSlot1 RigSlot2 RigSlot3 RigSlot4 RigSlot5 RigSlot6 RigSlot7 ShipHangar SpecializedFuelBay SpecializedOreHold SpecializedGasHold SpecializedMineralHold SpecializedSalvageHold SpecializedShipHold SpecializedSmallShipHold SpecializedMediumShipHold SpecializedLargeShipHold SpecializedIndustrialShipHold SpecializedAmmoHold SpecializedCommandCenterHold SpecializedPlanetaryCommoditiesHold SpecializedMaterialBay SubSystemSlot0 SubSystemSlot1 SubSystemSlot2 SubSystemSlot3 SubSystemSlot4 SubSystemSlot5 SubSystemSlot6 SubSystemSlot7 FighterBay FighterTube0 FighterTube1 FighterTube2 FighterTube3 FighterTube4 Module]
	LocationFlag *string `json:"location_flag"`

	// get_characters_character_id_blueprints_location_id
	//
	// References a solar system, station or item_id if this blueprint is located within a container. If the return value is an item_id, then the Character AssetList API must be queried to find the container using the given item_id to determine the correct location of the Blueprint.
	// Required: true
	LocationID *int64 `json:"location_id"`

	// get_characters_character_id_blueprints_material_efficiency
	//
	// Material Efficiency Level of the blueprint.
	// Required: true
	// Maximum: 25
	// Minimum: 0
	MaterialEfficiency *int32 `json:"material_efficiency"`

	// get_characters_character_id_blueprints_quantity
	//
	// A range of numbers with a minimum of -2 and no maximum value where -1 is an original and -2 is a copy. It can be a positive integer if it is a stack of blueprint originals fresh from the market (e.g. no activities performed on them yet).
	// Required: true
	// Minimum: -2
	Quantity *int32 `json:"quantity"`

	// get_characters_character_id_blueprints_runs
	//
	// Number of runs remaining if the blueprint is a copy, -1 if it is an original.
	// Required: true
	// Minimum: -1
	Runs *int32 `json:"runs"`

	// get_characters_character_id_blueprints_time_efficiency
	//
	// Time Efficiency Level of the blueprint.
	// Required: true
	// Maximum: 20
	// Minimum: 0
	TimeEfficiency *int32 `json:"time_efficiency"`

	// get_characters_character_id_blueprints_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this get characters character Id blueprints o k body items
func (m *GetCharactersCharacterIDBlueprintsOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItemID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaterialEfficiency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeEfficiency(formats); err != nil {
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

func (m *GetCharactersCharacterIDBlueprintsOKBodyItems) validateItemID(formats strfmt.Registry) error {

	if err := validate.Required("item_id", "body", m.ItemID); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdBlueprintsOKBodyItemsTypeLocationFlagPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AutoFit","Cargo","CorpseBay","DroneBay","FleetHangar","Deliveries","HiddenModifiers","Hangar","HangarAll","LoSlot0","LoSlot1","LoSlot2","LoSlot3","LoSlot4","LoSlot5","LoSlot6","LoSlot7","MedSlot0","MedSlot1","MedSlot2","MedSlot3","MedSlot4","MedSlot5","MedSlot6","MedSlot7","HiSlot0","HiSlot1","HiSlot2","HiSlot3","HiSlot4","HiSlot5","HiSlot6","HiSlot7","AssetSafety","Locked","Unlocked","Implant","QuafeBay","RigSlot0","RigSlot1","RigSlot2","RigSlot3","RigSlot4","RigSlot5","RigSlot6","RigSlot7","ShipHangar","SpecializedFuelBay","SpecializedOreHold","SpecializedGasHold","SpecializedMineralHold","SpecializedSalvageHold","SpecializedShipHold","SpecializedSmallShipHold","SpecializedMediumShipHold","SpecializedLargeShipHold","SpecializedIndustrialShipHold","SpecializedAmmoHold","SpecializedCommandCenterHold","SpecializedPlanetaryCommoditiesHold","SpecializedMaterialBay","SubSystemSlot0","SubSystemSlot1","SubSystemSlot2","SubSystemSlot3","SubSystemSlot4","SubSystemSlot5","SubSystemSlot6","SubSystemSlot7","FighterBay","FighterTube0","FighterTube1","FighterTube2","FighterTube3","FighterTube4","Module"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdBlueprintsOKBodyItemsTypeLocationFlagPropEnum = append(getCharactersCharacterIdBlueprintsOKBodyItemsTypeLocationFlagPropEnum, v)
	}
}

const (

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagAutoFit captures enum value "AutoFit"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagAutoFit string = "AutoFit"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagCargo captures enum value "Cargo"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagCargo string = "Cargo"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagCorpseBay captures enum value "CorpseBay"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagCorpseBay string = "CorpseBay"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagDroneBay captures enum value "DroneBay"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagDroneBay string = "DroneBay"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagFleetHangar captures enum value "FleetHangar"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagFleetHangar string = "FleetHangar"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagDeliveries captures enum value "Deliveries"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagDeliveries string = "Deliveries"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagHiddenModifiers captures enum value "HiddenModifiers"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagHiddenModifiers string = "HiddenModifiers"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagHangar captures enum value "Hangar"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagHangar string = "Hangar"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagHangarAll captures enum value "HangarAll"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagHangarAll string = "HangarAll"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagLoSlot0 captures enum value "LoSlot0"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagLoSlot0 string = "LoSlot0"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagLoSlot1 captures enum value "LoSlot1"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagLoSlot1 string = "LoSlot1"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagLoSlot2 captures enum value "LoSlot2"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagLoSlot2 string = "LoSlot2"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagLoSlot3 captures enum value "LoSlot3"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagLoSlot3 string = "LoSlot3"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagLoSlot4 captures enum value "LoSlot4"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagLoSlot4 string = "LoSlot4"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagLoSlot5 captures enum value "LoSlot5"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagLoSlot5 string = "LoSlot5"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagLoSlot6 captures enum value "LoSlot6"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagLoSlot6 string = "LoSlot6"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagLoSlot7 captures enum value "LoSlot7"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagLoSlot7 string = "LoSlot7"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagMedSlot0 captures enum value "MedSlot0"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagMedSlot0 string = "MedSlot0"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagMedSlot1 captures enum value "MedSlot1"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagMedSlot1 string = "MedSlot1"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagMedSlot2 captures enum value "MedSlot2"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagMedSlot2 string = "MedSlot2"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagMedSlot3 captures enum value "MedSlot3"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagMedSlot3 string = "MedSlot3"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagMedSlot4 captures enum value "MedSlot4"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagMedSlot4 string = "MedSlot4"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagMedSlot5 captures enum value "MedSlot5"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagMedSlot5 string = "MedSlot5"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagMedSlot6 captures enum value "MedSlot6"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagMedSlot6 string = "MedSlot6"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagMedSlot7 captures enum value "MedSlot7"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagMedSlot7 string = "MedSlot7"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagHiSlot0 captures enum value "HiSlot0"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagHiSlot0 string = "HiSlot0"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagHiSlot1 captures enum value "HiSlot1"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagHiSlot1 string = "HiSlot1"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagHiSlot2 captures enum value "HiSlot2"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagHiSlot2 string = "HiSlot2"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagHiSlot3 captures enum value "HiSlot3"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagHiSlot3 string = "HiSlot3"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagHiSlot4 captures enum value "HiSlot4"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagHiSlot4 string = "HiSlot4"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagHiSlot5 captures enum value "HiSlot5"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagHiSlot5 string = "HiSlot5"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagHiSlot6 captures enum value "HiSlot6"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagHiSlot6 string = "HiSlot6"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagHiSlot7 captures enum value "HiSlot7"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagHiSlot7 string = "HiSlot7"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagAssetSafety captures enum value "AssetSafety"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagAssetSafety string = "AssetSafety"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagLocked captures enum value "Locked"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagLocked string = "Locked"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagUnlocked captures enum value "Unlocked"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagUnlocked string = "Unlocked"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagImplant captures enum value "Implant"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagImplant string = "Implant"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagQuafeBay captures enum value "QuafeBay"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagQuafeBay string = "QuafeBay"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagRigSlot0 captures enum value "RigSlot0"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagRigSlot0 string = "RigSlot0"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagRigSlot1 captures enum value "RigSlot1"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagRigSlot1 string = "RigSlot1"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagRigSlot2 captures enum value "RigSlot2"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagRigSlot2 string = "RigSlot2"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagRigSlot3 captures enum value "RigSlot3"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagRigSlot3 string = "RigSlot3"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagRigSlot4 captures enum value "RigSlot4"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagRigSlot4 string = "RigSlot4"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagRigSlot5 captures enum value "RigSlot5"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagRigSlot5 string = "RigSlot5"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagRigSlot6 captures enum value "RigSlot6"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagRigSlot6 string = "RigSlot6"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagRigSlot7 captures enum value "RigSlot7"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagRigSlot7 string = "RigSlot7"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagShipHangar captures enum value "ShipHangar"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagShipHangar string = "ShipHangar"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedFuelBay captures enum value "SpecializedFuelBay"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedFuelBay string = "SpecializedFuelBay"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedOreHold captures enum value "SpecializedOreHold"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedOreHold string = "SpecializedOreHold"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedGasHold captures enum value "SpecializedGasHold"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedGasHold string = "SpecializedGasHold"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedMineralHold captures enum value "SpecializedMineralHold"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedMineralHold string = "SpecializedMineralHold"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedSalvageHold captures enum value "SpecializedSalvageHold"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedSalvageHold string = "SpecializedSalvageHold"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedShipHold captures enum value "SpecializedShipHold"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedShipHold string = "SpecializedShipHold"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedSmallShipHold captures enum value "SpecializedSmallShipHold"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedSmallShipHold string = "SpecializedSmallShipHold"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedMediumShipHold captures enum value "SpecializedMediumShipHold"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedMediumShipHold string = "SpecializedMediumShipHold"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedLargeShipHold captures enum value "SpecializedLargeShipHold"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedLargeShipHold string = "SpecializedLargeShipHold"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedIndustrialShipHold captures enum value "SpecializedIndustrialShipHold"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedIndustrialShipHold string = "SpecializedIndustrialShipHold"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedAmmoHold captures enum value "SpecializedAmmoHold"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedAmmoHold string = "SpecializedAmmoHold"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedCommandCenterHold captures enum value "SpecializedCommandCenterHold"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedCommandCenterHold string = "SpecializedCommandCenterHold"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedPlanetaryCommoditiesHold captures enum value "SpecializedPlanetaryCommoditiesHold"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedPlanetaryCommoditiesHold string = "SpecializedPlanetaryCommoditiesHold"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedMaterialBay captures enum value "SpecializedMaterialBay"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSpecializedMaterialBay string = "SpecializedMaterialBay"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot0 captures enum value "SubSystemSlot0"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot0 string = "SubSystemSlot0"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot1 captures enum value "SubSystemSlot1"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot1 string = "SubSystemSlot1"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot2 captures enum value "SubSystemSlot2"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot2 string = "SubSystemSlot2"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot3 captures enum value "SubSystemSlot3"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot3 string = "SubSystemSlot3"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot4 captures enum value "SubSystemSlot4"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot4 string = "SubSystemSlot4"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot5 captures enum value "SubSystemSlot5"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot5 string = "SubSystemSlot5"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot6 captures enum value "SubSystemSlot6"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot6 string = "SubSystemSlot6"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot7 captures enum value "SubSystemSlot7"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagSubSystemSlot7 string = "SubSystemSlot7"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagFighterBay captures enum value "FighterBay"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagFighterBay string = "FighterBay"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagFighterTube0 captures enum value "FighterTube0"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagFighterTube0 string = "FighterTube0"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagFighterTube1 captures enum value "FighterTube1"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagFighterTube1 string = "FighterTube1"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagFighterTube2 captures enum value "FighterTube2"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagFighterTube2 string = "FighterTube2"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagFighterTube3 captures enum value "FighterTube3"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagFighterTube3 string = "FighterTube3"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagFighterTube4 captures enum value "FighterTube4"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagFighterTube4 string = "FighterTube4"

	// GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagModule captures enum value "Module"
	GetCharactersCharacterIDBlueprintsOKBodyItemsLocationFlagModule string = "Module"
)

// prop value enum
func (m *GetCharactersCharacterIDBlueprintsOKBodyItems) validateLocationFlagEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCharactersCharacterIdBlueprintsOKBodyItemsTypeLocationFlagPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCharactersCharacterIDBlueprintsOKBodyItems) validateLocationFlag(formats strfmt.Registry) error {

	if err := validate.Required("location_flag", "body", m.LocationFlag); err != nil {
		return err
	}

	// value enum
	if err := m.validateLocationFlagEnum("location_flag", "body", *m.LocationFlag); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDBlueprintsOKBodyItems) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("location_id", "body", m.LocationID); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDBlueprintsOKBodyItems) validateMaterialEfficiency(formats strfmt.Registry) error {

	if err := validate.Required("material_efficiency", "body", m.MaterialEfficiency); err != nil {
		return err
	}

	if err := validate.MinimumInt("material_efficiency", "body", int64(*m.MaterialEfficiency), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("material_efficiency", "body", int64(*m.MaterialEfficiency), 25, false); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDBlueprintsOKBodyItems) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", m.Quantity); err != nil {
		return err
	}

	if err := validate.MinimumInt("quantity", "body", int64(*m.Quantity), -2, false); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDBlueprintsOKBodyItems) validateRuns(formats strfmt.Registry) error {

	if err := validate.Required("runs", "body", m.Runs); err != nil {
		return err
	}

	if err := validate.MinimumInt("runs", "body", int64(*m.Runs), -1, false); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDBlueprintsOKBodyItems) validateTimeEfficiency(formats strfmt.Registry) error {

	if err := validate.Required("time_efficiency", "body", m.TimeEfficiency); err != nil {
		return err
	}

	if err := validate.MinimumInt("time_efficiency", "body", int64(*m.TimeEfficiency), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("time_efficiency", "body", int64(*m.TimeEfficiency), 20, false); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDBlueprintsOKBodyItems) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", m.TypeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDBlueprintsOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDBlueprintsOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDBlueprintsOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
