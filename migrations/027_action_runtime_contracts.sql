-- =========================================================
-- ACTION RUNTIME CONTRACTS
-- Reconstructed from Apeiron recovery artifacts on 2026-06-22.
-- =========================================================

CREATE TABLE IF NOT EXISTS apeiron.movement_reconciliation_contract (
    id TEXT PRIMARY KEY,
    category TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    max_smooth_error_cm FLOAT NOT NULL DEFAULT 35.0,
    hard_snap_error_cm FLOAT NOT NULL DEFAULT 180.0,
    smoothing_time_ms INT NOT NULL DEFAULT 90,
    yaw_tolerance_deg FLOAT NOT NULL DEFAULT 8.0,
    owns_position BOOLEAN NOT NULL DEFAULT FALSE,
    owns_yaw BOOLEAN NOT NULL DEFAULT FALSE,
    allows_client_prediction BOOLEAN NOT NULL DEFAULT TRUE,
    input_policy TEXT NOT NULL DEFAULT 'normal',
    handoff_policy TEXT NOT NULL DEFAULT 'explicit',
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_reconciliation_error CHECK (max_smooth_error_cm >= 0 AND hard_snap_error_cm >= max_smooth_error_cm),
    CONSTRAINT chk_reconciliation_smoothing CHECK (smoothing_time_ms >= 0)
);

CREATE TABLE IF NOT EXISTS apeiron.movement_action_contract (
    id TEXT PRIMARY KEY,
    action_type TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    duration_ms INT NOT NULL,
    active_ms INT NOT NULL,
    recovery_ms INT NOT NULL DEFAULT 0,
    distance_cm FLOAT NOT NULL DEFAULT 0.0,
    base_speed_cm_s FLOAT NOT NULL DEFAULT 0.0,
    yaw_degrees FLOAT NOT NULL DEFAULT 0.0,
    phase_window_policy TEXT NOT NULL DEFAULT 'grounded_action',
    prediction_error_policy TEXT NOT NULL DEFAULT 'bounded_smooth_correction',
    reconciliation_contract_id TEXT,
    allow_windup_locomotion BOOLEAN NOT NULL DEFAULT FALSE,
    allow_active_locomotion BOOLEAN NOT NULL DEFAULT FALSE,
    allow_recovery_locomotion BOOLEAN NOT NULL DEFAULT FALSE,
    allow_yaw_adjustment BOOLEAN NOT NULL DEFAULT TRUE,
    root_motion_owner TEXT NOT NULL DEFAULT 'movement',
    contact_policy TEXT NOT NULL DEFAULT 'none',
    speed_curve JSONB NOT NULL DEFAULT '[]'::jsonb,
    vertical_curve JSONB NOT NULL DEFAULT '[]'::jsonb,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_movement_action_reconciliation
        FOREIGN KEY (reconciliation_contract_id)
        REFERENCES apeiron.movement_reconciliation_contract(id),
    CONSTRAINT chk_movement_action_time CHECK (duration_ms > 0 AND active_ms >= 0 AND recovery_ms >= 0),
    CONSTRAINT chk_movement_action_distance CHECK (distance_cm >= 0)
);

CREATE TABLE IF NOT EXISTS apeiron.skill_action_timing (
    skill_id TEXT PRIMARY KEY,
    windup_ms INT NOT NULL DEFAULT 0,
    active_ms INT NOT NULL DEFAULT 0,
    recovery_ms INT NOT NULL DEFAULT 0,
    cooldown_ms INT NOT NULL DEFAULT 0,
    combo_window_ms INT NOT NULL DEFAULT 0,
    movement_lock_policy TEXT NOT NULL DEFAULT 'contract',
    queue_policy TEXT NOT NULL DEFAULT 'none',
    cancel_policy TEXT NOT NULL DEFAULT 'none',
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_skill_action_timing_skill
        FOREIGN KEY (skill_id)
        REFERENCES apeiron.skill(id)
        ON DELETE CASCADE,
    CONSTRAINT chk_skill_action_timing CHECK (windup_ms >= 0 AND active_ms >= 0 AND recovery_ms >= 0 AND cooldown_ms >= 0)
);

CREATE TABLE IF NOT EXISTS apeiron.skill_movement_action_binding (
    skill_id TEXT PRIMARY KEY,
    movement_action_contract_id TEXT NOT NULL,
    starts_at_phase TEXT NOT NULL DEFAULT 'windup',
    handoff_policy TEXT NOT NULL DEFAULT 'explicit_recovery_handoff',
    normal_input_policy TEXT NOT NULL DEFAULT 'blocked_during_owned_root',
    target_policy TEXT NOT NULL DEFAULT 'aim_direction',
    contact_policy TEXT NOT NULL DEFAULT 'none',
    is_enabled BOOLEAN NOT NULL DEFAULT TRUE,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_skill_movement_action_skill
        FOREIGN KEY (skill_id)
        REFERENCES apeiron.skill(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_skill_movement_action_contract
        FOREIGN KEY (movement_action_contract_id)
        REFERENCES apeiron.movement_action_contract(id)
);

CREATE INDEX IF NOT EXISTS idx_movement_action_reconciliation
ON apeiron.movement_action_contract(reconciliation_contract_id);

CREATE INDEX IF NOT EXISTS idx_skill_movement_action_contract
ON apeiron.skill_movement_action_binding(movement_action_contract_id);
