package postgres

import (
	"context"
	"time"

	"db-apeiron/internal/database"
)

type ProfileRepository struct {
	db database.TxManager
}

func NewProfileRepository(db database.TxManager) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (r *ProfileRepository) GetSpawnProfileByID(ctx context.Context, id string) (SpawnProfile, error) {
	var p SpawnProfile

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			allowed_biomes,
			biome_weights,
			allowed_regions,
			region_weights,
			density_base,
			density_cap,
			respawn_time_seconds,
			spawn_batch_size,
			spawn_variance,
			territorial_bias,
			migration_allowed,
			clustering_strength,
			is_dynamic,
			anti_overcrowd,
			pvp_zone_modifier,
			created_at,
			updated_at
		FROM apeiron.spawn_profile
		WHERE id = $1
	`, id).Scan(
		&p.ID,
		&p.AllowedBiomes,
		&p.BiomeWeights,
		&p.AllowedRegions,
		&p.RegionWeights,
		&p.DensityBase,
		&p.DensityCap,
		&p.RespawnTimeSeconds,
		&p.SpawnBatchSize,
		&p.SpawnVariance,
		&p.TerritorialBias,
		&p.MigrationAllowed,
		&p.ClusteringStrength,
		&p.IsDynamic,
		&p.AntiOvercrowd,
		&p.PVPZoneModifier,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	return p, err
}

func (r *ProfileRepository) GetMovementProfileByID(ctx context.Context, id string) (MovementProfile, error) {
	var p MovementProfile

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			max_speed,
			acceleration,
			deceleration,
			friction,
			gravity_multiplier,
			mass,
			momentum_retention,
			turn_rate,
			air_control,
			strafe_efficiency,
			backpedal_penalty,
			dodge_distance,
			dodge_duration_ms,
			sprint_multiplier,
			slope_limit,
			slide_on_slope,
			is_airborne_enabled,
			created_at,
			updated_at
		FROM apeiron.movement_profile
		WHERE id = $1
	`, id).Scan(
		&p.ID,
		&p.MaxSpeed,
		&p.Acceleration,
		&p.Deceleration,
		&p.Friction,
		&p.GravityMultiplier,
		&p.Mass,
		&p.MomentumRetention,
		&p.TurnRate,
		&p.AirControl,
		&p.StrafeEfficiency,
		&p.BackpedalPenalty,
		&p.DodgeDistance,
		&p.DodgeDurationMS,
		&p.SprintMultiplier,
		&p.SlopeLimit,
		&p.SlideOnSlope,
		&p.IsAirborneEnabled,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	return p, err
}

func (r *ProfileRepository) GetCombatCoreProfileByID(ctx context.Context, id string) (CombatCoreProfile, error) {
	var p CombatCoreProfile

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			physical_defense,
			magic_defense,
			critical_chance,
			critical_multiplier,
			damage_taken_multiplier,
			damage_dealt_multiplier,
			max_stamina,
			stamina_regen_per_sec,
			dodge_stamina_cost,
			sprint_stamina_cost_per_sec,
			block_stamina_cost_per_sec,
			attack_stamina_cost,
			stamina_exhaustion_threshold,
			stamina_zero_regen_multiplier,
			max_posture,
			posture_recovery_rate,
			posture_damage_multiplier,
			posture_break_duration_ms,
			block_damage_reduction,
			parry_window_ms,
			parry_reward_multiplier,
			can_block,
			can_parry,
			dodge_iframe_ms,
			dodge_cooldown_ms,
			stun_resistance,
			root_resistance,
			knockback_resistance,
			cc_duration_multiplier,
			is_boss,
			is_pvp_immune,
			created_at,
			updated_at
		FROM apeiron.combat_core_profile
		WHERE id = $1
	`, id).Scan(
		&p.ID,
		&p.PhysicalDefense,
		&p.MagicDefense,
		&p.CriticalChance,
		&p.CriticalMultiplier,
		&p.DamageTakenMultiplier,
		&p.DamageDealtMultiplier,
		&p.MaxStamina,
		&p.StaminaRegenPerSec,
		&p.DodgeStaminaCost,
		&p.SprintStaminaCostPerSec,
		&p.BlockStaminaCostPerSec,
		&p.AttackStaminaCost,
		&p.StaminaExhaustionThreshold,
		&p.StaminaZeroRegenMultiplier,
		&p.MaxPosture,
		&p.PostureRecoveryRate,
		&p.PostureDamageMultiplier,
		&p.PostureBreakDurationMS,
		&p.BlockDamageReduction,
		&p.ParryWindowMS,
		&p.ParryRewardMultiplier,
		&p.CanBlock,
		&p.CanParry,
		&p.DodgeIframeMS,
		&p.DodgeCooldownMS,
		&p.StunResistance,
		&p.RootResistance,
		&p.KnockbackResistance,
		&p.CCDurationMultiplier,
		&p.IsBoss,
		&p.IsPVPImmune,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	return p, err
}

func (r *ProfileRepository) GetCombatDefenseContractByID(ctx context.Context, id string) (CombatDefenseContract, error) {
	var c CombatDefenseContract

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			name,
			description,
			defense_type,
			frontal_arc_deg,
			defender_margin_left_ratio,
			defender_margin_right_ratio,
			stamina_damage_only_on_block,
			health_damage_on_unblocked_hit,
			posture_damage_on_block,
			perfect_block_window_ms,
			parry_window_ms,
			guard_damage_multiplier,
			block_stamina_drain_per_second,
			metadata::text,
			created_at,
			updated_at
		FROM apeiron.combat_defense_contract
		WHERE id = $1
	`, id).Scan(
		&c.ID,
		&c.Name,
		&c.Description,
		&c.DefenseType,
		&c.FrontalArcDeg,
		&c.DefenderMarginLeftRatio,
		&c.DefenderMarginRightRatio,
		&c.StaminaDamageOnlyOnBlock,
		&c.HealthDamageOnUnblockedHit,
		&c.PostureDamageOnBlock,
		&c.PerfectBlockWindowMS,
		&c.ParryWindowMS,
		&c.GuardDamageMultiplier,
		&c.BlockStaminaDrainPerSecond,
		&c.MetadataJSON,
		&c.CreatedAt,
		&c.UpdatedAt,
	)

	return c, err
}

func (r *ProfileRepository) GetCombatStyleProfileByID(ctx context.Context, id string) (CombatStyleProfile, error) {
	var p CombatStyleProfile

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			archetype,
			aggression,
			defense,
			patience,
			risk_tolerance,
			preferred_range,
			chase_tendency,
			disengage_threshold,
			combo_preference,
			feint_usage,
			punish_window_awareness,
			aggression_spike_chance,
			dodge_frequency,
			block_frequency,
			parry_aggressiveness,
			panic_threshold,
			strafe_usage,
			circle_target,
			backstep_usage,
			reposition_frequency,
			randomness,
			consistency,
			target_switching,
			focus_fire,
			retaliation_bias,
			is_elite,
			is_coward,
			is_duelist,
			created_at,
			updated_at
		FROM apeiron.combat_style_profile
		WHERE id = $1
	`, id).Scan(
		&p.ID,
		&p.Archetype,
		&p.Aggression,
		&p.Defense,
		&p.Patience,
		&p.RiskTolerance,
		&p.PreferredRange,
		&p.ChaseTendency,
		&p.DisengageThreshold,
		&p.ComboPreference,
		&p.FeintUsage,
		&p.PunishWindowAwareness,
		&p.AggressionSpikeChance,
		&p.DodgeFrequency,
		&p.BlockFrequency,
		&p.ParryAggressiveness,
		&p.PanicThreshold,
		&p.StrafeUsage,
		&p.CircleTarget,
		&p.BackstepUsage,
		&p.RepositionFrequency,
		&p.Randomness,
		&p.Consistency,
		&p.TargetSwitching,
		&p.FocusFire,
		&p.RetaliationBias,
		&p.IsElite,
		&p.IsCoward,
		&p.IsDuelist,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	return p, err
}

func (r *ProfileRepository) GetNeedsProfileByID(ctx context.Context, id string) (NeedsProfile, error) {
	var p NeedsProfile

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			hunger_enabled,
			thirst_enabled,
			fatigue_enabled,
			hunger_decay_per_hour,
			thirst_decay_per_hour,
			fatigue_decay_per_hour,
			hunger_threshold,
			thirst_threshold,
			fatigue_threshold,
			hunger_damage_threshold,
			stamina_regen_penalty_at_low_needs,
			movement_penalty_at_low_needs,
			combat_performance_penalty,
			food_saturation_rate,
			water_saturation_rate,
			rest_recovery_rate,
			panic_threshold,
			aggression_increase_when_starving,
			fear_increase_when_starving,
			needs_enabled,
			created_at,
			updated_at
		FROM apeiron.needs_profile
		WHERE id = $1
	`, id).Scan(
		&p.ID,
		&p.HungerEnabled,
		&p.ThirstEnabled,
		&p.FatigueEnabled,
		&p.HungerDecayPerHour,
		&p.ThirstDecayPerHour,
		&p.FatigueDecayPerHour,
		&p.HungerThreshold,
		&p.ThirstThreshold,
		&p.FatigueThreshold,
		&p.HungerDamageThreshold,
		&p.StaminaRegenPenaltyAtLowNeeds,
		&p.MovementPenaltyAtLowNeeds,
		&p.CombatPerformancePenalty,
		&p.FoodSaturationRate,
		&p.WaterSaturationRate,
		&p.RestRecoveryRate,
		&p.PanicThreshold,
		&p.AggressionIncreaseWhenStarving,
		&p.FearIncreaseWhenStarving,
		&p.NeedsEnabled,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	return p, err
}

func (r *ProfileRepository) GetAIPersonalityProfileByID(ctx context.Context, id string) (AIPersonalityProfile, error) {
	var p AIPersonalityProfile

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			courage,
			curiosity,
			discipline,
			aggression_baseline,
			fear_sensitivity,
			dominance,
			submission,
			loyalty,
			empathy,
			temperament_stability,
			adaptability,
			predictability,
			is_pack_animal,
			is_solo,
			is_predator,
			created_at,
			updated_at
		FROM apeiron.ai_personality_profile
		WHERE id = $1
	`, id).Scan(
		&p.ID,
		&p.Courage,
		&p.Curiosity,
		&p.Discipline,
		&p.AggressionBaseline,
		&p.FearSensitivity,
		&p.Dominance,
		&p.Submission,
		&p.Loyalty,
		&p.Empathy,
		&p.TemperamentStability,
		&p.Adaptability,
		&p.Predictability,
		&p.IsPackAnimal,
		&p.IsSolo,
		&p.IsPredator,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	return p, err
}

func (r *ProfileRepository) GetAIDecisionProfileByID(ctx context.Context, id string) (AIDecisionProfile, error) {
	var p AIDecisionProfile

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			reaction_time_ms,
			decision_interval_ms,
			combat_engage_delay_ms,
			input_delay_variance_ms,
			target_switch_delay_ms,
			target_priority_bias,
			threat_evaluation_speed,
			combo_interrupt_tolerance,
			greed_vs_safety,
			punish_recognition_speed,
			mistake_chance,
			hesitation_factor,
			overcommit_risk,
			assist_ally_priority,
			pack_sync_factor,
			focus_fire_coordination,
			is_predictive,
			is_adaptive,
			created_at,
			updated_at
		FROM apeiron.ai_decision_profile
		WHERE id = $1
	`, id).Scan(
		&p.ID,
		&p.ReactionTimeMS,
		&p.DecisionIntervalMS,
		&p.CombatEngageDelayMS,
		&p.InputDelayVarianceMS,
		&p.TargetSwitchDelayMS,
		&p.TargetPriorityBias,
		&p.ThreatEvaluationSpeed,
		&p.ComboInterruptTolerance,
		&p.GreedVsSafety,
		&p.PunishRecognitionSpeed,
		&p.MistakeChance,
		&p.HesitationFactor,
		&p.OvercommitRisk,
		&p.AssistAllyPriority,
		&p.PackSyncFactor,
		&p.FocusFireCoordination,
		&p.IsPredictive,
		&p.IsAdaptive,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	return p, err
}

func (r *ProfileRepository) GetSensoryProfileByID(ctx context.Context, id string) (SensoryProfile, error) {
	var p SensoryProfile

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			vision_range,
			vision_angle,
			night_vision_modifier,
			can_detect_stealth,
			stealth_detection_modifier,
			hearing_range,
			noise_sensitivity,
			smell_range,
			blood_detection_range,
			tracking_persistence_ms,
			target_memory_ms,
			last_known_position_memory_ms,
			alertness_gain_rate,
			alertness_decay_rate,
			surprise_threshold,
			created_at,
			updated_at
		FROM apeiron.sensory_profile
		WHERE id = $1
	`, id).Scan(
		&p.ID,
		&p.VisionRange,
		&p.VisionAngle,
		&p.NightVisionModifier,
		&p.CanDetectStealth,
		&p.StealthDetectionModifier,
		&p.HearingRange,
		&p.NoiseSensitivity,
		&p.SmellRange,
		&p.BloodDetectionRange,
		&p.TrackingPersistenceMS,
		&p.TargetMemoryMS,
		&p.LastKnownPositionMemoryMS,
		&p.AlertnessGainRate,
		&p.AlertnessDecayRate,
		&p.SurpriseThreshold,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	return p, err
}

type SpawnProfile struct {
	ID string

	AllowedBiomes  []byte
	BiomeWeights   []byte
	AllowedRegions []byte
	RegionWeights  []byte

	DensityBase float64
	DensityCap  *int

	RespawnTimeSeconds int
	SpawnBatchSize     int
	SpawnVariance      float64

	TerritorialBias    float64
	MigrationAllowed   bool
	ClusteringStrength float64

	IsDynamic       bool
	AntiOvercrowd   bool
	PVPZoneModifier float64

	CreatedAt time.Time
	UpdatedAt time.Time
}

type MovementProfile struct {
	ID string

	MaxSpeed          float64
	Acceleration      float64
	Deceleration      float64
	Friction          float64
	GravityMultiplier float64
	Mass              float64
	MomentumRetention float64

	TurnRate         float64
	AirControl       float64
	StrafeEfficiency float64
	BackpedalPenalty float64

	DodgeDistance   float64
	DodgeDurationMS int

	SprintMultiplier float64

	SlopeLimit   float64
	SlideOnSlope bool

	IsAirborneEnabled bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

type CombatCoreProfile struct {
	ID string

	PhysicalDefense float64
	MagicDefense    float64

	CriticalChance     float64
	CriticalMultiplier float64

	DamageTakenMultiplier float64
	DamageDealtMultiplier float64

	MaxStamina         float64
	StaminaRegenPerSec float64

	DodgeStaminaCost        float64
	SprintStaminaCostPerSec float64
	BlockStaminaCostPerSec  float64
	AttackStaminaCost       float64

	StaminaExhaustionThreshold float64
	StaminaZeroRegenMultiplier float64

	MaxPosture              float64
	PostureRecoveryRate     float64
	PostureDamageMultiplier float64
	PostureBreakDurationMS  int

	BlockDamageReduction  float64
	ParryWindowMS         int
	ParryRewardMultiplier float64

	CanBlock bool
	CanParry bool

	DodgeIframeMS   int
	DodgeCooldownMS int

	StunResistance       float64
	RootResistance       float64
	KnockbackResistance  float64
	CCDurationMultiplier float64

	IsBoss      bool
	IsPVPImmune bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

type CombatDefenseContract struct {
	ID string

	Name        string
	Description string
	DefenseType string

	FrontalArcDeg            float64
	DefenderMarginLeftRatio  float64
	DefenderMarginRightRatio float64

	StaminaDamageOnlyOnBlock   bool
	HealthDamageOnUnblockedHit bool
	PostureDamageOnBlock       bool

	PerfectBlockWindowMS int
	ParryWindowMS        int

	GuardDamageMultiplier      float64
	BlockStaminaDrainPerSecond float64
	MetadataJSON               string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type CombatStyleProfile struct {
	ID string

	Archetype string

	Aggression    float64
	Defense       float64
	Patience      float64
	RiskTolerance float64

	PreferredRange     float64
	ChaseTendency      float64
	DisengageThreshold float64

	ComboPreference       float64
	FeintUsage            float64
	PunishWindowAwareness float64
	AggressionSpikeChance float64

	DodgeFrequency      float64
	BlockFrequency      float64
	ParryAggressiveness float64
	PanicThreshold      float64

	StrafeUsage         float64
	CircleTarget        float64
	BackstepUsage       float64
	RepositionFrequency float64

	Randomness  float64
	Consistency float64

	TargetSwitching float64
	FocusFire       float64
	RetaliationBias float64

	IsElite   bool
	IsCoward  bool
	IsDuelist bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

type NeedsProfile struct {
	ID string

	HungerEnabled  bool
	ThirstEnabled  bool
	FatigueEnabled bool

	HungerDecayPerHour  float64
	ThirstDecayPerHour  float64
	FatigueDecayPerHour float64

	HungerThreshold  float64
	ThirstThreshold  float64
	FatigueThreshold float64

	HungerDamageThreshold float64

	StaminaRegenPenaltyAtLowNeeds float64
	MovementPenaltyAtLowNeeds     float64
	CombatPerformancePenalty      float64

	FoodSaturationRate  float64
	WaterSaturationRate float64
	RestRecoveryRate    float64

	PanicThreshold                 float64
	AggressionIncreaseWhenStarving float64
	FearIncreaseWhenStarving       float64

	NeedsEnabled bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

type AIPersonalityProfile struct {
	ID string

	Courage            float64
	Curiosity          float64
	Discipline         float64
	AggressionBaseline float64
	FearSensitivity    float64

	Dominance  float64
	Submission float64
	Loyalty    float64
	Empathy    float64

	TemperamentStability float64
	Adaptability         float64
	Predictability       float64

	IsPackAnimal bool
	IsSolo       bool
	IsPredator   bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

type AIDecisionProfile struct {
	ID string

	ReactionTimeMS       int
	DecisionIntervalMS   int
	CombatEngageDelayMS  int
	InputDelayVarianceMS int

	TargetSwitchDelayMS   int
	TargetPriorityBias    float64
	ThreatEvaluationSpeed float64

	ComboInterruptTolerance float64
	GreedVsSafety           float64
	PunishRecognitionSpeed  float64

	MistakeChance    float64
	HesitationFactor float64
	OvercommitRisk   float64

	AssistAllyPriority    float64
	PackSyncFactor        float64
	FocusFireCoordination float64

	IsPredictive bool
	IsAdaptive   bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

type SensoryProfile struct {
	ID string

	VisionRange         float64
	VisionAngle         float64
	NightVisionModifier float64

	CanDetectStealth         bool
	StealthDetectionModifier float64

	HearingRange     float64
	NoiseSensitivity float64

	SmellRange            float64
	BloodDetectionRange   float64
	TrackingPersistenceMS int

	TargetMemoryMS            int
	LastKnownPositionMemoryMS int

	AlertnessGainRate  float64
	AlertnessDecayRate float64
	SurpriseThreshold  float64

	CreatedAt time.Time
	UpdatedAt time.Time
}
