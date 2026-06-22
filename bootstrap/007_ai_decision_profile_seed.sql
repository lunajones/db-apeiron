-- =========================================================
-- DEFAULT AI DECISION PROFILE
-- APEIRON MMO
-- STEPPE WOLF - FAST PACK PREDATOR DECISION MODEL
-- =========================================================

INSERT INTO apeiron.ai_decision_profile (
    id,

    reaction_time_ms,
    decision_interval_ms,
    combat_engage_delay_ms,
    input_delay_variance_ms,

    target_switch_delay_ms,
    target_priority_bias,
    threat_evaluation_speed,

    combo_interrupt_tolerance,
    greed_vs_safety,
    punish_recognition_speed,

    mistake_chance,
    hesitation_factor,
    overcommit_risk,

    assist_ally_priority,
    pack_sync_factor,
    focus_fire_coordination,

    is_predictive,
    is_adaptive
)
VALUES (
    'ai_decision_steppe_wolf',

    180,
    260,
    160,
    60,

    450,
    0.68,
    0.72,

    0.48,
    0.58,
    0.62,

    0.08,
    0.16,
    0.42,

    0.74,
    0.82,
    0.76,

    FALSE,
    TRUE
)
ON CONFLICT (id) DO UPDATE SET
    reaction_time_ms = EXCLUDED.reaction_time_ms,
    decision_interval_ms = EXCLUDED.decision_interval_ms,
    combat_engage_delay_ms = EXCLUDED.combat_engage_delay_ms,
    input_delay_variance_ms = EXCLUDED.input_delay_variance_ms,

    target_switch_delay_ms = EXCLUDED.target_switch_delay_ms,
    target_priority_bias = EXCLUDED.target_priority_bias,
    threat_evaluation_speed = EXCLUDED.threat_evaluation_speed,

    combo_interrupt_tolerance = EXCLUDED.combo_interrupt_tolerance,
    greed_vs_safety = EXCLUDED.greed_vs_safety,
    punish_recognition_speed = EXCLUDED.punish_recognition_speed,

    mistake_chance = EXCLUDED.mistake_chance,
    hesitation_factor = EXCLUDED.hesitation_factor,
    overcommit_risk = EXCLUDED.overcommit_risk,

    assist_ally_priority = EXCLUDED.assist_ally_priority,
    pack_sync_factor = EXCLUDED.pack_sync_factor,
    focus_fire_coordination = EXCLUDED.focus_fire_coordination,

    is_predictive = EXCLUDED.is_predictive,
    is_adaptive = EXCLUDED.is_adaptive,

    updated_at = NOW();