// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GetCharactersCharacterIDStatsOKBodyItemsMining get_characters_character_id_stats_mining
//
// mining object
// swagger:model getCharactersCharacterIdStatsOKBodyItemsMining
type GetCharactersCharacterIDStatsOKBodyItemsMining struct {

	// get_characters_character_id_stats_drone_mine
	//
	// drone_mine integer
	DroneMine int64 `json:"drone_mine,omitempty"`

	// get_characters_character_id_stats_ore_arkonor
	//
	// ore_arkonor integer
	OreArkonor int64 `json:"ore_arkonor,omitempty"`

	// get_characters_character_id_stats_ore_bistot
	//
	// ore_bistot integer
	OreBistot int64 `json:"ore_bistot,omitempty"`

	// get_characters_character_id_stats_ore_crokite
	//
	// ore_crokite integer
	OreCrokite int64 `json:"ore_crokite,omitempty"`

	// get_characters_character_id_stats_ore_dark_ochre
	//
	// ore_dark_ochre integer
	OreDarkOchre int64 `json:"ore_dark_ochre,omitempty"`

	// get_characters_character_id_stats_ore_gneiss
	//
	// ore_gneiss integer
	OreGneiss int64 `json:"ore_gneiss,omitempty"`

	// get_characters_character_id_stats_ore_harvestable_cloud
	//
	// ore_harvestable_cloud integer
	OreHarvestableCloud int64 `json:"ore_harvestable_cloud,omitempty"`

	// get_characters_character_id_stats_ore_hedbergite
	//
	// ore_hedbergite integer
	OreHedbergite int64 `json:"ore_hedbergite,omitempty"`

	// get_characters_character_id_stats_ore_hemorphite
	//
	// ore_hemorphite integer
	OreHemorphite int64 `json:"ore_hemorphite,omitempty"`

	// get_characters_character_id_stats_ore_ice
	//
	// ore_ice integer
	OreIce int64 `json:"ore_ice,omitempty"`

	// get_characters_character_id_stats_ore_jaspet
	//
	// ore_jaspet integer
	OreJaspet int64 `json:"ore_jaspet,omitempty"`

	// get_characters_character_id_stats_ore_kernite
	//
	// ore_kernite integer
	OreKernite int64 `json:"ore_kernite,omitempty"`

	// get_characters_character_id_stats_ore_mercoxit
	//
	// ore_mercoxit integer
	OreMercoxit int64 `json:"ore_mercoxit,omitempty"`

	// get_characters_character_id_stats_ore_omber
	//
	// ore_omber integer
	OreOmber int64 `json:"ore_omber,omitempty"`

	// get_characters_character_id_stats_ore_plagioclase
	//
	// ore_plagioclase integer
	OrePlagioclase int64 `json:"ore_plagioclase,omitempty"`

	// get_characters_character_id_stats_ore_pyroxeres
	//
	// ore_pyroxeres integer
	OrePyroxeres int64 `json:"ore_pyroxeres,omitempty"`

	// get_characters_character_id_stats_ore_scordite
	//
	// ore_scordite integer
	OreScordite int64 `json:"ore_scordite,omitempty"`

	// get_characters_character_id_stats_ore_spodumain
	//
	// ore_spodumain integer
	OreSpodumain int64 `json:"ore_spodumain,omitempty"`

	// get_characters_character_id_stats_ore_veldspar
	//
	// ore_veldspar integer
	OreVeldspar int64 `json:"ore_veldspar,omitempty"`
}

// Validate validates this get characters character Id stats o k body items mining
func (m *GetCharactersCharacterIDStatsOKBodyItemsMining) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDStatsOKBodyItemsMining) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDStatsOKBodyItemsMining) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDStatsOKBodyItemsMining
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
