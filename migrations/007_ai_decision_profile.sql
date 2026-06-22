-- =========================================================
-- AI DECISION PROFILE
-- APEIRON MMO - COGNITION / TIMING LAYER
-- =========================================================

CREATE TABLE IF NOT EXISTS apeiron.ai_decision_profile (
    id TEXT PRIMARY KEY,

    -- =========================
    -- REACTION & TIMING
    -- =========================

    reaction_time_ms INT NOT NULL DEFAULT 250,
    decision_interval_ms INT NOT NULL DEFAULT 300,
    combat_engage_delay_ms INT NOT NULL DEFAULT 200,
    input_delay_variance_ms INT NOT NULL DEFAULT 50,

    -- =========================
    -- TARGET LOGIC
    -- =========================

    target_switch_delay_ms INT NOT NULL DEFAULT 400,
    target_priority_bias FLOAT NOT NULL DEFAULT 0.5,
    threat_evaluation_speed FLOAT NOT NULL DEFAULT 0.5,

    -- =========================
    -- COMBAT FLOW
    -- =========================

    combo_interrupt_tolerance FLOAT NOT NULL DEFAULT 0.5,
    greed_vs_safety FLOAT NOT NULL DEFAULT 0.5,
    punish_recognition_speed FLOAT NOT NULL DEFAULT 0.5,

    -- =========================
    -- HUMANIZATION
    -- =========================

    mistake_chance FLOAT NOT NULL DEFAULT 0.05,
    hesitation_factor FLOAT NOT NULL DEFAULT 0.2,
    overcommit_risk FLOAT NOT NULL DEFAULT 0.3,

    -- =========================
    -- GROUP / COORDINATION
    -- =========================

    assist_ally_priority FLOAT NOT NULL DEFAULT 0.5,
    pack_sync_factor FLOAT NOT NULL DEFAULT 0.5,
    focus_fire_coordination FLOAT NOT NULL DEFAULT 0.5,

    -- =========================
    -- FLAGS
    -- =========================

    is_predictive BOOLEAN NOT NULL DEFAULT FALSE,
    is_adaptive BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
