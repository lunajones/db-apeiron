CREATE TABLE IF NOT EXISTS apeiron.skill_set (
    id TEXT PRIMARY KEY,

    -- =========================
    -- IDENTITY
    -- =========================

    name TEXT NOT NULL,
    -- nome visual do conjunto

    description TEXT NOT NULL DEFAULT '',
    -- debug / designer notes

    -- =========================
    -- META FLAGS
    -- =========================

    is_player_usable BOOLEAN NOT NULL DEFAULT FALSE,
    -- pode ser usado por player

    is_npc_usable BOOLEAN NOT NULL DEFAULT TRUE,
    -- pode ser usado por AI

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
