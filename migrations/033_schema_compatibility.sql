-- =========================================================
-- RECOVERY SCHEMA COMPATIBILITY
-- Applies additive schema fixes for partially recovered databases where
-- earlier migrations were marked as applied before newer columns existed.
-- =========================================================

ALTER TABLE IF EXISTS apeiron.movement_profile
    ADD COLUMN IF NOT EXISTS max_speed FLOAT NOT NULL DEFAULT 5.0,
    ADD COLUMN IF NOT EXISTS acceleration FLOAT NOT NULL DEFAULT 10.0,
    ADD COLUMN IF NOT EXISTS deceleration FLOAT NOT NULL DEFAULT 12.0,
    ADD COLUMN IF NOT EXISTS friction FLOAT NOT NULL DEFAULT 8.0,
    ADD COLUMN IF NOT EXISTS gravity_multiplier FLOAT NOT NULL DEFAULT 1.0,
    ADD COLUMN IF NOT EXISTS mass FLOAT NOT NULL DEFAULT 1.0,
    ADD COLUMN IF NOT EXISTS momentum_retention FLOAT NOT NULL DEFAULT 0.6,
    ADD COLUMN IF NOT EXISTS turn_rate FLOAT NOT NULL DEFAULT 720.0,
    ADD COLUMN IF NOT EXISTS air_control FLOAT NOT NULL DEFAULT 0.3,
    ADD COLUMN IF NOT EXISTS strafe_efficiency FLOAT NOT NULL DEFAULT 0.8,
    ADD COLUMN IF NOT EXISTS backpedal_penalty FLOAT NOT NULL DEFAULT 0.7,
    ADD COLUMN IF NOT EXISTS dodge_distance FLOAT NOT NULL DEFAULT 4.0,
    ADD COLUMN IF NOT EXISTS dodge_duration_ms INT NOT NULL DEFAULT 250,
    ADD COLUMN IF NOT EXISTS sprint_multiplier FLOAT NOT NULL DEFAULT 1.4,
    ADD COLUMN IF NOT EXISTS slope_limit FLOAT NOT NULL DEFAULT 45.0,
    ADD COLUMN IF NOT EXISTS slide_on_slope BOOLEAN NOT NULL DEFAULT TRUE,
    ADD COLUMN IF NOT EXISTS is_airborne_enabled BOOLEAN NOT NULL DEFAULT TRUE,
    ADD COLUMN IF NOT EXISTS created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP NOT NULL DEFAULT NOW();
