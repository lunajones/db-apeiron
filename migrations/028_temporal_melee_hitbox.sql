-- =========================================================
-- TEMPORAL MELEE HITBOXES
-- Reconstructed from Apeiron temporal melee hit volume roadmap.
-- =========================================================




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


CREATE INDEX IF NOT EXISTS idx_skill_hitbox_motion_skill
ON apeiron.skill_hitbox_motion_profile(skill_id);

CREATE INDEX IF NOT EXISTS idx_skill_hitbox_motion_sample_profile
ON apeiron.skill_hitbox_motion_sample(motion_profile_id);
