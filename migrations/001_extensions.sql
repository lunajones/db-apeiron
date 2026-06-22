-- =========================================
-- APEIRON MMO - BASE EXTENSIONS
-- =========================================

-- UUID support (essencial pra MMO distribuído)
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Melhor performance pra texto/lookup
CREATE EXTENSION IF NOT EXISTS "pg_trgm";

-- suporte a JSONB index avançado
CREATE EXTENSION IF NOT EXISTS "btree_gin";

-- =========================================
-- FUTURO MMO NOTES
-- =========================================
-- uuid = entidades globais (player, creature, world objects)
-- jsonb = profiles data-driven (behavior, combat, needs)
-- trigram = search de nomes/skills/items