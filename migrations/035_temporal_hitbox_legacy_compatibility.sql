-- =========================================================
-- TEMPORAL HITBOX LEGACY COMPATIBILITY
-- The final model is:
--   skill_hitbox_profile.motion_profile_id -> skill_hitbox_motion_profile.id
--   skill_hitbox_motion_sample.motion_profile_id -> skill_hitbox_motion_profile.id
--
-- Some recovered databases can still contain an old inverse column named
-- hitbox_profile_id on skill_hitbox_motion_profile. Preserve it only as
-- nullable compatibility data; fresh databases must not require it.
-- =========================================================

DO $$
BEGIN
    IF EXISTS (
        SELECT 1
        FROM information_schema.columns
        WHERE table_schema = 'apeiron'
          AND table_name = 'skill_hitbox_motion_profile'
          AND column_name = 'hitbox_profile_id'
    ) THEN
        ALTER TABLE apeiron.skill_hitbox_motion_profile
            ALTER COLUMN hitbox_profile_id DROP NOT NULL;
    END IF;
END $$;
