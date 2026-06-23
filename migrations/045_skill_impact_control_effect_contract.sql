-- =========================================================
-- SKILL IMPACT CONTROL EFFECT CONTRACT
-- =========================================================
--
-- Explicit control-effect fields for impact profiles. These fields keep
-- push/stagger/carry semantics DB-owned instead of deriving control runtime
-- from hitstop, strings in movement code, or per-skill branches.

ALTER TABLE apeiron.skill_impact_profile
    ADD COLUMN IF NOT EXISTS control_type TEXT NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS control_effect_duration_ms INT NOT NULL DEFAULT 0,
    ADD COLUMN IF NOT EXISTS control_release_policy_id TEXT NOT NULL DEFAULT '';

ALTER TABLE apeiron.skill_impact_profile
    DROP CONSTRAINT IF EXISTS chk_skill_impact_control_duration;

ALTER TABLE apeiron.skill_impact_profile
    ADD CONSTRAINT chk_skill_impact_control_duration
    CHECK (control_effect_duration_ms >= 0);

CREATE INDEX IF NOT EXISTS idx_skill_impact_control_type
ON apeiron.skill_impact_profile(control_type);
