package postgres

import (
	"context"
	"database/sql"
	"time"

	"db-apeiron/internal/database"
)

type SkillRepository struct {
	db database.TxManager
}

func NewSkillRepository(db database.TxManager) *SkillRepository {
	return &SkillRepository{db: db}
}

func (r *SkillRepository) GetSkillByID(ctx context.Context, id string) (Skill, error) {
	var s Skill

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			name,
			description,
			archetype,
			skill_type,
			stamina_cost,
			mana_cost,
			health_cost,
			windup_ms,
			active_frames_ms,
			recovery_ms,
			cast_time_ms,
			cooldown_ms,
			cancel_window_start_ms,
			cancel_window_end_ms,
			iframe_start_ms,
			iframe_end_ms,
			min_range,
			max_range,
			cone_angle,
			max_targets,
			target_type,
			requires_target,
			base_damage,
			damage_type,
			elemental_type,
			posture_damage,
			armor_penetration,
			damage_multiplier,
			critical_bonus_multiplier,
			stun_duration_ms,
			root_duration_ms,
			knockback_force,
			movement_multiplier,
			locks_movement,
			movement_distance,
			combo_group,
			combo_index,
			combo_window_ms,
			is_interruptible,
			is_blockable,
			is_parryable,
			ignores_line_of_sight,
			ignores_collision,
			created_at,
			updated_at
		FROM apeiron.skill
		WHERE id = $1
	`, id).Scan(
		&s.ID,
		&s.Name,
		&s.Description,
		&s.Archetype,
		&s.SkillType,
		&s.StaminaCost,
		&s.ManaCost,
		&s.HealthCost,
		&s.WindupMS,
		&s.ActiveFramesMS,
		&s.RecoveryMS,
		&s.CastTimeMS,
		&s.CooldownMS,
		&s.CancelWindowStartMS,
		&s.CancelWindowEndMS,
		&s.IFrameStartMS,
		&s.IFrameEndMS,
		&s.MinRange,
		&s.MaxRange,
		&s.ConeAngle,
		&s.MaxTargets,
		&s.TargetType,
		&s.RequiresTarget,
		&s.BaseDamage,
		&s.DamageType,
		&s.ElementalType,
		&s.PostureDamage,
		&s.ArmorPenetration,
		&s.DamageMultiplier,
		&s.CriticalBonusMultiplier,
		&s.StunDurationMS,
		&s.RootDurationMS,
		&s.KnockbackForce,
		&s.MovementMultiplier,
		&s.LocksMovement,
		&s.MovementDistance,
		&s.ComboGroup,
		&s.ComboIndex,
		&s.ComboWindowMS,
		&s.IsInterruptible,
		&s.IsBlockable,
		&s.IsParryable,
		&s.IgnoresLineOfSight,
		&s.IgnoresCollision,
		&s.CreatedAt,
		&s.UpdatedAt,
	)

	return s, err
}

func (r *SkillRepository) GetSkillSetByID(ctx context.Context, id string) (SkillSet, error) {
	var ss SkillSet

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			name,
			description,
			is_player_usable,
			is_npc_usable,
			created_at,
			updated_at
		FROM apeiron.skill_set
		WHERE id = $1
	`, id).Scan(
		&ss.ID,
		&ss.Name,
		&ss.Description,
		&ss.IsPlayerUsable,
		&ss.IsNPCUsable,
		&ss.CreatedAt,
		&ss.UpdatedAt,
	)

	return ss, err
}

func (r *SkillRepository) GetSlotsBySkillSetID(ctx context.Context, skillSetID string) ([]SkillSlot, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			id,
			skill_set_id,
			skill_id,
			slot_index,
			is_enabled,
			priority,
			usage_weight,
			cooldown_override_ms,
			min_target_hp_percent,
			max_target_hp_percent,
			min_self_hp_percent,
			max_self_hp_percent,
			required_distance_min,
			required_distance_max,
			requires_line_of_sight,
			opener_weight,
			finisher_weight,
			shared_cooldown_group,
			use_only_in_combat,
			created_at,
			updated_at
		FROM apeiron.skill_slot
		WHERE skill_set_id = $1
		ORDER BY slot_index ASC
	`, skillSetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []SkillSlot

	for rows.Next() {
		var slot SkillSlot

		if err := rows.Scan(
			&slot.ID,
			&slot.SkillSetID,
			&slot.SkillID,
			&slot.SlotIndex,
			&slot.IsEnabled,
			&slot.Priority,
			&slot.UsageWeight,
			&slot.CooldownOverrideMS,
			&slot.MinTargetHPPercent,
			&slot.MaxTargetHPPercent,
			&slot.MinSelfHPPercent,
			&slot.MaxSelfHPPercent,
			&slot.RequiredDistanceMin,
			&slot.RequiredDistanceMax,
			&slot.RequiresLineOfSight,
			&slot.OpenerWeight,
			&slot.FinisherWeight,
			&slot.SharedCooldownGroup,
			&slot.UseOnlyInCombat,
			&slot.CreatedAt,
			&slot.UpdatedAt,
		); err != nil {
			return nil, err
		}

		out = append(out, slot)
	}

	return out, rows.Err()
}

func (r *SkillRepository) GetSkillSetLoadout(ctx context.Context, skillSetID string) ([]SkillLoadoutItem, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			sslot.id,
			sslot.skill_set_id,
			sslot.skill_id,
			sslot.slot_index,
			sslot.is_enabled,
			sslot.priority,
			sslot.usage_weight,
			sslot.cooldown_override_ms,
			sslot.min_target_hp_percent,
			sslot.max_target_hp_percent,
			sslot.min_self_hp_percent,
			sslot.max_self_hp_percent,
			sslot.required_distance_min,
			sslot.required_distance_max,
			sslot.requires_line_of_sight,
			sslot.opener_weight,
			sslot.finisher_weight,
			sslot.shared_cooldown_group,
			sslot.use_only_in_combat,
			sslot.created_at,
			sslot.updated_at,

			s.id,
			s.name,
			s.description,
			s.archetype,
			s.skill_type,
			s.stamina_cost,
			s.mana_cost,
			s.health_cost,
			s.windup_ms,
			s.active_frames_ms,
			s.recovery_ms,
			s.cast_time_ms,
			s.cooldown_ms,
			s.cancel_window_start_ms,
			s.cancel_window_end_ms,
			s.iframe_start_ms,
			s.iframe_end_ms,
			s.min_range,
			s.max_range,
			s.cone_angle,
			s.max_targets,
			s.target_type,
			s.requires_target,
			s.base_damage,
			s.damage_type,
			s.elemental_type,
			s.posture_damage,
			s.armor_penetration,
			s.damage_multiplier,
			s.critical_bonus_multiplier,
			s.stun_duration_ms,
			s.root_duration_ms,
			s.knockback_force,
			s.movement_multiplier,
			s.locks_movement,
			s.movement_distance,
			s.combo_group,
			s.combo_index,
			s.combo_window_ms,
			s.is_interruptible,
			s.is_blockable,
			s.is_parryable,
			s.ignores_line_of_sight,
			s.ignores_collision,
			s.created_at,
			s.updated_at
		FROM apeiron.skill_slot sslot
		INNER JOIN apeiron.skill s
			ON s.id = sslot.skill_id
		WHERE sslot.skill_set_id = $1
		ORDER BY sslot.slot_index ASC
	`, skillSetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []SkillLoadoutItem

	for rows.Next() {
		var item SkillLoadoutItem

		if err := rows.Scan(
			&item.Slot.ID,
			&item.Slot.SkillSetID,
			&item.Slot.SkillID,
			&item.Slot.SlotIndex,
			&item.Slot.IsEnabled,
			&item.Slot.Priority,
			&item.Slot.UsageWeight,
			&item.Slot.CooldownOverrideMS,
			&item.Slot.MinTargetHPPercent,
			&item.Slot.MaxTargetHPPercent,
			&item.Slot.MinSelfHPPercent,
			&item.Slot.MaxSelfHPPercent,
			&item.Slot.RequiredDistanceMin,
			&item.Slot.RequiredDistanceMax,
			&item.Slot.RequiresLineOfSight,
			&item.Slot.OpenerWeight,
			&item.Slot.FinisherWeight,
			&item.Slot.SharedCooldownGroup,
			&item.Slot.UseOnlyInCombat,
			&item.Slot.CreatedAt,
			&item.Slot.UpdatedAt,

			&item.Skill.ID,
			&item.Skill.Name,
			&item.Skill.Description,
			&item.Skill.Archetype,
			&item.Skill.SkillType,
			&item.Skill.StaminaCost,
			&item.Skill.ManaCost,
			&item.Skill.HealthCost,
			&item.Skill.WindupMS,
			&item.Skill.ActiveFramesMS,
			&item.Skill.RecoveryMS,
			&item.Skill.CastTimeMS,
			&item.Skill.CooldownMS,
			&item.Skill.CancelWindowStartMS,
			&item.Skill.CancelWindowEndMS,
			&item.Skill.IFrameStartMS,
			&item.Skill.IFrameEndMS,
			&item.Skill.MinRange,
			&item.Skill.MaxRange,
			&item.Skill.ConeAngle,
			&item.Skill.MaxTargets,
			&item.Skill.TargetType,
			&item.Skill.RequiresTarget,
			&item.Skill.BaseDamage,
			&item.Skill.DamageType,
			&item.Skill.ElementalType,
			&item.Skill.PostureDamage,
			&item.Skill.ArmorPenetration,
			&item.Skill.DamageMultiplier,
			&item.Skill.CriticalBonusMultiplier,
			&item.Skill.StunDurationMS,
			&item.Skill.RootDurationMS,
			&item.Skill.KnockbackForce,
			&item.Skill.MovementMultiplier,
			&item.Skill.LocksMovement,
			&item.Skill.MovementDistance,
			&item.Skill.ComboGroup,
			&item.Skill.ComboIndex,
			&item.Skill.ComboWindowMS,
			&item.Skill.IsInterruptible,
			&item.Skill.IsBlockable,
			&item.Skill.IsParryable,
			&item.Skill.IgnoresLineOfSight,
			&item.Skill.IgnoresCollision,
			&item.Skill.CreatedAt,
			&item.Skill.UpdatedAt,
		); err != nil {
			return nil, err
		}

		out = append(out, item)
	}

	return out, rows.Err()
}

func (r *SkillRepository) GetProjectileProfileBySkillID(ctx context.Context, skillID string) (SkillProjectileProfile, error) {
	var p SkillProjectileProfile

	err := r.db.QueryRow(ctx, `
		SELECT
			skill_id,
			trajectory_type,
			projectile_speed,
			projectile_radius,
			max_lifetime_ms,
			gravity_multiplier,
			drag_multiplier,
			collision_mode,
			can_be_blocked,
			can_be_parried,
			can_be_dodged,
			requires_server_confirmation,
			can_pierce,
			max_pierce_count,
			can_home,
			homing_strength,
			homing_turn_rate,
			can_ricochet,
			max_ricochet_count,
			spawn_offset_x,
			spawn_offset_y,
			spawn_offset_z,
			created_at,
			updated_at
		FROM apeiron.skill_projectile_profile
		WHERE skill_id = $1
	`, skillID).Scan(
		&p.SkillID,
		&p.TrajectoryType,
		&p.ProjectileSpeed,
		&p.ProjectileRadius,
		&p.MaxLifetimeMS,
		&p.GravityMultiplier,
		&p.DragMultiplier,
		&p.CollisionMode,
		&p.CanBeBlocked,
		&p.CanBeParried,
		&p.CanBeDodged,
		&p.RequiresServerConfirmation,
		&p.CanPierce,
		&p.MaxPierceCount,
		&p.CanHome,
		&p.HomingStrength,
		&p.HomingTurnRate,
		&p.CanRicochet,
		&p.MaxRicochetCount,
		&p.SpawnOffsetX,
		&p.SpawnOffsetY,
		&p.SpawnOffsetZ,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	return p, err
}

func (r *SkillRepository) GetHitboxProfilesBySkillID(ctx context.Context, skillID string) ([]SkillHitboxProfile, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			id,
			skill_id,
			hitbox_index,
			hitbox_shape,
			hitbox_start_ms,
			hitbox_end_ms,
			offset_x,
			offset_y,
			offset_z,
			size_x,
			size_y,
			size_z,
			radius,
			length,
			angle,
			follows_caster,
			follows_projectile,
			can_multi_hit,
			max_hits_per_target,
			hit_interval_ms,
			friendly_fire,
			created_at,
			updated_at
		FROM apeiron.skill_hitbox_profile
		WHERE skill_id = $1
		ORDER BY hitbox_index ASC
	`, skillID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []SkillHitboxProfile

	for rows.Next() {
		var h SkillHitboxProfile

		if err := rows.Scan(
			&h.ID,
			&h.SkillID,
			&h.HitboxIndex,
			&h.HitboxShape,
			&h.HitboxStartMS,
			&h.HitboxEndMS,
			&h.OffsetX,
			&h.OffsetY,
			&h.OffsetZ,
			&h.SizeX,
			&h.SizeY,
			&h.SizeZ,
			&h.Radius,
			&h.Length,
			&h.Angle,
			&h.FollowsCaster,
			&h.FollowsProjectile,
			&h.CanMultiHit,
			&h.MaxHitsPerTarget,
			&h.HitIntervalMS,
			&h.FriendlyFire,
			&h.CreatedAt,
			&h.UpdatedAt,
		); err != nil {
			return nil, err
		}

		out = append(out, h)
	}

	return out, rows.Err()
}

func (r *SkillRepository) GetAreaEffectProfileBySkillID(ctx context.Context, skillID string) (SkillAreaEffectProfile, error) {
	var a SkillAreaEffectProfile

	err := r.db.QueryRow(ctx, `
		SELECT
			skill_id,
			area_shape,
			radius,
			length,
			width,
			height,
			angle,
			duration_ms,
			tick_interval_ms,
			damage_falloff_type,
			min_falloff_multiplier,
			applies_on_impact,
			persists_after_impact,
			max_targets,
			friendly_fire,
			status_effect_id,
			created_at,
			updated_at
		FROM apeiron.skill_area_effect_profile
		WHERE skill_id = $1
	`, skillID).Scan(
		&a.SkillID,
		&a.AreaShape,
		&a.Radius,
		&a.Length,
		&a.Width,
		&a.Height,
		&a.Angle,
		&a.DurationMS,
		&a.TickIntervalMS,
		&a.DamageFalloffType,
		&a.MinFalloffMultiplier,
		&a.AppliesOnImpact,
		&a.PersistsAfterImpact,
		&a.MaxTargets,
		&a.FriendlyFire,
		&a.StatusEffectID,
		&a.CreatedAt,
		&a.UpdatedAt,
	)

	return a, err
}

func (r *SkillRepository) GetImpactProfileBySkillID(ctx context.Context, skillID string) (SkillImpactProfile, error) {
	var i SkillImpactProfile

	err := r.db.QueryRow(ctx, `
		SELECT
			skill_id,
			impact_type,
			hit_reaction,
			poise_damage,
			stagger_power,
			interrupt_power,
			guard_damage_multiplier,
			bounce_on_shield,
			destroy_on_hit,
			stick_on_hit,
			hitstop_ms,
			screenshake_strength,
			knockback_force,
			knockback_upward_force,
			pull_force,
			applies_status_effect,
			status_effect_id,
			status_effect_chance,
			created_at,
			updated_at
		FROM apeiron.skill_impact_profile
		WHERE skill_id = $1
	`, skillID).Scan(
		&i.SkillID,
		&i.ImpactType,
		&i.HitReaction,
		&i.PoiseDamage,
		&i.StaggerPower,
		&i.InterruptPower,
		&i.GuardDamageMultiplier,
		&i.BounceOnShield,
		&i.DestroyOnHit,
		&i.StickOnHit,
		&i.HitstopMS,
		&i.ScreenshakeStrength,
		&i.KnockbackForce,
		&i.KnockbackUpwardForce,
		&i.PullForce,
		&i.AppliesStatusEffect,
		&i.StatusEffectID,
		&i.StatusEffectChance,
		&i.CreatedAt,
		&i.UpdatedAt,
	)

	return i, err
}

func (r *SkillRepository) GetStatusEffectByID(ctx context.Context, id string) (StatusEffect, error) {
	var s StatusEffect

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			name,
			description,
			effect_type,
			stacking_mode,
			max_stacks,
			duration_ms,
			tick_interval_ms,
			is_dispellable,
			is_pvp_enabled,
			movement_modifier,
			damage_dealt_modifier,
			damage_taken_modifier,
			healing_received_modifier,
			stamina_regen_modifier,
			blocks_movement,
			blocks_actions,
			blocks_skills,
			created_at,
			updated_at
		FROM apeiron.status_effect
		WHERE id = $1
	`, id).Scan(
		&s.ID,
		&s.Name,
		&s.Description,
		&s.EffectType,
		&s.StackingMode,
		&s.MaxStacks,
		&s.DurationMS,
		&s.TickIntervalMS,
		&s.IsDispellable,
		&s.IsPVPEnabled,
		&s.MovementModifier,
		&s.DamageDealtModifier,
		&s.DamageTakenModifier,
		&s.HealingReceivedModifier,
		&s.StaminaRegenModifier,
		&s.BlocksMovement,
		&s.BlocksActions,
		&s.BlocksSkills,
		&s.CreatedAt,
		&s.UpdatedAt,
	)

	return s, err
}

type Skill struct {
	ID          string
	Name        string
	Description string
	Archetype   string
	SkillType   string

	StaminaCost float64
	ManaCost    float64
	HealthCost  float64

	WindupMS            int
	ActiveFramesMS      int
	RecoveryMS          int
	CastTimeMS          int
	CooldownMS          int
	CancelWindowStartMS int
	CancelWindowEndMS   int
	IFrameStartMS       int
	IFrameEndMS         int

	MinRange       float64
	MaxRange       float64
	ConeAngle      float64
	MaxTargets     int
	TargetType     string
	RequiresTarget bool

	BaseDamage              float64
	DamageType              string
	ElementalType           sql.NullString
	PostureDamage           float64
	ArmorPenetration        float64
	DamageMultiplier        float64
	CriticalBonusMultiplier float64

	StunDurationMS int
	RootDurationMS int
	KnockbackForce float64

	MovementMultiplier float64
	LocksMovement      bool
	MovementDistance   float64

	ComboGroup    sql.NullString
	ComboIndex    sql.NullInt64
	ComboWindowMS int

	IsInterruptible    bool
	IsBlockable        bool
	IsParryable        bool
	IgnoresLineOfSight bool
	IgnoresCollision   bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

type SkillSet struct {
	ID          string
	Name        string
	Description string

	IsPlayerUsable bool
	IsNPCUsable    bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

type SkillSlot struct {
	ID         int64
	SkillSetID string
	SkillID    string

	SlotIndex int
	IsEnabled bool

	Priority           int
	UsageWeight        float64
	CooldownOverrideMS sql.NullInt64

	MinTargetHPPercent sql.NullFloat64
	MaxTargetHPPercent sql.NullFloat64
	MinSelfHPPercent   sql.NullFloat64
	MaxSelfHPPercent   sql.NullFloat64

	RequiredDistanceMin sql.NullFloat64
	RequiredDistanceMax sql.NullFloat64

	RequiresLineOfSight bool

	OpenerWeight   float64
	FinisherWeight float64

	SharedCooldownGroup sql.NullString
	UseOnlyInCombat     bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

type SkillProjectileProfile struct {
	SkillID string

	TrajectoryType string

	ProjectileSpeed  float64
	ProjectileRadius float64
	MaxLifetimeMS    int

	GravityMultiplier float64
	DragMultiplier    float64

	CollisionMode string

	CanBeBlocked bool
	CanBeParried bool
	CanBeDodged  bool

	RequiresServerConfirmation bool

	CanPierce      bool
	MaxPierceCount int

	CanHome        bool
	HomingStrength float64
	HomingTurnRate float64

	CanRicochet      bool
	MaxRicochetCount int

	SpawnOffsetX float64
	SpawnOffsetY float64
	SpawnOffsetZ float64

	CreatedAt time.Time
	UpdatedAt time.Time
}

type SkillHitboxProfile struct {
	ID      string
	SkillID string

	HitboxIndex int
	HitboxShape string

	HitboxStartMS int
	HitboxEndMS   int

	OffsetX float64
	OffsetY float64
	OffsetZ float64

	SizeX float64
	SizeY float64
	SizeZ float64

	Radius float64
	Length float64
	Angle  float64

	FollowsCaster     bool
	FollowsProjectile bool

	CanMultiHit      bool
	MaxHitsPerTarget int
	HitIntervalMS    int

	FriendlyFire bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

type SkillAreaEffectProfile struct {
	SkillID string

	AreaShape string

	Radius float64
	Length float64
	Width  float64
	Height float64
	Angle  float64

	DurationMS     int
	TickIntervalMS int

	DamageFalloffType    string
	MinFalloffMultiplier float64

	AppliesOnImpact     bool
	PersistsAfterImpact bool

	MaxTargets int

	FriendlyFire bool

	StatusEffectID sql.NullString

	CreatedAt time.Time
	UpdatedAt time.Time
}

type SkillImpactProfile struct {
	SkillID string

	ImpactType  string
	HitReaction string

	PoiseDamage  float64
	StaggerPower float64

	InterruptPower        float64
	GuardDamageMultiplier float64

	BounceOnShield bool
	DestroyOnHit   bool
	StickOnHit     bool

	HitstopMS           int
	ScreenshakeStrength float64

	KnockbackForce       float64
	KnockbackUpwardForce float64
	PullForce            float64

	AppliesStatusEffect bool
	StatusEffectID      sql.NullString
	StatusEffectChance  float64

	CreatedAt time.Time
	UpdatedAt time.Time
}

type SkillLoadoutItem struct {
	Slot  SkillSlot
	Skill Skill
}

type StatusEffect struct {
	ID          string
	Name        string
	Description string

	EffectType   string
	StackingMode string

	MaxStacks      int
	DurationMS     int
	TickIntervalMS int

	IsDispellable bool
	IsPVPEnabled  bool

	MovementModifier        float64
	DamageDealtModifier     float64
	DamageTakenModifier     float64
	HealingReceivedModifier float64
	StaminaRegenModifier    float64

	BlocksMovement bool
	BlocksActions  bool
	BlocksSkills   bool

	CreatedAt time.Time
	UpdatedAt time.Time
}
