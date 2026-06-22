package postgres

import (
	"context"
	"time"
)

type CreatureBehaviorRuntimeContract struct {
	ID                  string
	CreatureTemplateID  string
	Description         string
	AggressionCurveJSON string
	RangePolicyJSON     string
	OrbitPolicyJSON     string
	PressurePolicyJSON  string
	StaminaPolicyJSON   string
	MetadataJSON        string
	CreatedAt           time.Time
	UpdatedAt           time.Time
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
		&contract.CreatedAt,
		&contract.UpdatedAt,
	)
	return contract, err
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
