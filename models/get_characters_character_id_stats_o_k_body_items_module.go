// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GetCharactersCharacterIDStatsOKBodyItemsModule get_characters_character_id_stats_module
//
// module object
// swagger:model getCharactersCharacterIdStatsOKBodyItemsModule
type GetCharactersCharacterIDStatsOKBodyItemsModule struct {

	// get_characters_character_id_stats_activations_armor_hardener
	//
	// activations_armor_hardener integer
	ActivationsArmorHardener int64 `json:"activations_armor_hardener,omitempty"`

	// get_characters_character_id_stats_activations_armor_repair_unit
	//
	// activations_armor_repair_unit integer
	ActivationsArmorRepairUnit int64 `json:"activations_armor_repair_unit,omitempty"`

	// get_characters_character_id_stats_activations_armor_resistance_shift_hardener
	//
	// activations_armor_resistance_shift_hardener integer
	ActivationsArmorResistanceShiftHardener int64 `json:"activations_armor_resistance_shift_hardener,omitempty"`

	// get_characters_character_id_stats_activations_automated_targeting_system
	//
	// activations_automated_targeting_system integer
	ActivationsAutomatedTargetingSystem int64 `json:"activations_automated_targeting_system,omitempty"`

	// get_characters_character_id_stats_activations_bastion
	//
	// activations_bastion integer
	ActivationsBastion int64 `json:"activations_bastion,omitempty"`

	// get_characters_character_id_stats_activations_bomb_launcher
	//
	// activations_bomb_launcher integer
	ActivationsBombLauncher int64 `json:"activations_bomb_launcher,omitempty"`

	// get_characters_character_id_stats_activations_capacitor_booster
	//
	// activations_capacitor_booster integer
	ActivationsCapacitorBooster int64 `json:"activations_capacitor_booster,omitempty"`

	// get_characters_character_id_stats_activations_cargo_scanner
	//
	// activations_cargo_scanner integer
	ActivationsCargoScanner int64 `json:"activations_cargo_scanner,omitempty"`

	// get_characters_character_id_stats_activations_cloaking_device
	//
	// activations_cloaking_device integer
	ActivationsCloakingDevice int64 `json:"activations_cloaking_device,omitempty"`

	// get_characters_character_id_stats_activations_clone_vat_bay
	//
	// activations_clone_vat_bay integer
	ActivationsCloneVatBay int64 `json:"activations_clone_vat_bay,omitempty"`

	// get_characters_character_id_stats_activations_cynosural_field
	//
	// activations_cynosural_field integer
	ActivationsCynosuralField int64 `json:"activations_cynosural_field,omitempty"`

	// get_characters_character_id_stats_activations_damage_control
	//
	// activations_damage_control integer
	ActivationsDamageControl int64 `json:"activations_damage_control,omitempty"`

	// get_characters_character_id_stats_activations_data_miners
	//
	// activations_data_miners integer
	ActivationsDataMiners int64 `json:"activations_data_miners,omitempty"`

	// get_characters_character_id_stats_activations_drone_control_unit
	//
	// activations_drone_control_unit integer
	ActivationsDroneControlUnit int64 `json:"activations_drone_control_unit,omitempty"`

	// get_characters_character_id_stats_activations_drone_tracking_modules
	//
	// activations_drone_tracking_modules integer
	ActivationsDroneTrackingModules int64 `json:"activations_drone_tracking_modules,omitempty"`

	// get_characters_character_id_stats_activations_eccm
	//
	// activations_eccm integer
	ActivationsEccm int64 `json:"activations_eccm,omitempty"`

	// get_characters_character_id_stats_activations_ecm
	//
	// activations_ecm integer
	ActivationsEcm int64 `json:"activations_ecm,omitempty"`

	// get_characters_character_id_stats_activations_ecm_burst
	//
	// activations_ecm_burst integer
	ActivationsEcmBurst int64 `json:"activations_ecm_burst,omitempty"`

	// get_characters_character_id_stats_activations_energy_destabilizer
	//
	// activations_energy_destabilizer integer
	ActivationsEnergyDestabilizer int64 `json:"activations_energy_destabilizer,omitempty"`

	// get_characters_character_id_stats_activations_energy_vampire
	//
	// activations_energy_vampire integer
	ActivationsEnergyVampire int64 `json:"activations_energy_vampire,omitempty"`

	// get_characters_character_id_stats_activations_energy_weapon
	//
	// activations_energy_weapon integer
	ActivationsEnergyWeapon int64 `json:"activations_energy_weapon,omitempty"`

	// get_characters_character_id_stats_activations_festival_launcher
	//
	// activations_festival_launcher integer
	ActivationsFestivalLauncher int64 `json:"activations_festival_launcher,omitempty"`

	// get_characters_character_id_stats_activations_frequency_mining_laser
	//
	// activations_frequency_mining_laser integer
	ActivationsFrequencyMiningLaser int64 `json:"activations_frequency_mining_laser,omitempty"`

	// get_characters_character_id_stats_activations_fueled_armor_repairer
	//
	// activations_fueled_armor_repairer integer
	ActivationsFueledArmorRepairer int64 `json:"activations_fueled_armor_repairer,omitempty"`

	// get_characters_character_id_stats_activations_fueled_shield_booster
	//
	// activations_fueled_shield_booster integer
	ActivationsFueledShieldBooster int64 `json:"activations_fueled_shield_booster,omitempty"`

	// get_characters_character_id_stats_activations_gang_coordinator
	//
	// activations_gang_coordinator integer
	ActivationsGangCoordinator int64 `json:"activations_gang_coordinator,omitempty"`

	// get_characters_character_id_stats_activations_gas_cloud_harvester
	//
	// activations_gas_cloud_harvester integer
	ActivationsGasCloudHarvester int64 `json:"activations_gas_cloud_harvester,omitempty"`

	// get_characters_character_id_stats_activations_hull_repair_unit
	//
	// activations_hull_repair_unit integer
	ActivationsHullRepairUnit int64 `json:"activations_hull_repair_unit,omitempty"`

	// get_characters_character_id_stats_activations_hybrid_weapon
	//
	// activations_hybrid_weapon integer
	ActivationsHybridWeapon int64 `json:"activations_hybrid_weapon,omitempty"`

	// get_characters_character_id_stats_activations_industrial_core
	//
	// activations_industrial_core integer
	ActivationsIndustrialCore int64 `json:"activations_industrial_core,omitempty"`

	// get_characters_character_id_stats_activations_interdiction_sphere_launcher
	//
	// activations_interdiction_sphere_launcher integer
	ActivationsInterdictionSphereLauncher int64 `json:"activations_interdiction_sphere_launcher,omitempty"`

	// get_characters_character_id_stats_activations_micro_jump_drive
	//
	// activations_micro_jump_drive integer
	ActivationsMicroJumpDrive int64 `json:"activations_micro_jump_drive,omitempty"`

	// get_characters_character_id_stats_activations_mining_laser
	//
	// activations_mining_laser integer
	ActivationsMiningLaser int64 `json:"activations_mining_laser,omitempty"`

	// get_characters_character_id_stats_activations_missile_launcher
	//
	// activations_missile_launcher integer
	ActivationsMissileLauncher int64 `json:"activations_missile_launcher,omitempty"`

	// get_characters_character_id_stats_activations_passive_targeting_system
	//
	// activations_passive_targeting_system integer
	ActivationsPassiveTargetingSystem int64 `json:"activations_passive_targeting_system,omitempty"`

	// get_characters_character_id_stats_activations_probe_launcher
	//
	// activations_probe_launcher integer
	ActivationsProbeLauncher int64 `json:"activations_probe_launcher,omitempty"`

	// get_characters_character_id_stats_activations_projected_eccm
	//
	// activations_projected_eccm integer
	ActivationsProjectedEccm int64 `json:"activations_projected_eccm,omitempty"`

	// get_characters_character_id_stats_activations_projectile_weapon
	//
	// activations_projectile_weapon integer
	ActivationsProjectileWeapon int64 `json:"activations_projectile_weapon,omitempty"`

	// get_characters_character_id_stats_activations_propulsion_module
	//
	// activations_propulsion_module integer
	ActivationsPropulsionModule int64 `json:"activations_propulsion_module,omitempty"`

	// get_characters_character_id_stats_activations_remote_armor_repairer
	//
	// activations_remote_armor_repairer integer
	ActivationsRemoteArmorRepairer int64 `json:"activations_remote_armor_repairer,omitempty"`

	// get_characters_character_id_stats_activations_remote_capacitor_transmitter
	//
	// activations_remote_capacitor_transmitter integer
	ActivationsRemoteCapacitorTransmitter int64 `json:"activations_remote_capacitor_transmitter,omitempty"`

	// get_characters_character_id_stats_activations_remote_ecm_burst
	//
	// activations_remote_ecm_burst integer
	ActivationsRemoteEcmBurst int64 `json:"activations_remote_ecm_burst,omitempty"`

	// get_characters_character_id_stats_activations_remote_hull_repairer
	//
	// activations_remote_hull_repairer integer
	ActivationsRemoteHullRepairer int64 `json:"activations_remote_hull_repairer,omitempty"`

	// get_characters_character_id_stats_activations_remote_sensor_booster
	//
	// activations_remote_sensor_booster integer
	ActivationsRemoteSensorBooster int64 `json:"activations_remote_sensor_booster,omitempty"`

	// get_characters_character_id_stats_activations_remote_sensor_damper
	//
	// activations_remote_sensor_damper integer
	ActivationsRemoteSensorDamper int64 `json:"activations_remote_sensor_damper,omitempty"`

	// get_characters_character_id_stats_activations_remote_shield_booster
	//
	// activations_remote_shield_booster integer
	ActivationsRemoteShieldBooster int64 `json:"activations_remote_shield_booster,omitempty"`

	// get_characters_character_id_stats_activations_remote_tracking_computer
	//
	// activations_remote_tracking_computer integer
	ActivationsRemoteTrackingComputer int64 `json:"activations_remote_tracking_computer,omitempty"`

	// get_characters_character_id_stats_activations_salvager
	//
	// activations_salvager integer
	ActivationsSalvager int64 `json:"activations_salvager,omitempty"`

	// get_characters_character_id_stats_activations_sensor_booster
	//
	// activations_sensor_booster integer
	ActivationsSensorBooster int64 `json:"activations_sensor_booster,omitempty"`

	// get_characters_character_id_stats_activations_shield_booster
	//
	// activations_shield_booster integer
	ActivationsShieldBooster int64 `json:"activations_shield_booster,omitempty"`

	// get_characters_character_id_stats_activations_shield_hardener
	//
	// activations_shield_hardener integer
	ActivationsShieldHardener int64 `json:"activations_shield_hardener,omitempty"`

	// get_characters_character_id_stats_activations_ship_scanner
	//
	// activations_ship_scanner integer
	ActivationsShipScanner int64 `json:"activations_ship_scanner,omitempty"`

	// get_characters_character_id_stats_activations_siege
	//
	// activations_siege integer
	ActivationsSiege int64 `json:"activations_siege,omitempty"`

	// get_characters_character_id_stats_activations_smart_bomb
	//
	// activations_smart_bomb integer
	ActivationsSmartBomb int64 `json:"activations_smart_bomb,omitempty"`

	// get_characters_character_id_stats_activations_stasis_web
	//
	// activations_stasis_web integer
	ActivationsStasisWeb int64 `json:"activations_stasis_web,omitempty"`

	// get_characters_character_id_stats_activations_strip_miner
	//
	// activations_strip_miner integer
	ActivationsStripMiner int64 `json:"activations_strip_miner,omitempty"`

	// get_characters_character_id_stats_activations_super_weapon
	//
	// activations_super_weapon integer
	ActivationsSuperWeapon int64 `json:"activations_super_weapon,omitempty"`

	// get_characters_character_id_stats_activations_survey_scanner
	//
	// activations_survey_scanner integer
	ActivationsSurveyScanner int64 `json:"activations_survey_scanner,omitempty"`

	// get_characters_character_id_stats_activations_target_breaker
	//
	// activations_target_breaker integer
	ActivationsTargetBreaker int64 `json:"activations_target_breaker,omitempty"`

	// get_characters_character_id_stats_activations_target_painter
	//
	// activations_target_painter integer
	ActivationsTargetPainter int64 `json:"activations_target_painter,omitempty"`

	// get_characters_character_id_stats_activations_tracking_computer
	//
	// activations_tracking_computer integer
	ActivationsTrackingComputer int64 `json:"activations_tracking_computer,omitempty"`

	// get_characters_character_id_stats_activations_tracking_disruptor
	//
	// activations_tracking_disruptor integer
	ActivationsTrackingDisruptor int64 `json:"activations_tracking_disruptor,omitempty"`

	// get_characters_character_id_stats_activations_tractor_beam
	//
	// activations_tractor_beam integer
	ActivationsTractorBeam int64 `json:"activations_tractor_beam,omitempty"`

	// get_characters_character_id_stats_activations_triage
	//
	// activations_triage integer
	ActivationsTriage int64 `json:"activations_triage,omitempty"`

	// get_characters_character_id_stats_activations_warp_disrupt_field_generator
	//
	// activations_warp_disrupt_field_generator integer
	ActivationsWarpDisruptFieldGenerator int64 `json:"activations_warp_disrupt_field_generator,omitempty"`

	// get_characters_character_id_stats_activations_warp_scrambler
	//
	// activations_warp_scrambler integer
	ActivationsWarpScrambler int64 `json:"activations_warp_scrambler,omitempty"`

	// get_characters_character_id_stats_link_weapons
	//
	// link_weapons integer
	LinkWeapons int64 `json:"link_weapons,omitempty"`

	// get_characters_character_id_stats_overload
	//
	// overload integer
	Overload int64 `json:"overload,omitempty"`

	// get_characters_character_id_stats_repairs
	//
	// repairs integer
	Repairs int64 `json:"repairs,omitempty"`
}

// Validate validates this get characters character Id stats o k body items module
func (m *GetCharactersCharacterIDStatsOKBodyItemsModule) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDStatsOKBodyItemsModule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDStatsOKBodyItemsModule) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDStatsOKBodyItemsModule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
