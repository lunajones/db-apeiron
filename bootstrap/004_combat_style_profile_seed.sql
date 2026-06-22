-- =========================================================
-- DEFAULT COMBAT STYLE PROFILE
-- APEIRON MMO
-- STEPPE WOLF - PACK HUNTER / OPEN FIELD PRESSURE
-- =========================================================

INSERT INTO apeiron.combat_style_profile (
    id,

    archetype,

    aggression,
    defense,
    patience,
    risk_tolerance,

    preferred_range,
    chase_tendency,
    disengage_threshold,

    combo_preference,
    feint_usage,
    punish_window_awareness,
    aggression_spike_chance,

    dodge_frequency,
    block_frequency,
    parry_aggressiveness,
    panic_threshold,

    strafe_usage,
    circle_target,
    backstep_usage,
    reposition_frequency,

    randomness,
    consistency,

    target_switching,
    focus_fire,
    retaliation_bias,

    is_elite,
    is_coward,
    is_duelist
)
VALUES (
    'combat_style_steppe_wolf',

    'pack_predator',

    0.72,
    0.28,
    0.55,
    0.62,

    1.8,
    0.82,
    0.22,

    0.58,
    0.12,
    0.68,
    0.24,

    0.66,
    0.0,
    0.0,
    0.34,

    0.72,
    0.78,
    0.18,
    0.70,

    0.22,
    0.74,

    0.38,
    0.72,
    0.68,

    FALSE,
    FALSE,
    FALSE
)
ON CONFLICT (id) DO UPDATE SET
    archetype = EXCLUDED.archetype,

    aggression = EXCLUDED.aggression,
    defense = EXCLUDED.defense,
    patience = EXCLUDED.patience,
    risk_tolerance = EXCLUDED.risk_tolerance,

    preferred_range = EXCLUDED.preferred_range,
    chase_tendency = EXCLUDED.chase_tendency,
    disengage_threshold = EXCLUDED.disengage_threshold,

    combo_preference = EXCLUDED.combo_preference,
    feint_usage = EXCLUDED.feint_usage,
    punish_window_awareness = EXCLUDED.punish_window_awareness,
    aggression_spike_chance = EXCLUDED.aggression_spike_chance,

    dodge_frequency = EXCLUDED.dodge_frequency,
    block_frequency = EXCLUDED.block_frequency,
    parry_aggressiveness = EXCLUDED.parry_aggressiveness,
    panic_threshold = EXCLUDED.panic_threshold,

    strafe_usage = EXCLUDED.strafe_usage,
    circle_target = EXCLUDED.circle_target,
    backstep_usage = EXCLUDED.backstep_usage,
    reposition_frequency = EXCLUDED.reposition_frequency,

    randomness = EXCLUDED.randomness,
    consistency = EXCLUDED.consistency,

    target_switching = EXCLUDED.target_switching,
    focus_fire = EXCLUDED.focus_fire,
    retaliation_bias = EXCLUDED.retaliation_bias,

    is_elite = EXCLUDED.is_elite,
    is_coward = EXCLUDED.is_coward,
    is_duelist = EXCLUDED.is_duelist,

    updated_at = NOW();