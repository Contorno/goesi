/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * OpenAPI spec version: 0.8.5
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package esi

/* A list of GetCharactersCharacterIdStatsModule. */
//easyjson:json
type GetCharactersCharacterIdStatsModuleList []GetCharactersCharacterIdStatsModule

/* module object */
//easyjson:json
type GetCharactersCharacterIdStatsModule struct {
	ActivationsArmorHardener                int64 `json:"activations_armor_hardener,omitempty"`                  /* activations_armor_hardener integer */
	ActivationsArmorRepairUnit              int64 `json:"activations_armor_repair_unit,omitempty"`               /* activations_armor_repair_unit integer */
	ActivationsArmorResistanceShiftHardener int64 `json:"activations_armor_resistance_shift_hardener,omitempty"` /* activations_armor_resistance_shift_hardener integer */
	ActivationsAutomatedTargetingSystem     int64 `json:"activations_automated_targeting_system,omitempty"`      /* activations_automated_targeting_system integer */
	ActivationsBastion                      int64 `json:"activations_bastion,omitempty"`                         /* activations_bastion integer */
	ActivationsBombLauncher                 int64 `json:"activations_bomb_launcher,omitempty"`                   /* activations_bomb_launcher integer */
	ActivationsCapacitorBooster             int64 `json:"activations_capacitor_booster,omitempty"`               /* activations_capacitor_booster integer */
	ActivationsCargoScanner                 int64 `json:"activations_cargo_scanner,omitempty"`                   /* activations_cargo_scanner integer */
	ActivationsCloakingDevice               int64 `json:"activations_cloaking_device,omitempty"`                 /* activations_cloaking_device integer */
	ActivationsCloneVatBay                  int64 `json:"activations_clone_vat_bay,omitempty"`                   /* activations_clone_vat_bay integer */
	ActivationsCynosuralField               int64 `json:"activations_cynosural_field,omitempty"`                 /* activations_cynosural_field integer */
	ActivationsDamageControl                int64 `json:"activations_damage_control,omitempty"`                  /* activations_damage_control integer */
	ActivationsDataMiners                   int64 `json:"activations_data_miners,omitempty"`                     /* activations_data_miners integer */
	ActivationsDroneControlUnit             int64 `json:"activations_drone_control_unit,omitempty"`              /* activations_drone_control_unit integer */
	ActivationsDroneTrackingModules         int64 `json:"activations_drone_tracking_modules,omitempty"`          /* activations_drone_tracking_modules integer */
	ActivationsEccm                         int64 `json:"activations_eccm,omitempty"`                            /* activations_eccm integer */
	ActivationsEcm                          int64 `json:"activations_ecm,omitempty"`                             /* activations_ecm integer */
	ActivationsEcmBurst                     int64 `json:"activations_ecm_burst,omitempty"`                       /* activations_ecm_burst integer */
	ActivationsEnergyDestabilizer           int64 `json:"activations_energy_destabilizer,omitempty"`             /* activations_energy_destabilizer integer */
	ActivationsEnergyVampire                int64 `json:"activations_energy_vampire,omitempty"`                  /* activations_energy_vampire integer */
	ActivationsEnergyWeapon                 int64 `json:"activations_energy_weapon,omitempty"`                   /* activations_energy_weapon integer */
	ActivationsFestivalLauncher             int64 `json:"activations_festival_launcher,omitempty"`               /* activations_festival_launcher integer */
	ActivationsFrequencyMiningLaser         int64 `json:"activations_frequency_mining_laser,omitempty"`          /* activations_frequency_mining_laser integer */
	ActivationsFueledArmorRepairer          int64 `json:"activations_fueled_armor_repairer,omitempty"`           /* activations_fueled_armor_repairer integer */
	ActivationsFueledShieldBooster          int64 `json:"activations_fueled_shield_booster,omitempty"`           /* activations_fueled_shield_booster integer */
	ActivationsGangCoordinator              int64 `json:"activations_gang_coordinator,omitempty"`                /* activations_gang_coordinator integer */
	ActivationsGasCloudHarvester            int64 `json:"activations_gas_cloud_harvester,omitempty"`             /* activations_gas_cloud_harvester integer */
	ActivationsHullRepairUnit               int64 `json:"activations_hull_repair_unit,omitempty"`                /* activations_hull_repair_unit integer */
	ActivationsHybridWeapon                 int64 `json:"activations_hybrid_weapon,omitempty"`                   /* activations_hybrid_weapon integer */
	ActivationsIndustrialCore               int64 `json:"activations_industrial_core,omitempty"`                 /* activations_industrial_core integer */
	ActivationsInterdictionSphereLauncher   int64 `json:"activations_interdiction_sphere_launcher,omitempty"`    /* activations_interdiction_sphere_launcher integer */
	ActivationsMicroJumpDrive               int64 `json:"activations_micro_jump_drive,omitempty"`                /* activations_micro_jump_drive integer */
	ActivationsMiningLaser                  int64 `json:"activations_mining_laser,omitempty"`                    /* activations_mining_laser integer */
	ActivationsMissileLauncher              int64 `json:"activations_missile_launcher,omitempty"`                /* activations_missile_launcher integer */
	ActivationsPassiveTargetingSystem       int64 `json:"activations_passive_targeting_system,omitempty"`        /* activations_passive_targeting_system integer */
	ActivationsProbeLauncher                int64 `json:"activations_probe_launcher,omitempty"`                  /* activations_probe_launcher integer */
	ActivationsProjectedEccm                int64 `json:"activations_projected_eccm,omitempty"`                  /* activations_projected_eccm integer */
	ActivationsProjectileWeapon             int64 `json:"activations_projectile_weapon,omitempty"`               /* activations_projectile_weapon integer */
	ActivationsPropulsionModule             int64 `json:"activations_propulsion_module,omitempty"`               /* activations_propulsion_module integer */
	ActivationsRemoteArmorRepairer          int64 `json:"activations_remote_armor_repairer,omitempty"`           /* activations_remote_armor_repairer integer */
	ActivationsRemoteCapacitorTransmitter   int64 `json:"activations_remote_capacitor_transmitter,omitempty"`    /* activations_remote_capacitor_transmitter integer */
	ActivationsRemoteEcmBurst               int64 `json:"activations_remote_ecm_burst,omitempty"`                /* activations_remote_ecm_burst integer */
	ActivationsRemoteHullRepairer           int64 `json:"activations_remote_hull_repairer,omitempty"`            /* activations_remote_hull_repairer integer */
	ActivationsRemoteSensorBooster          int64 `json:"activations_remote_sensor_booster,omitempty"`           /* activations_remote_sensor_booster integer */
	ActivationsRemoteSensorDamper           int64 `json:"activations_remote_sensor_damper,omitempty"`            /* activations_remote_sensor_damper integer */
	ActivationsRemoteShieldBooster          int64 `json:"activations_remote_shield_booster,omitempty"`           /* activations_remote_shield_booster integer */
	ActivationsRemoteTrackingComputer       int64 `json:"activations_remote_tracking_computer,omitempty"`        /* activations_remote_tracking_computer integer */
	ActivationsSalvager                     int64 `json:"activations_salvager,omitempty"`                        /* activations_salvager integer */
	ActivationsSensorBooster                int64 `json:"activations_sensor_booster,omitempty"`                  /* activations_sensor_booster integer */
	ActivationsShieldBooster                int64 `json:"activations_shield_booster,omitempty"`                  /* activations_shield_booster integer */
	ActivationsShieldHardener               int64 `json:"activations_shield_hardener,omitempty"`                 /* activations_shield_hardener integer */
	ActivationsShipScanner                  int64 `json:"activations_ship_scanner,omitempty"`                    /* activations_ship_scanner integer */
	ActivationsSiege                        int64 `json:"activations_siege,omitempty"`                           /* activations_siege integer */
	ActivationsSmartBomb                    int64 `json:"activations_smart_bomb,omitempty"`                      /* activations_smart_bomb integer */
	ActivationsStasisWeb                    int64 `json:"activations_stasis_web,omitempty"`                      /* activations_stasis_web integer */
	ActivationsStripMiner                   int64 `json:"activations_strip_miner,omitempty"`                     /* activations_strip_miner integer */
	ActivationsSuperWeapon                  int64 `json:"activations_super_weapon,omitempty"`                    /* activations_super_weapon integer */
	ActivationsSurveyScanner                int64 `json:"activations_survey_scanner,omitempty"`                  /* activations_survey_scanner integer */
	ActivationsTargetBreaker                int64 `json:"activations_target_breaker,omitempty"`                  /* activations_target_breaker integer */
	ActivationsTargetPainter                int64 `json:"activations_target_painter,omitempty"`                  /* activations_target_painter integer */
	ActivationsTrackingComputer             int64 `json:"activations_tracking_computer,omitempty"`               /* activations_tracking_computer integer */
	ActivationsTrackingDisruptor            int64 `json:"activations_tracking_disruptor,omitempty"`              /* activations_tracking_disruptor integer */
	ActivationsTractorBeam                  int64 `json:"activations_tractor_beam,omitempty"`                    /* activations_tractor_beam integer */
	ActivationsTriage                       int64 `json:"activations_triage,omitempty"`                          /* activations_triage integer */
	ActivationsWarpDisruptFieldGenerator    int64 `json:"activations_warp_disrupt_field_generator,omitempty"`    /* activations_warp_disrupt_field_generator integer */
	ActivationsWarpScrambler                int64 `json:"activations_warp_scrambler,omitempty"`                  /* activations_warp_scrambler integer */
	LinkWeapons                             int64 `json:"link_weapons,omitempty"`                                /* link_weapons integer */
	Overload                                int64 `json:"overload,omitempty"`                                    /* overload integer */
	Repairs                                 int64 `json:"repairs,omitempty"`                                     /* repairs integer */
}
