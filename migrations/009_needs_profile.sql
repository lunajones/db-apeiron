-- =========================================================
-- NEEDS PROFILE
-- APEIRON MMO - SURVIVAL / LIFE SYSTEM LAYER
-- =========================================================

CREATE TABLE IF NOT EXISTS apeiron.needs_profile (
    id TEXT PRIMARY KEY,

    -- =========================
    -- BASIC SURVIVAL NEEDS
    -- =========================

    hunger_enabled BOOLEAN NOT NULL DEFAULT FALSE,
    thirst_enabled BOOLEAN NOT NULL DEFAULT FALSE,
    fatigue_enabled BOOLEAN NOT NULL DEFAULT FALSE,

    hunger_decay_per_hour FLOAT NOT NULL DEFAULT 0.0,
    thirst_decay_per_hour FLOAT NOT NULL DEFAULT 0.0,
    fatigue_decay_per_hour FLOAT NOT NULL DEFAULT 0.0,

    hunger_threshold FLOAT NOT NULL DEFAULT 0.3,
    thirst_threshold FLOAT NOT NULL DEFAULT 0.3,
    fatigue_threshold FLOAT NOT NULL DEFAULT 0.3,
    -- abaixo disso: penalidades começam

    -- =========================
    -- EFFECTS (GAMEPLAY IMPACT)
    -- =========================

    hunger_damage_threshold FLOAT NOT NULL DEFAULT 0.0,
    -- se chegar nisso, começa perder HP

    stamina_regen_penalty_at_low_needs FLOAT NOT NULL DEFAULT 0.5,
    -- reduz regen de stamina quando necessidades baixas

    movement_penalty_at_low_needs FLOAT NOT NULL DEFAULT 0.5,
    -- redução de velocidade geral

    combat_performance_penalty FLOAT NOT NULL DEFAULT 0.5,
    -- precisão / reação / agressividade reduzida

    -- =========================
    -- RECOVERY RULES
    -- =========================

    food_saturation_rate FLOAT NOT NULL DEFAULT 1.0,
    water_saturation_rate FLOAT NOT NULL DEFAULT 1.0,
    rest_recovery_rate FLOAT NOT NULL DEFAULT 1.0,

    -- =========================
    -- AI INFLUENCE (DECISION ONLY)
    -- =========================

    panic_threshold FLOAT NOT NULL DEFAULT 0.2,
    -- fome/sede/fadiga alta pode gerar comportamento errático

    aggression_increase_when_starving FLOAT NOT NULL DEFAULT 0.3,
    fear_increase_when_starving FLOAT NOT NULL DEFAULT 0.3,

    -- =========================
    -- FLAGS
    -- =========================

    needs_enabled BOOLEAN NOT NULL DEFAULT FALSE,
    -- liga/desliga sistema inteiro para criatura

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
