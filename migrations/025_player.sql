-- =========================================================
-- PLAYER
-- APEIRON MMO - HUMAN RUNTIME ENTITY
-- =========================================================

CREATE TABLE IF NOT EXISTS apeiron.player (
    id TEXT PRIMARY KEY,

    -- =========================
    -- ACCOUNT LINK
    -- =========================

    account_id TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL UNIQUE,

    -- =========================
    -- CREATURE RUNTIME BASE
    -- =========================

    -- Nullable: the persistent player record does not require a permanent live-world body row.
    -- The runtime body (creature_instance) is spawned at login; see aaa-character-progression-roadmap.md.
    creature_instance_id TEXT UNIQUE,

    -- =========================
    -- PROGRESSION
    -- =========================

    level INT NOT NULL DEFAULT 1,
    experience BIGINT NOT NULL DEFAULT 0,
    attribute_points INT NOT NULL DEFAULT 0,

    -- =========================
    -- CORE ATTRIBUTES
    -- =========================

    strength FLOAT NOT NULL DEFAULT 1.0,
    dexterity FLOAT NOT NULL DEFAULT 1.0,
    intelligence FLOAT NOT NULL DEFAULT 1.0,
    endurance FLOAT NOT NULL DEFAULT 1.0,

    -- =========================
    -- COMBAT FLAGS
    -- =========================

    pvp_enabled BOOLEAN NOT NULL DEFAULT TRUE,
    is_in_safe_zone BOOLEAN NOT NULL DEFAULT FALSE,

    -- =========================
    -- SOCIAL
    -- =========================

    guild_id TEXT,
    party_id TEXT,

    reputation FLOAT NOT NULL DEFAULT 0.0,

    -- =========================
    -- ECONOMY
    -- =========================

    coin BIGINT NOT NULL DEFAULT 0,

    -- =========================
    -- LIFECYCLE
    -- =========================

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_player_creature
        FOREIGN KEY (creature_instance_id)
        REFERENCES apeiron.creature_instance(id)
);

CREATE INDEX IF NOT EXISTS idx_player_account
ON apeiron.player(account_id);

CREATE INDEX IF NOT EXISTS idx_player_creature
ON apeiron.player(creature_instance_id);
