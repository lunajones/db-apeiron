package postgres

import (
	"context"
	"database/sql"
	"time"

	"db-apeiron/internal/database"
)

type PlayerRepository struct {
	db database.TxManager
}

func NewPlayerRepository(db database.TxManager) *PlayerRepository {
	return &PlayerRepository{db: db}
}

func (r *PlayerRepository) GetByID(ctx context.Context, id string) (Player, error) {
	var p Player

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			account_id,
			name,
			creature_instance_id,
			level,
			experience,
			attribute_points,
			strength,
			dexterity,
			intelligence,
			endurance,
			pvp_enabled,
			is_in_safe_zone,
			guild_id,
			party_id,
			reputation,
			coin,
			created_at,
			updated_at
		FROM apeiron.player
		WHERE id = $1
	`, id).Scan(
		&p.ID,
		&p.AccountID,
		&p.Name,
		&p.CreatureInstanceID,
		&p.Level,
		&p.Experience,
		&p.AttributePoints,
		&p.Strength,
		&p.Dexterity,
		&p.Intelligence,
		&p.Endurance,
		&p.PVPEnabled,
		&p.IsInSafeZone,
		&p.GuildID,
		&p.PartyID,
		&p.Reputation,
		&p.Coin,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	return p, err
}

func (r *PlayerRepository) Create(ctx context.Context, p *Player) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO apeiron.player (
			id,
			account_id,
			name,
			creature_instance_id,
			level,
			experience,
			attribute_points,
			strength,
			dexterity,
			intelligence,
			endurance,
			pvp_enabled,
			is_in_safe_zone,
			guild_id,
			party_id,
			reputation,
			coin,
			created_at,
			updated_at
		)
		VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,
			$11,$12,$13,$14,$15,$16,$17,$18,$19
		)
	`,
		p.ID,
		p.AccountID,
		p.Name,
		p.CreatureInstanceID,
		p.Level,
		p.Experience,
		p.AttributePoints,
		p.Strength,
		p.Dexterity,
		p.Intelligence,
		p.Endurance,
		p.PVPEnabled,
		p.IsInSafeZone,
		p.GuildID,
		p.PartyID,
		p.Reputation,
		p.Coin,
		p.CreatedAt,
		p.UpdatedAt,
	)

	return err
}

func (r *PlayerRepository) UpdateProgression(
	ctx context.Context,
	id string,
	level int,
	experience int64,
	attributePoints int,
) error {
	_, err := r.db.Exec(ctx, `
		UPDATE apeiron.player
		SET
			level = $1,
			experience = $2,
			attribute_points = $3,
			updated_at = NOW()
		WHERE id = $4
	`,
		level,
		experience,
		attributePoints,
		id,
	)

	return err
}

func (r *PlayerRepository) UpdateAttributes(
	ctx context.Context,
	id string,
	strength float64,
	dexterity float64,
	intelligence float64,
	endurance float64,
) error {
	_, err := r.db.Exec(ctx, `
		UPDATE apeiron.player
		SET
			strength = $1,
			dexterity = $2,
			intelligence = $3,
			endurance = $4,
			updated_at = NOW()
		WHERE id = $5
	`,
		strength,
		dexterity,
		intelligence,
		endurance,
		id,
	)

	return err
}

func (r *PlayerRepository) UpdateCombatFlags(
	ctx context.Context,
	id string,
	pvpEnabled bool,
	isInSafeZone bool,
) error {
	_, err := r.db.Exec(ctx, `
		UPDATE apeiron.player
		SET
			pvp_enabled = $1,
			is_in_safe_zone = $2,
			updated_at = NOW()
		WHERE id = $3
	`,
		pvpEnabled,
		isInSafeZone,
		id,
	)

	return err
}

func (r *PlayerRepository) UpdateSocial(
	ctx context.Context,
	id string,
	guildID sql.NullString,
	partyID sql.NullString,
) error {
	_, err := r.db.Exec(ctx, `
		UPDATE apeiron.player
		SET
			guild_id = $1,
			party_id = $2,
			updated_at = NOW()
		WHERE id = $3
	`,
		guildID,
		partyID,
		id,
	)

	return err
}

func (r *PlayerRepository) UpdateReputation(ctx context.Context, id string, reputation float64) error {
	_, err := r.db.Exec(ctx, `
		UPDATE apeiron.player
		SET
			reputation = $1,
			updated_at = NOW()
		WHERE id = $2
	`, reputation, id)

	return err
}

func (r *PlayerRepository) AddCoin(ctx context.Context, id string, amount int64) error {
	_, err := r.db.Exec(ctx, `
		UPDATE apeiron.player
		SET
			coin = coin + $1,
			updated_at = NOW()
		WHERE id = $2
	`, amount, id)

	return err
}

type Player struct {
	ID                 string
	AccountID          string
	Name               string
	CreatureInstanceID string

	Level           int
	Experience      int64
	AttributePoints int

	Strength     float64
	Dexterity    float64
	Intelligence float64
	Endurance    float64

	PVPEnabled   bool
	IsInSafeZone bool

	GuildID sql.NullString
	PartyID sql.NullString

	Reputation float64
	Coin       int64

	CreatedAt time.Time
	UpdatedAt time.Time
}
