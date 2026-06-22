-- =========================================================
-- COMBAT STYLE PROFILE
-- APEIRON MMO - BEHAVIOR / TENDENCY LAYER
-- =========================================================

CREATE TABLE apeiron.combat_style_profile (
    id TEXT PRIMARY KEY,

    -- =========================
    -- CORE TENDENCIES
    -- =========================

    archetype TEXT NOT NULL,

    aggression FLOAT NOT NULL DEFAULT 0.5,
    defense FLOAT NOT NULL DEFAULT 0.5,
    patience FLOAT NOT NULL DEFAULT 0.5,
    risk_tolerance FLOAT NOT NULL DEFAULT 0.5,

    -- =========================
    -- ENGAGEMENT STYLE
    -- =========================

    preferred_range FLOAT NOT NULL DEFAULT 2.0,
    chase_tendency FLOAT NOT NULL DEFAULT 0.5,
    disengage_threshold FLOAT NOT NULL DEFAULT 0.3,

    -- =========================
    -- ATTACK BEHAVIOR
    -- =========================

    combo_preference FLOAT NOT NULL DEFAULT 0.5,
    feint_usage FLOAT NOT NULL DEFAULT 0.2,
    punish_window_awareness FLOAT NOT NULL DEFAULT 0.5,
    aggression_spike_chance FLOAT NOT NULL DEFAULT 0.1,

    -- =========================
    -- DEFENSIVE BEHAVIOR
    -- =========================

    dodge_frequency FLOAT NOT NULL DEFAULT 0.5,
    block_frequency FLOAT NOT NULL DEFAULT 0.5,
    parry_aggressiveness FLOAT NOT NULL DEFAULT 0.3,
    panic_threshold FLOAT NOT NULL DEFAULT 0.2,

    -- =========================
    -- MOVEMENT STYLE
    -- =========================

    strafe_usage FLOAT NOT NULL DEFAULT 0.5,
    circle_target FLOAT NOT NULL DEFAULT 0.4,
    backstep_usage FLOAT NOT NULL DEFAULT 0.3,
    reposition_frequency FLOAT NOT NULL DEFAULT 0.5,

    -- =========================
    -- VARIATION
    -- =========================

    randomness FLOAT NOT NULL DEFAULT 0.1,
    consistency FLOAT NOT NULL DEFAULT 0.8,

    -- =========================
    -- TACTICAL MODIFIERS
    -- =========================

    target_switching FLOAT NOT NULL DEFAULT 0.2,
    focus_fire FLOAT NOT NULL DEFAULT 0.5,
    retaliation_bias FLOAT NOT NULL DEFAULT 0.6,

    -- =========================
    -- FLAGS
    -- =========================

    is_elite BOOLEAN NOT NULL DEFAULT FALSE,
    is_coward BOOLEAN NOT NULL DEFAULT FALSE,
    is_duelist BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);