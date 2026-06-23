-- =========================================================
-- PLAYER SWORD AND SHIELD SKILLS
-- Reconstructed from Apeiron chat/runtime recovery.
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
('player_basic_attack_1','Opening Cut','Sword-and-shield light combo opener. Forward temporal cut from body to one and a half player cylinders.', 'sword_shield','basic_attack',
 0,0,0, 90,140,120,0,0, 0,0,0,0, 0,1.8,0,1,'enemy',FALSE, 8,'physical',NULL,10,0,1,1.2, 0,0,0.35, 1,FALSE,0, 'sword_shield_light_combo',1,2000, TRUE,TRUE,TRUE,FALSE,FALSE),
('player_basic_attack_2','Cross Cut','Right-to-left sword cut using a temporal 90-degree sweep.', 'sword_shield','basic_attack',
 0,0,0, 100,150,120,0,0, 0,0,0,0, 0,1.8,90,2,'enemy',FALSE, 7,'physical',NULL,9,0,1,1.15, 0,0,0.25, 1,FALSE,0, 'sword_shield_light_combo',2,2000, TRUE,TRUE,TRUE,FALSE,FALSE),
('player_basic_attack_3','Shield Drive','Short shield drive finisher that carries contact and pushes enemies.', 'sword_shield','basic_attack',
 0,0,0, 180,260,180,0,0, 0,0,0,0, 0,2.2,0,3,'enemy',FALSE, 6,'physical',NULL,18,0,1,1.1, 0,0,2.1, 1,FALSE,2.0, 'sword_shield_light_combo',3,2000, TRUE,TRUE,TRUE,FALSE,FALSE),
('player_shield_bash','Shield Bash','Bulwark close shield strike: short committed step, frontal temporal hit, push and stun.', 'sword_shield','shield_attack',
 18,0,0, 110,170,120,0,2600, 0,0,0,0, 0,1.6,0,4,'enemy',FALSE, 10,'physical',NULL,26,0,1,1.25, 1500,0,1.1, 1,FALSE,0.95, NULL,NULL,0, TRUE,TRUE,TRUE,FALSE,FALSE),
('player_shield_rush','Shield Rush','Bulwark committed forward rush. Damage begins close to body contact and follows the player front.', 'sword_shield','shield_rush',
 26,0,0, 160,720,260,0,5200, 0,0,0,0, 0,10.2,0,5,'enemy',FALSE, 14,'physical',NULL,34,0,1,1.35, 0,0,5.0, 1,FALSE,9.6, NULL,NULL,0, TRUE,TRUE,TRUE,FALSE,FALSE)
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
    min_range = EXCLUDED.min_range,
    max_range = EXCLUDED.max_range,
    cone_angle = EXCLUDED.cone_angle,
    max_targets = EXCLUDED.max_targets,
    requires_target = EXCLUDED.requires_target,
    base_damage = EXCLUDED.base_damage,
    posture_damage = EXCLUDED.posture_damage,
    knockback_force = EXCLUDED.knockback_force,
    movement_multiplier = EXCLUDED.movement_multiplier,
    locks_movement = EXCLUDED.locks_movement,
    movement_distance = EXCLUDED.movement_distance,
    combo_group = EXCLUDED.combo_group,
    combo_index = EXCLUDED.combo_index,
    combo_window_ms = EXCLUDED.combo_window_ms,
    updated_at = NOW();

INSERT INTO apeiron.skill_set (id, name, description, is_player_usable, is_npc_usable)
VALUES ('skillset_player_sword_shield', 'Player Sword and Shield', 'Player sword-and-shield selectable combat kit.', TRUE, FALSE)
ON CONFLICT (id) DO UPDATE SET
    name = EXCLUDED.name,
    description = EXCLUDED.description,
    is_player_usable = EXCLUDED.is_player_usable,
    is_npc_usable = EXCLUDED.is_npc_usable,
    updated_at = NOW();
