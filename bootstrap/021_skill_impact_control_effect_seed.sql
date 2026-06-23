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
('impact_shield_bash_push','Shield Bash Push','Bulwark frontal shield bash push/control.', 'crowd_control','refresh',1,220,0,FALSE,TRUE,1,1,1,1,1,TRUE,FALSE,FALSE),
('impact_shield_rush_carry_push','Shield Rush Carry Push','Committed rush contact carry and push control.', 'crowd_control','refresh',1,430,0,FALSE,TRUE,1,1,1,1,1,TRUE,TRUE,TRUE)
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

UPDATE apeiron.skill_impact_profile
SET
    applies_status_effect = TRUE,
    status_effect_id = 'impact_shield_drive_push',
    status_effect_chance = 1.0,
    control_type = 'push',
    control_effect_duration_ms = 180,
    control_release_policy_id = 'carry_contact_forward_release',
    control_distance_cm = COALESCE((SELECT movement_distance * 100.0 FROM apeiron.skill WHERE id = 'player_basic_attack_3'), 0.0),
    control_speed_cm_s = CASE
        WHEN 180 > 0 THEN COALESCE((SELECT movement_distance * 100.0 FROM apeiron.skill WHERE id = 'player_basic_attack_3'), 0.0) / (180.0 / 1000.0)
        ELSE 0.0
    END,
    control_direction_policy = 'source_forward'
WHERE skill_id = 'player_basic_attack_3';

UPDATE apeiron.skill_impact_profile
SET
    applies_status_effect = TRUE,
    status_effect_id = 'impact_shield_bash_push',
    status_effect_chance = 1.0,
    control_type = 'push',
    control_effect_duration_ms = 220,
    control_release_policy_id = 'multi_target_push_forward_release',
    control_distance_cm = COALESCE((SELECT movement_distance * 100.0 FROM apeiron.skill WHERE id = 'player_shield_bash'), 0.0),
    control_speed_cm_s = CASE
        WHEN 220 > 0 THEN COALESCE((SELECT movement_distance * 100.0 FROM apeiron.skill WHERE id = 'player_shield_bash'), 0.0) / (220.0 / 1000.0)
        ELSE 0.0
    END,
    control_direction_policy = 'source_forward'
WHERE skill_id = 'player_shield_bash';

UPDATE apeiron.skill_impact_profile
SET
    applies_status_effect = TRUE,
    status_effect_id = 'impact_shield_rush_carry_push',
    status_effect_chance = 1.0,
    control_type = 'carry_push',
    control_effect_duration_ms = 430,
    control_release_policy_id = 'multi_target_carry_push_forward_release',
    control_distance_cm = COALESCE((SELECT movement_distance * 100.0 FROM apeiron.skill WHERE id = 'player_shield_rush'), 0.0),
    control_speed_cm_s = CASE
        WHEN 430 > 0 THEN COALESCE((SELECT movement_distance * 100.0 FROM apeiron.skill WHERE id = 'player_shield_rush'), 0.0) / (430.0 / 1000.0)
        ELSE 0.0
    END,
    control_direction_policy = 'source_forward'
WHERE skill_id = 'player_shield_rush';

DO $$
BEGIN
    IF EXISTS (
        SELECT 1
        FROM apeiron.skill_impact_profile
        WHERE status_effect_id IN (
            'impact_shield_drive_push',
            'impact_shield_bash_push',
            'impact_shield_rush_carry_push'
        )
        AND (
            applies_status_effect IS DISTINCT FROM TRUE
            OR status_effect_chance <= 0
            OR COALESCE(control_type, '') = ''
            OR control_effect_duration_ms <= 0
            OR COALESCE(control_release_policy_id, '') = ''
            OR control_distance_cm <= 0
            OR control_speed_cm_s <= 0
            OR COALESCE(control_direction_policy, '') = ''
        )
    ) THEN
        RAISE EXCEPTION 'Apeiron bootstrap produced incomplete skill impact control motion contract';
    END IF;
END $$;
