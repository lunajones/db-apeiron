-- =========================================================
-- SKILL HITBOX PROFILE
-- APEIRON MMO - MELEE / PROJECTILE HITBOX GEOMETRY
-- =========================================================

CREATE TABLE IF NOT EXISTS apeiron.skill_hitbox_profile (
    id TEXT PRIMARY KEY,

    skill_id TEXT NOT NULL,

    hitbox_index INT NOT NULL DEFAULT 0,
    -- permite múltiplas hitboxes na mesma skill

    hitbox_shape TEXT NOT NULL,
    -- sphere | capsule | box | oriented_box | cone | arc | ray | asymmetric_arc | capsule_strip | temporal_sweep

    hitbox_start_ms INT NOT NULL,
    hitbox_end_ms INT NOT NULL,

    offset_x FLOAT NOT NULL DEFAULT 0.0,
    offset_y FLOAT NOT NULL DEFAULT 1.0,
    offset_z FLOAT NOT NULL DEFAULT 0.0,

    size_x FLOAT NOT NULL DEFAULT 1.0,
    size_y FLOAT NOT NULL DEFAULT 1.0,
    size_z FLOAT NOT NULL DEFAULT 1.0,

    radius FLOAT NOT NULL DEFAULT 0.5,
    length FLOAT NOT NULL DEFAULT 1.0,
    angle FLOAT NOT NULL DEFAULT 0.0,
    min_angle_deg FLOAT NOT NULL DEFAULT 0.0,
    max_angle_deg FLOAT NOT NULL DEFAULT 0.0,
    start_radius FLOAT NOT NULL DEFAULT 0.0,
    end_radius FLOAT NOT NULL DEFAULT 0.0,

    motion_profile_id TEXT,
    damage_group_id TEXT,

    follows_caster BOOLEAN NOT NULL DEFAULT TRUE,
    follows_projectile BOOLEAN NOT NULL DEFAULT FALSE,

    can_multi_hit BOOLEAN NOT NULL DEFAULT FALSE,
    max_hits_per_target INT NOT NULL DEFAULT 1,
    hit_interval_ms INT NOT NULL DEFAULT 0,

    friendly_fire BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_skill_hitbox_profile_skill
        FOREIGN KEY (skill_id)
        REFERENCES apeiron.skill(id)
        ON DELETE CASCADE,

    CONSTRAINT uq_skill_hitbox_profile_skill_index
        UNIQUE (skill_id, hitbox_index),

    CONSTRAINT chk_skill_hitbox_shape
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
        )),

    CONSTRAINT chk_skill_hitbox_time
        CHECK (hitbox_end_ms >= hitbox_start_ms),

    CONSTRAINT chk_skill_hitbox_hits
        CHECK (max_hits_per_target >= 1)
);

CREATE INDEX IF NOT EXISTS idx_skill_hitbox_profile_skill
ON apeiron.skill_hitbox_profile(skill_id);

CREATE INDEX IF NOT EXISTS idx_skill_hitbox_profile_shape
ON apeiron.skill_hitbox_profile(hitbox_shape);
