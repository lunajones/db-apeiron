-- =========================================================
-- COMBAT DEFENSE CONTRACTS
-- Reconstructed from block/stamina regression thread.
-- =========================================================

CREATE TABLE IF NOT EXISTS apeiron.combat_defense_contract (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    defense_type TEXT NOT NULL,
    frontal_arc_deg FLOAT NOT NULL DEFAULT 120.0,
    defender_margin_left_ratio FLOAT NOT NULL DEFAULT 0.30,
    defender_margin_right_ratio FLOAT NOT NULL DEFAULT 0.30,
    stamina_damage_only_on_block BOOLEAN NOT NULL DEFAULT TRUE,
    health_damage_on_unblocked_hit BOOLEAN NOT NULL DEFAULT TRUE,
    posture_damage_on_block BOOLEAN NOT NULL DEFAULT TRUE,
    perfect_block_window_ms INT NOT NULL DEFAULT 0,
    parry_window_ms INT NOT NULL DEFAULT 0,
    guard_damage_multiplier FLOAT NOT NULL DEFAULT 1.0,
    block_stamina_drain_per_second FLOAT NOT NULL DEFAULT 0.0,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_combat_defense_arc CHECK (frontal_arc_deg >= 0 AND frontal_arc_deg <= 360),
    CONSTRAINT chk_combat_defense_margin CHECK (defender_margin_left_ratio >= 0 AND defender_margin_right_ratio >= 0),
    CONSTRAINT chk_combat_defense_windows CHECK (perfect_block_window_ms >= 0 AND parry_window_ms >= 0)
);

CREATE INDEX IF NOT EXISTS idx_combat_defense_type
ON apeiron.combat_defense_contract(defense_type);
