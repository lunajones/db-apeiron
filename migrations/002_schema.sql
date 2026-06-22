-- =========================================================
-- SCHEMA + PERMISSIONS
-- APEIRON MMO - DATABASE FOUNDATION
-- =========================================================

-- 1. CREATE SCHEMA
CREATE SCHEMA IF NOT EXISTS apeiron;

-- 2. SET OWNERSHIP
ALTER SCHEMA apeiron OWNER TO postgres;

-- 3. GRANT USAGE ON SCHEMA
GRANT USAGE ON SCHEMA apeiron TO postgres;

-- (opcional) se tiver usuário de aplicação separado
-- CREATE ROLE apeiron_app LOGIN PASSWORD 'strong_password_here';

-- 4. GRANT PERMISSIONS FOR APP USER
-- Troque "apeiron_app" pelo user real da aplicação
GRANT USAGE, CREATE ON SCHEMA apeiron TO postgres;

-- 5. DEFAULT PRIVILEGES (IMPORTANTÍSSIMO)
-- garante que tabelas futuras já herdam permissão

ALTER DEFAULT PRIVILEGES IN SCHEMA apeiron
GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES
TO postgres;

ALTER DEFAULT PRIVILEGES IN SCHEMA apeiron
GRANT USAGE, SELECT, UPDATE ON SEQUENCES
TO postgres;

-- 6. OPTIONAL: FORCE EVERYTHING INTO SCHEMA
-- (evita tabela solta no public)
SET search_path TO apeiron;

CREATE TABLE IF NOT EXISTS apeiron.schema_migrations (
    file_name TEXT PRIMARY KEY,
    applied_at TIMESTAMP NOT NULL DEFAULT NOW()
);