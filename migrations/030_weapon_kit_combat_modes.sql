-- =========================================================
-- PLAYER WEAPON KIT / COMBAT MODES
-- Reconstructed from sword-and-shield weapon kit roadmap.
-- =========================================================

CREATE TABLE IF NOT EXISTS apeiron.weapon_kit (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    primary_weapon_type TEXT NOT NULL,
    offhand_weapon_type TEXT,
    is_enabled BOOLEAN NOT NULL DEFAULT TRUE,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS apeiron.weapon_combat_mode (
    id TEXT PRIMARY KEY,
    weapon_kit_id TEXT NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    mode_index INT NOT NULL,
    switch_duration_ms INT NOT NULL DEFAULT 500,
    is_enabled BOOLEAN NOT NULL DEFAULT TRUE,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_weapon_mode_kit
        FOREIGN KEY (weapon_kit_id)
        REFERENCES apeiron.weapon_kit(id)
        ON DELETE CASCADE,
    CONSTRAINT uq_weapon_mode_index UNIQUE (weapon_kit_id, mode_index)
);

CREATE TABLE IF NOT EXISTS apeiron.weapon_combat_mode_skill_slot (
    id BIGSERIAL PRIMARY KEY,
    combat_mode_id TEXT NOT NULL,
    input_slot TEXT NOT NULL,
    skill_id TEXT,
    is_basic_attack BOOLEAN NOT NULL DEFAULT FALSE,
    is_fatality BOOLEAN NOT NULL DEFAULT FALSE,
    is_enabled BOOLEAN NOT NULL DEFAULT TRUE,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_weapon_mode_slot_mode
        FOREIGN KEY (combat_mode_id)
        REFERENCES apeiron.weapon_combat_mode(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_weapon_mode_slot_skill
        FOREIGN KEY (skill_id)
        REFERENCES apeiron.skill(id)
        ON DELETE SET NULL,
    CONSTRAINT uq_weapon_mode_input UNIQUE (combat_mode_id, input_slot)
);

CREATE TABLE IF NOT EXISTS apeiron.player_weapon_kit_loadout (
    player_id TEXT NOT NULL,
    weapon_kit_id TEXT NOT NULL,
    active_mode_id TEXT NOT NULL,
    secondary_mode_id TEXT,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (player_id, weapon_kit_id),
    CONSTRAINT fk_player_weapon_loadout_kit
        FOREIGN KEY (weapon_kit_id)
        REFERENCES apeiron.weapon_kit(id),
    CONSTRAINT fk_player_weapon_loadout_active_mode
        FOREIGN KEY (active_mode_id)
        REFERENCES apeiron.weapon_combat_mode(id),
    CONSTRAINT fk_player_weapon_loadout_secondary_mode
        FOREIGN KEY (secondary_mode_id)
        REFERENCES apeiron.weapon_combat_mode(id)
);
