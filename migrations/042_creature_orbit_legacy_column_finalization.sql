-- =========================================================
-- CREATURE ORBIT LEGACY COLUMN FINALIZATION
-- Some recovered databases already contain an older creature_orbit_policy table
-- with tuning columns that the modern behavior contract no longer treats as
-- authoritative. Keep the columns for compatibility, but make them non-blocking
-- for reconstructed seeds.
-- =========================================================

DO $$
DECLARE
    legacy_column TEXT;
BEGIN
    FOREACH legacy_column IN ARRAY ARRAY[
        'preferred_radius_cm',
        'min_radius_cm',
        'max_radius_cm'
    ]
    LOOP
        IF EXISTS (
            SELECT 1
            FROM information_schema.columns
            WHERE table_schema = 'apeiron'
              AND table_name = 'creature_orbit_policy'
              AND column_name = legacy_column
        ) THEN
            EXECUTE format(
                'ALTER TABLE apeiron.creature_orbit_policy ALTER COLUMN %I SET DEFAULT 0',
                legacy_column
            );
            EXECUTE format(
                'ALTER TABLE apeiron.creature_orbit_policy ALTER COLUMN %I DROP NOT NULL',
                legacy_column
            );
        END IF;
    END LOOP;
END $$;
