-- =========================================================
-- TEMPORAL MELEE HITBOXES
-- Reconstructed from Apeiron temporal melee hit volume roadmap.
-- =========================================================

ALTER TABLE apeiron.skill_hitbox_profile
    DROP CONSTRAINT IF EXISTS chk_skill_hitbox_shape;

ALTER TABLE apeiron.skill_hitbox_profile
    ADD CONSTRAINT chk_skill_hitbox_shape
    CHECK (hitbox_shape IN (
        'sphere',
        'capsule',
        'box',
        'oriented_box',
        'cone',
        'arc',
        'ray',
        'asymmetric_arc',
        'capsule_strip',
        'temporal_sweep'
    ));

ALTER TABLE apeiron.skill_hitbox_profile
    ADD COLUMN IF NOT EXISTS motion_profile_id TEXT,
    ADD COLUMN IF NOT EXISTS damage_group_id TEXT,
    ADD COLUMN IF NOT EXISTS min_angle_deg FLOAT NOT NULL DEFAULT 0.0,
    ADD COLUMN IF NOT EXISTS max_angle_deg FLOAT NOT NULL DEFAULT 0.0,
    ADD COLUMN IF NOT EXISTS start_radius FLOAT NOT NULL DEFAULT 0.0,
    ADD COLUMN IF NOT EXISTS end_radius FLOAT NOT NULL DEFAULT 0.0;

CREATE TABLE IF NOT EXISTS apeiron.skill_hitbox_damage_group (
    id TEXT PRIMARY KEY,
    skill_id TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    max_hits_per_target INT NOT NULL DEFAULT 1,
    hit_interval_ms INT NOT NULL DEFAULT 0,
    can_multi_hit BOOLEAN NOT NULL DEFAULT FALSE,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_hitbox_damage_group_skill
        FOREIGN KEY (skill_id)
        REFERENCES apeiron.skill(id)
        ON DELETE CASCADE,
    CONSTRAINT chk_hitbox_damage_group_hits CHECK (max_hits_per_target >= 1)
);

ALTER TABLE IF EXISTS apeiron.skill_hitbox_damage_group
    ADD COLUMN IF NOT EXISTS skill_id TEXT NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS description TEXT NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS max_hits_per_target INT NOT NULL DEFAULT 1,
    ADD COLUMN IF NOT EXISTS hit_interval_ms INT NOT NULL DEFAULT 0,
    ADD COLUMN IF NOT EXISTS can_multi_hit BOOLEAN NOT NULL DEFAULT FALSE,
    ADD COLUMN IF NOT EXISTS metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    ADD COLUMN IF NOT EXISTS created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP NOT NULL DEFAULT NOW();

CREATE TABLE IF NOT EXISTS apeiron.skill_hitbox_motion_profile (
    id TEXT PRIMARY KEY,
    skill_id TEXT NOT NULL,
    motion_type TEXT NOT NULL DEFAULT 'timeline_sweep',
    time_basis TEXT NOT NULL DEFAULT 'hitbox_window_normalized',
    description TEXT NOT NULL DEFAULT '',
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_hitbox_motion_skill
        FOREIGN KEY (skill_id)
        REFERENCES apeiron.skill(id)
        ON DELETE CASCADE
);

ALTER TABLE IF EXISTS apeiron.skill_hitbox_motion_profile
    ADD COLUMN IF NOT EXISTS skill_id TEXT NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS motion_type TEXT NOT NULL DEFAULT 'timeline_sweep',
    ADD COLUMN IF NOT EXISTS time_basis TEXT NOT NULL DEFAULT 'hitbox_window_normalized',
    ADD COLUMN IF NOT EXISTS description TEXT NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    ADD COLUMN IF NOT EXISTS created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP NOT NULL DEFAULT NOW();

CREATE TABLE IF NOT EXISTS apeiron.skill_hitbox_motion_sample (
    id BIGSERIAL PRIMARY KEY,
    motion_profile_id TEXT NOT NULL,
    sample_index INT NOT NULL,
    t FLOAT NOT NULL,
    shape TEXT NOT NULL,
    offset_x FLOAT NOT NULL DEFAULT 0.0,
    offset_y FLOAT NOT NULL DEFAULT 0.0,
    offset_z FLOAT NOT NULL DEFAULT 0.0,
    size_x FLOAT NOT NULL DEFAULT 0.0,
    size_y FLOAT NOT NULL DEFAULT 0.0,
    size_z FLOAT NOT NULL DEFAULT 0.0,
    radius FLOAT NOT NULL DEFAULT 0.0,
    length FLOAT NOT NULL DEFAULT 0.0,
    min_angle_deg FLOAT NOT NULL DEFAULT 0.0,
    max_angle_deg FLOAT NOT NULL DEFAULT 0.0,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_hitbox_motion_sample_profile
        FOREIGN KEY (motion_profile_id)
        REFERENCES apeiron.skill_hitbox_motion_profile(id)
        ON DELETE CASCADE,
    CONSTRAINT uq_hitbox_motion_sample UNIQUE (motion_profile_id, sample_index),
    CONSTRAINT chk_hitbox_motion_t CHECK (t >= 0.0 AND t <= 1.0)
);

ALTER TABLE IF EXISTS apeiron.skill_hitbox_motion_sample
    ADD COLUMN IF NOT EXISTS motion_profile_id TEXT NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS sample_index INT NOT NULL DEFAULT 0,
    ADD COLUMN IF NOT EXISTS t FLOAT NOT NULL DEFAULT 0.0,
    ADD COLUMN IF NOT EXISTS shape TEXT NOT NULL DEFAULT 'capsule_strip',
    ADD COLUMN IF NOT EXISTS offset_x FLOAT NOT NULL DEFAULT 0.0,
    ADD COLUMN IF NOT EXISTS offset_y FLOAT NOT NULL DEFAULT 0.0,
    ADD COLUMN IF NOT EXISTS offset_z FLOAT NOT NULL DEFAULT 0.0,
    ADD COLUMN IF NOT EXISTS size_x FLOAT NOT NULL DEFAULT 0.0,
    ADD COLUMN IF NOT EXISTS size_y FLOAT NOT NULL DEFAULT 0.0,
    ADD COLUMN IF NOT EXISTS size_z FLOAT NOT NULL DEFAULT 0.0,
    ADD COLUMN IF NOT EXISTS radius FLOAT NOT NULL DEFAULT 0.0,
    ADD COLUMN IF NOT EXISTS length FLOAT NOT NULL DEFAULT 0.0,
    ADD COLUMN IF NOT EXISTS min_angle_deg FLOAT NOT NULL DEFAULT 0.0,
    ADD COLUMN IF NOT EXISTS max_angle_deg FLOAT NOT NULL DEFAULT 0.0,
    ADD COLUMN IF NOT EXISTS metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    ADD COLUMN IF NOT EXISTS created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP NOT NULL DEFAULT NOW();

CREATE INDEX IF NOT EXISTS idx_skill_hitbox_motion_skill
ON apeiron.skill_hitbox_motion_profile(skill_id);

CREATE INDEX IF NOT EXISTS idx_skill_hitbox_motion_sample_profile
ON apeiron.skill_hitbox_motion_sample(motion_profile_id);
