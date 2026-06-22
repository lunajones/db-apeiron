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

UPDATE apeiron.skill
SET
    base_damage = base_damage * 0.5,
    posture_damage = posture_damage * 0.8,
    updated_at = NOW()
WHERE id IN ('bite', 'lunge', 'maul');

INSERT INTO apeiron.skill_action_timing (
    skill_id, windup_ms, active_ms, recovery_ms, cooldown_ms,
    combo_window_ms, movement_lock_policy, queue_policy, cancel_policy, metadata
)
VALUES
('wolf_dodge',0,420,100,0,0,'contract','none','none','{"full_iframe":true,"source":"reconstructed"}'),
('maul',180,260,360,5200,0,'contract','none','none','{"pressureCounter":true,"source":"reconstructed"}')
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
('wolf_dodge','wolf_dodge_lateral_leap_v1','active','explicit_recovery_handoff','blocked_during_owned_root','evasion_direction','iframe',TRUE,'{"source":"reconstructed"}')
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
    '{"recoverDistanceWithRun":true,"chaseLungeSetup":true,"biteAtCloseRange":true,"preferredMinCm":130,"preferredMaxCm":520}',
    '{"preferLongSideCommit":true,"sideFlipChanceMultiplier":0.35,"minimumSideCommitMs":3200,"flankBeforeLunge":true}',
    '{"dodgeUnderPressure":true,"maulCounterUnderPressure":true,"maulCounterChance":0.22,"dodgeRetreatMultiplier":0.70,"globalDodgeMultiplier":0.85}',
    '{"max":100,"dodgeCostMultiplier":0.50,"runDrainEnabled":true,"zeroStaminaLockoutUntilFull":true}',
    '{"source":"reconstructed"}'
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
    '{"source":"reconstructed"}'
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
