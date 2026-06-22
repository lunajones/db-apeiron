-- Recovery compatibility for the legacy skill_movement_effect endpoint.
-- Runtime skill movement should prefer movement_action_contract bindings; these
-- columns keep GetSkillMovementEffect(skill_id) and recovered legacy seeds usable.

ALTER TABLE IF EXISTS apeiron.skill_movement_effect
    ADD COLUMN IF NOT EXISTS windup_lock_ms INT NOT NULL DEFAULT 0,
    ADD COLUMN IF NOT EXISTS recovery_lock_ms INT NOT NULL DEFAULT 0,
    ADD COLUMN IF NOT EXISTS can_rotate BOOLEAN NOT NULL DEFAULT TRUE,
    ADD COLUMN IF NOT EXISTS ignores_collision BOOLEAN NOT NULL DEFAULT FALSE,
    ADD COLUMN IF NOT EXISTS metadata JSONB NOT NULL DEFAULT '{}'::jsonb;
