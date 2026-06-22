-- =========================================================
-- LEGACY SKILL MOVEMENT EFFECT SEED
-- Kept for compatibility with GetSkillMovementEffect(skill_id).
--
-- Recovered databases may already have skill_movement_effect rows referenced
-- by skill_slot through older IDs. Preserve existing IDs by updating by
-- skill_id first; only insert the reconstructed ID when no row exists.
-- =========================================================

WITH desired (
    id, skill_id, movement_type, distance, speed, duration_ms,
    windup_lock_ms, recovery_lock_ms, can_rotate, ignores_collision, metadata
) AS (
    VALUES
    ('leap_default','lunge','leap',620::DOUBLE PRECISION,760::DOUBLE PRECISION,980,3600,500,TRUE,FALSE,'{"source":"chat_recovery","thread":"DB","compatibility":"GetSkillMovementEffect(lunge)","prefer":"movement_action_contract","movement_action_contract_id":"wolf_lunge_airborne_v1","post_landing_inertia_multiplier":1.1,"airborne_passthrough":true}'::jsonb),
    ('player_shield_rush_legacy','player_shield_rush','charge',340::DOUBLE PRECISION,470::DOUBLE PRECISION,830,160,240,TRUE,FALSE,'{"source":"chat_recovery","prefer":"movement_action_contract"}'::jsonb),
    ('player_shield_bash_legacy','player_shield_bash','dash',130::DOUBLE PRECISION,280::DOUBLE PRECISION,520,120,180,TRUE,FALSE,'{"source":"chat_recovery","prefer":"movement_action_contract","legacy_semantics":"short_charge"}'::jsonb)
)
UPDATE apeiron.skill_movement_effect existing
SET
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
    ('leap_default','lunge','leap',620::DOUBLE PRECISION,760::DOUBLE PRECISION,980,3600,500,TRUE,FALSE,'{"source":"chat_recovery","thread":"DB","compatibility":"GetSkillMovementEffect(lunge)","prefer":"movement_action_contract","movement_action_contract_id":"wolf_lunge_airborne_v1","post_landing_inertia_multiplier":1.1,"airborne_passthrough":true}'::jsonb),
    ('player_shield_rush_legacy','player_shield_rush','charge',340::DOUBLE PRECISION,470::DOUBLE PRECISION,830,160,240,TRUE,FALSE,'{"source":"chat_recovery","prefer":"movement_action_contract"}'::jsonb),
    ('player_shield_bash_legacy','player_shield_bash','dash',130::DOUBLE PRECISION,280::DOUBLE PRECISION,520,120,180,TRUE,FALSE,'{"source":"chat_recovery","prefer":"movement_action_contract","legacy_semantics":"short_charge"}'::jsonb)
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
