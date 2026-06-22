-- =========================================================
-- TEMPORAL SAMPLE ID DEFAULT COMPATIBILITY
-- Older recovered databases can store temporal sample ids as TEXT without
-- a default. Current seeds intentionally let the database assign sample ids.
-- =========================================================

ALTER TABLE IF EXISTS apeiron.skill_hitbox_motion_sample
    ALTER COLUMN id SET DEFAULT uuid_generate_v4()::text;
