-- =========================================================
-- AI PERSONALITY PROFILE
-- APEIRON MMO - CORE PERSONALITY DNA
-- =========================================================

CREATE TABLE IF NOT EXISTS apeiron.ai_personality_profile (
    id TEXT PRIMARY KEY,

    -- =========================
    -- CORE PERSONALITY TRAITS
    -- =========================

    courage FLOAT NOT NULL DEFAULT 0.5,
    -- coragem base diante de risco

    curiosity FLOAT NOT NULL DEFAULT 0.5,
    -- tendência a explorar / investigar

    discipline FLOAT NOT NULL DEFAULT 0.5,
    -- consistência de comportamento

    aggression_baseline FLOAT NOT NULL DEFAULT 0.5,
    -- agressividade NATIVA (não estilo de combate)

    fear_sensitivity FLOAT NOT NULL DEFAULT 0.5,
    -- sensibilidade ao perigo

    -- =========================
    -- SOCIAL / GROUP DNA
    -- =========================

    dominance FLOAT NOT NULL DEFAULT 0.5,
    -- tendência a liderar / impor

    submission FLOAT NOT NULL DEFAULT 0.5,
    -- tendência a ceder / evitar conflito social

    loyalty FLOAT NOT NULL DEFAULT 0.5,
    -- fidelidade ao grupo/fação

    empathy FLOAT NOT NULL DEFAULT 0.5,
    -- resposta a aliados feridos / contexto social

    -- =========================
    -- BEHAVIORAL STABILITY
    -- =========================

    temperament_stability FLOAT NOT NULL DEFAULT 0.5,
    -- variação emocional (baixo = instável)

    adaptability FLOAT NOT NULL DEFAULT 0.5,
    -- capacidade de mudar comportamento fora de combate

    predictability FLOAT NOT NULL DEFAULT 0.5,
    -- quão fácil é ler padrão da criatura

    -- =========================
    -- META FLAGS
    -- =========================

    is_pack_animal BOOLEAN NOT NULL DEFAULT FALSE,
    is_solo BOOLEAN NOT NULL DEFAULT TRUE,
    is_predator BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
