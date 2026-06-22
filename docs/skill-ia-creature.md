# APEIRON MMO - CREATURE TEMPLATE DESIGN (DATABASE CORE)

## OBJETIVO

Definir o modelo correto de `creature_template` para um MMORPG:

- PvP-first
- combate action (hitbox real)
- iFrame dodge
- IA emergente (não scripted)
- comportamento semelhante a players
- milhares de entidades simultâneas
- arquitetura data-driven
- zero lógica hardcoded por criatura

---

## PRINCÍPIO FUNDAMENTAL

`creature_template NÃO contém sistemas.`

Ele NÃO é:
- IA
- combate
- movimento detalhado
- comportamento
- estados runtime

Ele é apenas:

> "um agregador de profiles (DNA da criatura)"

---

## ARQUITETURA CORRETA

### Separação obrigatória

Cada sistema é isolado:

- movement_profile → movimentação física e feel
- combat_profile → combate e stats
- behavior_profile → decisões e combate AI
- needs_profile → sobrevivência e necessidades
- personality_profile → tendências base
- sensory_profile → percepção
- world_profile → interação com mundo
- skills → lista de habilidades

---

## CREATURE TEMPLATE (CORRETO)

```sql
-- =========================================================
-- CREATURE TEMPLATE (NORMALIZADO + COMENTADO)
-- APEIRON MMO - PvP FIRST / DATA DRIVEN
-- =========================================================

CREATE TABLE creature_template (
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

    behavior_profile_id TEXT NOT NULL,
    -- Regras de comportamento macro:
    -- flanco, circle, chase, retreat, pack logic, território

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

    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
    -- atualização de balanceamento
);