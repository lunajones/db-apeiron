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
		"('mode_sword_shield_vanguard','M1','player_basic_attack_1'",
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
