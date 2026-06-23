package handlers

import (
	"database/sql"
	"testing"

	"db-apeiron/internal/repository/postgres"
)

func TestMapSkillImpactProfileCarriesControlEffects(t *testing.T) {
	out := mapSkillImpactProfile(postgres.SkillImpactProfile{
		SkillID:                 "player_shield_rush",
		ImpactType:              "heavy",
		HitReaction:             "stagger",
		PoiseDamage:             20,
		StaggerPower:            0.4,
		InterruptPower:          0.8,
		GuardDamageMultiplier:   1.2,
		HitstopMS:               180,
		AppliesStatusEffect:     true,
		StatusEffectID:          sql.NullString{String: "shield_rush_stagger", Valid: true},
		StatusEffectChance:      1,
		ControlType:             "carry_push",
		ControlEffectDurationMS: 430,
		ControlReleasePolicyID:  "multi_target_carry_push_forward_release",
	})

	if len(out.GetControlEffects()) != 1 {
		t.Fatalf("control effects = %d", len(out.GetControlEffects()))
	}
	effect := out.GetControlEffects()[0]
	if effect.GetStatusEffectId() != "shield_rush_stagger" {
		t.Fatalf("status effect id = %q", effect.GetStatusEffectId())
	}
	if effect.GetControlType() != "carry_push" {
		t.Fatalf("control type = %q", effect.GetControlType())
	}
	if effect.GetDurationMs() != 430 {
		t.Fatalf("duration = %d", effect.GetDurationMs())
	}
	if effect.GetReleasePolicyId() != "multi_target_carry_push_forward_release" {
		t.Fatalf("release policy = %q", effect.GetReleasePolicyId())
	}
}
