package handlers

import (
	"testing"

	"db-apeiron/internal/repository/postgres"
)

func TestMapCreatureBehaviorRuntimeContractPreservesPolicyLinks(t *testing.T) {
	out := mapCreatureBehaviorRuntimeContract(postgres.CreatureBehaviorRuntimeContract{
		ID:                        "contract_wolf_pack_harasser_v1",
		CreatureTemplateID:        "steppe_wolf",
		TargetOpportunityPolicyID: "opportunity_wolf_harasser_v1",
		OrbitPolicyID:             "orbit_wolf_harasser_combat_walk_v1",
		AggressionCurveJSON:       `{"state":"recovered"}`,
		RangePolicyJSON:           `{"desiredRangeCm":420}`,
		OrbitPolicyJSON:           `{"side":"stable"}`,
		PressurePolicyJSON:        `{"maulCounter":true}`,
		StaminaPolicyJSON:         `{"usesCreatureStamina":true}`,
	})

	if out.GetId() != "contract_wolf_pack_harasser_v1" {
		t.Fatalf("contract id = %q", out.GetId())
	}
	if out.GetCreatureTemplateId() != "steppe_wolf" {
		t.Fatalf("template id = %q", out.GetCreatureTemplateId())
	}
	if out.GetTargetOpportunityPolicyId() != "opportunity_wolf_harasser_v1" {
		t.Fatalf("target opportunity policy id = %q", out.GetTargetOpportunityPolicyId())
	}
	if out.GetOrbitPolicyId() != "orbit_wolf_harasser_combat_walk_v1" {
		t.Fatalf("orbit policy id = %q", out.GetOrbitPolicyId())
	}
	if out.GetPressurePolicyJson() == "" || out.GetStaminaPolicyJson() == "" {
		t.Fatalf("behavior JSON policies were not preserved: pressure=%q stamina=%q", out.GetPressurePolicyJson(), out.GetStaminaPolicyJson())
	}
}

func TestMapCreatureTargetOpportunityPolicyPreservesRuntimeDecisionFields(t *testing.T) {
	out := mapCreatureTargetOpportunityPolicy(postgres.CreatureTargetOpportunityPolicy{
		ID:                          "opportunity_wolf_harasser_v1",
		CommitAngleMaxDeg:           180,
		MinCommitDistanceCM:         120,
		MaxCommitDistanceCM:         760,
		ApproachMinDistanceCM:       260,
		ApproachMaxDistanceCM:       760,
		BiteRangeCM:                 230,
		LungeMinRangeCM:             180,
		LungeMaxRangeCM:             700,
		MaulPressureThreshold:       0.72,
		TargetMemoryMS:              1200,
		NoReadySkillMemoryPolicy:    "observe_only",
		CandidateCooldownVisibility: true,
		AllowBacksideCommit:         true,
	})

	if out.GetCommitAngleMaxDeg() != 180 {
		t.Fatalf("commit angle = %.0f", out.GetCommitAngleMaxDeg())
	}
	if out.GetLungeMinRangeCm() != 180 || out.GetLungeMaxRangeCm() != 700 {
		t.Fatalf("lunge range = %.0f..%.0f", out.GetLungeMinRangeCm(), out.GetLungeMaxRangeCm())
	}
	if out.GetNoReadySkillMemoryPolicy() != "observe_only" {
		t.Fatalf("memory policy = %q", out.GetNoReadySkillMemoryPolicy())
	}
	if !out.GetCandidateCooldownVisibility() || !out.GetAllowBacksideCommit() {
		t.Fatalf("diagnostic/backside flags lost: diagnostic=%v backside=%v", out.GetCandidateCooldownVisibility(), out.GetAllowBacksideCommit())
	}
}

func TestMapCreatureOrbitPolicyPreservesAntiThrashFields(t *testing.T) {
	out := mapCreatureOrbitPolicy(postgres.CreatureOrbitPolicy{
		ID:                             "orbit_wolf_harasser_combat_walk_v1",
		BehaviorContractID:             "contract_wolf_pack_harasser_v1",
		OrbitLocomotionMode:            "combat_walk",
		OrbitSpeedScale:                0.55,
		MinOrbitDurationMS:             700,
		SideSwitchCooldownMS:           900,
		AllowSideSwitchWhenTargetFaces: true,
		PreferLongSideCommit:           true,
		SideFlipChanceMultiplier:       0.35,
		LockSideDuringSetup:            true,
	})

	if out.GetOrbitLocomotionMode() != "combat_walk" {
		t.Fatalf("orbit locomotion mode = %q", out.GetOrbitLocomotionMode())
	}
	if out.GetMinOrbitDurationMs() != 700 || out.GetSideSwitchCooldownMs() != 900 {
		t.Fatalf("orbit timing = %d/%d", out.GetMinOrbitDurationMs(), out.GetSideSwitchCooldownMs())
	}
	if out.GetSideFlipChanceMultiplier() != 0.35 {
		t.Fatalf("side flip multiplier = %.2f", out.GetSideFlipChanceMultiplier())
	}
	if !out.GetLockSideDuringSetup() {
		t.Fatal("lock side during setup was not preserved")
	}
}

func TestMapCreatureSkillBehaviorBindingPreservesDecisionBinding(t *testing.T) {
	out := mapCreatureSkillBehaviorBinding(postgres.CreatureSkillBehaviorBinding{
		ID:                  "wolf_lunge_circle_reposition_v1",
		BehaviorContractID:  "contract_wolf_pack_harasser_v1",
		SkillID:             "lunge",
		TacticalState:       "circle",
		DecisionPhase:       "reposition",
		SetupPolicyID:       "wolf_lunge_flank_windup_v1",
		MinRangeCM:          180,
		MaxRangeCM:          700,
		Priority:            90,
		UsageWeight:         0.85,
		CooldownGroup:       "wolf_lunge",
		RequiresLineOfSight: true,
		IsEnabled:           true,
	})

	if out.GetSkillId() != "lunge" || out.GetTacticalState() != "circle" || out.GetDecisionPhase() != "reposition" {
		t.Fatalf("binding identity lost: %#v", out)
	}
	if out.GetSetupPolicyId() != "wolf_lunge_flank_windup_v1" {
		t.Fatalf("setup policy id = %q", out.GetSetupPolicyId())
	}
	if out.GetMinRangeCm() != 180 || out.GetMaxRangeCm() != 700 {
		t.Fatalf("binding range = %.0f..%.0f", out.GetMinRangeCm(), out.GetMaxRangeCm())
	}
	if !out.GetRequiresLineOfSight() || !out.GetIsEnabled() {
		t.Fatalf("binding flags lost: los=%v enabled=%v", out.GetRequiresLineOfSight(), out.GetIsEnabled())
	}
}
