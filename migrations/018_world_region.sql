-- =========================================================
-- WORLD REGION
-- APEIRON MMO - WORLD SHARD / MACRO ZONE
-- =========================================================

CREATE TABLE apeiron.world_region (
    id TEXT PRIMARY KEY,

    name TEXT NOT NULL,

    region_type TEXT NOT NULL,
    -- pvp_zone, safe_zone, dungeon, city, wild

    world_scale INT NOT NULL DEFAULT 1,
    -- escala lógica (não física)

    is_instanced BOOLEAN NOT NULL DEFAULT FALSE,
    -- região única ou instanciada

    max_players INT NOT NULL DEFAULT 100,

    center_x DOUBLE PRECISION NOT NULL DEFAULT 0,
    center_y DOUBLE PRECISION NOT NULL DEFAULT 0,
    center_z DOUBLE PRECISION NOT NULL DEFAULT 0,

    size_x DOUBLE PRECISION NOT NULL DEFAULT 1000,
    size_y DOUBLE PRECISION NOT NULL DEFAULT 1000,
    size_z DOUBLE PRECISION NOT NULL DEFAULT 1000,

    danger_level FLOAT NOT NULL DEFAULT 0.5,
    -- influencia AI global

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);