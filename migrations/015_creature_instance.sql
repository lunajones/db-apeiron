-- =========================================================
-- CREATURE INSTANCE
-- APEIRON MMO - RUNTIME ENTITY (LIVE WORLD STATE)
-- =========================================================

CREATE TABLE apeiron.creature_instance (
    id TEXT PRIMARY KEY,

    -- =========================
    -- TEMPLATE LINK
    -- =========================

    template_id TEXT NOT NULL,

    -- =========================
    -- WORLD POSITION
    -- =========================

    region_id TEXT NOT NULL,
    biome_id TEXT NOT NULL,
    zone_id TEXT NOT NULL,

    pos_x DOUBLE PRECISION NOT NULL DEFAULT 0,
    pos_y DOUBLE PRECISION NOT NULL DEFAULT 0,
    pos_z DOUBLE PRECISION NOT NULL DEFAULT 0,

    rot_y DOUBLE PRECISION NOT NULL DEFAULT 0,

    -- =========================
    -- CORE RUNTIME STATS
    -- =========================

    hp_current FLOAT NOT NULL DEFAULT 100.0,
    stamina_current FLOAT NOT NULL DEFAULT 100.0,
    posture_current FLOAT NOT NULL DEFAULT 100.0,

    is_alive BOOLEAN NOT NULL DEFAULT TRUE,

    -- =========================
    -- COMBAT STATE
    -- =========================

    in_combat BOOLEAN NOT NULL DEFAULT FALSE,
    combat_target_id TEXT,
    last_damage_taken_ms BIGINT DEFAULT 0,

    -- =========================
    -- AI STATE (RUNTIME ONLY)
    -- =========================

    current_emotion TEXT,
    aggression_state FLOAT NOT NULL DEFAULT 0.5,
    fear_state FLOAT NOT NULL DEFAULT 0.5,

    last_decision_ms BIGINT DEFAULT 0,

    -- =========================
    -- NEEDS STATE (RUNTIME ONLY)
    -- =========================

    hunger_current FLOAT NOT NULL DEFAULT 1.0,
    thirst_current FLOAT NOT NULL DEFAULT 1.0,
    fatigue_current FLOAT NOT NULL DEFAULT 1.0,

    last_eat_ms BIGINT DEFAULT 0,
    last_drink_ms BIGINT DEFAULT 0,
    last_rest_ms BIGINT DEFAULT 0,

    -- =========================
    -- MOVEMENT STATE
    -- =========================

    velocity_x DOUBLE PRECISION NOT NULL DEFAULT 0,
    velocity_y DOUBLE PRECISION NOT NULL DEFAULT 0,
    velocity_z DOUBLE PRECISION NOT NULL DEFAULT 0,

    is_moving BOOLEAN NOT NULL DEFAULT FALSE,

    -- =========================
    -- SKILL RUNTIME STATE
    -- =========================

    skill_set_id TEXT NOT NULL,

    -- =========================
    -- WORLD BEHAVIOR STATE
    -- =========================

    home_region_id TEXT,
    home_zone_id TEXT,

    leash_center_x DOUBLE PRECISION,
    leash_center_y DOUBLE PRECISION,
    leash_center_z DOUBLE PRECISION,
    leash_distance FLOAT DEFAULT 25.0,

    -- =========================
    -- STATUS EFFECTS
    -- =========================

    is_stunned BOOLEAN NOT NULL DEFAULT FALSE,
    is_rooted BOOLEAN NOT NULL DEFAULT FALSE,
    is_silenced BOOLEAN NOT NULL DEFAULT FALSE,

    cc_end_ms BIGINT DEFAULT 0,

    -- =========================
    -- LIFECYCLE
    -- =========================

    spawn_time TIMESTAMP NOT NULL DEFAULT NOW(),
    last_update TIMESTAMP NOT NULL DEFAULT NOW(),

    -- =========================
    -- FOREIGN KEYS
    -- =========================

    CONSTRAINT fk_creature_instance_template
        FOREIGN KEY (template_id)
        REFERENCES apeiron.creature_template(id),

    CONSTRAINT fk_creature_instance_skill_set
        FOREIGN KEY (skill_set_id)
        REFERENCES apeiron.skill_set(id)
);

CREATE INDEX idx_creature_instance_region
ON apeiron.creature_instance(region_id);

CREATE INDEX idx_creature_instance_zone
ON apeiron.creature_instance(zone_id);

CREATE INDEX idx_creature_instance_combat
ON apeiron.creature_instance(in_combat);

CREATE INDEX idx_creature_instance_template
ON apeiron.creature_instance(template_id);