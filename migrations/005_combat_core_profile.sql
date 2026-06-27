-- =========================================================
-- COMBAT CORE PROFILE
-- APEIRON MMO - PvP CORE RULESET (PHYSICS ONLY)
-- =========================================================

CREATE TABLE IF NOT EXISTS apeiron.combat_core_profile (
    id TEXT PRIMARY KEY,

    -- =========================
    -- DAMAGE MODEL
    -- =========================

    physical_defense FLOAT NOT NULL DEFAULT 0.0,
    magic_defense FLOAT NOT NULL DEFAULT 0.0,

    -- 3-resistance model (ratings, mitigated via rating/(rating+K) curve). Replaces the two
    -- defenses above (kept until callers are migrated). See
    -- server-apeiron/docs/roadmap/aaa-damage-types-resistances-weapons-roadmap.md
    physical_resistance_rating FLOAT NOT NULL DEFAULT 0.0,
    chemical_resistance_rating FLOAT NOT NULL DEFAULT 0.0,
    biological_resistance_rating FLOAT NOT NULL DEFAULT 0.0,
    resistance_cap FLOAT NOT NULL DEFAULT 0.85,

    critical_chance FLOAT NOT NULL DEFAULT 0.05,
    critical_multiplier FLOAT NOT NULL DEFAULT 1.5,

    damage_taken_multiplier FLOAT NOT NULL DEFAULT 1.0,
    damage_dealt_multiplier FLOAT NOT NULL DEFAULT 1.0,

    -- =========================
    -- STAMINA SYSTEM
    -- =========================

    max_stamina FLOAT NOT NULL DEFAULT 100.0,
    stamina_regen_per_sec FLOAT NOT NULL DEFAULT 10.0,

    dodge_stamina_cost FLOAT NOT NULL DEFAULT 25.0,
    sprint_stamina_cost_per_sec FLOAT NOT NULL DEFAULT 6.0,
    block_stamina_cost_per_sec FLOAT NOT NULL DEFAULT 8.0,
    attack_stamina_cost FLOAT NOT NULL DEFAULT 15.0,

    stamina_exhaustion_threshold FLOAT NOT NULL DEFAULT 0.15,
    stamina_zero_regen_multiplier FLOAT NOT NULL DEFAULT 0.5,

    -- =========================
    -- POSTURE SYSTEM
    -- =========================

    max_posture FLOAT NOT NULL DEFAULT 100.0,
    posture_recovery_rate FLOAT NOT NULL DEFAULT 12.0,
    posture_damage_multiplier FLOAT NOT NULL DEFAULT 1.0,
    posture_break_duration_ms INT NOT NULL DEFAULT 2500,

    -- =========================
    -- BLOCK / PARRY
    -- =========================

    block_damage_reduction FLOAT NOT NULL DEFAULT 0.6,
    parry_window_ms INT NOT NULL DEFAULT 300,
    parry_reward_multiplier FLOAT NOT NULL DEFAULT 1.8,

    can_block BOOLEAN NOT NULL DEFAULT TRUE,
    can_parry BOOLEAN NOT NULL DEFAULT TRUE,

    -- =========================
    -- DODGE RULES
    -- =========================

    dodge_iframe_ms INT NOT NULL DEFAULT 250,
    dodge_cooldown_ms INT NOT NULL DEFAULT 900,

    -- =========================
    -- CONTROL RESISTANCES
    -- =========================

    stun_resistance FLOAT NOT NULL DEFAULT 1.0,
    root_resistance FLOAT NOT NULL DEFAULT 1.0,
    knockback_resistance FLOAT NOT NULL DEFAULT 1.0,
    cc_duration_multiplier FLOAT NOT NULL DEFAULT 1.0,

    -- =========================
    -- FLAGS
    -- =========================

    is_boss BOOLEAN NOT NULL DEFAULT FALSE,
    is_pvp_immune BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
