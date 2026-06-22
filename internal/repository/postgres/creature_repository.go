package postgres

import (
	"context"
	stdsql "database/sql"
	"time"

	"db-apeiron/internal/database"
)

type CreatureRepository struct {
	db database.TxManager
}

func NewCreatureRepository(db database.TxManager) *CreatureRepository {
	return &CreatureRepository{db: db}
}

func (r *CreatureRepository) GetTemplateByID(ctx context.Context, id string) (CreatureTemplate, error) {
	var t CreatureTemplate

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			name,
			faction,
			tier,
			archetype,
			movement_profile_id,
			combat_core_profile_id,
			combat_style_profile_id,
			ai_decision_profile_id,
			personality_profile_id,
			sensory_profile_id,
			needs_profile_id,
			skill_set_id,
			spawn_profile_id,
			created_at,
			updated_at
		FROM apeiron.creature_template
		WHERE id = $1
	`, id).Scan(
		&t.ID,
		&t.Name,
		&t.Faction,
		&t.Tier,
		&t.Archetype,
		&t.MovementProfileID,
		&t.CombatCoreProfileID,
		&t.CombatStyleProfileID,
		&t.AIDecisionProfileID,
		&t.PersonalityProfileID,
		&t.SensoryProfileID,
		&t.NeedsProfileID,
		&t.SkillSetID,
		&t.SpawnProfileID,
		&t.CreatedAt,
		&t.UpdatedAt,
	)

	return t, err
}

func (r *CreatureRepository) GetInstanceByID(ctx context.Context, id string) (CreatureInstance, error) {
	var c CreatureInstance

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			template_id,
			region_id,
			biome_id,
			zone_id,
			pos_x,
			pos_y,
			pos_z,
			rot_y,
			hp_current,
			stamina_current,
			posture_current,
			is_alive,
			in_combat,
			combat_target_id,
			last_damage_taken_ms,
			current_emotion,
			aggression_state,
			fear_state,
			last_decision_ms,
			hunger_current,
			thirst_current,
			fatigue_current,
			last_eat_ms,
			last_drink_ms,
			last_rest_ms,
			velocity_x,
			velocity_y,
			velocity_z,
			is_moving,
			skill_set_id,
			home_region_id,
			home_zone_id,
			leash_center_x,
			leash_center_y,
			leash_center_z,
			leash_distance,
			is_stunned,
			is_rooted,
			is_silenced,
			cc_end_ms,
			spawn_time,
			last_update
		FROM apeiron.creature_instance
		WHERE id = $1
	`, id).Scan(
		&c.ID,
		&c.TemplateID,
		&c.RegionID,
		&c.BiomeID,
		&c.ZoneID,
		&c.PosX,
		&c.PosY,
		&c.PosZ,
		&c.RotY,
		&c.HPCurrent,
		&c.StaminaCurrent,
		&c.PostureCurrent,
		&c.IsAlive,
		&c.InCombat,
		&c.CombatTargetID,
		&c.LastDamageTakenMS,
		&c.CurrentEmotion,
		&c.AggressionState,
		&c.FearState,
		&c.LastDecisionMS,
		&c.HungerCurrent,
		&c.ThirstCurrent,
		&c.FatigueCurrent,
		&c.LastEatMS,
		&c.LastDrinkMS,
		&c.LastRestMS,
		&c.VelocityX,
		&c.VelocityY,
		&c.VelocityZ,
		&c.IsMoving,
		&c.SkillSetID,
		&c.HomeRegionID,
		&c.HomeZoneID,
		&c.LeashCenterX,
		&c.LeashCenterY,
		&c.LeashCenterZ,
		&c.LeashDistance,
		&c.IsStunned,
		&c.IsRooted,
		&c.IsSilenced,
		&c.CCEndMS,
		&c.SpawnTime,
		&c.LastUpdate,
	)

	return c, err
}

func (r *CreatureRepository) CreateInstance(ctx context.Context, c *CreatureInstance) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO apeiron.creature_instance (
			id,
			template_id,
			region_id,
			biome_id,
			zone_id,
			pos_x,
			pos_y,
			pos_z,
			rot_y,
			hp_current,
			stamina_current,
			posture_current,
			is_alive,
			in_combat,
			combat_target_id,
			last_damage_taken_ms,
			current_emotion,
			aggression_state,
			fear_state,
			last_decision_ms,
			hunger_current,
			thirst_current,
			fatigue_current,
			last_eat_ms,
			last_drink_ms,
			last_rest_ms,
			velocity_x,
			velocity_y,
			velocity_z,
			is_moving,
			skill_set_id,
			home_region_id,
			home_zone_id,
			leash_center_x,
			leash_center_y,
			leash_center_z,
			leash_distance,
			is_stunned,
			is_rooted,
			is_silenced,
			cc_end_ms,
			spawn_time,
			last_update
		)
		VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,
			$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,
			$21,$22,$23,$24,$25,$26,$27,$28,$29,$30,
			$31,$32,$33,$34,$35,$36,$37,$38,$39,$40,
			$41,$42,$43
		)
	`,
		c.ID,
		c.TemplateID,
		c.RegionID,
		c.BiomeID,
		c.ZoneID,
		c.PosX,
		c.PosY,
		c.PosZ,
		c.RotY,
		c.HPCurrent,
		c.StaminaCurrent,
		c.PostureCurrent,
		c.IsAlive,
		c.InCombat,
		c.CombatTargetID,
		c.LastDamageTakenMS,
		c.CurrentEmotion,
		c.AggressionState,
		c.FearState,
		c.LastDecisionMS,
		c.HungerCurrent,
		c.ThirstCurrent,
		c.FatigueCurrent,
		c.LastEatMS,
		c.LastDrinkMS,
		c.LastRestMS,
		c.VelocityX,
		c.VelocityY,
		c.VelocityZ,
		c.IsMoving,
		c.SkillSetID,
		c.HomeRegionID,
		c.HomeZoneID,
		c.LeashCenterX,
		c.LeashCenterY,
		c.LeashCenterZ,
		c.LeashDistance,
		c.IsStunned,
		c.IsRooted,
		c.IsSilenced,
		c.CCEndMS,
		c.SpawnTime,
		c.LastUpdate,
	)

	return err
}

func (r *CreatureRepository) GetByRegion(ctx context.Context, regionID string) ([]CreatureInstance, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id
		FROM apeiron.creature_instance
		WHERE region_id = $1
	`, regionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []CreatureInstance

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}

		c, err := r.GetInstanceByID(ctx, id)
		if err != nil {
			return nil, err
		}

		out = append(out, c)
	}

	return out, rows.Err()
}

func (r *CreatureRepository) UpdatePosition(ctx context.Context, id string, posX, posY, posZ, rotY float64) error {
	_, err := r.db.Exec(ctx, `
		UPDATE apeiron.creature_instance
		SET
			pos_x = $1,
			pos_y = $2,
			pos_z = $3,
			rot_y = $4,
			last_update = NOW()
		WHERE id = $5
	`, posX, posY, posZ, rotY, id)

	return err
}

func (r *CreatureRepository) UpdateRuntimeStats(ctx context.Context, id string, hp, stamina, posture float64, isAlive bool) error {
	_, err := r.db.Exec(ctx, `
		UPDATE apeiron.creature_instance
		SET
			hp_current = $1,
			stamina_current = $2,
			posture_current = $3,
			is_alive = $4,
			last_update = NOW()
		WHERE id = $5
	`, hp, stamina, posture, isAlive, id)

	return err
}

func (r *CreatureRepository) UpdateCombatState(ctx context.Context, id string, inCombat bool, targetID stdsql.NullString, lastDamageTakenMS stdsql.NullInt64) error {
	_, err := r.db.Exec(ctx, `
		UPDATE apeiron.creature_instance
		SET
			in_combat = $1,
			combat_target_id = $2,
			last_damage_taken_ms = $3,
			last_update = NOW()
		WHERE id = $4
	`, inCombat, targetID, lastDamageTakenMS, id)

	return err
}

func (r *CreatureRepository) UpdateAIState(ctx context.Context, id string, emotion stdsql.NullString, aggression, fear float64, lastDecisionMS stdsql.NullInt64) error {
	_, err := r.db.Exec(ctx, `
		UPDATE apeiron.creature_instance
		SET
			current_emotion = $1,
			aggression_state = $2,
			fear_state = $3,
			last_decision_ms = $4,
			last_update = NOW()
		WHERE id = $5
	`, emotion, aggression, fear, lastDecisionMS, id)

	return err
}

func (r *CreatureRepository) UpdateNeedsState(ctx context.Context, id string, hunger, thirst, fatigue float64, lastEatMS, lastDrinkMS, lastRestMS stdsql.NullInt64) error {
	_, err := r.db.Exec(ctx, `
		UPDATE apeiron.creature_instance
		SET
			hunger_current = $1,
			thirst_current = $2,
			fatigue_current = $3,
			last_eat_ms = $4,
			last_drink_ms = $5,
			last_rest_ms = $6,
			last_update = NOW()
		WHERE id = $7
	`, hunger, thirst, fatigue, lastEatMS, lastDrinkMS, lastRestMS, id)

	return err
}

func (r *CreatureRepository) UpdateMovementState(ctx context.Context, id string, vx, vy, vz float64, isMoving bool) error {
	_, err := r.db.Exec(ctx, `
		UPDATE apeiron.creature_instance
		SET
			velocity_x = $1,
			velocity_y = $2,
			velocity_z = $3,
			is_moving = $4,
			last_update = NOW()
		WHERE id = $5
	`, vx, vy, vz, isMoving, id)

	return err
}

func (r *CreatureRepository) UpdateStatusEffects(ctx context.Context, id string, stunned, rooted, silenced bool, ccEndMS stdsql.NullInt64) error {
	_, err := r.db.Exec(ctx, `
		UPDATE apeiron.creature_instance
		SET
			is_stunned = $1,
			is_rooted = $2,
			is_silenced = $3,
			cc_end_ms = $4,
			last_update = NOW()
		WHERE id = $5
	`, stunned, rooted, silenced, ccEndMS, id)

	return err
}

func (r *CreatureRepository) GetSkillStates(ctx context.Context, creatureInstanceID string) ([]CreatureSkillState, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			creature_instance_id,
			skill_id,
			cooldown_end_ms,
			last_used_ms,
			charges,
			is_locked,
			created_at,
			updated_at
		FROM apeiron.creature_instance_skill_state
		WHERE creature_instance_id = $1
	`, creatureInstanceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []CreatureSkillState

	for rows.Next() {
		var s CreatureSkillState

		if err := rows.Scan(
			&s.CreatureInstanceID,
			&s.SkillID,
			&s.CooldownEndMS,
			&s.LastUsedMS,
			&s.Charges,
			&s.IsLocked,
			&s.CreatedAt,
			&s.UpdatedAt,
		); err != nil {
			return nil, err
		}

		out = append(out, s)
	}

	return out, rows.Err()
}

func (r *CreatureRepository) UpsertSkillState(ctx context.Context, s *CreatureSkillState) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO apeiron.creature_instance_skill_state (
			creature_instance_id,
			skill_id,
			cooldown_end_ms,
			last_used_ms,
			charges,
			is_locked,
			created_at,
			updated_at
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
		ON CONFLICT (creature_instance_id, skill_id)
		DO UPDATE SET
			cooldown_end_ms = EXCLUDED.cooldown_end_ms,
			last_used_ms = EXCLUDED.last_used_ms,
			charges = EXCLUDED.charges,
			is_locked = EXCLUDED.is_locked,
			updated_at = NOW()
	`,
		s.CreatureInstanceID,
		s.SkillID,
		s.CooldownEndMS,
		s.LastUsedMS,
		s.Charges,
		s.IsLocked,
		s.CreatedAt,
		s.UpdatedAt,
	)

	return err
}

type CreatureTemplate struct {
	ID                   string
	Name                 string
	Faction              string
	Tier                 int
	Archetype            string
	MovementProfileID    string
	CombatCoreProfileID  string
	CombatStyleProfileID string
	AIDecisionProfileID  string
	PersonalityProfileID string
	SensoryProfileID     string
	NeedsProfileID       string
	SkillSetID           string
	SpawnProfileID       string
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

type CreatureInstance struct {
	ID         string
	TemplateID string

	RegionID string
	BiomeID  string
	ZoneID   string

	PosX float64
	PosY float64
	PosZ float64
	RotY float64

	HPCurrent      float64
	StaminaCurrent float64
	PostureCurrent float64
	IsAlive        bool

	InCombat          bool
	CombatTargetID    stdsql.NullString
	LastDamageTakenMS stdsql.NullInt64

	CurrentEmotion  stdsql.NullString
	AggressionState float64
	FearState       float64
	LastDecisionMS  stdsql.NullInt64

	HungerCurrent  float64
	ThirstCurrent  float64
	FatigueCurrent float64

	LastEatMS   stdsql.NullInt64
	LastDrinkMS stdsql.NullInt64
	LastRestMS  stdsql.NullInt64

	VelocityX float64
	VelocityY float64
	VelocityZ float64
	IsMoving  bool

	SkillSetID string

	HomeRegionID stdsql.NullString
	HomeZoneID   stdsql.NullString

	LeashCenterX  stdsql.NullFloat64
	LeashCenterY  stdsql.NullFloat64
	LeashCenterZ  stdsql.NullFloat64
	LeashDistance stdsql.NullFloat64

	IsStunned  bool
	IsRooted   bool
	IsSilenced bool
	CCEndMS    stdsql.NullInt64

	SpawnTime  time.Time
	LastUpdate time.Time
}

type CreatureSkillState struct {
	CreatureInstanceID string
	SkillID            string

	CooldownEndMS int64
	LastUsedMS    int64
	Charges       int
	IsLocked      bool

	CreatedAt time.Time
	UpdatedAt time.Time
}
