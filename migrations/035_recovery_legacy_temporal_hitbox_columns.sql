-- =========================================================
-- LEGACY TEMPORAL HITBOX COLUMN COMPATIBILITY
-- Older recovered databases can contain historical columns on temporal
-- hitbox tables. Preserve them, but keep the reconstructed motion-profile
-- model authoritative.
-- =========================================================

ALTER TABLE IF EXISTS apeiron.skill_hitbox_motion_profile
    ALTER COLUMN hitbox_profile_id DROP NOT NULL;
