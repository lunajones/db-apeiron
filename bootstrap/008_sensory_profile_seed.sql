-- =========================================================
-- DEFAULT SENSORY PROFILE
-- APEIRON MMO
-- STEPPE WOLF - PREDATOR PERCEPTION
-- =========================================================

INSERT INTO apeiron.sensory_profile (
    id,

    vision_range,
    vision_angle,
    night_vision_modifier,

    can_detect_stealth,
    stealth_detection_modifier,

    hearing_range,
    noise_sensitivity,

    smell_range,
    blood_detection_range,
    tracking_persistence_ms,

    target_memory_ms,
    last_known_position_memory_ms,

    alertness_gain_rate,
    alertness_decay_rate,
    surprise_threshold
)
VALUES (
    'sensory_steppe_wolf',

    32.0,
    140.0,
    1.15,

    FALSE,
    0.85,

    26.0,
    1.25,

    34.0,
    45.0,
    14000,

    9000,
    16000,

    1.25,
    0.35,
    0.18
)
ON CONFLICT (id) DO UPDATE SET
    vision_range = EXCLUDED.vision_range,
    vision_angle = EXCLUDED.vision_angle,
    night_vision_modifier = EXCLUDED.night_vision_modifier,

    can_detect_stealth = EXCLUDED.can_detect_stealth,
    stealth_detection_modifier = EXCLUDED.stealth_detection_modifier,

    hearing_range = EXCLUDED.hearing_range,
    noise_sensitivity = EXCLUDED.noise_sensitivity,

    smell_range = EXCLUDED.smell_range,
    blood_detection_range = EXCLUDED.blood_detection_range,
    tracking_persistence_ms = EXCLUDED.tracking_persistence_ms,

    target_memory_ms = EXCLUDED.target_memory_ms,
    last_known_position_memory_ms = EXCLUDED.last_known_position_memory_ms,

    alertness_gain_rate = EXCLUDED.alertness_gain_rate,
    alertness_decay_rate = EXCLUDED.alertness_decay_rate,
    surprise_threshold = EXCLUDED.surprise_threshold,

    updated_at = NOW();