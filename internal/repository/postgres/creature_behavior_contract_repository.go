package postgres

import (
	"context"
	"time"
)

type CreatureBehaviorRuntimeContract struct {
	ID                        string
	CreatureTemplateID        string
	Description               string
	AggressionCurveJSON       string
	RangePolicyJSON           string
	OrbitPolicyJSON           string
	PressurePolicyJSON        string
	StaminaPolicyJSON         string
	MetadataJSON              string
	TargetOpportunityPolicyID string
	OrbitPolicyID             string
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
}

type CreatureEvasionPolicy struct {
	ID                      string
	BehaviorContractID      string
	Description             string
	DodgeSkillID            string
	MaxChainCount           int
	StaminaCostMultiplier   float64
	RetreatChanceMultiplier float64
	LateralBias             float64
	BackstepBias            float64
	PressureThreshold       float64
	CooldownMS              int
	MetadataJSON            string
	CreatedAt               time.Time
	UpdatedAt               time.Time
}

type CreatureSkillSetupPolicy struct {
	ID                  string
	BehaviorContractID  string
	SkillID             string
	SetupType           string
	MinSetupMS          int
	MaxSetupMS          int
	CommitDistanceCM    float64
	PreferredMinRangeCM float64
	PreferredMaxRangeCM float64
	MovementTactic      string
	LockSideDuringSetup bool
	IsEnabled           bool
	MetadataJSON        string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type CreatureTargetOpportunityPolicy struct {
	ID                          string
	Description                 string
	CommitAngleMaxDeg           float64
	MinCommitDistanceCM         float64
	MaxCommitDistanceCM         float64
	ApproachMinDistanceCM       float64
	ApproachMaxDistanceCM       float64
	BiteRangeCM                 float64
	LungeMinRangeCM             float64
	LungeMaxRangeCM             float64
	MaulPressureThreshold       float64
	TargetMemoryMS              int
	NoReadySkillMemoryPolicy    string
	CandidateCooldownVisibility bool
	AllowBacksideCommit         bool
	MetadataJSON                string
	CreatedAt                   time.Time
	UpdatedAt                   time.Time
}

type CreatureOrbitPolicy struct {
	ID                             string
	BehaviorContractID             string
	Description                    string
	OrbitLocomotionMode            string
	OrbitSpeedScale                float64
	MinOrbitDurationMS             int
	SideSwitchCooldownMS           int
	AllowSideSwitchWhenTargetFaces bool
	PreferLongSideCommit           bool
	SideFlipChanceMultiplier       float64
	LockSideDuringSetup            bool
	MetadataJSON                   string
	CreatedAt                      time.Time
	UpdatedAt                      time.Time
}

type CreatureSkillBehaviorBinding struct {
	ID                  string
	BehaviorContractID  string
	SkillID             string
	TacticalState       string
	DecisionPhase       string
	SetupPolicyID       string
	MinRangeCM          float64
	MaxRangeCM          float64
	Priority            int
	UsageWeight         float64
	CooldownGroup       string
	RequiresLineOfSight bool
	IsEnabled           bool
	MetadataJSON        string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func (r *ProfileRepository) GetCreatureBehaviorRuntimeContractByID(ctx context.Context, id string) (CreatureBehaviorRuntimeContract, error) {
	var contract CreatureBehaviorRuntimeContract
	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			COALESCE(creature_template_id, ''),
			description,
			aggression_curve::TEXT,
			range_policy::TEXT,
			orbit_policy::TEXT,
			pressure_policy::TEXT,
			stamina_policy::TEXT,
			metadata::TEXT,
			COALESCE(target_opportunity_policy_id, ''),
			COALESCE(orbit_policy_id, ''),
			created_at,
			updated_at
		FROM apeiron.creature_behavior_runtime_contract
		WHERE id = $1
	`, id).Scan(
		&contract.ID,
		&contract.CreatureTemplateID,
		&contract.Description,
		&contract.AggressionCurveJSON,
		&contract.RangePolicyJSON,
		&contract.OrbitPolicyJSON,
		&contract.PressurePolicyJSON,
		&contract.StaminaPolicyJSON,
		&contract.MetadataJSON,
		&contract.TargetOpportunityPolicyID,
		&contract.OrbitPolicyID,
		&contract.CreatedAt,
		&contract.UpdatedAt,
	)
	return contract, err
}

func (r *ProfileRepository) GetCreatureBehaviorRuntimeContractByTemplateID(ctx context.Context, templateID string) (CreatureBehaviorRuntimeContract, error) {
	var contract CreatureBehaviorRuntimeContract
	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			COALESCE(creature_template_id, ''),
			description,
			aggression_curve::TEXT,
			range_policy::TEXT,
			orbit_policy::TEXT,
			pressure_policy::TEXT,
			stamina_policy::TEXT,
			metadata::TEXT,
			COALESCE(target_opportunity_policy_id, ''),
			COALESCE(orbit_policy_id, ''),
			created_at,
			updated_at
		FROM apeiron.creature_behavior_runtime_contract
		WHERE creature_template_id = $1
		ORDER BY id ASC
		LIMIT 1
	`, templateID).Scan(
		&contract.ID,
		&contract.CreatureTemplateID,
		&contract.Description,
		&contract.AggressionCurveJSON,
		&contract.RangePolicyJSON,
		&contract.OrbitPolicyJSON,
		&contract.PressurePolicyJSON,
		&contract.StaminaPolicyJSON,
		&contract.MetadataJSON,
		&contract.TargetOpportunityPolicyID,
		&contract.OrbitPolicyID,
		&contract.CreatedAt,
		&contract.UpdatedAt,
	)
	return contract, err
}

func (r *ProfileRepository) GetCreatureTargetOpportunityPolicyByID(ctx context.Context, id string) (CreatureTargetOpportunityPolicy, error) {
	var policy CreatureTargetOpportunityPolicy
	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			description,
			commit_angle_max_deg,
			min_commit_distance_cm,
			max_commit_distance_cm,
			approach_min_distance_cm,
			approach_max_distance_cm,
			bite_range_cm,
			lunge_min_range_cm,
			lunge_max_range_cm,
			maul_pressure_threshold,
			target_memory_ms,
			no_ready_skill_memory_policy,
			candidate_cooldown_visibility,
			allow_backside_commit,
			metadata::TEXT,
			created_at,
			updated_at
		FROM apeiron.creature_target_opportunity_policy
		WHERE id = $1
	`, id).Scan(
		&policy.ID,
		&policy.Description,
		&policy.CommitAngleMaxDeg,
		&policy.MinCommitDistanceCM,
		&policy.MaxCommitDistanceCM,
		&policy.ApproachMinDistanceCM,
		&policy.ApproachMaxDistanceCM,
		&policy.BiteRangeCM,
		&policy.LungeMinRangeCM,
		&policy.LungeMaxRangeCM,
		&policy.MaulPressureThreshold,
		&policy.TargetMemoryMS,
		&policy.NoReadySkillMemoryPolicy,
		&policy.CandidateCooldownVisibility,
		&policy.AllowBacksideCommit,
		&policy.MetadataJSON,
		&policy.CreatedAt,
		&policy.UpdatedAt,
	)
	return policy, err
}

func (r *ProfileRepository) GetCreatureOrbitPolicyByID(ctx context.Context, id string) (CreatureOrbitPolicy, error) {
	var policy CreatureOrbitPolicy
	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			COALESCE(behavior_contract_id, ''),
			description,
			orbit_locomotion_mode,
			orbit_speed_scale,
			min_orbit_duration_ms,
			side_switch_cooldown_ms,
			allow_side_switch_when_target_faces,
			prefer_long_side_commit,
			side_flip_chance_multiplier,
			lock_side_during_setup,
			metadata::TEXT,
			created_at,
			updated_at
		FROM apeiron.creature_orbit_policy
		WHERE id = $1
	`, id).Scan(
		&policy.ID,
		&policy.BehaviorContractID,
		&policy.Description,
		&policy.OrbitLocomotionMode,
		&policy.OrbitSpeedScale,
		&policy.MinOrbitDurationMS,
		&policy.SideSwitchCooldownMS,
		&policy.AllowSideSwitchWhenTargetFaces,
		&policy.PreferLongSideCommit,
		&policy.SideFlipChanceMultiplier,
		&policy.LockSideDuringSetup,
		&policy.MetadataJSON,
		&policy.CreatedAt,
		&policy.UpdatedAt,
	)
	return policy, err
}

func (r *ProfileRepository) GetCreatureSkillBehaviorBindingsByBehaviorContractID(ctx context.Context, behaviorContractID string) ([]CreatureSkillBehaviorBinding, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			id,
			behavior_contract_id,
			skill_id,
			tactical_state,
			decision_phase,
			COALESCE(setup_policy_id, ''),
			min_range_cm,
			max_range_cm,
			priority,
			usage_weight,
			cooldown_group,
			requires_line_of_sight,
			is_enabled,
			metadata::TEXT,
			created_at,
			updated_at
		FROM apeiron.creature_skill_behavior_binding
		WHERE behavior_contract_id = $1
		ORDER BY tactical_state ASC, decision_phase ASC, priority DESC, id ASC
	`, behaviorContractID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []CreatureSkillBehaviorBinding
	for rows.Next() {
		var binding CreatureSkillBehaviorBinding
		if err := rows.Scan(
			&binding.ID,
			&binding.BehaviorContractID,
			&binding.SkillID,
			&binding.TacticalState,
			&binding.DecisionPhase,
			&binding.SetupPolicyID,
			&binding.MinRangeCM,
			&binding.MaxRangeCM,
			&binding.Priority,
			&binding.UsageWeight,
			&binding.CooldownGroup,
			&binding.RequiresLineOfSight,
			&binding.IsEnabled,
			&binding.MetadataJSON,
			&binding.CreatedAt,
			&binding.UpdatedAt,
		); err != nil {
			return nil, err
		}
		out = append(out, binding)
	}
	return out, rows.Err()
}

func (r *ProfileRepository) GetCreatureEvasionPoliciesByBehaviorContractID(ctx context.Context, behaviorContractID string) ([]CreatureEvasionPolicy, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			id,
			behavior_contract_id,
			description,
			COALESCE(dodge_skill_id, ''),
			max_chain_count,
			stamina_cost_multiplier,
			retreat_chance_multiplier,
			lateral_bias,
			backstep_bias,
			pressure_threshold,
			cooldown_ms,
			metadata::TEXT,
			created_at,
			updated_at
		FROM apeiron.creature_evasion_policy
		WHERE behavior_contract_id = $1
		ORDER BY id ASC
	`, behaviorContractID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []CreatureEvasionPolicy
	for rows.Next() {
		var policy CreatureEvasionPolicy
		if err := rows.Scan(
			&policy.ID,
			&policy.BehaviorContractID,
			&policy.Description,
			&policy.DodgeSkillID,
			&policy.MaxChainCount,
			&policy.StaminaCostMultiplier,
			&policy.RetreatChanceMultiplier,
			&policy.LateralBias,
			&policy.BackstepBias,
			&policy.PressureThreshold,
			&policy.CooldownMS,
			&policy.MetadataJSON,
			&policy.CreatedAt,
			&policy.UpdatedAt,
		); err != nil {
			return nil, err
		}
		out = append(out, policy)
	}
	return out, rows.Err()
}

func (r *ProfileRepository) GetCreatureSkillSetupPoliciesByBehaviorContractID(ctx context.Context, behaviorContractID string) ([]CreatureSkillSetupPolicy, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			id,
			behavior_contract_id,
			skill_id,
			setup_type,
			min_setup_ms,
			max_setup_ms,
			commit_distance_cm,
			preferred_min_range_cm,
			preferred_max_range_cm,
			movement_tactic,
			lock_side_during_setup,
			is_enabled,
			metadata::TEXT,
			created_at,
			updated_at
		FROM apeiron.creature_skill_setup_policy
		WHERE behavior_contract_id = $1
		ORDER BY skill_id ASC, id ASC
	`, behaviorContractID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []CreatureSkillSetupPolicy
	for rows.Next() {
		var policy CreatureSkillSetupPolicy
		if err := rows.Scan(
			&policy.ID,
			&policy.BehaviorContractID,
			&policy.SkillID,
			&policy.SetupType,
			&policy.MinSetupMS,
			&policy.MaxSetupMS,
			&policy.CommitDistanceCM,
			&policy.PreferredMinRangeCM,
			&policy.PreferredMaxRangeCM,
			&policy.MovementTactic,
			&policy.LockSideDuringSetup,
			&policy.IsEnabled,
			&policy.MetadataJSON,
			&policy.CreatedAt,
			&policy.UpdatedAt,
		); err != nil {
			return nil, err
		}
		out = append(out, policy)
	}
	return out, rows.Err()
}
