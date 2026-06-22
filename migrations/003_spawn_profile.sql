-- =========================================================
-- SPAWN PROFILE
-- APEIRON MMO - WORLD SIMULATION LAYER
-- =========================================================

CREATE TABLE IF NOT EXISTS apeiron.spawn_profile (
    id TEXT PRIMARY KEY,

    -- =========================
    -- BIOME RULES
    -- =========================

    allowed_biomes JSONB NOT NULL DEFAULT '[]'::jsonb,
    -- biomas onde essa criatura pode spawnar
    -- regra dura do sistema de mundo
    -- ex:
    -- ["forest", "swamp", "mountain"]

    biome_weights JSONB NOT NULL DEFAULT '[]'::jsonb,
    -- distribuição relativa dentro dos biomas permitidos
    -- ex:
    -- [
    --   {"biome_id": "forest", "weight": 1.0},
    --   {"biome_id": "swamp", "weight": 0.4}
    -- ]

    -- =========================
    -- REGION CONTROL
    -- =========================

    allowed_regions JSONB NOT NULL DEFAULT '[]'::jsonb,
    -- regiões macro do mundo onde pode existir

    region_weights JSONB NOT NULL DEFAULT '[]'::jsonb,
    -- peso de distribuição por região

    -- =========================
    -- POPULATION CONTROL
    -- =========================

    density_base FLOAT NOT NULL DEFAULT 1.0,
    -- densidade base global (spawn rate relativa)

    density_cap INT,
    -- limite absoluto de entidades simultâneas por área
    -- NULL = sem limite (cuidado com isso)

    respawn_time_seconds INT NOT NULL DEFAULT 300,
    -- tempo base de respawn

    spawn_batch_size INT NOT NULL DEFAULT 1,
    -- quantas entidades podem spawnar por evento

    spawn_variance FLOAT NOT NULL DEFAULT 0.2,
    -- aleatoriedade no spawn (evita padrão previsível)

    -- =========================
    -- WORLD BEHAVIOR RULES
    -- =========================

    territorial_bias FLOAT NOT NULL DEFAULT 0.0,
    -- tendência a respeitar território existente (0 = ignora, 1 = forte bloqueio)

    migration_allowed BOOLEAN NOT NULL DEFAULT FALSE,
    -- pode migrar entre regiões automaticamente

    clustering_strength FLOAT NOT NULL DEFAULT 0.5,
    -- tendência a spawnar em grupos vs disperso

    -- =========================
    -- SYSTEM FLAGS
    -- =========================

    is_dynamic BOOLEAN NOT NULL DEFAULT TRUE,
    -- se TRUE: o mundo pode ajustar spawn com base em pressão do jogador

    anti_overcrowd BOOLEAN NOT NULL DEFAULT TRUE,
    -- evita superpopulação automática

    pvp_zone_modifier FLOAT NOT NULL DEFAULT 1.0,
    -- multiplica spawn em zonas de PvP aberto

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);