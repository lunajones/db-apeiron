-- Recovery compatibility for legacy creature setup-policy columns.
-- The current behavior runtime uses setup_type plus movement_tactic; older
-- recovered schemas may still require setup_tactic even when no runtime reads it.

ALTER TABLE IF EXISTS apeiron.creature_skill_setup_policy
    ALTER COLUMN setup_tactic SET DEFAULT 'none',
    ALTER COLUMN setup_tactic DROP NOT NULL;
