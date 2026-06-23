-- =========================================================
-- SKILL IMPACT CONTROL MOTION CONTRACT
-- Physical control responses must be DB-driven. A control effect can mark
-- the target as controlled, but the server also needs contract-backed motion
-- parameters to execute push/carry/knockback without per-skill literals.
-- =========================================================

ALTER TABLE apeiron.skill_impact_profile
    ADD COLUMN IF NOT EXISTS control_distance_cm FLOAT NOT NULL DEFAULT 0.0,
    ADD COLUMN IF NOT EXISTS control_speed_cm_s FLOAT NOT NULL DEFAULT 0.0,
    ADD COLUMN IF NOT EXISTS control_direction_policy TEXT NOT NULL DEFAULT '';

ALTER TABLE apeiron.skill_impact_profile
    DROP CONSTRAINT IF EXISTS chk_skill_impact_control_distance;

ALTER TABLE apeiron.skill_impact_profile
    ADD CONSTRAINT chk_skill_impact_control_distance
    CHECK (control_distance_cm >= 0 AND control_speed_cm_s >= 0);

CREATE INDEX IF NOT EXISTS idx_skill_impact_control_direction_policy
ON apeiron.skill_impact_profile(control_direction_policy);
