-- =========================================================
-- DEFAULT MOVEMENT PROFILE
-- APEIRON MMO
-- STEPPE WOLF - OPEN FIELD PREDATOR
-- =========================================================

INSERT INTO apeiron.movement_profile (
    id,

    max_speed,
    acceleration,
    deceleration,
    friction,
    gravity_multiplier,
    mass,
    momentum_retention,

    turn_rate,
    air_control,
    strafe_efficiency,
    backpedal_penalty,

    dodge_distance,
    dodge_duration_ms,

    sprint_multiplier,

    slope_limit,
    slide_on_slope,

    is_airborne_enabled
)
VALUES (
    'movement_steppe_wolf',

    3.4,
    14.0,
    11.0,
    7.5,
    1.15,
    0.85,
    0.72,

    860.0,
    0.22,
    0.92,
    0.55,

    3.8,
    220,

    1.55,

    42.0,
    TRUE,

    TRUE
)
ON CONFLICT (id) DO UPDATE SET
    max_speed = EXCLUDED.max_speed,
    acceleration = EXCLUDED.acceleration,
    deceleration = EXCLUDED.deceleration,
    friction = EXCLUDED.friction,
    gravity_multiplier = EXCLUDED.gravity_multiplier,
    mass = EXCLUDED.mass,
    momentum_retention = EXCLUDED.momentum_retention,

    turn_rate = EXCLUDED.turn_rate,
    air_control = EXCLUDED.air_control,
    strafe_efficiency = EXCLUDED.strafe_efficiency,
    backpedal_penalty = EXCLUDED.backpedal_penalty,

    dodge_distance = EXCLUDED.dodge_distance,
    dodge_duration_ms = EXCLUDED.dodge_duration_ms,

    sprint_multiplier = EXCLUDED.sprint_multiplier,

    slope_limit = EXCLUDED.slope_limit,
    slide_on_slope = EXCLUDED.slide_on_slope,

    is_airborne_enabled = EXCLUDED.is_airborne_enabled,

    updated_at = NOW();
