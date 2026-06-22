-- =========================================================
-- DEFAULT NEEDS PROFILE
-- APEIRON MMO
-- STEPPE WOLF - SURVIVAL PREDATOR
-- =========================================================

INSERT INTO apeiron.needs_profile (
    id,

    hunger_enabled,
    thirst_enabled,
    fatigue_enabled,

    hunger_decay_per_hour,
    thirst_decay_per_hour,
    fatigue_decay_per_hour,

    hunger_threshold,
    thirst_threshold,
    fatigue_threshold,

    hunger_damage_threshold,

    stamina_regen_penalty_at_low_needs,
    movement_penalty_at_low_needs,
    combat_performance_penalty,

    food_saturation_rate,
    water_saturation_rate,
    rest_recovery_rate,

    panic_threshold,
    aggression_increase_when_starving,
    fear_increase_when_starving,

    needs_enabled
)
VALUES (
    'needs_steppe_wolf',

    TRUE,
    TRUE,
    TRUE,

    0.08,
    0.10,
    0.06,

    0.35,
    0.30,
    0.25,

    0.12,

    0.35,
    0.25,
    0.30,

    1.20,
    1.10,
    0.85,

    0.18,
    0.45,
    0.20,

    TRUE
)
ON CONFLICT (id) DO UPDATE SET
    hunger_enabled = EXCLUDED.hunger_enabled,
    thirst_enabled = EXCLUDED.thirst_enabled,
    fatigue_enabled = EXCLUDED.fatigue_enabled,

    hunger_decay_per_hour = EXCLUDED.hunger_decay_per_hour,
    thirst_decay_per_hour = EXCLUDED.thirst_decay_per_hour,
    fatigue_decay_per_hour = EXCLUDED.fatigue_decay_per_hour,

    hunger_threshold = EXCLUDED.hunger_threshold,
    thirst_threshold = EXCLUDED.thirst_threshold,
    fatigue_threshold = EXCLUDED.fatigue_threshold,

    hunger_damage_threshold = EXCLUDED.hunger_damage_threshold,

    stamina_regen_penalty_at_low_needs = EXCLUDED.stamina_regen_penalty_at_low_needs,
    movement_penalty_at_low_needs = EXCLUDED.movement_penalty_at_low_needs,
    combat_performance_penalty = EXCLUDED.combat_performance_penalty,

    food_saturation_rate = EXCLUDED.food_saturation_rate,
    water_saturation_rate = EXCLUDED.water_saturation_rate,
    rest_recovery_rate = EXCLUDED.rest_recovery_rate,

    panic_threshold = EXCLUDED.panic_threshold,
    aggression_increase_when_starving = EXCLUDED.aggression_increase_when_starving,
    fear_increase_when_starving = EXCLUDED.fear_increase_when_starving,

    needs_enabled = EXCLUDED.needs_enabled,
    updated_at = NOW();