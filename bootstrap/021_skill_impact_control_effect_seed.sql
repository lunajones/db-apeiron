-- =========================================================
-- SKILL IMPACT CONTROL EFFECT SEED
-- =========================================================
--
-- Control effects are gameplay contracts consumed by the combat impact
-- pipeline. They are intentionally separate from movement_action_contract:
-- movement contracts own root/position, impact control owns what happens to
-- the target when a damaging contact lands.

INSERT INTO apeiron.status_effect (
    id,
    name,
    description,
    effect_type,
    stacking_mode,
    max_stacks,
    duration_ms,
    tick_interval_ms,
    is_dispellable,
    is_pvp_enabled,
    movement_modifier,
    damage_dealt_modifier,
    damage_taken_modifier,
    healing_received_modifier,
    stamina_regen_modifier,
    blocks_movement,
    blocks_actions,
    blocks_skills
)
VALUES
('impact_shield_drive_push','Shield Drive Push','Short shield-drive contact displacement/control.', 'crowd_control','refresh',1,180,0,FALSE,TRUE,1,1,1,1,1,TRUE,FALSE,FALSE),
('impact_shield_bash_push','Shield Bash Stun','Bulwark frontal shield bash push and short stun.', 'crowd_control','refresh',1,1500,0,FALSE,TRUE,1,1,1,1,1,TRUE,TRUE,TRUE),
('impact_shield_rush_carry_push','Shield Rush Carry Push','Committed rush contact carry and push control.', 'crowd_control','refresh',1,720,0,FALSE,TRUE,1,1,1,1,1,TRUE,TRUE,TRUE),
('impact_wolf_maul_lateral_grab','Wolf Maul Lateral Grab','Maul catches the target in the lateral counter path and staggers movement/actions until release.', 'crowd_control','refresh',1,520,0,FALSE,TRUE,1,1,1,1,1,TRUE,TRUE,TRUE)
ON CONFLICT (id) DO UPDATE SET
    name = EXCLUDED.name,
    description = EXCLUDED.description,
    effect_type = EXCLUDED.effect_type,
    stacking_mode = EXCLUDED.stacking_mode,
    max_stacks = EXCLUDED.max_stacks,
    duration_ms = EXCLUDED.duration_ms,
    tick_interval_ms = EXCLUDED.tick_interval_ms,
    is_dispellable = EXCLUDED.is_dispellable,
    is_pvp_enabled = EXCLUDED.is_pvp_enabled,
    movement_modifier = EXCLUDED.movement_modifier,
    damage_dealt_modifier = EXCLUDED.damage_dealt_modifier,
    damage_taken_modifier = EXCLUDED.damage_taken_modifier,
    healing_received_modifier = EXCLUDED.healing_received_modifier,
    stamina_regen_modifier = EXCLUDED.stamina_regen_modifier,
    blocks_movement = EXCLUDED.blocks_movement,
    blocks_actions = EXCLUDED.blocks_actions,
    blocks_skills = EXCLUDED.blocks_skills,
    updated_at = NOW();

INSERT INTO apeiron.skill_impact_profile (
    skill_id,
    impact_type,
    hit_reaction,
    poise_damage,
    stagger_power,
    interrupt_power,
    guard_damage_multiplier,
    bounce_on_shield,
    destroy_on_hit,
    stick_on_hit,
    hitstop_ms,
    screenshake_strength,
    knockback_force,
    knockback_upward_force,
    pull_force,
    applies_status_effect,
    status_effect_id,
    status_effect_chance,
    control_type,
    control_effect_duration_ms,
    control_release_policy_id,
    control_distance_cm,
    control_speed_cm_s,
    control_direction_policy
)
VALUES
(
    'player_basic_attack_3',
    'normal',
    'stagger',
    COALESCE((SELECT posture_damage FROM apeiron.skill WHERE id = 'player_basic_attack_3'), 0.0),
    0.25,
    0.20,
    1.0,
    FALSE,
    TRUE,
    FALSE,
    0,
    0.70,
    0.75,
    0.0,
    0.0,
    TRUE,
    'impact_shield_drive_push',
    1.0,
    'push',
    180,
    'carry_contact_forward_release',
    COALESCE((SELECT movement_distance * 100.0 FROM apeiron.skill WHERE id = 'player_basic_attack_3'), 0.0),
    CASE
        WHEN 180 > 0 THEN COALESCE((SELECT movement_distance * 100.0 FROM apeiron.skill WHERE id = 'player_basic_attack_3'), 0.0) / (180.0 / 1000.0)
        ELSE 0.0
    END,
    'source_forward'
),
(
    'player_shield_bash',
    'normal',
    'stagger',
    COALESCE((SELECT posture_damage FROM apeiron.skill WHERE id = 'player_shield_bash'), 0.0),
    0.70,
    0.80,
    1.0,
    FALSE,
    TRUE,
    FALSE,
    0,
    0.45,
    0.55,
    0.0,
    0.0,
    TRUE,
    'impact_shield_bash_push',
    1.0,
    'push',
    170,
    'multi_target_push_forward_release',
    COALESCE((SELECT movement_distance * 100.0 FROM apeiron.skill WHERE id = 'player_shield_bash'), 0.0),
    CASE
        WHEN 170 > 0 THEN COALESCE((SELECT movement_distance * 100.0 FROM apeiron.skill WHERE id = 'player_shield_bash'), 0.0) / (170.0 / 1000.0)
        ELSE 0.0
    END,
    'source_forward'
),
(
    'player_shield_rush',
    'normal',
    'stagger',
    COALESCE((SELECT posture_damage FROM apeiron.skill WHERE id = 'player_shield_rush'), 0.0),
    0.90,
    1.00,
    1.0,
    FALSE,
    TRUE,
    FALSE,
    0,
    0.0,
    0.0,
    0.0,
    0.0,
    TRUE,
    'impact_shield_rush_carry_push',
    1.0,
    'carry_push',
    720,
    'multi_target_carry_push_forward_release',
    COALESCE((SELECT movement_distance * 100.0 FROM apeiron.skill WHERE id = 'player_shield_rush'), 0.0),
    CASE
        WHEN 720 > 0 THEN COALESCE((SELECT movement_distance * 100.0 FROM apeiron.skill WHERE id = 'player_shield_rush'), 0.0) / (720.0 / 1000.0)
        ELSE 0.0
    END,
    'source_forward'
),
(
    'maul',
    'normal',
    'stagger',
    COALESCE((SELECT posture_damage FROM apeiron.skill WHERE id = 'maul'), 0.0),
    0.85,
    0.95,
    1.0,
    FALSE,
    TRUE,
    FALSE,
    0,
    0.35,
    0.0,
    0.0,
    0.0,
    TRUE,
    'impact_wolf_maul_lateral_grab',
    1.0,
    'grab',
    COALESCE((SELECT active_ms FROM apeiron.movement_action_contract WHERE id = 'wolf_maul_lateral_counter_v1'), 0),
    'lateral_grab_release',
    COALESCE((SELECT distance_cm FROM apeiron.movement_action_contract WHERE id = 'wolf_maul_lateral_counter_v1'), 0.0),
    COALESCE((SELECT base_speed_cm_s FROM apeiron.movement_action_contract WHERE id = 'wolf_maul_lateral_counter_v1'), 0.0),
    'source_action_direction'
)
ON CONFLICT (skill_id) DO UPDATE SET
    impact_type = EXCLUDED.impact_type,
    hit_reaction = EXCLUDED.hit_reaction,
    poise_damage = EXCLUDED.poise_damage,
    stagger_power = EXCLUDED.stagger_power,
    interrupt_power = EXCLUDED.interrupt_power,
    guard_damage_multiplier = EXCLUDED.guard_damage_multiplier,
    bounce_on_shield = EXCLUDED.bounce_on_shield,
    destroy_on_hit = EXCLUDED.destroy_on_hit,
    stick_on_hit = EXCLUDED.stick_on_hit,
    hitstop_ms = EXCLUDED.hitstop_ms,
    screenshake_strength = EXCLUDED.screenshake_strength,
    knockback_force = EXCLUDED.knockback_force,
    knockback_upward_force = EXCLUDED.knockback_upward_force,
    pull_force = EXCLUDED.pull_force,
    applies_status_effect = EXCLUDED.applies_status_effect,
    status_effect_id = EXCLUDED.status_effect_id,
    status_effect_chance = EXCLUDED.status_effect_chance,
    control_type = EXCLUDED.control_type,
    control_effect_duration_ms = EXCLUDED.control_effect_duration_ms,
    control_release_policy_id = EXCLUDED.control_release_policy_id,
    control_distance_cm = EXCLUDED.control_distance_cm,
    control_speed_cm_s = EXCLUDED.control_speed_cm_s,
    control_direction_policy = EXCLUDED.control_direction_policy,
    updated_at = NOW();

DO $$
BEGIN
    IF (
        SELECT COUNT(*)
        FROM apeiron.skill_impact_profile
        WHERE skill_id IN ('player_basic_attack_3', 'player_shield_bash', 'player_shield_rush', 'maul')
            AND status_effect_id IN (
                'impact_shield_drive_push',
                'impact_shield_bash_push',
                'impact_shield_rush_carry_push',
                'impact_wolf_maul_lateral_grab'
            )
            AND applies_status_effect IS TRUE
            AND status_effect_chance > 0
            AND COALESCE(control_type, '') <> ''
            AND control_effect_duration_ms > 0
            AND COALESCE(control_release_policy_id, '') <> ''
            AND control_distance_cm > 0
            AND control_speed_cm_s > 0
            AND COALESCE(control_direction_policy, '') <> ''
    ) <> 4 THEN
        RAISE EXCEPTION 'Apeiron bootstrap produced incomplete skill impact control motion contract';
    END IF;

    IF EXISTS (
        SELECT 1
        FROM apeiron.skill_impact_profile
        WHERE skill_id IN ('player_basic_attack_3', 'player_shield_bash', 'player_shield_rush', 'maul')
        AND (
            (skill_id = 'player_basic_attack_3' AND status_effect_id <> 'impact_shield_drive_push')
            OR (skill_id = 'player_shield_bash' AND status_effect_id <> 'impact_shield_bash_push')
            OR (skill_id = 'player_shield_rush' AND status_effect_id <> 'impact_shield_rush_carry_push')
            OR (skill_id = 'maul' AND status_effect_id <> 'impact_wolf_maul_lateral_grab')
        )
    ) THEN
        RAISE EXCEPTION 'Apeiron bootstrap produced mismatched skill impact control contract';
    END IF;
END $$;
