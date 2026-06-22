-- Recovery compatibility for stale creature setup check constraints.
-- Reconstructed behavior contracts use setup_type values such as
-- moving_windup, chase_windup, and pressure_counter. Legacy constraints came
-- from an older policy vocabulary and must not reject the current contract set.

ALTER TABLE IF EXISTS apeiron.creature_skill_setup_policy
    ALTER COLUMN setup_tactic SET DEFAULT 'any';

ALTER TABLE IF EXISTS apeiron.creature_skill_setup_policy
    DROP CONSTRAINT IF EXISTS chk_creature_skill_setup_tactic,
    DROP CONSTRAINT IF EXISTS chk_creature_skill_setup_type;
