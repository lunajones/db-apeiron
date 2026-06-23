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
		"wolf_lunge_airborne_v1",
		"wolf_dodge_lateral_leap_v1",
		"wolf_maul_lateral_counter_v1",
	} {
		if !strings.Contains(sql, "'"+contractID+"'") {
			t.Fatalf("bootstrap is missing required runtime contract %s", contractID)
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
		"lunge":                 "wolf_lunge_airborne_v1",
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
	requiredLegacyRows := map[string]string{
		"lunge":              "wolf_lunge_airborne_v1",
		"player_shield_rush": "shield_rush_front_contact_v1",
		"player_shield_bash": "shield_bash_front_push_v1",
	}
	for skillID, canonicalContractID := range requiredLegacyRows {
		if !strings.Contains(sql, "'"+skillID+"'") {
			t.Fatalf("legacy movement compatibility row missing skill %s", skillID)
		}
		if !strings.Contains(sql, `"prefer":"movement_action_contract"`) {
			t.Fatal("legacy skill movement seed must declare movement_action_contract as preferred authority")
		}
		if !strings.Contains(sql, "'"+canonicalContractID+"'") {
			t.Fatalf("canonical movement action contract missing for legacy skill %s -> %s", skillID, canonicalContractID)
		}
		if !strings.Contains(sql, "('"+skillID+"','"+canonicalContractID+"'") {
			t.Fatalf("legacy skill %s does not have canonical skill_movement_action_binding -> %s", skillID, canonicalContractID)
		}
	}
}

func TestSkillDataProtoDocumentsLegacyMovementEndpoint(t *testing.T) {
	proto, err := os.ReadFile(filepath.Join("..", "..", "proto", "apeiron", "v1", "skill_data_service.proto"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(proto)
	for _, required := range []string{
		"Legacy compatibility endpoint",
		"GetSkillMovementActionBinding",
		"GetMovementActionContract",
		"must not be used to tune or resolve root motion",
	} {
		if !strings.Contains(content, required) {
			t.Fatalf("skill_data_service.proto is missing legacy endpoint documentation fragment %q", required)
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
		`"front_contact_offset_cm":45`,
		"('motion_player_shield_rush_front_contact_v1',0,0.00,'capsule_strip',45,0,100,190,0,160,96,105",
		"('hitbox_player_shield_rush_0','player_shield_rush',0,'temporal_sweep'",
		"'motion_player_shield_rush_front_contact_v1','player_shield_rush_front_contact'",
	}
	for _, fragment := range required {
		if !strings.Contains(sql, fragment) {
			t.Fatalf("shield rush front-contact geometry fragment missing: %s", fragment)
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
		"('hitbox_lunge_0','lunge',0,'temporal_sweep',3600,4030,130,0,105,100,0,120,50,320",
		"('hitbox_maul_0','maul',0,'temporal_sweep',180,440,80,0,100,0,0,130,62,170",
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
		`"repeatSkillPenaltyMultiplier":0.65`,
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
