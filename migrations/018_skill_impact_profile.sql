-- =========================================================
-- SKILL IMPACT PROFILE
-- APEIRON MMO - HIT REACTION / IMPACT RESULT CONFIG
-- =========================================================

CREATE TABLE apeiron.skill_impact_profile (
    skill_id TEXT PRIMARY KEY,

    impact_type TEXT NOT NULL DEFAULT 'normal',
    -- normal | heavy | stagger | knockdown | launch | pull | grab | no_reaction

    hit_reaction TEXT NOT NULL DEFAULT 'flinch',
    -- none | flinch | stagger | knockdown | launch | interrupt

    poise_damage FLOAT NOT NULL DEFAULT 0.0,
    stagger_power FLOAT NOT NULL DEFAULT 0.0,

    interrupt_power FLOAT NOT NULL DEFAULT 0.0,
    guard_damage_multiplier FLOAT NOT NULL DEFAULT 1.0,

    bounce_on_shield BOOLEAN NOT NULL DEFAULT FALSE,
    destroy_on_hit BOOLEAN NOT NULL DEFAULT TRUE,
    stick_on_hit BOOLEAN NOT NULL DEFAULT FALSE,

    hitstop_ms INT NOT NULL DEFAULT 0,
    screenshake_strength FLOAT NOT NULL DEFAULT 0.0,

    knockback_force FLOAT NOT NULL DEFAULT 0.0,
    knockback_upward_force FLOAT NOT NULL DEFAULT 0.0,

    pull_force FLOAT NOT NULL DEFAULT 0.0,

    applies_status_effect BOOLEAN NOT NULL DEFAULT FALSE,
    status_effect_id TEXT,
    status_effect_chance FLOAT NOT NULL DEFAULT 0.0,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_skill_impact_profile_skill
        FOREIGN KEY (skill_id)
        REFERENCES apeiron.skill(id)
        ON DELETE CASCADE,
    
    CONSTRAINT fk_skill_impact_status_effect
        FOREIGN KEY (status_effect_id)
        REFERENCES apeiron.status_effect(id)
        ON DELETE CASCADE,

    CONSTRAINT chk_skill_impact_type
        CHECK (impact_type IN (
            'normal',
            'heavy',
            'stagger',
            'knockdown',
            'launch',
            'pull',
            'grab',
            'no_reaction'
        )),

    CONSTRAINT chk_skill_hit_reaction
        CHECK (hit_reaction IN (
            'none',
            'flinch',
            'stagger',
            'knockdown',
            'launch',
            'interrupt'
        )),

    CONSTRAINT chk_skill_status_chance
        CHECK (status_effect_chance >= 0.0 AND status_effect_chance <= 1.0),

    CONSTRAINT chk_skill_hitstop
        CHECK (hitstop_ms >= 0)
);

CREATE INDEX idx_skill_impact_type
ON apeiron.skill_impact_profile(impact_type);

CREATE INDEX idx_skill_impact_reaction
ON apeiron.skill_impact_profile(hit_reaction);

CREATE INDEX idx_skill_impact_status
ON apeiron.skill_impact_profile(status_effect_id);