-- =========================================================
-- TEMPORAL HITBOX SAMPLE ID FINALIZATION
-- Fresh Apeiron schemas use BIGSERIAL sample ids from
-- 028_temporal_melee_hitbox.sql. Some recovered schemas can have TEXT ids
-- without a default. Normalize only that legacy shape without changing the
-- final BIGSERIAL model.
-- =========================================================

DO $$
DECLARE
    id_type TEXT;
BEGIN
    SELECT data_type
    INTO id_type
    FROM information_schema.columns
    WHERE table_schema = 'apeiron'
      AND table_name = 'skill_hitbox_motion_sample'
      AND column_name = 'id';

    IF id_type = 'text' THEN
        ALTER TABLE apeiron.skill_hitbox_motion_sample
            ALTER COLUMN id SET DEFAULT uuid_generate_v4()::text;
    END IF;
END $$;
