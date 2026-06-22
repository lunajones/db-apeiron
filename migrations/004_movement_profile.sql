-- =========================================================
-- MOVEMENT PROFILE
-- APEIRON MMO - PHYSICS / FEEL LAYER (CLEAN)
-- =========================================================

CREATE TABLE IF NOT EXISTS apeiron.movement_profile (
    id TEXT PRIMARY KEY,

    -- =========================
    -- CORE PHYSICS (ESSENCIAL)
    -- =========================

    max_speed FLOAT NOT NULL DEFAULT 5.0,
    -- velocidade máxima base

    acceleration FLOAT NOT NULL DEFAULT 10.0,
    -- resposta ao input (quão rápido chega na velocidade)

    deceleration FLOAT NOT NULL DEFAULT 12.0,
    -- parada / controle de frenagem

    friction FLOAT NOT NULL DEFAULT 8.0,
    -- resistência contínua ao movimento (controle base do feel)

    gravity_multiplier FLOAT NOT NULL DEFAULT 1.0,
    -- peso da criatura no ar / queda

    mass FLOAT NOT NULL DEFAULT 1.0,
    -- impacto em knockback e colisões físicas

    momentum_retention FLOAT NOT NULL DEFAULT 0.6,
    -- inércia real: mantém movimento após parar input

    -- =========================
    -- ROTATION / CONTROL
    -- =========================

    turn_rate FLOAT NOT NULL DEFAULT 720.0,
    -- velocidade de rotação (responsividade PvP)

    air_control FLOAT NOT NULL DEFAULT 0.3,
    -- controle no ar (souls-like feel)

    strafe_efficiency FLOAT NOT NULL DEFAULT 0.8,
    -- eficiência lateral (strafe combat)

    backpedal_penalty FLOAT NOT NULL DEFAULT 0.7,
    -- penalidade de recuar (evita kite infinito forte)

    -- =========================
    -- DODGE (MOVEMENT-ONLY FEEL)
    -- =========================

    dodge_distance FLOAT NOT NULL DEFAULT 4.0,
    -- distância do dodge (dash roll feel)

    dodge_duration_ms INT NOT NULL DEFAULT 250,
    -- duração do deslocamento (iframe não entra aqui)

    -- =========================
    -- SPRINT / TRAVEL FEEL
    -- =========================

    sprint_multiplier FLOAT NOT NULL DEFAULT 1.4,
    -- multiplicador de velocidade ao sprint

    -- =========================
    -- TERRAIN / WORLD INTERACTION (BÁSICO)
    -- =========================

    slope_limit FLOAT NOT NULL DEFAULT 45.0,
    -- limite de subida

    slide_on_slope BOOLEAN NOT NULL DEFAULT TRUE,
    -- escorregar em inclinações fortes

    -- =========================
    -- SYSTEM FLAGS (SÓ MOVIMENTO)
    -- =========================

    is_airborne_enabled BOOLEAN NOT NULL DEFAULT TRUE,
    -- permite estados no ar (jump / knock-up / fall control)

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

ALTER TABLE IF EXISTS apeiron.movement_profile
    ADD COLUMN IF NOT EXISTS max_speed FLOAT NOT NULL DEFAULT 5.0,
    ADD COLUMN IF NOT EXISTS acceleration FLOAT NOT NULL DEFAULT 10.0,
    ADD COLUMN IF NOT EXISTS deceleration FLOAT NOT NULL DEFAULT 12.0,
    ADD COLUMN IF NOT EXISTS friction FLOAT NOT NULL DEFAULT 8.0,
    ADD COLUMN IF NOT EXISTS gravity_multiplier FLOAT NOT NULL DEFAULT 1.0,
    ADD COLUMN IF NOT EXISTS mass FLOAT NOT NULL DEFAULT 1.0,
    ADD COLUMN IF NOT EXISTS momentum_retention FLOAT NOT NULL DEFAULT 0.6,
    ADD COLUMN IF NOT EXISTS turn_rate FLOAT NOT NULL DEFAULT 720.0,
    ADD COLUMN IF NOT EXISTS air_control FLOAT NOT NULL DEFAULT 0.3,
    ADD COLUMN IF NOT EXISTS strafe_efficiency FLOAT NOT NULL DEFAULT 0.8,
    ADD COLUMN IF NOT EXISTS backpedal_penalty FLOAT NOT NULL DEFAULT 0.7,
    ADD COLUMN IF NOT EXISTS dodge_distance FLOAT NOT NULL DEFAULT 4.0,
    ADD COLUMN IF NOT EXISTS dodge_duration_ms INT NOT NULL DEFAULT 250,
    ADD COLUMN IF NOT EXISTS sprint_multiplier FLOAT NOT NULL DEFAULT 1.4,
    ADD COLUMN IF NOT EXISTS slope_limit FLOAT NOT NULL DEFAULT 45.0,
    ADD COLUMN IF NOT EXISTS slide_on_slope BOOLEAN NOT NULL DEFAULT TRUE,
    ADD COLUMN IF NOT EXISTS is_airborne_enabled BOOLEAN NOT NULL DEFAULT TRUE,
    ADD COLUMN IF NOT EXISTS created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP NOT NULL DEFAULT NOW();
