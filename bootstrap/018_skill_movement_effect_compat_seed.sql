-- =========================================================
-- SKILL MOVEMENT EFFECT COMPATIBILITY SEED
-- Kept for compatibility with GetSkillMovementEffect(skill_id).
--
-- Older databases may already have skill_movement_effect rows keyed by
-- generic IDs such as leap_default. No current table references those IDs, so
-- normalize them to the canonical compatibility IDs while updating by skill_id.
-- =========================================================

WITH desired (
    id, skill_id, movement_type, distance, speed, duration_ms,
    windup_lock_ms, recovery_lock_ms, can_rotate, ignores_collision, metadata
) AS (
    VALUES
    ('low_fast_lunge_effect_v1','lunge','leap',2148::DOUBLE PRECISION,2619.5::DOUBLE PRECISION,820,3600,200,TRUE,FALSE,'{"source":"canonical_bootstrap","thread":"DB","surface":"GetSkillMovementEffect(lunge)","prefer":"movement_action_contract","movement_action_contract_id":"low_fast_lunge_v1","orientation_policy_id":"orientation_lunge_flank_commit_v1","envelope_policy_id":"envelope_lunge_low_raking_100_520_200_v1","pre_commit_ms":100,"post_landing_inertia_multiplier":0.62,"airborne_duration_ms":520,"landing_inertia_ms":200,"airborne_passthrough":true,"vertical_arc":"very_low_long_raking"}'::jsonb),
    ('wolf_dodge_lateral_leap_effect_v1','wolf_dodge','dodge',210::DOUBLE PRECISION,520::DOUBLE PRECISION,520,0,100,TRUE,TRUE,'{"source":"canonical_bootstrap","surface":"GetSkillMovementEffect(wolf_dodge)","prefer":"movement_action_contract","movement_action_contract_id":"wolf_dodge_lateral_leap_v1","full_iframe":true}'::jsonb),
    ('wolf_maul_lateral_counter_effect_v1','maul','grounded_skill',420::DOUBLE PRECISION,690::DOUBLE PRECISION,920,220,220,TRUE,FALSE,'{"source":"canonical_bootstrap","surface":"GetSkillMovementEffect(maul)","prefer":"movement_action_contract","movement_action_contract_id":"wolf_maul_lateral_counter_v1","contact_policy":"lateral_counter_contact","control_effect":"impact_wolf_maul_lateral_grab"}'::jsonb),
    ('player_shield_rush_effect_v1','player_shield_rush','charge',864::DOUBLE PRECISION,1033.2::DOUBLE PRECISION,1100,160,260,TRUE,FALSE,'{"source":"canonical_bootstrap","prefer":"movement_action_contract","front_contact_offset_cm":8,"design_note":"matches shield_rush_front_contact_v1"}'::jsonb),
    ('player_shield_bash_effect_v1','player_shield_bash','dash',95::DOUBLE PRECISION,541::DOUBLE PRECISION,300,110,120,TRUE,FALSE,'{"source":"canonical_bootstrap","prefer":"movement_action_contract","compat_semantics":"short_charge"}'::jsonb)
)
UPDATE apeiron.skill_movement_effect existing
SET
    id = desired.id,
    movement_type = desired.movement_type,
    distance = desired.distance,
    speed = desired.speed,
    duration_ms = desired.duration_ms,
    windup_lock_ms = desired.windup_lock_ms,
    recovery_lock_ms = desired.recovery_lock_ms,
    can_rotate = desired.can_rotate,
    ignores_collision = desired.ignores_collision,
    metadata = desired.metadata,
    updated_at = NOW()
FROM desired
WHERE existing.skill_id = desired.skill_id;

WITH desired (
    id, skill_id, movement_type, distance, speed, duration_ms,
    windup_lock_ms, recovery_lock_ms, can_rotate, ignores_collision, metadata
) AS (
    VALUES
    ('low_fast_lunge_effect_v1','lunge','leap',2148::DOUBLE PRECISION,2619.5::DOUBLE PRECISION,820,3600,200,TRUE,FALSE,'{"source":"canonical_bootstrap","thread":"DB","surface":"GetSkillMovementEffect(lunge)","prefer":"movement_action_contract","movement_action_contract_id":"low_fast_lunge_v1","orientation_policy_id":"orientation_lunge_flank_commit_v1","envelope_policy_id":"envelope_lunge_low_raking_100_520_200_v1","pre_commit_ms":100,"post_landing_inertia_multiplier":0.62,"airborne_duration_ms":520,"landing_inertia_ms":200,"airborne_passthrough":true,"vertical_arc":"very_low_long_raking"}'::jsonb),
    ('wolf_dodge_lateral_leap_effect_v1','wolf_dodge','dodge',210::DOUBLE PRECISION,520::DOUBLE PRECISION,520,0,100,TRUE,TRUE,'{"source":"canonical_bootstrap","surface":"GetSkillMovementEffect(wolf_dodge)","prefer":"movement_action_contract","movement_action_contract_id":"wolf_dodge_lateral_leap_v1","full_iframe":true}'::jsonb),
    ('wolf_maul_lateral_counter_effect_v1','maul','grounded_skill',420::DOUBLE PRECISION,690::DOUBLE PRECISION,920,220,220,TRUE,FALSE,'{"source":"canonical_bootstrap","surface":"GetSkillMovementEffect(maul)","prefer":"movement_action_contract","movement_action_contract_id":"wolf_maul_lateral_counter_v1","contact_policy":"lateral_counter_contact","control_effect":"impact_wolf_maul_lateral_grab"}'::jsonb),
    ('player_shield_rush_effect_v1','player_shield_rush','charge',864::DOUBLE PRECISION,1033.2::DOUBLE PRECISION,1100,160,260,TRUE,FALSE,'{"source":"canonical_bootstrap","prefer":"movement_action_contract","front_contact_offset_cm":8,"design_note":"matches shield_rush_front_contact_v1"}'::jsonb),
    ('player_shield_bash_effect_v1','player_shield_bash','dash',95::DOUBLE PRECISION,541::DOUBLE PRECISION,300,110,120,TRUE,FALSE,'{"source":"canonical_bootstrap","prefer":"movement_action_contract","compat_semantics":"short_charge"}'::jsonb)
)
INSERT INTO apeiron.skill_movement_effect (
    id, skill_id, movement_type, distance, speed, duration_ms,
    windup_lock_ms, recovery_lock_ms, can_rotate, ignores_collision, metadata
)
SELECT
    desired.id,
    desired.skill_id,
    desired.movement_type,
    desired.distance,
    desired.speed,
    desired.duration_ms,
    desired.windup_lock_ms,
    desired.recovery_lock_ms,
    desired.can_rotate,
    desired.ignores_collision,
    desired.metadata
FROM desired
WHERE NOT EXISTS (
    SELECT 1
    FROM apeiron.skill_movement_effect existing
    WHERE existing.skill_id = desired.skill_id
)
ON CONFLICT (id) DO UPDATE SET
    skill_id = EXCLUDED.skill_id,
    movement_type = EXCLUDED.movement_type,
    distance = EXCLUDED.distance,
    speed = EXCLUDED.speed,
    duration_ms = EXCLUDED.duration_ms,
    windup_lock_ms = EXCLUDED.windup_lock_ms,
    recovery_lock_ms = EXCLUDED.recovery_lock_ms,
    can_rotate = EXCLUDED.can_rotate,
    ignores_collision = EXCLUDED.ignores_collision,
    metadata = EXCLUDED.metadata,
    updated_at = NOW();
