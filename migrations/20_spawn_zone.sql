-- =========================================================
-- SPAWN ZONE
-- APEIRON MMO - LIVING SPAWN SYSTEM
-- =========================================================

CREATE TABLE apeiron.spawn_zone (
    id TEXT PRIMARY KEY,

    region_id TEXT NOT NULL,
    biome_id TEXT NOT NULL,

    name TEXT NOT NULL,

    center_x DOUBLE PRECISION NOT NULL,
    center_y DOUBLE PRECISION NOT NULL,
    center_z DOUBLE PRECISION NOT NULL,

    radius FLOAT NOT NULL DEFAULT 50.0,

    max_entities INT NOT NULL DEFAULT 10,
    current_entities INT NOT NULL DEFAULT 0,

    respawn_time_ms BIGINT NOT NULL DEFAULT 60000,

    spawn_density FLOAT NOT NULL DEFAULT 0.5,

    allowed_archetypes TEXT,
    -- lista controlada no código (JSON ou CSV)
    -- ex: beast, undead

    aggression_level FLOAT NOT NULL DEFAULT 0.5,

    leash_enabled BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_spawn_region
        FOREIGN KEY (region_id)
        REFERENCES apeiron.world_region(id),

    CONSTRAINT fk_spawn_biome
        FOREIGN KEY (biome_id)
        REFERENCES apeiron.biome(id)
);