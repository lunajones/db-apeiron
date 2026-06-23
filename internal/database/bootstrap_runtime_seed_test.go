package database

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBootstrapSeedsCoverRequiredRuntimeSkills(t *testing.T) {
	sql := readBootstrapSQL(t)
	requiredSkills := []string{
		"player_basic_attack_1",
		"player_basic_attack_2",
		"player_basic_attack_3",
		"player_shield_bash",
		"player_shield_rush",
		"bite",
		"lunge",
		"wolf_dodge",
		"maul",
	}
	for _, skillID := range requiredSkills {
		quoted := "'" + skillID + "'"
		if !strings.Contains(sql, quoted) {
			t.Fatalf("bootstrap is missing required runtime skill %s", skillID)
		}
	}
	for _, contractID := range []string{
		"grounded_move_reconciliation",
		"grounded_skill_action_reconciliation",
		"dodge_reconciliation",
		"leap_reconciliation",
		"turn_reconciliation",
		"post_action_handoff_reconciliation",
		"grounded_move_v1",
		"turn_v1_rate_limited_contextual",
		"dodge_v1_full_iframe",
		"jump_v1_authoritative_grounded_handoff",
		"basic_attack_1_forward_cut_v1",
		"basic_attack_2_cross_cut_v1",
		"basic_attack_3_shield_drive_v1",
		"shield_bash_front_push_v1",
		"shield_rush_front_contact_v1",
		"wolf_bite_melee_commit_v1",
		"low_fast_lunge_v1",
		"wolf_dodge_lateral_leap_v1",
		"wolf_maul_lateral_counter_v1",
	} {
		if !strings.Contains(sql, "'"+contractID+"'") {
			t.Fatalf("bootstrap is missing required runtime contract %s", contractID)
		}
	}
}

func TestBootstrapSeedsMirrorServerRuntimeRequirementManifest(t *testing.T) {
	sql := readBootstrapSQL(t)

	manifest := []runtimeSeedRequirement{
		{
			category:    "movement_profile",
			id:          "player_default_movement_profile",
			description: "rich runtime movement/reconciliation profile consumed by server and Unreal",
			fragments: []string{
				"'player_default_movement_profile'",
				"grounded_transition_deadzone_min",
				"leap_landing_clamp_ignore_deadzone",
				"dodge_carry_handoff_ms",
				"rotation_rate_yaw",
				"movement_turn_resubmit_dot_threshold",
			},
		},
		{
			category:    "base_movement_action",
			id:          "grounded_move_v1",
			description: "normal grounded locomotion",
			fragments: []string{
				"('grounded_move_v1','move'",
				"'grounded_move_reconciliation'",
				"'grounded_move'",
				"'movement'",
			},
		},
		{
			category:    "base_movement_action",
			id:          "turn_v1_rate_limited_contextual",
			description: "camera/control yaw action",
			fragments: []string{
				"('turn_v1_rate_limited_contextual','turn'",
				"'turn_reconciliation'",
				"'turn'",
				`"yaw_rate_deg_per_sec":720`,
			},
		},
		{
			category:    "base_movement_action",
			id:          "dodge_v1_full_iframe",
			description: "protected full-iframe dodge baseline",
			fragments: []string{
				"('dodge_v1_full_iframe','dodge'",
				"'dodge_reconciliation'",
				"'iframe'",
				",260,812.5,0,",
				`"ability_key":"dodge"`,
				`"speed_semantics":"authored_base_speed_distance_over_duration"`,
			},
		},
		{
			category:    "base_movement_action",
			id:          "jump_v1_authoritative_grounded_handoff",
			description: "protected leap/jump baseline",
			fragments: []string{
				"('jump_v1_authoritative_grounded_handoff','leap'",
				"'leap_reconciliation'",
				"'grounded_handoff'",
				`"landing_detection_policy":"server_grounded_handoff"`,
			},
		},
		{
			category:    "combat_core_profile",
			id:          "combat_core_player_sword_shield_v1",
			description: "player sword-and-shield combat core",
			fragments: []string{
				"'combat_core_player_sword_shield_v1'",
				"block_stamina_cost_per_sec",
				"can_block",
				"can_parry",
			},
		},
		{
			category:    "combat_core_profile",
			id:          "combat_core_steppe_wolf",
			description: "steppe wolf combat core",
			fragments: []string{
				"'combat_core_steppe_wolf'",
				"dodge_stamina_cost",
				"stamina_regen_per_sec",
				"cc_duration_multiplier",
			},
		},
		{
			category:    "defense_contract",
			id:          "player_shield_guard_v1",
			description: "player shield guard defense contract",
			fragments: []string{
				"'player_shield_guard_v1'",
				"'shield_block'",
				`"frontFacing":"control_rotation_yaw"`,
			},
		},
		{
			category:    "defense_contract",
			id:          "wolf_attack_vs_guard_v1",
			description: "wolf hit-vs-guard interaction contract",
			fragments: []string{
				"'wolf_attack_vs_guard_v1'",
				"'incoming_melee'",
				"stamina_damage_only_on_block",
			},
		},
		{
			category:    "weapon_kit",
			id:          "weaponkit_sword_shield",
			description: "current player sword/shield combat mode slots",
			fragments: []string{
				"'weaponkit_sword_shield'",
				"'mode_sword_shield_bulwark'",
				"'mode_sword_shield_vanguard'",
				"('mode_sword_shield_bulwark','M1','player_basic_attack_1'",
				"('mode_sword_shield_bulwark','R','player_shield_bash'",
				"('mode_sword_shield_bulwark','F','player_shield_rush'",
				"('mode_sword_shield_vanguard','M1',NULL,FALSE,FALSE,FALSE",
			},
		},
		{
			category:    "wolf_brain_policy",
			id:          "contract_wolf_pack_harasser_v1",
			description: "wolf creature brain runtime contract",
			fragments: []string{
				"'contract_wolf_pack_harasser_v1'",
				"'opportunity_wolf_harasser_v1'",
				"'orbit_wolf_harasser_combat_walk_v1'",
				"'wolf_evasion_pressure_v1'",
				"'wolf_lunge_flank_windup_v1'",
				"'wolf_lunge_chase_windup_v1'",
				"'wolf_maul_pressure_counter_v1'",
				"'wolf_dodge_pressure_evasion_v1'",
			},
		},
	}

	for _, requirement := range manifest {
		for _, fragment := range requirement.fragments {
			if !strings.Contains(sql, fragment) {
				t.Fatalf("bootstrap runtime manifest missing %s %s (%s): %s", requirement.category, requirement.id, requirement.description, fragment)
			}
		}
	}
}

func TestBootstrapSeedsMirrorRequiredSkillActionManifest(t *testing.T) {
	sql := readBootstrapSQL(t)

	manifest := []skillActionSeedRequirement{
		{
			skillID:          "player_basic_attack_1",
			actionContractID: "basic_attack_1_forward_cut_v1",
			timing:           "'player_basic_attack_1',90,140,120,0,2000",
			hitbox:           "'hitbox_player_basic_attack_1_0','player_basic_attack_1',0,'temporal_sweep'",
			motionProfileID:  "motion_player_basic_attack_1_forward_v1",
			damageGroupID:    "player_basic_attack_1_damage",
			binding:          "('player_basic_attack_1','basic_attack_1_forward_cut_v1','active','explicit_recovery_handoff','blocked_during_owned_root'",
		},
		{
			skillID:          "player_basic_attack_2",
			actionContractID: "basic_attack_2_cross_cut_v1",
			timing:           "'player_basic_attack_2',100,150,120,0,2000",
			hitbox:           "'hitbox_player_basic_attack_2_0','player_basic_attack_2',0,'temporal_sweep'",
			motionProfileID:  "motion_player_basic_attack_2_right_to_left_v1",
			damageGroupID:    "player_basic_attack_2_damage",
			binding:          "('player_basic_attack_2','basic_attack_2_cross_cut_v1','active','explicit_recovery_handoff','blocked_during_owned_root'",
		},
		{
			skillID:          "player_basic_attack_3",
			actionContractID: "basic_attack_3_shield_drive_v1",
			timing:           "'player_basic_attack_3',180,260,180,0,2000",
			hitbox:           "'hitbox_player_basic_attack_3_0','player_basic_attack_3',0,'temporal_sweep'",
			motionProfileID:  "motion_player_basic_attack_3_shield_drive_v1",
			damageGroupID:    "player_basic_attack_3_damage",
			binding:          "('player_basic_attack_3','basic_attack_3_shield_drive_v1','active','explicit_recovery_handoff','blocked_during_owned_root'",
		},
		{
			skillID:          "player_shield_bash",
			actionContractID: "shield_bash_front_push_v1",
			timing:           "'player_shield_bash',110,170,120,2600",
			hitbox:           "'hitbox_player_shield_bash_0','player_shield_bash',0,'temporal_sweep'",
			motionProfileID:  "motion_player_shield_bash_front_push_v1",
			damageGroupID:    "player_shield_bash_front_push",
			binding:          "('player_shield_bash','shield_bash_front_push_v1','active','explicit_recovery_handoff','blocked_during_owned_root'",
		},
		{
			skillID:          "player_shield_rush",
			actionContractID: "shield_rush_front_contact_v1",
			timing:           "'player_shield_rush',160,720,260,5200",
			hitbox:           "'hitbox_player_shield_rush_0','player_shield_rush',0,'temporal_sweep'",
			motionProfileID:  "motion_player_shield_rush_front_contact_v1",
			damageGroupID:    "player_shield_rush_front_contact",
			binding:          "('player_shield_rush','shield_rush_front_contact_v1','active','explicit_recovery_handoff','blocked_during_owned_root'",
		},
		{
			skillID:          "bite",
			actionContractID: "wolf_bite_melee_commit_v1",
			timing:           "'bite',120,220,180,900",
			hitbox:           "'hitbox_bite_0','bite',0,'temporal_sweep'",
			motionProfileID:  "motion_wolf_bite_melee_v1",
			damageGroupID:    "wolf_bite_damage",
			binding:          "('bite','wolf_bite_melee_commit_v1','active','explicit_recovery_handoff','blocked_during_owned_root'",
		},
		{
			skillID:          "lunge",
			actionContractID: "low_fast_lunge_v1",
			timing:           "'lunge',3600,380,520,7200",
			hitbox:           "'hitbox_lunge_0','lunge',0,'temporal_sweep'",
			motionProfileID:  "motion_wolf_lunge_cross_v1",
			damageGroupID:    "wolf_lunge_damage",
			binding:          "('lunge','low_fast_lunge_v1','active','grounded_handoff','blocked_during_airborne'",
		},
		{
			skillID:          "wolf_dodge",
			actionContractID: "wolf_dodge_lateral_leap_v1",
			timing:           "'wolf_dodge',0,420,100,0",
			hitbox:           "",
			motionProfileID:  "",
			damageGroupID:    "",
			binding:          "('wolf_dodge','wolf_dodge_lateral_leap_v1','active','explicit_recovery_handoff','blocked_during_owned_root'",
		},
		{
			skillID:          "maul",
			actionContractID: "wolf_maul_lateral_counter_v1",
			timing:           "'maul',220,520,220,5200",
			hitbox:           "'hitbox_maul_0','maul',0,'temporal_sweep'",
			motionProfileID:  "motion_wolf_maul_lateral_counter_v1",
			damageGroupID:    "wolf_maul_damage",
			binding:          "('maul','wolf_maul_lateral_counter_v1','active','explicit_recovery_handoff','blocked_during_owned_root'",
		},
	}

	for _, requirement := range manifest {
		required := []string{
			"'" + requirement.skillID + "'",
			"'" + requirement.actionContractID + "'",
			requirement.timing,
			requirement.binding,
		}
		if requirement.hitbox != "" {
			required = append(required, requirement.hitbox)
		}
		if requirement.motionProfileID != "" {
			required = append(required, "'"+requirement.motionProfileID+"'")
		}
		if requirement.damageGroupID != "" {
			required = append(required, "'"+requirement.damageGroupID+"'")
		}
		for _, fragment := range required {
			if !strings.Contains(sql, fragment) {
				t.Fatalf("bootstrap skill action manifest missing %s -> %s: %s", requirement.skillID, requirement.actionContractID, fragment)
			}
		}
	}
}

func TestBootstrapSeedsCoverTemporalHitboxesForCombatRuntime(t *testing.T) {
	sql := readBootstrapSQL(t)
	requiredMotionProfiles := []string{
		"motion_player_basic_attack_1_forward_v1",
		"motion_player_basic_attack_2_right_to_left_v1",
		"motion_player_basic_attack_3_shield_drive_v1",
		"motion_player_shield_bash_front_push_v1",
		"motion_player_shield_rush_front_contact_v1",
		"motion_wolf_bite_melee_v1",
		"motion_wolf_lunge_cross_v1",
		"motion_wolf_maul_lateral_counter_v1",
	}
	for _, profileID := range requiredMotionProfiles {
		if !strings.Contains(sql, "'"+profileID+"'") {
			t.Fatalf("bootstrap is missing temporal hitbox motion profile %s", profileID)
		}
	}
}

func TestBootstrapSeedsPreservePlayerShieldKitMotionGeometry(t *testing.T) {
	sql := readBootstrapSQL(t)
	required := []string{
		"('basic_attack_1_forward_cut_v1','grounded_skill','Basic attack 1 short shield jab: one player cylinder forward.',350,140,120,84,240",
		"('basic_attack_2_cross_cut_v1','grounded_skill','Basic attack 2 short left-to-right shield sweep.',370,150,120,42,114",
		"('basic_attack_3_shield_drive_v1','grounded_skill','Basic attack 3 committed overhead shield punch with contact carry and interrupt.',620,260,180,252,406.4",
		"('motion_player_basic_attack_1_forward_v1',0,0.00,'box_strip',0,0,90,42,84,150,42,42",
		"('motion_player_basic_attack_1_forward_v1',2,1.00,'box_strip',0,0,90,84,84,150,42,84",
		"('motion_player_basic_attack_2_right_to_left_v1',0,0.00,'arc_slice',70,-35,95,0,0,150,50,125,15,45",
		"('motion_player_basic_attack_2_right_to_left_v1',2,1.00,'arc_slice',70,35,95,0,0,150,50,125,-45,-15",
		"('motion_player_basic_attack_3_shield_drive_v1',2,1.00,'capsule_strip',0,0,95,84,0,155,42,252",
		"('hitbox_player_basic_attack_1_0','player_basic_attack_1',0,'temporal_sweep',90,230,0,0,90,84,84,150,42,84",
		"('hitbox_player_basic_attack_3_0','player_basic_attack_3',0,'temporal_sweep',180,440,0,0,95,84,0,155,42,252",
	}
	for _, fragment := range required {
		if !strings.Contains(sql, fragment) {
			t.Fatalf("player shield kit temporal/movement geometry fragment missing: %s", fragment)
		}
	}
}

func TestBootstrapSeedsCoverCurrentSwordShieldCombatModes(t *testing.T) {
	sql := readBootstrapSQL(t)
	requiredSlots := []string{
		"('mode_sword_shield_vanguard','M1',NULL,FALSE,FALSE,FALSE",
		"('mode_sword_shield_bulwark','M1','player_basic_attack_1'",
		"('mode_sword_shield_bulwark','R','player_shield_bash'",
		"('mode_sword_shield_bulwark','F','player_shield_rush'",
	}
	for _, slot := range requiredSlots {
		if !strings.Contains(sql, slot) {
			t.Fatalf("bootstrap is missing combat mode slot %s", slot)
		}
	}
}

func TestBootstrapSeedsBindRequiredSkillsToCanonicalMovementActions(t *testing.T) {
	sql := readBootstrapSQL(t)
	requiredBindings := map[string]string{
		"player_basic_attack_1": "basic_attack_1_forward_cut_v1",
		"player_basic_attack_2": "basic_attack_2_cross_cut_v1",
		"player_basic_attack_3": "basic_attack_3_shield_drive_v1",
		"player_shield_bash":    "shield_bash_front_push_v1",
		"player_shield_rush":    "shield_rush_front_contact_v1",
		"lunge":                 "low_fast_lunge_v1",
		"wolf_dodge":            "wolf_dodge_lateral_leap_v1",
		"bite":                  "wolf_bite_melee_commit_v1",
		"maul":                  "wolf_maul_lateral_counter_v1",
	}
	for skillID, contractID := range requiredBindings {
		expected := "('" + skillID + "','" + contractID + "'"
		if !strings.Contains(sql, expected) {
			t.Fatalf("bootstrap is missing movement binding %s -> %s", skillID, contractID)
		}
	}
}

func TestLegacySkillMovementEffectsAreCompatibilityOnly(t *testing.T) {
	sql := readBootstrapSQL(t)
	requiredCompatRows := map[string]string{
		"lunge":              "low_fast_lunge_v1",
		"player_shield_rush": "shield_rush_front_contact_v1",
		"player_shield_bash": "shield_bash_front_push_v1",
	}
	for skillID, canonicalContractID := range requiredCompatRows {
		if !strings.Contains(sql, "'"+skillID+"'") {
			t.Fatalf("skill movement compatibility row missing skill %s", skillID)
		}
		if !strings.Contains(sql, `"prefer":"movement_action_contract"`) {
			t.Fatal("skill movement compatibility seed must declare movement_action_contract as preferred authority")
		}
		if !strings.Contains(sql, "'"+canonicalContractID+"'") {
			t.Fatalf("canonical movement action contract missing for compatibility skill %s -> %s", skillID, canonicalContractID)
		}
		if !strings.Contains(sql, "('"+skillID+"','"+canonicalContractID+"'") {
			t.Fatalf("compatibility skill %s does not have canonical skill_movement_action_binding -> %s", skillID, canonicalContractID)
		}
	}
}

func TestSkillDataProtoDocumentsSkillMovementCompatibilityEndpoint(t *testing.T) {
	proto, err := os.ReadFile(filepath.Join("..", "..", "proto", "apeiron", "v1", "skill_data_service.proto"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(proto)
	for _, required := range []string{
		"Compatibility endpoint",
		"GetSkillMovementActionBinding",
		"GetMovementActionContract",
		"must not be used to tune or resolve root motion",
	} {
		if !strings.Contains(content, required) {
			t.Fatalf("skill_data_service.proto is missing compatibility endpoint documentation fragment %q", required)
		}
	}
}

func TestBootstrapSeedsKeepUnimplementedVanguardSlotsEmpty(t *testing.T) {
	sql := readBootstrapSQL(t)
	requiredEmptySlots := []string{
		"('mode_sword_shield_vanguard','M1',NULL,FALSE,FALSE,FALSE",
		"('mode_sword_shield_vanguard','Q',NULL,FALSE,FALSE,FALSE",
		"('mode_sword_shield_vanguard','R',NULL,FALSE,FALSE,FALSE",
		"('mode_sword_shield_vanguard','F',NULL,FALSE,FALSE,FALSE",
	}
	for _, slot := range requiredEmptySlots {
		if !strings.Contains(sql, slot) {
			t.Fatalf("vanguard unimplemented slot should stay empty and disabled: %s", slot)
		}
	}
	for _, forbidden := range []string{
		"('mode_sword_shield_vanguard','M1','player_",
		"('mode_sword_shield_vanguard','Q','player_",
		"('mode_sword_shield_vanguard','R','player_",
		"('mode_sword_shield_vanguard','F','player_",
	} {
		if strings.Contains(sql, forbidden) {
			t.Fatalf("vanguard has a fake player skill bound to an unimplemented slot: %s", forbidden)
		}
	}
}

func TestBootstrapSeedsPreserveShieldRushFrontContactGeometry(t *testing.T) {
	sql := readBootstrapSQL(t)
	required := []string{
		"'player_shield_rush','Shield Rush'",
		"0,10.2,0,5,'enemy'",
		"FALSE,9.6, NULL,NULL",
		"('shield_rush_front_contact_v1','grounded_skill'",
		"1100,720,260,960,1148",
		"('player_shield_rush_effect_v1','player_shield_rush','charge',960::DOUBLE PRECISION,1148::DOUBLE PRECISION,1100,160,260",
		`"front_contact_offset_cm":8`,
		"('motion_player_shield_rush_front_contact_v1',0,0.00,'capsule_strip',8,0,100,224,0,160,112,96",
		"('hitbox_player_shield_rush_0','player_shield_rush',0,'temporal_sweep'",
		"'motion_player_shield_rush_front_contact_v1','player_shield_rush_front_contact'",
	}
	for _, fragment := range required {
		if !strings.Contains(sql, fragment) {
			t.Fatalf("shield rush front-contact geometry fragment missing: %s", fragment)
		}
	}
}

func TestBootstrapSeedsCoverPlayerImpactControlMotionContracts(t *testing.T) {
	sql := readBootstrapSQL(t)
	required := []struct {
		skillID       string
		statusID      string
		controlType   string
		durationMS    string
		releasePolicy string
	}{
		{
			skillID:       "player_basic_attack_3",
			statusID:      "impact_shield_drive_push",
			controlType:   "push",
			durationMS:    "180",
			releasePolicy: "carry_contact_forward_release",
		},
		{
			skillID:       "player_shield_bash",
			statusID:      "impact_shield_bash_push",
			controlType:   "push",
			durationMS:    "170",
			releasePolicy: "multi_target_push_forward_release",
		},
		{
			skillID:       "player_shield_rush",
			statusID:      "impact_shield_rush_carry_push",
			controlType:   "carry_push",
			durationMS:    "720",
			releasePolicy: "multi_target_carry_push_forward_release",
		},
	}

	for _, requirement := range required {
		fragments := []string{
			"'" + requirement.skillID + "'",
			"'" + requirement.statusID + "'",
			"'" + requirement.controlType + "'",
			requirement.durationMS,
			"'" + requirement.releasePolicy + "'",
			"COALESCE((SELECT movement_distance * 100.0 FROM apeiron.skill WHERE id = '" + requirement.skillID + "'), 0.0)",
			"COALESCE((SELECT movement_distance * 100.0 FROM apeiron.skill WHERE id = '" + requirement.skillID + "'), 0.0) / (" + requirement.durationMS + ".0 / 1000.0)",
			"'source_forward'",
		}
		for _, fragment := range fragments {
			if !strings.Contains(sql, fragment) {
				t.Fatalf("impact control motion contract missing for %s: %s", requirement.skillID, fragment)
			}
		}
	}
	guardFragments := []string{
		"INSERT INTO apeiron.skill_impact_profile",
		"ON CONFLICT (skill_id) DO UPDATE SET",
		"RAISE EXCEPTION 'Apeiron bootstrap produced incomplete skill impact control motion contract'",
		") <> 4 THEN",
		"control_distance_cm > 0",
		"control_speed_cm_s > 0",
		"COALESCE(control_direction_policy, '') <> ''",
		"RAISE EXCEPTION 'Apeiron bootstrap produced mismatched skill impact control contract'",
	}
	for _, fragment := range guardFragments {
		if !strings.Contains(sql, fragment) {
			t.Fatalf("impact control bootstrap completeness guard missing: %s", fragment)
		}
	}
}

func TestBootstrapSeedsCoverWolfMaulImpactControlMotionContract(t *testing.T) {
	sql := readBootstrapSQL(t)
	required := []string{
		"'impact_wolf_maul_lateral_grab'",
		"'maul'",
		"'grab'",
		"'lateral_grab_release'",
		"COALESCE((SELECT active_ms FROM apeiron.movement_action_contract WHERE id = 'wolf_maul_lateral_counter_v1'), 0)",
		"COALESCE((SELECT distance_cm FROM apeiron.movement_action_contract WHERE id = 'wolf_maul_lateral_counter_v1'), 0.0)",
		"COALESCE((SELECT base_speed_cm_s FROM apeiron.movement_action_contract WHERE id = 'wolf_maul_lateral_counter_v1'), 0.0)",
		"'source_action_direction'",
		"OR (skill_id = 'maul' AND status_effect_id <> 'impact_wolf_maul_lateral_grab')",
	}
	for _, fragment := range required {
		if !strings.Contains(sql, fragment) {
			t.Fatalf("wolf maul impact control contract missing fragment: %s", fragment)
		}
	}
}

func TestBootstrapSeedsUseForwardXAndLateralYForTemporalCreatureHitboxes(t *testing.T) {
	sql := readBootstrapSQL(t)
	required := []string{
		"('motion_wolf_bite_melee_v1',0,0.00,'capsule_strip',45,0,85,90,0,115,45,70",
		"('motion_wolf_lunge_cross_v1',2,1.00,'capsule_strip',210,0,90,100,0,120,50,320",
		"('motion_wolf_maul_lateral_counter_v1',0,0.00,'asymmetric_arc',65,40,95,0,0,125,58,120",
		"('hitbox_bite_0','bite',0,'temporal_sweep',120,340,80,0,90,95,0,115,48,145",
		"('hitbox_lunge_0','lunge',0,'temporal_sweep',3600,3980,130,0,105,100,0,120,50,320",
		"('hitbox_maul_0','maul',0,'temporal_sweep',220,740,80,0,100,0,0,130,62,170",
	}
	for _, fragment := range required {
		if !strings.Contains(sql, fragment) {
			t.Fatalf("temporal creature hitbox should use offset_x as forward and offset_y as lateral: %s", fragment)
		}
	}
}

func TestRuntimeMovementReconciliationProfileDoesNotDependOnClientFallbacks(t *testing.T) {
	sql := readBootstrapSQL(t)
	required := []string{
		"    34,\n    45,\n    65,\n    120,",
		"    180,\n    145,\n    145,\n    90,",
		"    600,\n    120,\n    120,\n    70,",
		"    0.92,\n    0.50,\n    0.75,\n    0.50,",
	}
	for _, fragment := range required {
		if !strings.Contains(sql, fragment) {
			t.Fatalf("runtime movement reconciliation profile still leaves fallback-owned values implicit: %s", fragment)
		}
	}

	migration, err := os.ReadFile(filepath.Join("..", "..", "migrations", "043_runtime_movement_reconciliation_profile.sql"))
	if err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{
		"grounded_transition_deadzone_min FLOAT NOT NULL DEFAULT 0",
		"leap_landing_clamp_ignore_deadzone FLOAT NOT NULL DEFAULT 0",
		"leap_landing_soft_snap_deadzone FLOAT NOT NULL DEFAULT 0",
		"dodge_carry_handoff_ms INT NOT NULL DEFAULT 0",
		"leap_landing_correction_grace_ms INT NOT NULL DEFAULT 0",
		"leap_grounded_carry_handoff_ms INT NOT NULL DEFAULT 0",
	} {
		if strings.Contains(string(migration), forbidden) {
			t.Fatalf("runtime movement migration keeps client-fallback default: %s", forbidden)
		}
	}
}

func TestBootstrapSeedsCoverWolfMaulCounterRuntime(t *testing.T) {
	sql := readBootstrapSQL(t)
	required := []string{
		"'maul','Maul','Pressure counter",
		"'motion_wolf_maul_lateral_counter_v1'",
		"'hitbox_maul_0','maul',0,'temporal_sweep'",
		"'wolf_maul_lateral_counter_v1','grounded_skill'",
		"920,520,220,420,690",
		`"drag_target_until_release":true`,
		`"randomize_lateral_side":true`,
		"('maul','wolf_maul_lateral_counter_v1'",
		"'wolf_maul_pressure_counter_v1','contract_wolf_pack_harasser_v1','maul'",
	}
	for _, fragment := range required {
		if !strings.Contains(sql, fragment) {
			t.Fatalf("wolf maul counter runtime fragment missing: %s", fragment)
		}
	}
}

func TestBootstrapSeedsCoverWolfBehaviorOpportunityRuntime(t *testing.T) {
	sql := readBootstrapSQL(t)
	required := []string{
		"'opportunity_wolf_harasser_v1'",
		"commit_angle_max_deg",
		"'orbit_wolf_harasser_combat_walk_v1'",
		"'combat_walk'",
		"'wolf_lunge_approach_acquire_v1'",
		"'wolf_lunge_circle_reposition_v1'",
		"'wolf_bite_approach_acquire_v1'",
		"'wolf_bite_circle_reposition_v1'",
		"'wolf_dodge_pressure_evasion_v1'",
		"'observe_only'",
		`"candidate_skills_diagnostics":true`,
		`"desiredRangeCm":560`,
		`"chaseRangeCm":860`,
		`"retreatRangeCm":340`,
		`"orbitSpeedCmS":150`,
		`"chaseSpeedCmS":310`,
		`"lungeSpeedCmS":380`,
		`"maulSpeedCmS":345`,
		`"retreatSpeedCmS":260`,
		`"repeatSkillPenaltyMultiplier":0.35`,
		`"dodgeUnderPressure":true`,
		`"maulCounterUnderPressure":true`,
		`"maulCounterChance":0.30`,
		`"dodgeRetreatMultiplier":0.70`,
		`"globalDodgeMultiplier":0.85`,
		"'wolf_evasion_pressure_v1'",
		"0.72,\n    0.28,\n    0.42,",
		`"commitThreatWeight":0.28`,
		`"closingThreatWeight":0.18`,
		`"defensiveBiteWeight":0.14`,
		`"fleeingLungeWeight":0.20`,
		`"lowResourceRiskFloor":0.16`,
		`"dodgeCommittedThreatMultiplier":1.12`,
		`"vulnerableBiteMultiplier":1.16`,
		`"vulnerableMaulMultiplier":1.16`,
		`"tacticalDestinationDistanceCm":180`,
		`"dodgeCostMultiplier":0.50`,
		`"regenPerSecond":12`,
		`"zeroStaminaLockoutUntilFull":true`,
	}
	for _, fragment := range required {
		if !strings.Contains(sql, fragment) {
			t.Fatalf("wolf behavior opportunity runtime fragment missing: %s", fragment)
		}
	}
}

func readBootstrapSQL(t *testing.T) string {
	t.Helper()
	files, err := loadSQLFiles(filepath.Join("..", "..", "bootstrap"))
	if err != nil {
		t.Fatal(err)
	}
	var builder strings.Builder
	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			t.Fatal(err)
		}
		builder.Write(content)
		builder.WriteByte('\n')
	}
	return builder.String()
}

type runtimeSeedRequirement struct {
	category    string
	id          string
	description string
	fragments   []string
}

type skillActionSeedRequirement struct {
	skillID          string
	actionContractID string
	timing           string
	hitbox           string
	motionProfileID  string
	damageGroupID    string
	binding          string
}
