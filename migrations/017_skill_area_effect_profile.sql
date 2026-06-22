-- =========================================================
-- SKILL AREA EFFECT PROFILE
-- APEIRON MMO - AOE / DOT / FIELD EFFECT CONFIG
-- =========================================================

CREATE TABLE IF NOT EXISTS apeiron.skill_area_effect_profile (
    skill_id TEXT PRIMARY KEY,

    area_shape TEXT NOT NULL DEFAULT 'sphere',
    -- sphere | circle | cone | box | line

    radius FLOAT NOT NULL DEFAULT 3.0,
    length FLOAT NOT NULL DEFAULT 0.0,
    width FLOAT NOT NULL DEFAULT 0.0,
    height FLOAT NOT NULL DEFAULT 0.0,
    angle FLOAT NOT NULL DEFAULT 360.0,

    duration_ms INT NOT NULL DEFAULT 0,
    tick_interval_ms INT NOT NULL DEFAULT 0,

    damage_falloff_type TEXT NOT NULL DEFAULT 'none',
    -- none | linear | quadratic

    min_falloff_multiplier FLOAT NOT NULL DEFAULT 1.0,

    applies_on_impact BOOLEAN NOT NULL DEFAULT TRUE,
    persists_after_impact BOOLEAN NOT NULL DEFAULT FALSE,

    max_targets INT NOT NULL DEFAULT 0,
    -- 0 = sem limite

    friendly_fire BOOLEAN NOT NULL DEFAULT FALSE,

    status_effect_id TEXT,
    -- futuro: FK quando existir tabela de status_effect

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_skill_area_effect_profile_skill
        FOREIGN KEY (skill_id)
        REFERENCES apeiron.skill(id)
        ON DELETE CASCADE,
   
    CONSTRAINT fk_skill_area_effect_status_effect
        FOREIGN KEY (status_effect_id)
        REFERENCES apeiron.status_effect(id)
        ON DELETE CASCADE,
   
    CONSTRAINT chk_skill_area_shape
        CHECK (area_shape IN (
            'sphere',
            'circle',
            'cone',
            'box',
            'line'
        )),

    CONSTRAINT chk_skill_area_damage_falloff
        CHECK (damage_falloff_type IN (
            'none',
            'linear',
            'quadratic'
        )),

    CONSTRAINT chk_skill_area_radius
        CHECK (radius >= 0),

    CONSTRAINT chk_skill_area_duration
        CHECK (duration_ms >= 0),

    CONSTRAINT chk_skill_area_tick_interval
        CHECK (tick_interval_ms >= 0),

    CONSTRAINT chk_skill_area_max_targets
        CHECK (max_targets >= 0)
);

CREATE INDEX IF NOT EXISTS idx_skill_area_effect_shape
ON apeiron.skill_area_effect_profile(area_shape);

CREATE INDEX IF NOT EXISTS idx_skill_area_effect_status
ON apeiron.skill_area_effect_profile(status_effect_id);
