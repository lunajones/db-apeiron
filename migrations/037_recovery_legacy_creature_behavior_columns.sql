-- Recovery compatibility for partially restored creature behavior schemas.
-- Older local schemas kept UI/catalog columns as required fields, while the
-- reconstructed runtime contract now stores authoritative behavior in JSON
-- policies keyed by creature_template_id.

ALTER TABLE IF EXISTS apeiron.creature_behavior_runtime_contract
    ALTER COLUMN display_name DROP NOT NULL,
    ALTER COLUMN combat_role_id DROP NOT NULL;
