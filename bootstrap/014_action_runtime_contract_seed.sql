-- =========================================================
-- ACTION RUNTIME / RECONCILIATION SEEDS
-- =========================================================

INSERT INTO apeiron.movement_reconciliation_contract (
    id, category, description, max_smooth_error_cm, hard_snap_error_cm,
    smoothing_time_ms, yaw_tolerance_deg, owns_position, owns_yaw,
    allows_client_prediction, input_policy, handoff_policy, metadata
)
VALUES
('grounded_move_reconciliation','grounded_move','Normal grounded walk/run/strafe reconciliation.',35,180,90,8,FALSE,FALSE,TRUE,'normal','continuous','{"source":"canonical_bootstrap"}'),
('grounded_skill_action_reconciliation','grounded_skill_action','Committed grounded skill movement reconciliation with explicit post-action handoff.',25,140,70,10,TRUE,TRUE,TRUE,'blocked_during_owned_root','explicit','{"source":"canonical_bootstrap"}'),
('dodge_reconciliation','dodge','Dodge owns position and iframe window until movement end.',22,130,60,14,TRUE,TRUE,TRUE,'blocked_during_owned_root','explicit','{"source":"canonical_bootstrap"}'),
('leap_reconciliation','leap','Leap owns vertical and horizontal movement until grounded handoff.',30,170,70,12,TRUE,TRUE,TRUE,'blocked_during_airborne','grounded_handoff','{"source":"canonical_bootstrap"}'),
('turn_reconciliation','turn','Rate-limited contextual turn reconciliation.',18,90,45,4,FALSE,TRUE,TRUE,'turn_only','continuous','{"source":"canonical_bootstrap"}'),
('post_action_handoff_reconciliation','post_action_handoff','Recovery phase contract that prevents normal locomotion from racing skill recovery.',20,110,55,8,TRUE,FALSE,TRUE,'buffer_until_handoff','explicit','{"source":"canonical_bootstrap"}')
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
('grounded_move_v1','move','Normal grounded walk/run/strafe authoritative command.',180,120,60,0,470,0,'grounded_move','bounded_smooth_correction','grounded_move_reconciliation',TRUE,TRUE,TRUE,TRUE,'movement','none','[]','[]','{"source":"canonical_bootstrap","ability_key":"move"}'),
('turn_v1_rate_limited_contextual','turn','Rate-limited contextual yaw update.',180,120,60,0,0,720,'turn','bounded_smooth_correction','turn_reconciliation',TRUE,TRUE,TRUE,TRUE,'movement','none','[]','[]','{"source":"canonical_bootstrap","ability_key":"turn","yaw_rate_deg_per_sec":720}'),
('dodge_v1_full_iframe','dodge','Player dodge owns position and iframe until movement handoff.',320,260,60,360,1125,0,'dodge','bounded_smooth_correction','dodge_reconciliation',FALSE,FALSE,TRUE,TRUE,'movement','iframe','[{"t":0,"v":0.35},{"t":0.35,"v":1.0},{"t":1,"v":0.2}]','[{"t":0,"z":0},{"t":0.4,"z":18},{"t":1,"z":0}]','{"source":"canonical_bootstrap","ability_key":"dodge","speed_semantics":"authored_base_speed_distance_over_duration"}'),
('jump_v1_authoritative_grounded_handoff','leap','Player leap/jump owns airborne movement until grounded handoff.',980,920,60,280,285.7,0,'leap','bounded_smooth_correction','leap_reconciliation',TRUE,FALSE,TRUE,TRUE,'movement','grounded_handoff','[{"t":0,"v":0.35},{"t":0.35,"v":0.95},{"t":1,"v":0.62}]','[{"t":0,"z":0},{"t":0.50,"z":118},{"t":1,"z":0}]','{"source":"canonical_bootstrap","ability_key":"jump","airborne_duration_ms":980,"vertical_motion_model":"ballistic","jump_z_velocity":480,"gravity_scale":1,"gravity_z_cm_s2":980,"expected_apex_ms":490,"landing_detection_policy":"server_grounded_handoff","ground_z_policy":"server_position_is_actor_root","capsule_base_offset":0,"allows_air_control":true,"air_control_modifier":0.35}'),
('basic_attack_1_forward_cut_v1','grounded_skill','Basic attack 1 short shield jab: one player cylinder forward.',350,140,120,84,240,0,'grounded_skill_action','bounded_smooth_correction','grounded_skill_action_reconciliation',FALSE,FALSE,TRUE,TRUE,'movement','none','[{"t":0,"v":0.35},{"t":0.35,"v":1.0},{"t":1,"v":0.2}]','[]','{"source":"canonical_bootstrap","design_note":"one-cylinder shield jab"}'),
('basic_attack_2_cross_cut_v1','grounded_skill','Basic attack 2 short left-to-right shield sweep.',370,150,120,42,114,0,'grounded_skill_action','bounded_smooth_correction','grounded_skill_action_reconciliation',FALSE,FALSE,TRUE,TRUE,'movement','none','[{"t":0,"v":0.25},{"t":0.5,"v":0.8},{"t":1,"v":0.15}]','[]','{"source":"canonical_bootstrap","design_note":"half-cylinder commit, temporal frontal sweep"}'),
('basic_attack_3_shield_drive_v1','grounded_skill','Basic attack 3 committed overhead shield punch with contact carry and interrupt.',620,260,180,252,406.4,0,'grounded_skill_action','bounded_smooth_correction','grounded_skill_action_reconciliation',FALSE,FALSE,TRUE,TRUE,'movement','carry_contact','[{"t":0,"v":0.2},{"t":0.25,"v":0.75},{"t":0.65,"v":1.0},{"t":1,"v":0.25}]','[]','{"source":"canonical_bootstrap","design_note":"double-length committed shield-tip punch"}'),
('shield_bash_front_push_v1','grounded_skill','Shield Bash short forward step with temporal frontal stun/push.',300,170,120,95,541,0,'grounded_skill_action','bounded_smooth_correction','grounded_skill_action_reconciliation',FALSE,FALSE,TRUE,TRUE,'movement','multi_target_push','[{"t":0,"v":0.2},{"t":0.4,"v":1.0},{"t":1,"v":0.15}]','[]','{"source":"canonical_bootstrap","design_note":"short shield step: one player cylinder forward"}'),
('shield_rush_front_contact_v1','grounded_skill','Shield Rush committed long rush with damage beginning at shield/body contact.',1100,720,260,864,1033.2,0,'grounded_skill_action','bounded_smooth_correction','grounded_skill_action_reconciliation',FALSE,FALSE,TRUE,TRUE,'movement','multi_target_carry_push','[{"t":0,"v":0.1},{"t":0.2,"v":0.85},{"t":0.75,"v":1.0},{"t":1,"v":0.25}]','[]','{"front_contact_offset_cm":8,"source":"canonical_bootstrap","design_note":"nine-cylinder committed shield rush with close curved shield-front contact"}'),
('wolf_dodge_lateral_leap_v1','dodge','Wolf low fast lateral/back diagonal dodge with full iframe.',520,420,100,210,520,0,'dodge','bounded_smooth_correction','dodge_reconciliation',FALSE,FALSE,TRUE,TRUE,'movement','iframe','[{"t":0,"v":0.4},{"t":0.35,"v":1.0},{"t":1,"v":0.2}]','[{"t":0,"z":0},{"t":0.4,"z":28},{"t":1,"z":0}]','{"source":"canonical_bootstrap"}'),
('low_fast_lunge_v1','leap','Wolf low fast raking lunge; after setup it aligns for 100ms, spends 520ms in a low airborne pass-through, then carries 200ms of grounded inertia.',820,620,200,2148,2619.5,0,'leap','bounded_smooth_correction','leap_reconciliation',TRUE,FALSE,TRUE,TRUE,'movement','airborne_passthrough','[{"t":0,"v":0.22},{"t":0.12,"v":0.62},{"t":0.24,"v":0.98},{"t":0.68,"v":1.0},{"t":0.80,"v":0.48},{"t":1,"v":0.20}]','[{"t":0,"z":0},{"t":0.12,"z":0},{"t":0.30,"z":7},{"t":0.52,"z":8},{"t":0.76,"z":0},{"t":1,"z":0}]','{"post_landing_inertia_multiplier":0.62,"source":"canonical_bootstrap","canonical_id":"low_fast_lunge_v1","pre_commit_ms":100,"airborne_duration_ms":520,"landing_inertia_ms":200,"vertical_motion_model":"low_raking_arc","jump_z_velocity":0,"gravity_scale":1,"gravity_z_cm_s2":980,"expected_apex_ms":260,"landing_detection_policy":"server_grounded_handoff","ground_z_policy":"server_position_is_actor_root","capsule_base_offset":0,"allows_air_control":false,"air_control_modifier":0,"orientation_policy_id":"orientation_lunge_flank_commit_v1","envelope_policy_id":"envelope_lunge_low_raking_100_520_200_v1"}')
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

INSERT INTO apeiron.action_orientation_policy (
    id, owner_kind, description,
    body_yaw_source, focus_yaw_source, attack_yaw_source,
    body_turn_rate_deg_s, focus_turn_rate_deg_s, attack_turn_rate_deg_s,
    commit_align_ms, attack_yaw_latch_policy,
    allow_head_look_while_strafing, allow_body_side_on_movement, metadata
)
VALUES
('orientation_lunge_flank_commit_v1','shared','Flank/circle keeps focus on target while root/body follows movement, then body aligns to a latched attack yaw during pre-lunge commit.',
 'movement_direction_until_commit','target','commit_target_snapshot',
 420,900,720,
 100,'latch_at_takeoff',
 TRUE,TRUE,'{"source":"canonical_bootstrap","supports":"creature_and_player_actions","phase_model":"tactical_setup_to_pre_commit_to_airborne"}'),
('orientation_forward_commit_v1','shared','Forward committed attack orientation: body and attack yaw follow aim/control direction at commit, then active frames use latched attack yaw.',
 'aim_direction','aim_direction','commit_aim_snapshot',
 540,720,720,
 0,'latch_at_active_start',
 FALSE,FALSE,'{"source":"canonical_bootstrap","supports":"player_shield_bash_rush_basic"}'),
('orientation_dodge_exit_v1','shared','Evasive burst orientation: movement direction owns the burst, focus may remain on target/camera, attack yaw is none.',
 'movement_direction','focus_or_camera','none',
 720,900,0,
 0,'none',
 TRUE,TRUE,'{"source":"canonical_bootstrap","supports":"dodge_exit_transition"}')
ON CONFLICT (id) DO UPDATE SET
    owner_kind = EXCLUDED.owner_kind,
    description = EXCLUDED.description,
    body_yaw_source = EXCLUDED.body_yaw_source,
    focus_yaw_source = EXCLUDED.focus_yaw_source,
    attack_yaw_source = EXCLUDED.attack_yaw_source,
    body_turn_rate_deg_s = EXCLUDED.body_turn_rate_deg_s,
    focus_turn_rate_deg_s = EXCLUDED.focus_turn_rate_deg_s,
    attack_turn_rate_deg_s = EXCLUDED.attack_turn_rate_deg_s,
    commit_align_ms = EXCLUDED.commit_align_ms,
    attack_yaw_latch_policy = EXCLUDED.attack_yaw_latch_policy,
    allow_head_look_while_strafing = EXCLUDED.allow_head_look_while_strafing,
    allow_body_side_on_movement = EXCLUDED.allow_body_side_on_movement,
    metadata = EXCLUDED.metadata,
    updated_at = NOW();

INSERT INTO apeiron.action_envelope_policy (
    id, owner_kind, description,
    pre_commit_ms, airborne_ms, landing_inertia_ms,
    pre_commit_direction_policy, airborne_direction_policy,
    inertia_direction_policy, tactical_reentry_policy,
    speed_curve, vertical_curve, metadata
)
VALUES
('envelope_lunge_low_raking_100_520_200_v1','shared','Lunge envelope with short post-flank alignment run, low airborne pass-through, and explicit landing inertia before tactical reentry.',
 100,520,200,
 'curve_from_setup_to_attack_yaw','latched_attack_yaw','preserve_landing_exit','blend_body_yaw_to_next_tactic',
 '[{"t":0,"v":0.22},{"t":0.12,"v":0.62},{"t":0.24,"v":0.98},{"t":0.68,"v":1.0},{"t":0.80,"v":0.48},{"t":1,"v":0.20}]',
 '[{"t":0,"z":0},{"t":0.12,"z":0},{"t":0.30,"z":7},{"t":0.52,"z":8},{"t":0.76,"z":0},{"t":1,"z":0}]',
 '{"source":"canonical_bootstrap","pre_commit_meaning":"post-flank body alignment and straightening run","airborne_shape":"low_long_raking","landing_inertia_is_action_transition":true}'),
('envelope_grounded_short_commit_v1','shared','Short grounded commit envelope for bash/basic actions that need compact owned movement and explicit handoff.',
 0,0,120,
 'none','none','preserve_exit_direction','explicit_handoff_to_grounded_move',
 '[{"t":0,"v":0.25},{"t":0.45,"v":1.0},{"t":1,"v":0.18}]',
 '[]',
 '{"source":"canonical_bootstrap","supports":"basic_attack_1_2_shield_bash"}'),
('envelope_grounded_rush_contact_v1','shared','Long grounded rush/contact envelope with close front contact and carry push release.',
 0,0,260,
 'none','none','preserve_contact_push_direction','explicit_handoff_to_grounded_move',
 '[{"t":0,"v":0.10},{"t":0.20,"v":0.85},{"t":0.75,"v":1.0},{"t":1,"v":0.25}]',
 '[]',
 '{"source":"canonical_bootstrap","supports":"shield_rush_and_basic_attack_3_contact_drive"}')
ON CONFLICT (id) DO UPDATE SET
    owner_kind = EXCLUDED.owner_kind,
    description = EXCLUDED.description,
    pre_commit_ms = EXCLUDED.pre_commit_ms,
    airborne_ms = EXCLUDED.airborne_ms,
    landing_inertia_ms = EXCLUDED.landing_inertia_ms,
    pre_commit_direction_policy = EXCLUDED.pre_commit_direction_policy,
    airborne_direction_policy = EXCLUDED.airborne_direction_policy,
    inertia_direction_policy = EXCLUDED.inertia_direction_policy,
    tactical_reentry_policy = EXCLUDED.tactical_reentry_policy,
    speed_curve = EXCLUDED.speed_curve,
    vertical_curve = EXCLUDED.vertical_curve,
    metadata = EXCLUDED.metadata,
    updated_at = NOW();

INSERT INTO apeiron.skill_action_timing (skill_id, windup_ms, active_ms, recovery_ms, cooldown_ms, combo_window_ms, movement_lock_policy, queue_policy, cancel_policy, metadata)
VALUES
('player_basic_attack_1',90,140,120,0,2000,'contract','basic_combo','short_recovery','{"source":"canonical_bootstrap"}'),
('player_basic_attack_2',100,150,120,0,2000,'contract','basic_combo','short_recovery','{"source":"canonical_bootstrap"}'),
('player_basic_attack_3',180,260,180,0,2000,'contract','basic_combo','short_recovery','{"source":"canonical_bootstrap"}'),
('player_shield_bash',110,170,120,26000,0,'contract','none','after_recovery','{"source":"canonical_bootstrap"}'),
('player_shield_rush',160,720,260,32000,0,'contract','none','after_recovery','{"source":"canonical_bootstrap"}'),
('lunge',3600,620,200,9000,0,'contract','none','none','{"windup_movement":"circle_or_chase_setup","source":"canonical_bootstrap","pre_commit_ms":100,"airborne_damage_window":true,"grounded_inertia_tail":true}')
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
('player_basic_attack_1','basic_attack_1_forward_cut_v1','active','explicit_recovery_handoff','blocked_during_owned_root','aim_direction','none',TRUE,'{"source":"canonical_bootstrap"}'),
('player_basic_attack_2','basic_attack_2_cross_cut_v1','active','explicit_recovery_handoff','blocked_during_owned_root','aim_direction','none',TRUE,'{"source":"canonical_bootstrap"}'),
('player_basic_attack_3','basic_attack_3_shield_drive_v1','active','explicit_recovery_handoff','blocked_during_owned_root','aim_direction','carry_contact',TRUE,'{"source":"canonical_bootstrap"}'),
('player_shield_bash','shield_bash_front_push_v1','active','explicit_recovery_handoff','blocked_during_owned_root','aim_direction','multi_target_push',TRUE,'{"source":"canonical_bootstrap"}'),
('player_shield_rush','shield_rush_front_contact_v1','active','explicit_recovery_handoff','blocked_during_owned_root','aim_direction','multi_target_carry_push',TRUE,'{"source":"canonical_bootstrap"}'),
('lunge','low_fast_lunge_v1','active','grounded_handoff','blocked_during_airborne','target_cross_through','airborne_passthrough',TRUE,'{"source":"canonical_bootstrap","canonical_movement":"low_fast_lunge_v1"}')
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

INSERT INTO apeiron.skill_action_policy_binding (
    skill_id, action_orientation_policy_id, action_envelope_policy_id, is_enabled, metadata
)
VALUES
('player_basic_attack_1','orientation_forward_commit_v1','envelope_grounded_short_commit_v1',TRUE,'{"source":"canonical_bootstrap"}'),
('player_basic_attack_2','orientation_forward_commit_v1','envelope_grounded_short_commit_v1',TRUE,'{"source":"canonical_bootstrap"}'),
('player_basic_attack_3','orientation_forward_commit_v1','envelope_grounded_rush_contact_v1',TRUE,'{"source":"canonical_bootstrap"}'),
('player_shield_bash','orientation_forward_commit_v1','envelope_grounded_short_commit_v1',TRUE,'{"source":"canonical_bootstrap"}'),
('player_shield_rush','orientation_forward_commit_v1','envelope_grounded_rush_contact_v1',TRUE,'{"source":"canonical_bootstrap"}'),
('lunge','orientation_lunge_flank_commit_v1','envelope_lunge_low_raking_100_520_200_v1',TRUE,'{"source":"canonical_bootstrap","post_flank_pre_lunge_commit":true}')
ON CONFLICT (skill_id) DO UPDATE SET
    action_orientation_policy_id = EXCLUDED.action_orientation_policy_id,
    action_envelope_policy_id = EXCLUDED.action_envelope_policy_id,
    is_enabled = EXCLUDED.is_enabled,
    metadata = EXCLUDED.metadata,
    updated_at = NOW();
