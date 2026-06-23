-- =========================================================
-- TEMPORAL MELEE HITBOX SEEDS
-- =========================================================

INSERT INTO apeiron.skill_hitbox_damage_group (id, skill_id, description, max_hits_per_target, hit_interval_ms, can_multi_hit, metadata)
VALUES
('player_basic_attack_1_damage','player_basic_attack_1','One damage application per opening cut.',1,0,FALSE,'{}'),
('player_basic_attack_2_damage','player_basic_attack_2','One damage application per cross cut.',1,0,FALSE,'{}'),
('player_basic_attack_3_damage','player_basic_attack_3','Shield drive contact push damage group.',1,0,FALSE,'{}'),
('player_shield_bash_front_push','player_shield_bash','Shield Bash can push all targets in the path once.',1,0,FALSE,'{"multi_target_push":true}'),
('player_shield_rush_front_contact','player_shield_rush','Shield Rush can carry/push all targets in the front path once.',1,0,FALSE,'{"multi_target_push":true}'),
('wolf_bite_damage','bite','Bite applies one close-range forward contact hit.',1,0,FALSE,'{"target_direction":true}'),
('wolf_lunge_damage','lunge','Lunge deals damage once while crossing target.',1,0,FALSE,'{"airborne_passthrough":true}')
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
('motion_player_basic_attack_1_forward_v1','player_basic_attack_1','timeline_sweep','hitbox_window_normalized','Shield jab uses a compact temporal contact patch that advances with the shield.','{"shape_intent":"compact_forward_box"}'),
('motion_player_basic_attack_2_right_to_left_v1','player_basic_attack_2','timeline_sweep','hitbox_window_normalized','Left-to-right shield sweep across a 90-degree frontal cone.','{"sweep_direction":"left_to_right","stable_id":"kept for runtime compatibility"}'),
('motion_player_basic_attack_3_shield_drive_v1','player_basic_attack_3','timeline_sweep','hitbox_window_normalized','Overhead shield punch keeps a constant narrow width and grows only in forward depth during the committed drive.','{}'),
('motion_player_shield_bash_front_push_v1','player_shield_bash','timeline_sweep','hitbox_window_normalized','Narrower short front push that only hits when the temporal shield contact reaches the target.','{}'),
('motion_player_shield_rush_front_contact_v1','player_shield_rush','timeline_sweep','hitbox_window_normalized','Short shield-face contact stays attached close to the player and follows the rush before carry starts.','{"front_contact_offset_cm":8,"front_arc_width_cm":224,"front_contact_depth_cm":54}'),
('motion_wolf_bite_melee_v1','bite','timeline_sweep','hitbox_window_normalized','Forward bite contact follows wolf target direction.','{}'),
('motion_wolf_lunge_cross_v1','lunge','timeline_sweep','hitbox_window_normalized','Airborne lunge hit volume travels through target.','{}')
ON CONFLICT (id) DO UPDATE SET
    skill_id = EXCLUDED.skill_id,
    motion_type = EXCLUDED.motion_type,
    time_basis = EXCLUDED.time_basis,
    description = EXCLUDED.description,
    metadata = EXCLUDED.metadata,
    updated_at = NOW();

DELETE FROM apeiron.skill_hitbox_motion_sample
WHERE motion_profile_id IN (
    'motion_player_basic_attack_1_forward_v1',
    'motion_player_basic_attack_2_right_to_left_v1',
    'motion_player_basic_attack_3_shield_drive_v1',
    'motion_player_shield_bash_front_push_v1',
    'motion_player_shield_rush_front_contact_v1',
    'motion_wolf_bite_melee_v1',
    'motion_wolf_lunge_cross_v1'
);

INSERT INTO apeiron.skill_hitbox_motion_sample (
    motion_profile_id, sample_index, t, shape,
    offset_x, offset_y, offset_z, size_x, size_y, size_z,
    radius, length, min_angle_deg, max_angle_deg, metadata
)
VALUES
('motion_player_basic_attack_1_forward_v1',0,0.00,'box_strip',0,0,90,28,52,150,26,28,0,0,'{}'),
('motion_player_basic_attack_1_forward_v1',1,0.50,'box_strip',0,0,90,46,52,150,26,46,0,0,'{}'),
('motion_player_basic_attack_1_forward_v1',2,1.00,'box_strip',0,0,90,64,52,150,26,64,0,0,'{}'),
('motion_player_basic_attack_2_right_to_left_v1',0,0.00,'arc_slice',70,-35,95,0,0,150,50,125,15,45,'{}'),
('motion_player_basic_attack_2_right_to_left_v1',1,0.50,'arc_slice',80,0,95,0,0,150,52,135,-15,15,'{}'),
('motion_player_basic_attack_2_right_to_left_v1',2,1.00,'arc_slice',70,35,95,0,0,150,50,125,-45,-15,'{}'),
('motion_player_basic_attack_3_shield_drive_v1',0,0.00,'capsule_strip',0,0,95,84,0,155,42,42,0,0,'{}'),
('motion_player_basic_attack_3_shield_drive_v1',1,0.55,'capsule_strip',0,0,95,84,0,155,42,140,0,0,'{}'),
('motion_player_basic_attack_3_shield_drive_v1',2,1.00,'capsule_strip',0,0,95,84,0,155,42,252,0,0,'{}'),
('motion_player_shield_bash_front_push_v1',0,0.00,'capsule_strip',45,0,95,132,0,160,66,75,0,0,'{}'),
('motion_player_shield_bash_front_push_v1',1,0.50,'capsule_strip',72,0,95,132,0,160,66,120,0,0,'{}'),
('motion_player_shield_bash_front_push_v1',2,1.00,'capsule_strip',92,0,95,132,0,160,66,160,0,0,'{}'),
('motion_player_shield_rush_front_contact_v1',0,0.00,'box_strip',8,0,100,34,224,160,112,34,0,0,'{"contact_face":"shield_front","contact_depth_cm":34}'),
('motion_player_shield_rush_front_contact_v1',1,0.50,'box_strip',10,0,100,44,224,160,112,44,0,0,'{"contact_face":"shield_front","contact_depth_cm":44}'),
('motion_player_shield_rush_front_contact_v1',2,1.00,'box_strip',12,0,100,54,224,160,112,54,0,0,'{"contact_face":"shield_front","contact_depth_cm":54}'),
('motion_wolf_bite_melee_v1',0,0.00,'capsule_strip',45,0,85,90,0,115,45,70,0,0,'{}'),
('motion_wolf_bite_melee_v1',1,0.55,'capsule_strip',80,0,90,95,0,115,48,125,0,0,'{}'),
('motion_wolf_bite_melee_v1',2,1.00,'capsule_strip',95,0,85,90,0,115,45,145,0,0,'{}'),
('motion_wolf_lunge_cross_v1',0,0.00,'capsule_strip',60,0,90,100,0,120,50,100,0,0,'{}'),
('motion_wolf_lunge_cross_v1',1,0.55,'capsule_strip',140,0,110,100,0,120,50,230,0,0,'{}'),
('motion_wolf_lunge_cross_v1',2,1.00,'capsule_strip',210,0,90,100,0,120,50,320,0,0,'{}');

INSERT INTO apeiron.skill_hitbox_profile (
    id, skill_id, hitbox_index, hitbox_shape, hitbox_start_ms, hitbox_end_ms,
    offset_x, offset_y, offset_z, size_x, size_y, size_z, radius, length, angle,
    follows_caster, follows_projectile, can_multi_hit, max_hits_per_target,
    hit_interval_ms, friendly_fire, motion_profile_id, damage_group_id,
    min_angle_deg, max_angle_deg, start_radius, end_radius
)
VALUES
('hitbox_player_basic_attack_1_0','player_basic_attack_1',0,'temporal_sweep',90,230,0,0,90,64,52,150,26,64,0,TRUE,FALSE,FALSE,1,0,FALSE,'motion_player_basic_attack_1_forward_v1','player_basic_attack_1_damage',0,0,26,26),
('hitbox_player_basic_attack_2_0','player_basic_attack_2',0,'temporal_sweep',100,250,80,0,95,0,104,150,52,135,90,TRUE,FALSE,FALSE,1,0,FALSE,'motion_player_basic_attack_2_right_to_left_v1','player_basic_attack_2_damage',-45,45,50,52),
('hitbox_player_basic_attack_3_0','player_basic_attack_3',0,'temporal_sweep',180,440,0,0,95,84,0,155,42,252,0,TRUE,FALSE,FALSE,1,0,FALSE,'motion_player_basic_attack_3_shield_drive_v1','player_basic_attack_3_damage',0,0,42,42),
('hitbox_player_shield_bash_0','player_shield_bash',0,'temporal_sweep',110,280,90,0,95,132,0,160,66,160,0,TRUE,FALSE,FALSE,1,0,FALSE,'motion_player_shield_bash_front_push_v1','player_shield_bash_front_push',0,0,66,66),
('hitbox_player_shield_rush_0','player_shield_rush',0,'temporal_sweep',160,880,8,0,100,54,224,160,112,54,0,TRUE,FALSE,FALSE,1,0,FALSE,'motion_player_shield_rush_front_contact_v1','player_shield_rush_front_contact',0,0,112,112),
('hitbox_bite_0','bite',0,'temporal_sweep',120,340,80,0,90,95,0,115,48,145,0,TRUE,FALSE,FALSE,1,0,FALSE,'motion_wolf_bite_melee_v1','wolf_bite_damage',0,0,45,48),
('hitbox_lunge_0','lunge',0,'temporal_sweep',3600,3980,130,0,105,100,0,120,50,320,0,TRUE,FALSE,FALSE,1,0,FALSE,'motion_wolf_lunge_cross_v1','wolf_lunge_damage',0,0,50,50)
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
