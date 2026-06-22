-- =========================================================
-- SKILL SLOT
-- APEIRON MMO - SKILL LOADOUT CONFIGURATION
-- =========================================================

CREATE TABLE IF NOT EXISTS apeiron.skill_slot (
    id BIGSERIAL PRIMARY KEY,

    -- =========================
    -- REFERENCES
    -- =========================

    skill_set_id TEXT NOT NULL,
    -- conjunto ao qual pertence

    skill_id TEXT NOT NULL,
    -- skill equipada

    -- =========================
    -- SLOT CONTROL
    -- =========================

    slot_index INT NOT NULL,
    -- ordem lógica da skill no set
    -- ex:
    -- wolf:
    -- 1 bite
    -- 2 leap
    -- 3 dodge
    --
    -- player:
    -- 1 left_click
    -- 2 right_click
    -- 3 Q
    -- 4 E

    is_enabled BOOLEAN NOT NULL DEFAULT TRUE,
    -- permite desativar skill
    -- sem remover do set

    -- =========================
    -- AI EXECUTION CONFIG
    -- =========================

    priority INT NOT NULL DEFAULT 1,
    -- importância da skill
    -- maior = mais prioridade

    usage_weight FLOAT NOT NULL DEFAULT 1.0,
    -- peso relativo de escolha
    -- evita IA previsível

    cooldown_override_ms INT,
    -- override opcional
    -- ex:
    -- elite wolf usa leap mais vezes

    -- =========================
    -- CONTEXT RULES
    -- =========================

    min_target_hp_percent FLOAT,
    max_target_hp_percent FLOAT,
    -- ex:
    -- execute:
    -- usar apenas abaixo de 30%

    min_self_hp_percent FLOAT,
    max_self_hp_percent FLOAT,
    -- ex:
    -- dodge/heal panic
    -- usar abaixo de 40%

    required_distance_min FLOAT,
    required_distance_max FLOAT,
    -- gating de uso
    -- ex:
    -- leap só entre 4m e 10m

    requires_line_of_sight BOOLEAN NOT NULL DEFAULT TRUE,

    -- =========================
    -- COMBAT FLOW
    -- =========================

    opener_weight FLOAT NOT NULL DEFAULT 0.0,
    -- chance de abrir combate com essa skill

    finisher_weight FLOAT NOT NULL DEFAULT 0.0,
    -- tendência a usar como finalização

    shared_cooldown_group TEXT,
    -- ex:
    -- dodge_group
    -- potion_group
    -- stance_group

    use_only_in_combat BOOLEAN NOT NULL DEFAULT TRUE,

    -- =========================
    -- METADATA
    -- =========================

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    -- =========================
    -- FOREIGN KEYS
    -- =========================

    CONSTRAINT fk_skill_slot_skill_set
    FOREIGN KEY (skill_set_id)
    REFERENCES skill_set(id),

    CONSTRAINT fk_skill_slot_skill
    FOREIGN KEY (skill_id)
    REFERENCES skill(id),

    -- =========================
    -- UNIQUENESS
    -- =========================

    CONSTRAINT uq_skill_set_slot
    UNIQUE (skill_set_id, slot_index),

    CONSTRAINT uq_skill_set_skill
    UNIQUE (skill_set_id, skill_id),

    -- =========================
    -- CONSTRAINTS
    -- =========================

    CONSTRAINT chk_skill_slot_priority
    CHECK (priority >= 0),

    CONSTRAINT chk_skill_slot_weight
    CHECK (usage_weight >= 0.0),

    CONSTRAINT chk_skill_slot_hp_percent
    CHECK (
        (
            min_target_hp_percent IS NULL
            OR (
                min_target_hp_percent >= 0.0
                AND min_target_hp_percent <= 1.0
            )
        )
        AND (
            max_target_hp_percent IS NULL
            OR (
                max_target_hp_percent >= 0.0
                AND max_target_hp_percent <= 1.0
            )
        )
        AND (
            min_self_hp_percent IS NULL
            OR (
                min_self_hp_percent >= 0.0
                AND min_self_hp_percent <= 1.0
            )
        )
        AND (
            max_self_hp_percent IS NULL
            OR (
                max_self_hp_percent >= 0.0
                AND max_self_hp_percent <= 1.0
            )
        )
    ),

    CONSTRAINT chk_skill_slot_distance
    CHECK (
        required_distance_min IS NULL
        OR required_distance_max IS NULL
        OR required_distance_min <= required_distance_max
    )
);

CREATE INDEX IF NOT EXISTS idx_skill_slot_skill_set
ON apeiron.skill_slot(skill_set_id);

CREATE INDEX IF NOT EXISTS idx_skill_slot_skill
ON apeiron.skill_slot(skill_id);
