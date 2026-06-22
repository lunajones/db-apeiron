-- =========================================================
-- LEGACY MOVEMENT ACTION COLUMN COMPATIBILITY
-- Older recovered databases can contain historical columns on
-- movement_action_contract. They are preserved, but must not block the
-- reconstructed canonical contract rows that use id/action_type fields.
-- =========================================================

ALTER TABLE IF EXISTS apeiron.movement_action_contract
    ALTER COLUMN ability_key DROP NOT NULL,
    ALTER COLUMN movement_type DROP NOT NULL,
    ALTER COLUMN contract_version DROP NOT NULL,
    ALTER COLUMN contract_hash DROP NOT NULL;
