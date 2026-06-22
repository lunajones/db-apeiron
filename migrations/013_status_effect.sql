-- =========================================================
-- STATUS EFFECT
-- APEIRON MMO - STATIC STATUS EFFECT DEFINITION
-- =========================================================

CREATE TABLE IF NOT EXISTS apeiron.status_effect (
    id TEXT PRIMARY KEY,

    name TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',

    effect_type TEXT NOT NULL,
    -- buff | debuff | dot | hot | crowd_control | utility

    stacking_mode TEXT NOT NULL DEFAULT 'refresh',
    -- none | refresh | stack_intensity | stack_duration

    max_stacks INT NOT NULL DEFAULT 1,

    duration_ms INT NOT NULL DEFAULT 0,
    tick_interval_ms INT NOT NULL DEFAULT 0,

    is_dispellable BOOLEAN NOT NULL DEFAULT TRUE,
    is_pvp_enabled BOOLEAN NOT NULL DEFAULT TRUE,

    movement_modifier FLOAT NOT NULL DEFAULT 1.0,
    damage_dealt_modifier FLOAT NOT NULL DEFAULT 1.0,
    damage_taken_modifier FLOAT NOT NULL DEFAULT 1.0,
    healing_received_modifier FLOAT NOT NULL DEFAULT 1.0,

    stamina_regen_modifier FLOAT NOT NULL DEFAULT 1.0,

    blocks_movement BOOLEAN NOT NULL DEFAULT FALSE,
    blocks_actions BOOLEAN NOT NULL DEFAULT FALSE,
    blocks_skills BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT chk_status_effect_type
        CHECK (effect_type IN (
            'buff',
            'debuff',
            'dot',
            'hot',
            'crowd_control',
            'utility'
        )),

    CONSTRAINT chk_status_effect_stacking_mode
        CHECK (stacking_mode IN (
            'none',
            'refresh',
            'stack_intensity',
            'stack_duration'
        )),

    CONSTRAINT chk_status_effect_max_stacks
        CHECK (max_stacks >= 1),

    CONSTRAINT chk_status_effect_duration
        CHECK (duration_ms >= 0),

    CONSTRAINT chk_status_effect_tick_interval
        CHECK (tick_interval_ms >= 0)
);

CREATE INDEX IF NOT EXISTS idx_status_effect_type
ON apeiron.status_effect(effect_type);

CREATE INDEX IF NOT EXISTS idx_status_effect_stacking_mode
ON apeiron.status_effect(stacking_mode);
