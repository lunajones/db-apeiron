-- =========================================================
-- BIOME
-- APEIRON MMO - ENVIRONMENTAL BEHAVIOR LAYER
-- =========================================================

CREATE TABLE apeiron.biome (
    id TEXT PRIMARY KEY,

    name TEXT NOT NULL,

    region_id TEXT NOT NULL,

    biome_type TEXT NOT NULL,
    -- forest, desert, swamp, mountain, tundra

    temperature FLOAT NOT NULL DEFAULT 0.5,
    humidity FLOAT NOT NULL DEFAULT 0.5,

    visibility_modifier FLOAT NOT NULL DEFAULT 1.0,
    -- neblina, luz, etc

    movement_modifier FLOAT NOT NULL DEFAULT 1.0,
    -- dificuldade de locomoção

    stealth_modifier FLOAT NOT NULL DEFAULT 1.0,

    aggression_modifier FLOAT NOT NULL DEFAULT 1.0,
    -- afeta AI (mais agressivo em certos biomas)

    resource_richness FLOAT NOT NULL DEFAULT 1.0,

    is_safe BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_biome_region
        FOREIGN KEY (region_id)
        REFERENCES apeiron.world_region(id)
);