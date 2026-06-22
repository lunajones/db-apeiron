package postgres

import (
	"context"
	"database/sql"
	"time"

	"db-apeiron/internal/database"
)

type WorldRepository struct {
	db database.TxManager
}

func NewWorldRepository(db database.TxManager) *WorldRepository {
	return &WorldRepository{db: db}
}

//
// =========================
// WORLD REGION
// =========================
//

func (r *WorldRepository) GetRegionByID(ctx context.Context, id string) (WorldRegion, error) {
	var wr WorldRegion

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			name,
			region_type,
			world_scale,
			is_instanced,
			max_players,
			center_x,
			center_y,
			center_z,
			size_x,
			size_y,
			size_z,
			danger_level,
			created_at,
			updated_at
		FROM apeiron.world_region
		WHERE id = $1
	`, id).Scan(
		&wr.ID,
		&wr.Name,
		&wr.RegionType,
		&wr.WorldScale,
		&wr.IsInstanced,
		&wr.MaxPlayers,
		&wr.CenterX,
		&wr.CenterY,
		&wr.CenterZ,
		&wr.SizeX,
		&wr.SizeY,
		&wr.SizeZ,
		&wr.DangerLevel,
		&wr.CreatedAt,
		&wr.UpdatedAt,
	)

	return wr, err
}

func (r *WorldRepository) ListRegions(ctx context.Context) ([]WorldRegion, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			id,
			name,
			region_type,
			world_scale,
			is_instanced,
			max_players,
			center_x,
			center_y,
			center_z,
			size_x,
			size_y,
			size_z,
			danger_level,
			created_at,
			updated_at
		FROM apeiron.world_region
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []WorldRegion

	for rows.Next() {
		var wr WorldRegion

		if err := rows.Scan(
			&wr.ID,
			&wr.Name,
			&wr.RegionType,
			&wr.WorldScale,
			&wr.IsInstanced,
			&wr.MaxPlayers,
			&wr.CenterX,
			&wr.CenterY,
			&wr.CenterZ,
			&wr.SizeX,
			&wr.SizeY,
			&wr.SizeZ,
			&wr.DangerLevel,
			&wr.CreatedAt,
			&wr.UpdatedAt,
		); err != nil {
			return nil, err
		}

		out = append(out, wr)
	}

	return out, rows.Err()
}

func (r *WorldRepository) CreateRegion(ctx context.Context, wr *WorldRegion) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO apeiron.world_region (
			id,
			name,
			region_type,
			world_scale,
			is_instanced,
			max_players,
			center_x,
			center_y,
			center_z,
			size_x,
			size_y,
			size_z,
			danger_level,
			created_at,
			updated_at
		)
		VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,
			$11,$12,$13,$14,$15
		)
	`,
		wr.ID,
		wr.Name,
		wr.RegionType,
		wr.WorldScale,
		wr.IsInstanced,
		wr.MaxPlayers,
		wr.CenterX,
		wr.CenterY,
		wr.CenterZ,
		wr.SizeX,
		wr.SizeY,
		wr.SizeZ,
		wr.DangerLevel,
		wr.CreatedAt,
		wr.UpdatedAt,
	)

	return err
}

//
// =========================
// BIOME
// =========================
//

func (r *WorldRepository) GetBiomeByID(ctx context.Context, id string) (Biome, error) {
	var b Biome

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			name,
			region_id,
			biome_type,
			temperature,
			humidity,
			visibility_modifier,
			movement_modifier,
			stealth_modifier,
			aggression_modifier,
			resource_richness,
			is_safe,
			created_at,
			updated_at
		FROM apeiron.biome
		WHERE id = $1
	`, id).Scan(
		&b.ID,
		&b.Name,
		&b.RegionID,
		&b.BiomeType,
		&b.Temperature,
		&b.Humidity,
		&b.VisibilityModifier,
		&b.MovementModifier,
		&b.StealthModifier,
		&b.AggressionModifier,
		&b.ResourceRichness,
		&b.IsSafe,
		&b.CreatedAt,
		&b.UpdatedAt,
	)

	return b, err
}

func (r *WorldRepository) GetBiomesByRegion(ctx context.Context, regionID string) ([]Biome, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			id,
			name,
			region_id,
			biome_type,
			temperature,
			humidity,
			visibility_modifier,
			movement_modifier,
			stealth_modifier,
			aggression_modifier,
			resource_richness,
			is_safe,
			created_at,
			updated_at
		FROM apeiron.biome
		WHERE region_id = $1
		ORDER BY id
	`, regionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []Biome

	for rows.Next() {
		var b Biome

		if err := rows.Scan(
			&b.ID,
			&b.Name,
			&b.RegionID,
			&b.BiomeType,
			&b.Temperature,
			&b.Humidity,
			&b.VisibilityModifier,
			&b.MovementModifier,
			&b.StealthModifier,
			&b.AggressionModifier,
			&b.ResourceRichness,
			&b.IsSafe,
			&b.CreatedAt,
			&b.UpdatedAt,
		); err != nil {
			return nil, err
		}

		out = append(out, b)
	}

	return out, rows.Err()
}

func (r *WorldRepository) CreateBiome(ctx context.Context, b *Biome) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO apeiron.biome (
			id,
			name,
			region_id,
			biome_type,
			temperature,
			humidity,
			visibility_modifier,
			movement_modifier,
			stealth_modifier,
			aggression_modifier,
			resource_richness,
			is_safe,
			created_at,
			updated_at
		)
		VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,
			$11,$12,$13,$14
		)
	`,
		b.ID,
		b.Name,
		b.RegionID,
		b.BiomeType,
		b.Temperature,
		b.Humidity,
		b.VisibilityModifier,
		b.MovementModifier,
		b.StealthModifier,
		b.AggressionModifier,
		b.ResourceRichness,
		b.IsSafe,
		b.CreatedAt,
		b.UpdatedAt,
	)

	return err
}

//
// =========================
// SPAWN ZONE
// =========================
//

func (r *WorldRepository) GetSpawnZoneByID(ctx context.Context, id string) (SpawnZone, error) {
	var sz SpawnZone

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			region_id,
			biome_id,
			name,
			center_x,
			center_y,
			center_z,
			radius,
			max_entities,
			current_entities,
			respawn_time_ms,
			spawn_density,
			allowed_archetypes,
			aggression_level,
			leash_enabled,
			created_at,
			updated_at
		FROM apeiron.spawn_zone
		WHERE id = $1
	`, id).Scan(
		&sz.ID,
		&sz.RegionID,
		&sz.BiomeID,
		&sz.Name,
		&sz.CenterX,
		&sz.CenterY,
		&sz.CenterZ,
		&sz.Radius,
		&sz.MaxEntities,
		&sz.CurrentEntities,
		&sz.RespawnTimeMS,
		&sz.SpawnDensity,
		&sz.AllowedArchetypes,
		&sz.AggressionLevel,
		&sz.LeashEnabled,
		&sz.CreatedAt,
		&sz.UpdatedAt,
	)

	return sz, err
}

func (r *WorldRepository) GetSpawnZonesByRegion(ctx context.Context, regionID string) ([]SpawnZone, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			id,
			region_id,
			biome_id,
			name,
			center_x,
			center_y,
			center_z,
			radius,
			max_entities,
			current_entities,
			respawn_time_ms,
			spawn_density,
			allowed_archetypes,
			aggression_level,
			leash_enabled,
			created_at,
			updated_at
		FROM apeiron.spawn_zone
		WHERE region_id = $1
		ORDER BY id
	`, regionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []SpawnZone

	for rows.Next() {
		var sz SpawnZone

		if err := rows.Scan(
			&sz.ID,
			&sz.RegionID,
			&sz.BiomeID,
			&sz.Name,
			&sz.CenterX,
			&sz.CenterY,
			&sz.CenterZ,
			&sz.Radius,
			&sz.MaxEntities,
			&sz.CurrentEntities,
			&sz.RespawnTimeMS,
			&sz.SpawnDensity,
			&sz.AllowedArchetypes,
			&sz.AggressionLevel,
			&sz.LeashEnabled,
			&sz.CreatedAt,
			&sz.UpdatedAt,
		); err != nil {
			return nil, err
		}

		out = append(out, sz)
	}

	return out, rows.Err()
}

func (r *WorldRepository) GetSpawnZonesByBiome(ctx context.Context, biomeID string) ([]SpawnZone, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			id,
			region_id,
			biome_id,
			name,
			center_x,
			center_y,
			center_z,
			radius,
			max_entities,
			current_entities,
			respawn_time_ms,
			spawn_density,
			allowed_archetypes,
			aggression_level,
			leash_enabled,
			created_at,
			updated_at
		FROM apeiron.spawn_zone
		WHERE biome_id = $1
		ORDER BY id
	`, biomeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []SpawnZone

	for rows.Next() {
		var sz SpawnZone

		if err := rows.Scan(
			&sz.ID,
			&sz.RegionID,
			&sz.BiomeID,
			&sz.Name,
			&sz.CenterX,
			&sz.CenterY,
			&sz.CenterZ,
			&sz.Radius,
			&sz.MaxEntities,
			&sz.CurrentEntities,
			&sz.RespawnTimeMS,
			&sz.SpawnDensity,
			&sz.AllowedArchetypes,
			&sz.AggressionLevel,
			&sz.LeashEnabled,
			&sz.CreatedAt,
			&sz.UpdatedAt,
		); err != nil {
			return nil, err
		}

		out = append(out, sz)
	}

	return out, rows.Err()
}

func (r *WorldRepository) CreateSpawnZone(ctx context.Context, sz *SpawnZone) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO apeiron.spawn_zone (
			id,
			region_id,
			biome_id,
			name,
			center_x,
			center_y,
			center_z,
			radius,
			max_entities,
			current_entities,
			respawn_time_ms,
			spawn_density,
			allowed_archetypes,
			aggression_level,
			leash_enabled,
			created_at,
			updated_at
		)
		VALUES (
			$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,
			$11,$12,$13,$14,$15,$16,$17
		)
	`,
		sz.ID,
		sz.RegionID,
		sz.BiomeID,
		sz.Name,
		sz.CenterX,
		sz.CenterY,
		sz.CenterZ,
		sz.Radius,
		sz.MaxEntities,
		sz.CurrentEntities,
		sz.RespawnTimeMS,
		sz.SpawnDensity,
		sz.AllowedArchetypes,
		sz.AggressionLevel,
		sz.LeashEnabled,
		sz.CreatedAt,
		sz.UpdatedAt,
	)

	return err
}

func (r *WorldRepository) IncrementSpawnZoneEntities(ctx context.Context, id string, amount int) error {
	_, err := r.db.Exec(ctx, `
		UPDATE apeiron.spawn_zone
		SET
			current_entities = current_entities + $1,
			updated_at = NOW()
		WHERE id = $2
	`, amount, id)

	return err
}

func (r *WorldRepository) DecrementSpawnZoneEntities(ctx context.Context, id string, amount int) error {
	_, err := r.db.Exec(ctx, `
		UPDATE apeiron.spawn_zone
		SET
			current_entities = GREATEST(current_entities - $1, 0),
			updated_at = NOW()
		WHERE id = $2
	`, amount, id)

	return err
}

//
// =========================
// MODELS
// =========================
//

type WorldRegion struct {
	ID          string
	Name        string
	RegionType  string
	WorldScale  int
	IsInstanced bool
	MaxPlayers  int

	CenterX float64
	CenterY float64
	CenterZ float64

	SizeX float64
	SizeY float64
	SizeZ float64

	DangerLevel float64

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Biome struct {
	ID       string
	Name     string
	RegionID string

	BiomeType string

	Temperature float64
	Humidity    float64

	VisibilityModifier float64
	MovementModifier   float64
	StealthModifier    float64
	AggressionModifier float64
	ResourceRichness   float64

	IsSafe bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

type SpawnZone struct {
	ID       string
	RegionID string
	BiomeID  string
	Name     string

	CenterX float64
	CenterY float64
	CenterZ float64

	Radius          float64
	MaxEntities     int
	CurrentEntities int
	RespawnTimeMS   int64
	SpawnDensity    float64

	AllowedArchetypes sql.NullString
	AggressionLevel   float64
	LeashEnabled      bool

	CreatedAt time.Time
	UpdatedAt time.Time
}
