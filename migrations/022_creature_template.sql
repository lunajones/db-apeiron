-- =========================================================
-- CREATURE TEMPLATE (NORMALIZADO + COMENTADO)
-- APEIRON MMO - PvP FIRST / DATA DRIVEN
-- =========================================================

CREATE TABLE apeiron.creature_template (
    id TEXT PRIMARY KEY,

    -- =========================
    -- IDENTITY
    -- =========================

    name TEXT NOT NULL,
    -- Nome visual da criatura (UI / debug / logs)

    faction TEXT NOT NULL,
    -- Grupo de hostilidade / aliança (ex: wild, undead, player_faction_x)

    tier INT NOT NULL DEFAULT 1,
    -- Escala de poder base (balanceamento global do MMO)

    archetype TEXT NOT NULL,
    -- Tipo macro (beast, humanoid, brute, assassin, caster etc)

    -- =========================
    -- PROFILE REFERENCES (CORE SYSTEMS)
    -- =========================

    movement_profile_id TEXT NOT NULL,
    -- Define movimentação física:
    -- aceleração, turn rate, inertia, strafing, feel "player-like"

    combat_core_profile_id TEXT NOT NULL,
    -- Regras físicas do combate:
    -- stamina, dodge iframe, block, parry, posture, ranges

    combat_style_profile_id TEXT NOT NULL,
    -- Estilo de luta:
    -- agressivo, passivo, poke, combo, feint, pressão, bait

    ai_decision_profile_id TEXT NOT NULL,
    -- "Cérebro":
    -- reaction time, decision interval, noise, switching delay

    personality_profile_id TEXT NOT NULL,
    -- Personalidade base FIXA:
    -- coragem, agressividade base, curiosidade, disciplina, medo

    sensory_profile_id TEXT NOT NULL,
    -- Percepção do mundo:
    -- visão, FOV, audição, cheiro, detecção stealth

    needs_profile_id TEXT NOT NULL,
    -- Sistema de sobrevivência:
    -- fome, sede, fadiga, thresholds e decay rates

    -- =========================
    -- SKILLS
    -- =========================

    skill_set_id TEXT NOT NULL,
    -- Conjunto de skills reutilizável:
    -- lista de habilidades + pesos + prioridades

    -- =========================
    -- WORLD CONTROL (SEPARAÇÃO LIMPA)
    -- =========================

    spawn_profile_id TEXT NOT NULL,
    -- DEFINE REGRAS DO AMBIENTE (NÃO DNA DA CRIATURA)
    -- controla:
    -- - biome constraints
    -- - spawn rules
    -- - roaming policy
    -- - population density behavior
    -- - migration rules globais
    -- IMPORTANTE: NÃO é comportamento da criatura, é regra do mundo

    -- =========================
    -- METADATA
    -- =========================

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    -- criação do template

    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    -- atualização de balanceamento

    -- =========================
    -- FOREIGN KEYS
    -- =========================

    CONSTRAINT fk_creature_template_movement_profile
        FOREIGN KEY (movement_profile_id)
        REFERENCES apeiron.movement_profile(id),

    CONSTRAINT fk_creature_template_combat_core_profile
        FOREIGN KEY (combat_core_profile_id)
        REFERENCES apeiron.combat_core_profile(id),

    CONSTRAINT fk_creature_template_combat_style_profile
        FOREIGN KEY (combat_style_profile_id)
        REFERENCES apeiron.combat_style_profile(id),

    CONSTRAINT fk_creature_template_ai_decision_profile
        FOREIGN KEY (ai_decision_profile_id)
        REFERENCES apeiron.ai_decision_profile(id),

    CONSTRAINT fk_creature_template_personality_profile
        FOREIGN KEY (personality_profile_id)
        REFERENCES apeiron.ai_personality_profile(id),

    CONSTRAINT fk_creature_template_sensory_profile
        FOREIGN KEY (sensory_profile_id)
        REFERENCES apeiron.sensory_profile(id),

    CONSTRAINT fk_creature_template_needs_profile
        FOREIGN KEY (needs_profile_id)
        REFERENCES apeiron.needs_profile(id),

    CONSTRAINT fk_creature_template_skill_set
        FOREIGN KEY (skill_set_id)
        REFERENCES apeiron.skill_set(id),

    CONSTRAINT fk_creature_template_spawn_profile
        FOREIGN KEY (spawn_profile_id)
        REFERENCES apeiron.spawn_profile(id)
);

CREATE INDEX idx_creature_template_movement_profile
ON apeiron.creature_template(movement_profile_id);

CREATE INDEX idx_creature_template_combat_core_profile
ON apeiron.creature_template(combat_core_profile_id);

CREATE INDEX idx_creature_template_combat_style_profile
ON apeiron.creature_template(combat_style_profile_id);

CREATE INDEX idx_creature_template_ai_decision_profile
ON apeiron.creature_template(ai_decision_profile_id);

CREATE INDEX idx_creature_template_personality_profile
ON apeiron.creature_template(personality_profile_id);

CREATE INDEX idx_creature_template_sensory_profile
ON apeiron.creature_template(sensory_profile_id);

CREATE INDEX idx_creature_template_needs_profile
ON apeiron.creature_template(needs_profile_id);

CREATE INDEX idx_creature_template_skill_set
ON apeiron.creature_template(skill_set_id);

CREATE INDEX idx_creature_template_spawn_profile
ON apeiron.creature_template(spawn_profile_id);