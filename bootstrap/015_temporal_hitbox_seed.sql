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
('motion_player_basic_attack_1_forward_v1','player_basic_attack_1','timeline_sweep','hitbox_window_normalized','Forward cut grows from body to one and a half player cylinders.','{}'),
('motion_player_basic_attack_2_right_to_left_v1','player_basic_attack_2','timeline_sweep','hitbox_window_normalized','Right-to-left sword cut across 90 degrees.','{}'),
('motion_player_basic_attack_3_shield_drive_v1','player_basic_attack_3','timeline_sweep','hitbox_window_normalized','Shield drive front strip follows player movement.','{}'),
('motion_player_shield_bash_front_push_v1','player_shield_bash','timeline_sweep','hitbox_window_normalized','Wide front path push, two-cylinder width.','{}'),
('motion_player_shield_rush_front_contact_v1','player_shield_rush','timeline_sweep','hitbox_window_normalized','Front contact starts close to body and follows the rush.','{"front_contact_offset_cm":45}'),
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
('motion_player_basic_attack_1_forward_v1',0,0.00,'capsule_strip',0,35,90,0,95,150,48,70,0,0,'{}'),
('motion_player_basic_attack_1_forward_v1',1,0.50,'capsule_strip',0,85,90,0,100,150,50,150,0,0,'{}'),
('motion_player_basic_attack_1_forward_v1',2,1.00,'capsule_strip',0,130,90,0,100,150,50,210,0,0,'{}'),
('motion_player_basic_attack_2_right_to_left_v1',0,0.00,'asymmetric_arc',35,70,95,0,0,150,55,155,-45,-15,'{}'),
('motion_player_basic_attack_2_right_to_left_v1',1,0.50,'asymmetric_arc',0,85,95,0,0,150,58,165,-15,15,'{}'),
('motion_player_basic_attack_2_right_to_left_v1',2,1.00,'asymmetric_arc',-35,70,95,0,0,150,55,155,15,45,'{}'),
('motion_player_basic_attack_3_shield_drive_v1',0,0.00,'capsule_strip',0,45,95,0,120,155,60,90,0,0,'{}'),
('motion_player_basic_attack_3_shield_drive_v1',1,0.55,'capsule_strip',0,90,95,0,120,155,60,175,0,0,'{}'),
('motion_player_basic_attack_3_shield_drive_v1',2,1.00,'capsule_strip',0,115,95,0,120,155,60,210,0,0,'{}'),
('motion_player_shield_bash_front_push_v1',0,0.00,'capsule_strip',0,45,95,0,190,160,95,95,0,0,'{}'),
('motion_player_shield_bash_front_push_v1',1,0.50,'capsule_strip',0,85,95,0,190,160,95,160,0,0,'{}'),
('motion_player_shield_bash_front_push_v1',2,1.00,'capsule_strip',0,120,95,0,190,160,95,210,0,0,'{}'),
('motion_player_shield_rush_front_contact_v1',0,0.00,'capsule_strip',0,45,100,0,190,160,96,105,0,0,'{}'),
('motion_player_shield_rush_front_contact_v1',1,0.50,'capsule_strip',0,105,100,0,190,160,96,220,0,0,'{}'),
('motion_player_shield_rush_front_contact_v1',2,1.00,'capsule_strip',0,145,100,0,190,160,96,290,0,0,'{}'),
('motion_wolf_bite_melee_v1',0,0.00,'capsule_strip',0,45,85,0,90,115,45,70,0,0,'{}'),
('motion_wolf_bite_melee_v1',1,0.55,'capsule_strip',0,80,90,0,95,115,48,125,0,0,'{}'),
('motion_wolf_bite_melee_v1',2,1.00,'capsule_strip',0,95,85,0,90,115,45,145,0,0,'{}'),
('motion_wolf_lunge_cross_v1',0,0.00,'capsule_strip',0,60,90,0,100,120,50,100,0,0,'{}'),
('motion_wolf_lunge_cross_v1',1,0.55,'capsule_strip',0,140,110,0,100,120,50,230,0,0,'{}'),
('motion_wolf_lunge_cross_v1',2,1.00,'capsule_strip',0,210,90,0,100,120,50,320,0,0,'{}');

INSERT INTO apeiron.skill_hitbox_profile (
    id, skill_id, hitbox_index, hitbox_shape, hitbox_start_ms, hitbox_end_ms,
    offset_x, offset_y, offset_z, size_x, size_y, size_z, radius, length, angle,
    follows_caster, follows_projectile, can_multi_hit, max_hits_per_target,
    hit_interval_ms, friendly_fire, motion_profile_id, damage_group_id,
    min_angle_deg, max_angle_deg, start_radius, end_radius
)
VALUES
('hitbox_player_basic_attack_1_0','player_basic_attack_1',0,'temporal_sweep',90,230,0,90,90,0,100,150,50,210,0,TRUE,FALSE,FALSE,1,0,FALSE,'motion_player_basic_attack_1_forward_v1','player_basic_attack_1_damage',0,0,48,50),
('hitbox_player_basic_attack_2_0','player_basic_attack_2',0,'temporal_sweep',100,250,0,85,95,0,110,150,58,165,90,TRUE,FALSE,FALSE,1,0,FALSE,'motion_player_basic_attack_2_right_to_left_v1','player_basic_attack_2_damage',-45,45,55,58),
('hitbox_player_basic_attack_3_0','player_basic_attack_3',0,'temporal_sweep',180,440,0,95,95,0,120,155,60,210,0,TRUE,FALSE,FALSE,1,0,FALSE,'motion_player_basic_attack_3_shield_drive_v1','player_basic_attack_3_damage',0,0,60,60),
('hitbox_player_shield_bash_0','player_shield_bash',0,'temporal_sweep',120,340,0,90,95,0,190,160,95,210,0,TRUE,FALSE,FALSE,1,0,FALSE,'motion_player_shield_bash_front_push_v1','player_shield_bash_front_push',0,0,95,95),
('hitbox_player_shield_rush_0','player_shield_rush',0,'temporal_sweep',160,590,0,120,100,0,190,160,96,290,0,TRUE,FALSE,FALSE,1,0,FALSE,'motion_player_shield_rush_front_contact_v1','player_shield_rush_front_contact',0,0,96,96),
('hitbox_bite_0','bite',0,'temporal_sweep',120,340,0,80,90,0,95,115,48,145,0,TRUE,FALSE,FALSE,1,0,FALSE,'motion_wolf_bite_melee_v1','wolf_bite_damage',0,0,45,48),
('hitbox_lunge_0','lunge',0,'temporal_sweep',3600,4030,0,130,105,0,100,120,50,320,0,TRUE,FALSE,FALSE,1,0,FALSE,'motion_wolf_lunge_cross_v1','wolf_lunge_damage',0,0,50,50)
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
