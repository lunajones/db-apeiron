-- =========================================================
-- ACTION RUNTIME / RECONCILIATION SEEDS
-- =========================================================

INSERT INTO apeiron.movement_reconciliation_contract (
    id, category, description, max_smooth_error_cm, hard_snap_error_cm,
    smoothing_time_ms, yaw_tolerance_deg, owns_position, owns_yaw,
    allows_client_prediction, input_policy, handoff_policy, metadata
)
VALUES
('grounded_move_reconciliation','grounded_move','Normal grounded walk/run/strafe reconciliation.',35,180,90,8,FALSE,FALSE,TRUE,'normal','continuous','{"source":"reconstructed"}'),
('grounded_skill_action_reconciliation','grounded_skill_action','Committed grounded skill movement reconciliation with explicit post-action handoff.',25,140,70,10,TRUE,TRUE,TRUE,'blocked_during_owned_root','explicit','{"source":"reconstructed"}'),
('dodge_reconciliation','dodge','Dodge owns position and iframe window until movement end.',22,130,60,14,TRUE,TRUE,TRUE,'blocked_during_owned_root','explicit','{"source":"reconstructed"}'),
('leap_reconciliation','leap','Leap owns vertical and horizontal movement until grounded handoff.',30,170,70,12,TRUE,TRUE,TRUE,'blocked_during_airborne','grounded_handoff','{"source":"reconstructed"}'),
('turn_reconciliation','turn','Rate-limited contextual turn reconciliation.',18,90,45,4,FALSE,TRUE,TRUE,'turn_only','continuous','{"source":"reconstructed"}'),
('post_action_handoff_reconciliation','post_action_handoff','Recovery phase contract that prevents normal locomotion from racing skill recovery.',20,110,55,8,TRUE,FALSE,TRUE,'buffer_until_handoff','explicit','{"source":"reconstructed"}')
ON CONFLICT (id) DO UPDATE SET
    category = EXCLUDED.category,
    description = EXCLUDED.description,
    max_smooth_error_cm = EXCLUDED.max_smooth_error_cm,
    hard_snap_error_cm = EXCLUDED.hard_snap_error_cm,
    smoothing_time_ms = EXCLUDED.smoothing_time_ms,
    yaw_tolerance_deg = EXCLUDED.yaw_tolerance_deg,
    owns_position = EXCLUDED.owns_position,
    owns_yaw = EXCLUDED.owns_yaw,
    allows_client_prediction = EXCLUDED.allows_client_prediction,
    input_policy = EXCLUDED.input_policy,
    handoff_policy = EXCLUDED.handoff_policy,
    metadata = EXCLUDED.metadata,
    updated_at = NOW();

INSERT INTO apeiron.movement_action_contract (
    id, action_type, description, duration_ms, active_ms, recovery_ms,
    distance_cm, base_speed_cm_s, yaw_degrees, phase_window_policy,
    prediction_error_policy, reconciliation_contract_id,
    allow_windup_locomotion, allow_active_locomotion, allow_recovery_locomotion,
    allow_yaw_adjustment, root_motion_owner, contact_policy, speed_curve, vertical_curve, metadata
)
VALUES
('grounded_move_v1','move','Normal grounded walk/run/strafe authoritative command.',180,120,60,0,470,0,'grounded_move','bounded_smooth_correction','grounded_move_reconciliation',TRUE,TRUE,TRUE,TRUE,'movement','none','[]','[]','{"source":"reconstructed","ability_key":"move"}'),
('turn_v1_rate_limited_contextual','turn','Rate-limited contextual yaw update.',180,120,60,0,0,0,'turn','bounded_smooth_correction','turn_reconciliation',TRUE,TRUE,TRUE,TRUE,'movement','none','[]','[]','{"source":"reconstructed","ability_key":"turn"}'),
('dodge_v1_full_iframe','dodge','Player dodge owns position and iframe until movement handoff.',320,260,60,260,812.5,0,'dodge','bounded_smooth_correction','dodge_reconciliation',FALSE,FALSE,TRUE,TRUE,'movement','iframe','[{"t":0,"v":0.35},{"t":0.35,"v":1.0},{"t":1,"v":0.2}]','[{"t":0,"z":0},{"t":0.4,"z":18},{"t":1,"z":0}]','{"source":"reconstructed","ability_key":"dodge"}'),
('jump_v1_authoritative_grounded_handoff','leap','Player leap/jump owns airborne movement until grounded handoff.',620,560,60,280,452,0,'leap','bounded_smooth_correction','leap_reconciliation',TRUE,FALSE,TRUE,TRUE,'movement','grounded_handoff','[{"t":0,"v":0.25},{"t":0.35,"v":0.85},{"t":1,"v":0.35}]','[{"t":0,"z":0},{"t":0.5,"z":180},{"t":1,"z":0}]','{"source":"reconstructed","ability_key":"jump"}'),
('basic_attack_1_forward_cut_v1','grounded_skill','Basic attack 1 micro-commit forward cut.',350,140,120,55,190,0,'grounded_skill_action','bounded_smooth_correction','grounded_skill_action_reconciliation',FALSE,FALSE,TRUE,TRUE,'movement','none','[{"t":0,"v":0.35},{"t":0.35,"v":1.0},{"t":1,"v":0.2}]','[]','{"source":"reconstructed"}'),
('basic_attack_2_cross_cut_v1','grounded_skill','Basic attack 2 short lateral sword cut.',370,150,120,35,140,0,'grounded_skill_action','bounded_smooth_correction','grounded_skill_action_reconciliation',FALSE,FALSE,TRUE,TRUE,'movement','none','[{"t":0,"v":0.25},{"t":0.5,"v":0.8},{"t":1,"v":0.15}]','[]','{"source":"reconstructed"}'),
('basic_attack_3_shield_drive_v1','grounded_skill','Basic attack 3 short shield drive with contact carry.',620,260,180,200,330,0,'grounded_skill_action','bounded_smooth_correction','grounded_skill_action_reconciliation',FALSE,FALSE,TRUE,TRUE,'movement','carry_contact','[{"t":0,"v":0.2},{"t":0.25,"v":0.75},{"t":0.65,"v":1.0},{"t":1,"v":0.25}]','[]','{"source":"reconstructed"}'),
('shield_bash_front_push_v1','grounded_skill','Shield Bash frontal push action.',520,220,180,130,280,0,'grounded_skill_action','bounded_smooth_correction','grounded_skill_action_reconciliation',FALSE,FALSE,TRUE,TRUE,'movement','multi_target_push','[{"t":0,"v":0.2},{"t":0.4,"v":1.0},{"t":1,"v":0.15}]','[]','{"source":"reconstructed"}'),
('shield_rush_front_contact_v1','grounded_skill','Shield Rush committed rush with damage beginning half-cylinder in front.',830,430,240,340,470,0,'grounded_skill_action','bounded_smooth_correction','grounded_skill_action_reconciliation',FALSE,FALSE,TRUE,TRUE,'movement','multi_target_carry_push','[{"t":0,"v":0.1},{"t":0.2,"v":0.85},{"t":0.75,"v":1.0},{"t":1,"v":0.25}]','[]','{"front_contact_offset_cm":45,"source":"reconstructed"}'),
('wolf_dodge_lateral_leap_v1','dodge','Wolf low fast lateral/back diagonal dodge with full iframe.',520,420,100,210,520,0,'dodge','bounded_smooth_correction','dodge_reconciliation',FALSE,FALSE,TRUE,TRUE,'movement','iframe','[{"t":0,"v":0.4},{"t":0.35,"v":1.0},{"t":1,"v":0.2}]','[{"t":0,"z":0},{"t":0.4,"z":28},{"t":1,"z":0}]','{"source":"reconstructed"}'),
('wolf_lunge_airborne_v1','leap','Wolf lunge airborne phase; crosses target without losing air speed unless hard-controlled.',980,430,260,620,760,0,'leap','bounded_smooth_correction','leap_reconciliation',TRUE,FALSE,TRUE,TRUE,'movement','airborne_passthrough','[{"t":0,"v":0.4},{"t":0.18,"v":1.0},{"t":0.72,"v":0.85},{"t":1,"v":0.35}]','[{"t":0,"z":0},{"t":0.36,"z":120},{"t":1,"z":0}]','{"post_landing_inertia_multiplier":1.1,"source":"reconstructed"}')
ON CONFLICT (id) DO UPDATE SET
    action_type = EXCLUDED.action_type,
    description = EXCLUDED.description,
    duration_ms = EXCLUDED.duration_ms,
    active_ms = EXCLUDED.active_ms,
    recovery_ms = EXCLUDED.recovery_ms,
    distance_cm = EXCLUDED.distance_cm,
    base_speed_cm_s = EXCLUDED.base_speed_cm_s,
    phase_window_policy = EXCLUDED.phase_window_policy,
    prediction_error_policy = EXCLUDED.prediction_error_policy,
    reconciliation_contract_id = EXCLUDED.reconciliation_contract_id,
    allow_windup_locomotion = EXCLUDED.allow_windup_locomotion,
    allow_active_locomotion = EXCLUDED.allow_active_locomotion,
    allow_recovery_locomotion = EXCLUDED.allow_recovery_locomotion,
    allow_yaw_adjustment = EXCLUDED.allow_yaw_adjustment,
    root_motion_owner = EXCLUDED.root_motion_owner,
    contact_policy = EXCLUDED.contact_policy,
    speed_curve = EXCLUDED.speed_curve,
    vertical_curve = EXCLUDED.vertical_curve,
    metadata = EXCLUDED.metadata,
    updated_at = NOW();

INSERT INTO apeiron.skill_action_timing (skill_id, windup_ms, active_ms, recovery_ms, cooldown_ms, combo_window_ms, movement_lock_policy, queue_policy, cancel_policy, metadata)
VALUES
('player_basic_attack_1',90,140,120,0,2000,'contract','basic_combo','short_recovery','{"source":"reconstructed"}'),
('player_basic_attack_2',100,150,120,0,2000,'contract','basic_combo','short_recovery','{"source":"reconstructed"}'),
('player_basic_attack_3',180,260,180,0,2000,'contract','basic_combo','short_recovery','{"source":"reconstructed"}'),
('player_shield_bash',120,220,180,2600,0,'contract','none','recovery_only','{"source":"reconstructed"}'),
('player_shield_rush',160,430,240,5200,0,'contract','none','recovery_only','{"source":"reconstructed"}'),
('lunge',3600,430,500,4200,0,'contract','none','none','{"windup_movement":"circle_or_chase_setup","source":"reconstructed"}')
ON CONFLICT (skill_id) DO UPDATE SET
    windup_ms = EXCLUDED.windup_ms,
    active_ms = EXCLUDED.active_ms,
    recovery_ms = EXCLUDED.recovery_ms,
    cooldown_ms = EXCLUDED.cooldown_ms,
    combo_window_ms = EXCLUDED.combo_window_ms,
    movement_lock_policy = EXCLUDED.movement_lock_policy,
    queue_policy = EXCLUDED.queue_policy,
    cancel_policy = EXCLUDED.cancel_policy,
    metadata = EXCLUDED.metadata,
    updated_at = NOW();

INSERT INTO apeiron.skill_movement_action_binding (skill_id, movement_action_contract_id, starts_at_phase, handoff_policy, normal_input_policy, target_policy, contact_policy, is_enabled, metadata)
VALUES
('player_basic_attack_1','basic_attack_1_forward_cut_v1','active','explicit_recovery_handoff','blocked_during_owned_root','aim_direction','none',TRUE,'{"source":"reconstructed"}'),
('player_basic_attack_2','basic_attack_2_cross_cut_v1','active','explicit_recovery_handoff','blocked_during_owned_root','aim_direction','none',TRUE,'{"source":"reconstructed"}'),
('player_basic_attack_3','basic_attack_3_shield_drive_v1','active','explicit_recovery_handoff','blocked_during_owned_root','aim_direction','carry_contact',TRUE,'{"source":"reconstructed"}'),
('player_shield_bash','shield_bash_front_push_v1','active','explicit_recovery_handoff','blocked_during_owned_root','aim_direction','multi_target_push',TRUE,'{"source":"reconstructed"}'),
('player_shield_rush','shield_rush_front_contact_v1','active','explicit_recovery_handoff','blocked_during_owned_root','aim_direction','multi_target_carry_push',TRUE,'{"source":"reconstructed"}'),
('lunge','wolf_lunge_airborne_v1','active','grounded_handoff','blocked_during_airborne','target_cross_through','airborne_passthrough',TRUE,'{"source":"reconstructed"}')
ON CONFLICT (skill_id) DO UPDATE SET
    movement_action_contract_id = EXCLUDED.movement_action_contract_id,
    starts_at_phase = EXCLUDED.starts_at_phase,
    handoff_policy = EXCLUDED.handoff_policy,
    normal_input_policy = EXCLUDED.normal_input_policy,
    target_policy = EXCLUDED.target_policy,
    contact_policy = EXCLUDED.contact_policy,
    is_enabled = EXCLUDED.is_enabled,
    metadata = EXCLUDED.metadata,
    updated_at = NOW();
