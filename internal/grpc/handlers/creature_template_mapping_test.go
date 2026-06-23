package handlers

import (
	"testing"

	"db-apeiron/internal/repository/postgres"
)

func TestMapCreatureTemplateCarriesImpactResponseProfile(t *testing.T) {
	out := mapCreatureTemplate(postgres.CreatureTemplate{
		ID:                    "steppe_wolf",
		Name:                  "Steppe Wolf",
		Faction:               "wildlife",
		Tier:                  1,
		Archetype:             "beast",
		MovementProfileID:     "movement_steppe_wolf",
		CombatCoreProfileID:   "combat_core_steppe_wolf",
		CombatStyleProfileID:  "combat_style_steppe_wolf",
		AIDecisionProfileID:   "ai_decision_steppe_wolf",
		PersonalityProfileID:  "personality_steppe_wolf",
		SensoryProfileID:      "sensory_steppe_wolf",
		NeedsProfileID:        "needs_steppe_wolf",
		SkillSetID:            "skillset_steppe_wolf",
		SpawnProfileID:        "spawn_steppe_wolf",
		ImpactResponseProfile: "creature_flesh_blood_red",
	})

	if out.GetImpactResponseProfile() != "creature_flesh_blood_red" {
		t.Fatalf("impact response profile = %q", out.GetImpactResponseProfile())
	}
}
