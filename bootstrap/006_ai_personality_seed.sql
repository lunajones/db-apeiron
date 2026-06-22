-- =========================================================
-- DEFAULT AI PERSONALITY PROFILE
-- APEIRON MMO
-- STEPPE WOLF - PACK PREDATOR PERSONALITY
-- =========================================================

INSERT INTO apeiron.ai_personality_profile (
    id,

    courage,
    curiosity,
    discipline,
    aggression_baseline,
    fear_sensitivity,

    dominance,
    submission,
    loyalty,
    empathy,

    temperament_stability,
    adaptability,
    predictability,

    is_pack_animal,
    is_solo,
    is_predator
)
VALUES (
    'personality_steppe_wolf',

    0.68,
    0.42,
    0.62,
    0.66,
    0.48,

    0.55,
    0.35,
    0.82,
    0.48,

    0.58,
    0.64,
    0.46,

    TRUE,
    FALSE,
    TRUE
)
ON CONFLICT (id) DO UPDATE SET
    courage = EXCLUDED.courage,
    curiosity = EXCLUDED.curiosity,
    discipline = EXCLUDED.discipline,
    aggression_baseline = EXCLUDED.aggression_baseline,
    fear_sensitivity = EXCLUDED.fear_sensitivity,

    dominance = EXCLUDED.dominance,
    submission = EXCLUDED.submission,
    loyalty = EXCLUDED.loyalty,
    empathy = EXCLUDED.empathy,

    temperament_stability = EXCLUDED.temperament_stability,
    adaptability = EXCLUDED.adaptability,
    predictability = EXCLUDED.predictability,

    is_pack_animal = EXCLUDED.is_pack_animal,
    is_solo = EXCLUDED.is_solo,
    is_predator = EXCLUDED.is_predator,

    updated_at = NOW();