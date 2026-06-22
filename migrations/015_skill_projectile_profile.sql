-- =========================================================
-- SKILL PROJECTILE PROFILE
-- APEIRON MMO - PROJECTILE CONFIG FOR SKILLS
-- =========================================================

CREATE TABLE apeiron.skill_projectile_profile (
    skill_id TEXT PRIMARY KEY,

    trajectory_type TEXT NOT NULL DEFAULT 'linear',
    -- linear | ballistic | lobbed | homing | beam

    projectile_speed FLOAT NOT NULL,
    projectile_radius FLOAT NOT NULL DEFAULT 0.15,
    max_lifetime_ms INT NOT NULL,

    gravity_multiplier FLOAT NOT NULL DEFAULT 1.0,
    drag_multiplier FLOAT NOT NULL DEFAULT 0.0,

    collision_mode TEXT NOT NULL DEFAULT 'hitbox',
    -- hitbox | capsule | sphere | raycast

    can_be_blocked BOOLEAN NOT NULL DEFAULT TRUE,
    can_be_parried BOOLEAN NOT NULL DEFAULT FALSE,
    can_be_dodged BOOLEAN NOT NULL DEFAULT TRUE,

    requires_server_confirmation BOOLEAN NOT NULL DEFAULT TRUE,

    can_pierce BOOLEAN NOT NULL DEFAULT FALSE,
    max_pierce_count INT NOT NULL DEFAULT 0,

    can_home BOOLEAN NOT NULL DEFAULT FALSE,
    homing_strength FLOAT NOT NULL DEFAULT 0.0,
    homing_turn_rate FLOAT NOT NULL DEFAULT 0.0,

    can_ricochet BOOLEAN NOT NULL DEFAULT FALSE,
    max_ricochet_count INT NOT NULL DEFAULT 0,

    spawn_offset_x FLOAT NOT NULL DEFAULT 0.0,
    spawn_offset_y FLOAT NOT NULL DEFAULT 1.2,
    spawn_offset_z FLOAT NOT NULL DEFAULT 0.8,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_skill_projectile_profile_skill
        FOREIGN KEY (skill_id)
        REFERENCES apeiron.skill(id)
        ON DELETE CASCADE,

    CONSTRAINT chk_skill_projectile_trajectory_type
        CHECK (trajectory_type IN (
            'linear',
            'ballistic',
            'lobbed',
            'homing',
            'beam'
        )),

    CONSTRAINT chk_skill_projectile_collision_mode
        CHECK (collision_mode IN (
            'hitbox',
            'capsule',
            'sphere',
            'raycast'
        )),

    CONSTRAINT chk_skill_projectile_speed_positive
        CHECK (projectile_speed > 0),

    CONSTRAINT chk_skill_projectile_radius_positive
        CHECK (projectile_radius > 0),

    CONSTRAINT chk_skill_projectile_lifetime_positive
        CHECK (max_lifetime_ms > 0),

    CONSTRAINT chk_skill_projectile_pierce_count
        CHECK (max_pierce_count >= 0),

    CONSTRAINT chk_skill_projectile_ricochet_count
        CHECK (max_ricochet_count >= 0)
);

CREATE INDEX idx_skill_projectile_trajectory_type
ON apeiron.skill_projectile_profile(trajectory_type);

CREATE INDEX idx_skill_projectile_collision_mode
ON apeiron.skill_projectile_profile(collision_mode);