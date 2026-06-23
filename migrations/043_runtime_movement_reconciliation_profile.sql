-- =========================================================
-- RUNTIME MOVEMENT RECONCILIATION PROFILE
-- Rich player/client reconciliation profile consumed by game-server snapshots.
-- This is distinct from movement_reconciliation_contract, which describes
-- per-action ownership categories.
-- =========================================================

CREATE TABLE IF NOT EXISTS apeiron.runtime_movement_reconciliation_profile (
    profile_id TEXT PRIMARY KEY,
    max_speed FLOAT NOT NULL,
    sprint_speed_multiplier FLOAT NOT NULL,
    acceleration FLOAT NOT NULL,
    deceleration FLOAT NOT NULL,
    ground_friction FLOAT NOT NULL,
    air_acceleration FLOAT NOT NULL,
    jump_height FLOAT NOT NULL,
    jump_duration_ms INT NOT NULL,
    rotation_rate_yaw FLOAT NOT NULL,
    gravity_scale FLOAT NOT NULL,
    braking_friction_factor FLOAT NOT NULL,
    max_slope_deg FLOAT NOT NULL,
    step_height FLOAT NOT NULL,
    base_deadzone FLOAT NOT NULL,
    grounded_speed_deadzone_factor FLOAT NOT NULL,
    grounded_speed_deadzone_min FLOAT NOT NULL,
    grounded_speed_deadzone_max FLOAT NOT NULL,
    grounded_transition_deadzone_min FLOAT NOT NULL DEFAULT 0,
    move_sustain_deadzone FLOAT NOT NULL,
    move_sustain_transition_deadzone FLOAT NOT NULL,
    airborne_deadzone FLOAT NOT NULL,
    leap_recent_deadzone FLOAT NOT NULL,
    leap_airborne_snapshot_deadzone FLOAT NOT NULL,
    leap_landing_deadzone_factor FLOAT NOT NULL,
    leap_landing_deadzone_min FLOAT NOT NULL,
    leap_landing_deadzone_max FLOAT NOT NULL,
    leap_landing_clamp_ignore_deadzone FLOAT NOT NULL DEFAULT 0,
    leap_landing_soft_snap_deadzone FLOAT NOT NULL DEFAULT 0,
    dodge_recent_deadzone FLOAT NOT NULL,
    dodge_active_deadzone FLOAT NOT NULL,
    dodge_exit_deadzone_factor FLOAT NOT NULL,
    dodge_exit_deadzone_min FLOAT NOT NULL,
    dodge_exit_deadzone_max FLOAT NOT NULL,
    post_action_grounded_deadzone FLOAT NOT NULL,
    correction_max_step FLOAT NOT NULL,
    hard_snap_distance FLOAT NOT NULL,
    severe_desync_distance FLOAT NOT NULL,
    visual_smoothing_ms INT NOT NULL,
    visual_smoothing_max_distance FLOAT NOT NULL,
    remote_visual_interpolation_ms INT NOT NULL,
    remote_visual_max_extrapolation_ms INT NOT NULL,
    remote_visual_hard_snap_distance FLOAT NOT NULL,
    dodge_carry_handoff_ms INT NOT NULL DEFAULT 0,
    leap_landing_correction_grace_ms INT NOT NULL DEFAULT 0,
    leap_grounded_carry_handoff_ms INT NOT NULL DEFAULT 0,
    movement_turn_resubmit_dot_threshold FLOAT NOT NULL,
    movement_turn_resubmit_min_interval_ms INT NOT NULL,
    movement_submit_interval_ms INT NOT NULL,
    snapshot_poll_interval_ms INT NOT NULL,
    strafe_speed_multiplier FLOAT NOT NULL,
    backpedal_speed_multiplier FLOAT NOT NULL,
    strafe_sprint_speed_multiplier FLOAT NOT NULL,
    backpedal_sprint_speed_multiplier FLOAT NOT NULL,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_runtime_movement_profile_speed CHECK (max_speed > 0 AND sprint_speed_multiplier > 0),
    CONSTRAINT chk_runtime_movement_profile_timing CHECK (
        jump_duration_ms >= 0
        AND visual_smoothing_ms >= 0
        AND remote_visual_interpolation_ms >= 0
        AND remote_visual_max_extrapolation_ms >= 0
        AND movement_submit_interval_ms > 0
        AND snapshot_poll_interval_ms > 0
    ),
    CONSTRAINT chk_runtime_movement_profile_deadzone CHECK (
        base_deadzone >= 0
        AND grounded_speed_deadzone_min >= 0
        AND grounded_speed_deadzone_max >= grounded_speed_deadzone_min
        AND hard_snap_distance >= correction_max_step
        AND severe_desync_distance >= hard_snap_distance
    )
);

ALTER TABLE IF EXISTS apeiron.runtime_movement_reconciliation_profile
    ADD COLUMN IF NOT EXISTS metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    ADD COLUMN IF NOT EXISTS created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP NOT NULL DEFAULT NOW();
