-- =========================================================
-- SKILL MOVEMENT EFFECT COMPATIBILITY
-- Compatibility surface for older GetSkillMovementEffect clients/tools.
-- =========================================================

CREATE TABLE IF NOT EXISTS apeiron.skill_movement_effect (
    id TEXT PRIMARY KEY,
    skill_id TEXT NOT NULL UNIQUE,
    movement_type TEXT NOT NULL,
    distance FLOAT NOT NULL DEFAULT 0.0,
    speed FLOAT NOT NULL DEFAULT 0.0,
    duration_ms INT NOT NULL DEFAULT 0,
    windup_lock_ms INT NOT NULL DEFAULT 0,
    recovery_lock_ms INT NOT NULL DEFAULT 0,
    can_rotate BOOLEAN NOT NULL DEFAULT TRUE,
    ignores_collision BOOLEAN NOT NULL DEFAULT FALSE,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_skill_movement_effect_skill
        FOREIGN KEY (skill_id)
        REFERENCES apeiron.skill(id)
        ON DELETE CASCADE,
    CONSTRAINT chk_skill_movement_effect_distance CHECK (distance >= 0),
    CONSTRAINT chk_skill_movement_effect_duration CHECK (duration_ms >= 0)
);

CREATE INDEX IF NOT EXISTS idx_skill_movement_effect_skill
ON apeiron.skill_movement_effect(skill_id);
