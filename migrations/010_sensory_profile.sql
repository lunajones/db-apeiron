-- =========================================================
-- SENSORY PROFILE
-- APEIRON MMO - CREATURE PERCEPTION CONFIG
-- =========================================================

CREATE TABLE IF NOT EXISTS apeiron.sensory_profile (
    id TEXT PRIMARY KEY,

    -- =========================
    -- VISION
    -- =========================

    vision_range FLOAT NOT NULL DEFAULT 25.0,
    vision_angle FLOAT NOT NULL DEFAULT 120.0,
    night_vision_modifier FLOAT NOT NULL DEFAULT 0.7,

    can_detect_stealth BOOLEAN NOT NULL DEFAULT FALSE,
    stealth_detection_modifier FLOAT NOT NULL DEFAULT 1.0,

    -- =========================
    -- HEARING
    -- =========================

    hearing_range FLOAT NOT NULL DEFAULT 18.0,
    noise_sensitivity FLOAT NOT NULL DEFAULT 1.0,

    -- =========================
    -- SMELL / TRACKING
    -- =========================

    smell_range FLOAT NOT NULL DEFAULT 12.0,
    blood_detection_range FLOAT NOT NULL DEFAULT 20.0,
    tracking_persistence_ms INT NOT NULL DEFAULT 8000,

    -- =========================
    -- TARGET MEMORY
    -- =========================

    target_memory_ms INT NOT NULL DEFAULT 6000,
    last_known_position_memory_ms INT NOT NULL DEFAULT 10000,

    -- =========================
    -- ALERTNESS
    -- =========================

    alertness_gain_rate FLOAT NOT NULL DEFAULT 1.0,
    alertness_decay_rate FLOAT NOT NULL DEFAULT 0.4,
    surprise_threshold FLOAT NOT NULL DEFAULT 0.25,

    -- =========================
    -- LIFECYCLE
    -- =========================

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT chk_sensory_vision_range
        CHECK (vision_range >= 0),

    CONSTRAINT chk_sensory_vision_angle
        CHECK (vision_angle >= 0 AND vision_angle <= 360),

    CONSTRAINT chk_sensory_hearing_range
        CHECK (hearing_range >= 0),

    CONSTRAINT chk_sensory_smell_range
        CHECK (smell_range >= 0),

    CONSTRAINT chk_sensory_tracking_persistence
        CHECK (tracking_persistence_ms >= 0),

    CONSTRAINT chk_sensory_target_memory
        CHECK (target_memory_ms >= 0),

    CONSTRAINT chk_sensory_last_known_position_memory
        CHECK (last_known_position_memory_ms >= 0)
);

CREATE INDEX IF NOT EXISTS idx_sensory_profile_vision_range
ON apeiron.sensory_profile(vision_range);

CREATE INDEX IF NOT EXISTS idx_sensory_profile_hearing_range
ON apeiron.sensory_profile(hearing_range);

CREATE INDEX IF NOT EXISTS idx_sensory_profile_smell_range
ON apeiron.sensory_profile(smell_range);
