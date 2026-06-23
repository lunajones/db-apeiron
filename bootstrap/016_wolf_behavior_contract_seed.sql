-- =========================================================
-- STEPPE WOLF MODERN BEHAVIOR / SKILL SEEDS
-- =========================================================

INSERT INTO apeiron.skill (
    id, name, description, archetype, skill_type,
    stamina_cost, mana_cost, health_cost,
    windup_ms, active_frames_ms, recovery_ms, cast_time_ms, cooldown_ms,
    cancel_window_start_ms, cancel_window_end_ms, iframe_start_ms, iframe_end_ms,
    min_range, max_range, cone_angle, max_targets, target_type, requires_target,
    base_damage, damage_type, elemental_type, posture_damage, armor_penetration,
    damage_multiplier, critical_bonus_multiplier,
    stun_duration_ms, root_duration_ms, knockback_force,
    movement_multiplier, locks_movement, movement_distance,
    combo_group, combo_index, combo_window_ms,
    is_interruptible, is_blockable, is_parryable, ignores_line_of_sight, ignores_collision
)
VALUES
('wolf_dodge','Wolf Dodge','Low fast lateral or backward hop with full invulnerability for the whole movement.', 'beast','dodge',
 6,0,0, 0,420,100,0,0, 0,0,0,520, 0,0,0,1,'self',FALSE, 0,'none',NULL,0,0,1,1, 0,0,0, 1,FALSE,2.1, NULL,NULL,0, FALSE,FALSE,FALSE,TRUE,TRUE),
('maul','Maul','Pressure counter: lateral mauling slash used when player overcommits.', 'beast','counter_attack',
 22,0,0, 180,260,360,0,5200, 0,0,0,180, 0,1.8,100,1,'enemy',TRUE, 13,'physical',NULL,24,0.05,1,1.25, 0,0,1.2, 1,FALSE,1.4, 'beast_counter',1,0, TRUE,TRUE,TRUE,FALSE,FALSE)
ON CONFLICT (id) DO UPDATE SET
    name = EXCLUDED.name,
    description = EXCLUDED.description,
    archetype = EXCLUDED.archetype,
    skill_type = EXCLUDED.skill_type,
    stamina_cost = EXCLUDED.stamina_cost,
    windup_ms = EXCLUDED.windup_ms,
    active_frames_ms = EXCLUDED.active_frames_ms,
    recovery_ms = EXCLUDED.recovery_ms,
    cooldown_ms = EXCLUDED.cooldown_ms,
    iframe_start_ms = EXCLUDED.iframe_start_ms,
    iframe_end_ms = EXCLUDED.iframe_end_ms,
    min_range = EXCLUDED.min_range,
    max_range = EXCLUDED.max_range,
    cone_angle = EXCLUDED.cone_angle,
    base_damage = EXCLUDED.base_damage,
    posture_damage = EXCLUDED.posture_damage,
    movement_distance = EXCLUDED.movement_distance,
    updated_at = NOW();

UPDATE apeiron.skill
SET
    windup_ms = 3600,
    active_frames_ms = 430,
    recovery_ms = 500,
    cooldown_ms = 4200,
    stamina_cost = 24,
    base_damage = 13,
    posture_damage = 24,
    movement_distance = 6.2,
    updated_at = NOW()
WHERE id = 'lunge';

INSERT INTO apeiron.skill_hitbox_damage_group (id, skill_id, description, max_hits_per_target, hit_interval_ms, can_multi_hit, metadata)
VALUES
('wolf_maul_damage','maul','Maul lateral counter applies one sweeping pressure hit.',1,0,FALSE,'{"target_lateral_cross":true}')
ON CONFLICT (id) DO UPDATE SET
    skill_id = EXCLUDED.skill_id,
    description = EXCLUDED.description,
    max_hits_per_target = EXCLUDED.max_hits_per_target,
    hit_interval_ms = EXCLUDED.hit_interval_ms,
    can_multi_hit = EXCLUDED.can_multi_hit,
    metadata = EXCLUDED.metadata,
    updated_at = NOW();

INSERT INTO apeiron.skill_hitbox_motion_profile (id, skill_id, motion_type, time_basis, description, metadata)
VALUES
('motion_wolf_maul_lateral_counter_v1','maul','timeline_sweep','hitbox_window_normalized','Maul sweeps laterally while crossing the target side-to-side.','{"target_lateral_cross":true}')
ON CONFLICT (id) DO UPDATE SET
    skill_id = EXCLUDED.skill_id,
    motion_type = EXCLUDED.motion_type,
    time_basis = EXCLUDED.time_basis,
    description = EXCLUDED.description,
    metadata = EXCLUDED.metadata,
    updated_at = NOW();

DELETE FROM apeiron.skill_hitbox_motion_sample
WHERE motion_profile_id = 'motion_wolf_maul_lateral_counter_v1';

INSERT INTO apeiron.skill_hitbox_motion_sample (
    motion_profile_id, sample_index, t, shape,
    offset_x, offset_y, offset_z, size_x, size_y, size_z,
    radius, length, min_angle_deg, max_angle_deg, metadata
)
VALUES
('motion_wolf_maul_lateral_counter_v1',0,0.00,'asymmetric_arc',65,40,95,0,0,125,58,120,-70,-25,'{}'),
('motion_wolf_maul_lateral_counter_v1',1,0.45,'asymmetric_arc',90,0,100,0,0,130,62,170,-25,25,'{}'),
('motion_wolf_maul_lateral_counter_v1',2,1.00,'asymmetric_arc',65,-40,95,0,0,125,58,120,25,70,'{}');

INSERT INTO apeiron.skill_hitbox_profile (
    id, skill_id, hitbox_index, hitbox_shape, hitbox_start_ms, hitbox_end_ms,
    offset_x, offset_y, offset_z, size_x, size_y, size_z, radius, length, angle,
    follows_caster, follows_projectile, can_multi_hit, max_hits_per_target,
    hit_interval_ms, friendly_fire, motion_profile_id, damage_group_id,
    min_angle_deg, max_angle_deg, start_radius, end_radius
)
VALUES
('hitbox_maul_0','maul',0,'temporal_sweep',180,440,80,0,100,0,0,130,62,170,140,TRUE,FALSE,FALSE,1,0,FALSE,'motion_wolf_maul_lateral_counter_v1','wolf_maul_damage',-70,70,58,62)
ON CONFLICT (id) DO UPDATE SET
    skill_id = EXCLUDED.skill_id,
    hitbox_index = EXCLUDED.hitbox_index,
    hitbox_shape = EXCLUDED.hitbox_shape,
    hitbox_start_ms = EXCLUDED.hitbox_start_ms,
    hitbox_end_ms = EXCLUDED.hitbox_end_ms,
    offset_x = EXCLUDED.offset_x,
    offset_y = EXCLUDED.offset_y,
    offset_z = EXCLUDED.offset_z,
    size_x = EXCLUDED.size_x,
    size_y = EXCLUDED.size_y,
    size_z = EXCLUDED.size_z,
    radius = EXCLUDED.radius,
    length = EXCLUDED.length,
    angle = EXCLUDED.angle,
    motion_profile_id = EXCLUDED.motion_profile_id,
    damage_group_id = EXCLUDED.damage_group_id,
    min_angle_deg = EXCLUDED.min_angle_deg,
    max_angle_deg = EXCLUDED.max_angle_deg,
    start_radius = EXCLUDED.start_radius,
    end_radius = EXCLUDED.end_radius,
    updated_at = NOW();

INSERT INTO apeiron.movement_action_contract (
    id, action_type, description, duration_ms, active_ms, recovery_ms,
    distance_cm, base_speed_cm_s, yaw_degrees, phase_window_policy,
    prediction_error_policy, reconciliation_contract_id,
    allow_windup_locomotion, allow_active_locomotion, allow_recovery_locomotion,
    allow_yaw_adjustment, root_motion_owner, contact_policy, speed_curve, vertical_curve, metadata
)
VALUES
('wolf_bite_melee_commit_v1','grounded_skill','Wolf bite close-range committed melee action without displacement.',520,220,180,0,0,0,'grounded_skill_action','bounded_smooth_correction','grounded_skill_action_reconciliation',FALSE,FALSE,TRUE,TRUE,'movement','melee_contact','[]','[]','{"source":"canonical_bootstrap","skill":"bite"}'),
('wolf_maul_lateral_counter_v1','grounded_skill','Wolf maul lateral counter dash that crosses the player side-to-side during pressure punish.',800,260,360,140,420,0,'grounded_skill_action','bounded_smooth_correction','grounded_skill_action_reconciliation',FALSE,FALSE,TRUE,TRUE,'movement','lateral_counter_contact','[{"t":0,"v":0.15},{"t":0.25,"v":0.85},{"t":0.62,"v":1.0},{"t":1,"v":0.20}]','[]','{"source":"canonical_bootstrap","setup_policy":"wolf_maul_pressure_counter_v1"}')
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

UPDATE apeiron.skill AS s
SET
    base_damage = tuned.base_damage,
    posture_damage = tuned.posture_damage,
    updated_at = NOW()
FROM (
    VALUES
        ('bite', 9.0::DOUBLE PRECISION, 17.6::DOUBLE PRECISION),
        ('lunge', 6.5::DOUBLE PRECISION, 19.2::DOUBLE PRECISION),
        ('maul', 6.5::DOUBLE PRECISION, 19.2::DOUBLE PRECISION)
) AS tuned(id, base_damage, posture_damage)
WHERE s.id = tuned.id;

INSERT INTO apeiron.skill_action_timing (
    skill_id, windup_ms, active_ms, recovery_ms, cooldown_ms,
    combo_window_ms, movement_lock_policy, queue_policy, cancel_policy, metadata
)
VALUES
('bite',120,220,180,900,0,'contract','none','none','{"source":"canonical_bootstrap"}'),
('wolf_dodge',0,420,100,0,0,'contract','none','none','{"full_iframe":true,"source":"canonical_bootstrap"}'),
('maul',180,260,360,5200,0,'contract','none','none','{"pressureCounter":true,"source":"canonical_bootstrap"}')
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

INSERT INTO apeiron.skill_movement_action_binding (
    skill_id, movement_action_contract_id, starts_at_phase, handoff_policy,
    normal_input_policy, target_policy, contact_policy, is_enabled, metadata
)
VALUES
('bite','wolf_bite_melee_commit_v1','active','explicit_recovery_handoff','blocked_during_owned_root','target_direction','melee_contact',TRUE,'{"source":"canonical_bootstrap"}'),
('wolf_dodge','wolf_dodge_lateral_leap_v1','active','explicit_recovery_handoff','blocked_during_owned_root','evasion_direction','iframe',TRUE,'{"source":"canonical_bootstrap"}'),
('maul','wolf_maul_lateral_counter_v1','active','explicit_recovery_handoff','blocked_during_owned_root','target_lateral_cross','lateral_counter_contact',TRUE,'{"source":"canonical_bootstrap"}')
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

DELETE FROM apeiron.skill_slot
WHERE skill_set_id = 'skillset_steppe_wolf';

INSERT INTO apeiron.skill_slot (
    skill_set_id, skill_id, slot_index, is_enabled,
    priority, usage_weight, cooldown_override_ms,
    min_target_hp_percent, max_target_hp_percent, min_self_hp_percent, max_self_hp_percent,
    required_distance_min, required_distance_max, requires_line_of_sight,
    opener_weight, finisher_weight, shared_cooldown_group, use_only_in_combat
)
VALUES
('skillset_steppe_wolf','bite',1,TRUE,70,1.15,NULL,NULL,NULL,0.10,NULL,0.0,1.8,TRUE,0.35,0.35,'wolf_melee',TRUE),
('skillset_steppe_wolf','lunge',2,TRUE,90,0.85,NULL,NULL,NULL,0.20,NULL,1.8,7.0,TRUE,0.70,0.25,'wolf_lunge',TRUE),
('skillset_steppe_wolf','maul',3,TRUE,82,0.55,NULL,NULL,NULL,0.15,NULL,0.0,1.9,TRUE,0.15,0.45,'wolf_counter',TRUE),
('skillset_steppe_wolf','wolf_dodge',4,TRUE,95,1.25,NULL,NULL,NULL,0.05,NULL,0.0,5.5,TRUE,0.05,0.05,'wolf_evasion',TRUE);

INSERT INTO apeiron.creature_behavior_runtime_contract (
    id, creature_template_id, description, aggression_curve, range_policy, orbit_policy, pressure_policy, stamina_policy, metadata
)
VALUES (
    'contract_wolf_pack_harasser_v1',
    'steppe_wolf',
    'Steppe wolf harasser: evasive, lateral, punishes overcommit, uses lunge setup with moving windup.',
    '{"base":0.68,"pressureRamp":0.22,"defensiveTargetBonus":0.12}',
    '{"recoverDistanceWithRun":true,"chaseLungeSetup":true,"biteAtCloseRange":true,"preferredMinCm":130,"preferredMaxCm":520,"desiredRangeCm":420,"chaseRangeCm":760,"retreatRangeCm":220,"orbitSpeedCmS":360,"chaseSpeedCmS":620,"lungeSpeedCmS":760,"maulSpeedCmS":420,"retreatSpeedCmS":520}',
    '{"preferLongSideCommit":true,"sideFlipChanceMultiplier":0.35,"minimumSideCommitMs":3200,"flankBeforeLunge":true}',
    '{"dodgeUnderPressure":true,"maulCounterUnderPressure":true,"maulCounterChance":0.22,"dodgeRetreatMultiplier":0.70,"globalDodgeMultiplier":0.85,"repeatSkillPenaltyMultiplier":0.65,"commitThreatWeight":0.28,"closingThreatWeight":0.18,"defensiveBiteWeight":0.14,"fleeingLungeWeight":0.20,"lowResourceRiskFloor":0.16,"dodgeCommittedThreatMultiplier":1.12,"vulnerableBiteMultiplier":1.16,"vulnerableMaulMultiplier":1.10,"tacticalDestinationDistanceCm":180}',
    '{"max":100,"dodgeCostMultiplier":0.50,"regenPerSecond":12,"runDrainEnabled":true,"zeroStaminaLockoutUntilFull":true}',
    '{"source":"canonical_bootstrap"}'
)
ON CONFLICT (id) DO UPDATE SET
    creature_template_id = EXCLUDED.creature_template_id,
    description = EXCLUDED.description,
    aggression_curve = EXCLUDED.aggression_curve,
    range_policy = EXCLUDED.range_policy,
    orbit_policy = EXCLUDED.orbit_policy,
    pressure_policy = EXCLUDED.pressure_policy,
    stamina_policy = EXCLUDED.stamina_policy,
    metadata = EXCLUDED.metadata,
    updated_at = NOW();

INSERT INTO apeiron.creature_target_opportunity_policy (
    id, description, commit_angle_max_deg, min_commit_distance_cm, max_commit_distance_cm,
    approach_min_distance_cm, approach_max_distance_cm, bite_range_cm,
    lunge_min_range_cm, lunge_max_range_cm, maul_pressure_threshold,
    target_memory_ms, no_ready_skill_memory_policy, candidate_cooldown_visibility,
    allow_backside_commit, metadata
)
VALUES (
    'opportunity_wolf_harasser_v1',
    'Steppe wolf opportunity window: can commit from behind, exposes candidate/cooldown diagnostics, and never records no-ready-skill as attack failure.',
    180,
    0,
    700,
    180,
    560,
    260,
    180,
    700,
    0.58,
    1800,
    'observe_only',
    TRUE,
    TRUE,
    '{"source":"canonical_bootstrap","candidate_skills_diagnostics":true,"cooldown_skills_diagnostics":true}'
)
ON CONFLICT (id) DO UPDATE SET
    description = EXCLUDED.description,
    commit_angle_max_deg = EXCLUDED.commit_angle_max_deg,
    min_commit_distance_cm = EXCLUDED.min_commit_distance_cm,
    max_commit_distance_cm = EXCLUDED.max_commit_distance_cm,
    approach_min_distance_cm = EXCLUDED.approach_min_distance_cm,
    approach_max_distance_cm = EXCLUDED.approach_max_distance_cm,
    bite_range_cm = EXCLUDED.bite_range_cm,
    lunge_min_range_cm = EXCLUDED.lunge_min_range_cm,
    lunge_max_range_cm = EXCLUDED.lunge_max_range_cm,
    maul_pressure_threshold = EXCLUDED.maul_pressure_threshold,
    target_memory_ms = EXCLUDED.target_memory_ms,
    no_ready_skill_memory_policy = EXCLUDED.no_ready_skill_memory_policy,
    candidate_cooldown_visibility = EXCLUDED.candidate_cooldown_visibility,
    allow_backside_commit = EXCLUDED.allow_backside_commit,
    metadata = EXCLUDED.metadata,
    updated_at = NOW();

INSERT INTO apeiron.creature_orbit_policy (
    id, behavior_contract_id, description, orbit_locomotion_mode, orbit_speed_scale,
    min_orbit_duration_ms, side_switch_cooldown_ms, allow_side_switch_when_target_faces,
    prefer_long_side_commit, side_flip_chance_multiplier, lock_side_during_setup, metadata
)
VALUES (
    'orbit_wolf_harasser_combat_walk_v1',
    'contract_wolf_pack_harasser_v1',
    'Wolf circles in combat-walk/run setup, keeps side long enough to read naturally, and avoids rapid left/right thrashing.',
    'combat_walk',
    0.55,
    700,
    900,
    TRUE,
    TRUE,
    0.35,
    TRUE,
    '{"source":"canonical_bootstrap","server_must_not_pick_side_by_target_id":true}'
)
ON CONFLICT (id) DO UPDATE SET
    behavior_contract_id = EXCLUDED.behavior_contract_id,
    description = EXCLUDED.description,
    orbit_locomotion_mode = EXCLUDED.orbit_locomotion_mode,
    orbit_speed_scale = EXCLUDED.orbit_speed_scale,
    min_orbit_duration_ms = EXCLUDED.min_orbit_duration_ms,
    side_switch_cooldown_ms = EXCLUDED.side_switch_cooldown_ms,
    allow_side_switch_when_target_faces = EXCLUDED.allow_side_switch_when_target_faces,
    prefer_long_side_commit = EXCLUDED.prefer_long_side_commit,
    side_flip_chance_multiplier = EXCLUDED.side_flip_chance_multiplier,
    lock_side_during_setup = EXCLUDED.lock_side_during_setup,
    metadata = EXCLUDED.metadata,
    updated_at = NOW();

UPDATE apeiron.creature_behavior_runtime_contract
SET
    target_opportunity_policy_id = 'opportunity_wolf_harasser_v1',
    orbit_policy_id = 'orbit_wolf_harasser_combat_walk_v1',
    updated_at = NOW()
WHERE id = 'contract_wolf_pack_harasser_v1';

INSERT INTO apeiron.creature_evasion_policy (
    id, behavior_contract_id, description, dodge_skill_id, max_chain_count,
    stamina_cost_multiplier, retreat_chance_multiplier, lateral_bias, backstep_bias,
    pressure_threshold, cooldown_ms, metadata
)
VALUES (
    'wolf_evasion_pressure_v1',
    'contract_wolf_pack_harasser_v1',
    'Allows up to four chained dodges, biased lateral, with lower retreat dodge frequency.',
    'wolf_dodge',
    4,
    0.50,
    0.70,
    0.72,
    0.28,
    0.42,
    180,
    '{"source":"canonical_bootstrap"}'
)
ON CONFLICT (id) DO UPDATE SET
    behavior_contract_id = EXCLUDED.behavior_contract_id,
    description = EXCLUDED.description,
    dodge_skill_id = EXCLUDED.dodge_skill_id,
    max_chain_count = EXCLUDED.max_chain_count,
    stamina_cost_multiplier = EXCLUDED.stamina_cost_multiplier,
    retreat_chance_multiplier = EXCLUDED.retreat_chance_multiplier,
    lateral_bias = EXCLUDED.lateral_bias,
    backstep_bias = EXCLUDED.backstep_bias,
    pressure_threshold = EXCLUDED.pressure_threshold,
    cooldown_ms = EXCLUDED.cooldown_ms,
    metadata = EXCLUDED.metadata,
    updated_at = NOW();

INSERT INTO apeiron.creature_skill_setup_policy (
    id, behavior_contract_id, skill_id, setup_type, min_setup_ms, max_setup_ms,
    commit_distance_cm, preferred_min_range_cm, preferred_max_range_cm,
    movement_tactic, lock_side_during_setup, is_enabled, metadata
)
VALUES
('wolf_lunge_flank_windup_v1','contract_wolf_pack_harasser_v1','lunge','moving_windup',3000,4200,520,180,700,'circle_then_curve_to_target',TRUE,TRUE,'{"airSpeedMultiplier":1.2,"postLandingInertiaMultiplier":1.1,"targetPassthrough":true}'),
('wolf_lunge_chase_windup_v1','contract_wolf_pack_harasser_v1','lunge','chase_windup',1200,2600,640,520,1200,'run_chase_then_jump',FALSE,TRUE,'{"whenTargetFlees":true,"airSpeedMultiplier":1.2}'),
('wolf_maul_pressure_counter_v1','contract_wolf_pack_harasser_v1','maul','pressure_counter',120,320,160,0,220,'lateral_counter_dash',TRUE,TRUE,'{"trigger":"player_overcommit"}')
ON CONFLICT (id) DO UPDATE SET
    behavior_contract_id = EXCLUDED.behavior_contract_id,
    skill_id = EXCLUDED.skill_id,
    setup_type = EXCLUDED.setup_type,
    min_setup_ms = EXCLUDED.min_setup_ms,
    max_setup_ms = EXCLUDED.max_setup_ms,
    commit_distance_cm = EXCLUDED.commit_distance_cm,
    preferred_min_range_cm = EXCLUDED.preferred_min_range_cm,
    preferred_max_range_cm = EXCLUDED.preferred_max_range_cm,
    movement_tactic = EXCLUDED.movement_tactic,
    lock_side_during_setup = EXCLUDED.lock_side_during_setup,
    is_enabled = EXCLUDED.is_enabled,
    metadata = EXCLUDED.metadata,
    updated_at = NOW();

DELETE FROM apeiron.creature_skill_behavior_binding
WHERE behavior_contract_id = 'contract_wolf_pack_harasser_v1';

INSERT INTO apeiron.creature_skill_behavior_binding (
    id, behavior_contract_id, skill_id, tactical_state, decision_phase, setup_policy_id,
    min_range_cm, max_range_cm, priority, usage_weight, cooldown_group,
    requires_line_of_sight, is_enabled, metadata
)
VALUES
('wolf_bite_approach_acquire_v1','contract_wolf_pack_harasser_v1','bite','approach','acquire',NULL,0,260,80,1.20,'wolf_melee',TRUE,TRUE,'{"close_range_pressure":true}'),
('wolf_bite_circle_reposition_v1','contract_wolf_pack_harasser_v1','bite','circle','reposition',NULL,0,260,62,0.85,'wolf_melee',TRUE,TRUE,'{"prevents_orbit_only_loop":true}'),
('wolf_lunge_approach_acquire_v1','contract_wolf_pack_harasser_v1','lunge','approach','acquire','wolf_lunge_chase_windup_v1',180,700,92,0.95,'wolf_lunge',TRUE,TRUE,'{"commit_attack":true,"any_target_state_wildcard":true}'),
('wolf_lunge_circle_reposition_v1','contract_wolf_pack_harasser_v1','lunge','circle','reposition','wolf_lunge_flank_windup_v1',180,700,90,0.85,'wolf_lunge',TRUE,TRUE,'{"flank_then_curve":true,"commit_angle_max_deg":180}'),
('wolf_maul_pressure_counter_v1','contract_wolf_pack_harasser_v1','maul','pressure','counter','wolf_maul_pressure_counter_v1',0,220,86,0.58,'wolf_counter',TRUE,TRUE,'{"counter_player_overcommit":true}'),
('wolf_dodge_pressure_evasion_v1','contract_wolf_pack_harasser_v1','wolf_dodge','pressure','evade',NULL,0,550,95,1.25,'wolf_evasion',TRUE,TRUE,'{"full_iframe":true,"chain_budget_from_evasion_policy":true}');
