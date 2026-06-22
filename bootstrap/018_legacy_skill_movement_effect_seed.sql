-- =========================================================
-- LEGACY SKILL MOVEMENT EFFECT SEED
-- Kept for compatibility with GetSkillMovementEffect(skill_id).
-- =========================================================

INSERT INTO apeiron.skill_movement_effect (
    id, skill_id, movement_type, distance, speed, duration_ms,
    windup_lock_ms, recovery_lock_ms, can_rotate, ignores_collision, metadata
)
VALUES
('leap_default','lunge','leap',620,760,980,3600,500,TRUE,FALSE,'{"source":"chat_recovery","thread":"DB","compatibility":"GetSkillMovementEffect(lunge)","prefer":"movement_action_contract","movement_action_contract_id":"wolf_lunge_airborne_v1","post_landing_inertia_multiplier":1.1,"airborne_passthrough":true}'),
('player_shield_rush_legacy','player_shield_rush','charge',340,470,830,160,240,TRUE,FALSE,'{"source":"chat_recovery","prefer":"movement_action_contract"}'),
('player_shield_bash_legacy','player_shield_bash','short_charge',130,280,520,120,180,TRUE,FALSE,'{"source":"chat_recovery","prefer":"movement_action_contract"}')
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
