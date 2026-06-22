-- =========================================================
-- SKILL
-- APEIRON MMO - UNIVERSAL SKILL SYSTEM
-- =========================================================

CREATE TABLE IF NOT EXISTS apeiron.skill (
    id TEXT PRIMARY KEY,

    -- =========================
    -- IDENTITY
    -- =========================

    name TEXT NOT NULL,
    -- nome da habilidade

    description TEXT NOT NULL DEFAULT '',
    -- tooltip / lore / debug

    archetype TEXT NOT NULL,
    -- ex:
    -- melee, ranged, spell, mobility,
    -- defensive, utility, summon

    skill_type TEXT NOT NULL,
    -- ex:
    -- attack, dodge, block, parry,
    -- buff, debuff, cc, movement

    -- =========================
    -- RESOURCE COST
    -- =========================

    stamina_cost FLOAT NOT NULL DEFAULT 0.0,
    mana_cost FLOAT NOT NULL DEFAULT 0.0,
    health_cost FLOAT NOT NULL DEFAULT 0.0,

    -- =========================
    -- TIMINGS (SOULS-LIKE CORE)
    -- =========================

    windup_ms INT NOT NULL DEFAULT 0,
    -- startup antes do hit

    active_frames_ms INT NOT NULL DEFAULT 100,
    -- janela ativa da hitbox

    recovery_ms INT NOT NULL DEFAULT 0,
    -- endlag após execução

    cast_time_ms INT NOT NULL DEFAULT 0,
    -- cast channel

    cooldown_ms INT NOT NULL DEFAULT 0,

    cancel_window_start_ms INT NOT NULL DEFAULT 0,
    cancel_window_end_ms INT NOT NULL DEFAULT 0,
    -- janela de cancel

    iframe_start_ms INT NOT NULL DEFAULT 0,
    iframe_end_ms INT NOT NULL DEFAULT 0,
    -- iframe window
    -- 0/0 = sem iframe

    -- =========================
    -- RANGE / TARGETING
    -- =========================

    min_range FLOAT NOT NULL DEFAULT 0.0,
    max_range FLOAT NOT NULL DEFAULT 2.0,

    cone_angle FLOAT NOT NULL DEFAULT 0.0,
    -- cleave / frontal cone

    max_targets INT NOT NULL DEFAULT 1,

    target_type TEXT NOT NULL DEFAULT 'enemy',
    -- enemy, ally, self, ground, area

    requires_target BOOLEAN NOT NULL DEFAULT TRUE,

    -- =========================
    -- DAMAGE MODEL
    -- =========================

    base_damage FLOAT NOT NULL DEFAULT 0.0,

    damage_type TEXT NOT NULL DEFAULT 'physical',
    -- physical, magical, elemental, true

    elemental_type TEXT,
    -- fire, frost, poison, holy, shadow etc

    posture_damage FLOAT NOT NULL DEFAULT 0.0,

    armor_penetration FLOAT NOT NULL DEFAULT 0.0,

    damage_multiplier FLOAT NOT NULL DEFAULT 1.0,
    -- multiplicador da skill

    critical_bonus_multiplier FLOAT NOT NULL DEFAULT 0.0,
    -- bônus sobre o crit do combat_core_profile
    -- ex:
    -- core = 1.5x
    -- skill bonus = 0.5
    -- final = 2.0x

    -- =========================
    -- CONTROL EFFECTS
    -- =========================

    stun_duration_ms INT NOT NULL DEFAULT 0,
    root_duration_ms INT NOT NULL DEFAULT 0,
    knockback_force FLOAT NOT NULL DEFAULT 0.0,

    -- =========================
    -- MOVEMENT INTERACTION
    -- =========================

    movement_multiplier FLOAT NOT NULL DEFAULT 1.0,
    -- quanto pode mover durante skill

    locks_movement BOOLEAN NOT NULL DEFAULT FALSE,
    -- trava locomoção

    movement_distance FLOAT NOT NULL DEFAULT 0.0,
    -- dash / leap / dodge

    -- =========================
    -- COMBO / FLOW
    -- =========================

    combo_group TEXT,
    -- light_chain, wolf_combo etc

    combo_index INT,
    -- ordem do combo

    combo_window_ms INT NOT NULL DEFAULT 0,

    -- =========================
    -- FLAGS
    -- =========================

    is_interruptible BOOLEAN NOT NULL DEFAULT TRUE,
    is_blockable BOOLEAN NOT NULL DEFAULT TRUE,
    is_parryable BOOLEAN NOT NULL DEFAULT TRUE,

    ignores_line_of_sight BOOLEAN NOT NULL DEFAULT FALSE,
    ignores_collision BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    -- =========================
    -- CONSTRAINTS
    -- =========================

    CONSTRAINT chk_skill_range
    CHECK (min_range <= max_range),

    CONSTRAINT chk_skill_targets
    CHECK (max_targets >= 1),

    CONSTRAINT chk_skill_combo_index
    CHECK (combo_index IS NULL OR combo_index >= 0),

    CONSTRAINT chk_skill_iframe_window
    CHECK (iframe_start_ms <= iframe_end_ms),

    CONSTRAINT chk_skill_cancel_window
    CHECK (cancel_window_start_ms <= cancel_window_end_ms)
);
