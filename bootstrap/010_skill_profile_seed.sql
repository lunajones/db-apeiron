-- =========================================================
-- DEFAULT SKILL PROFILES
-- APEIRON MMO
-- STEPPE WOLF - HITBOX / AREA / IMPACT CONFIG
-- =========================================================

-- =========================================================
-- BITE - HITBOX
-- =========================================================

INSERT INTO apeiron.skill_hitbox_profile (
    id,
    skill_id,

    hitbox_index,
    hitbox_shape,

    hitbox_start_ms,
    hitbox_end_ms,

    offset_x,
    offset_y,
    offset_z,

    size_x,
    size_y,
    size_z,

    radius,
    length,
    angle,

    follows_caster,
    follows_projectile,

    can_multi_hit,
    max_hits_per_target,
    hit_interval_ms,

    friendly_fire
)
VALUES (
    'hitbox_bite_0',
    'bite',

    0,
    'capsule',

    180,
    340,

    0.0,
    1.0,
    0.75,

    0.6,
    0.6,
    0.9,

    0.32,
    0.9,
    0.0,

    TRUE,
    FALSE,

    FALSE,
    1,
    0,

    FALSE
)
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
    follows_caster = EXCLUDED.follows_caster,
    follows_projectile = EXCLUDED.follows_projectile,
    can_multi_hit = EXCLUDED.can_multi_hit,
    max_hits_per_target = EXCLUDED.max_hits_per_target,
    hit_interval_ms = EXCLUDED.hit_interval_ms,
    friendly_fire = EXCLUDED.friendly_fire,
    updated_at = NOW();

-- =========================================================
-- BITE - IMPACT
-- =========================================================

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
    status_effect_chance
)
VALUES (
    'bite',

    'normal',
    'flinch',

    18.0,
    0.25,

    0.20,
    0.85,

    FALSE,
    TRUE,
    FALSE,

    45,
    0.10,

    0.65,
    0.05,
    0.0,

    FALSE,
    NULL,
    0.0
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
    updated_at = NOW();

-- =========================================================
-- LUNGE - HITBOX
-- =========================================================

INSERT INTO apeiron.skill_hitbox_profile (
    id,
    skill_id,

    hitbox_index,
    hitbox_shape,

    hitbox_start_ms,
    hitbox_end_ms,

    offset_x,
    offset_y,
    offset_z,

    size_x,
    size_y,
    size_z,

    radius,
    length,
    angle,

    follows_caster,
    follows_projectile,

    can_multi_hit,
    max_hits_per_target,
    hit_interval_ms,

    friendly_fire
)
VALUES (
    'hitbox_lunge_0',
    'lunge',

    0,
    'capsule',

    260,
    440,

    0.0,
    1.0,
    1.05,

    0.75,
    0.75,
    1.35,

    0.42,
    1.35,
    0.0,

    TRUE,
    FALSE,

    FALSE,
    1,
    0,

    FALSE
)
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
    follows_caster = EXCLUDED.follows_caster,
    follows_projectile = EXCLUDED.follows_projectile,
    can_multi_hit = EXCLUDED.can_multi_hit,
    max_hits_per_target = EXCLUDED.max_hits_per_target,
    hit_interval_ms = EXCLUDED.hit_interval_ms,
    friendly_fire = EXCLUDED.friendly_fire,
    updated_at = NOW();

-- =========================================================
-- LUNGE - IMPACT
-- =========================================================

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
    status_effect_chance
)
VALUES (
    'lunge',

    'heavy',
    'stagger',

    32.0,
    0.55,

    0.45,
    1.15,

    TRUE,
    TRUE,
    FALSE,

    70,
    0.22,

    2.0,
    0.15,
    0.0,

    FALSE,
    NULL,
    0.0
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
    updated_at = NOW();

-- =========================================================
-- PACK HOWL - AREA EFFECT
-- =========================================================

INSERT INTO apeiron.skill_area_effect_profile (
    skill_id,

    area_shape,

    radius,
    length,
    width,
    height,
    angle,

    duration_ms,
    tick_interval_ms,

    damage_falloff_type,
    min_falloff_multiplier,

    applies_on_impact,
    persists_after_impact,

    max_targets,

    friendly_fire,

    status_effect_id
)
VALUES (
    'pack_howl',

    'sphere',

    9.0,
    0.0,
    0.0,
    0.0,
    360.0,

    0,
    0,

    'none',
    1.0,

    TRUE,
    FALSE,

    6,

    FALSE,

    NULL
)
ON CONFLICT (skill_id) DO UPDATE SET
    area_shape = EXCLUDED.area_shape,
    radius = EXCLUDED.radius,
    length = EXCLUDED.length,
    width = EXCLUDED.width,
    height = EXCLUDED.height,
    angle = EXCLUDED.angle,
    duration_ms = EXCLUDED.duration_ms,
    tick_interval_ms = EXCLUDED.tick_interval_ms,
    damage_falloff_type = EXCLUDED.damage_falloff_type,
    min_falloff_multiplier = EXCLUDED.min_falloff_multiplier,
    applies_on_impact = EXCLUDED.applies_on_impact,
    persists_after_impact = EXCLUDED.persists_after_impact,
    max_targets = EXCLUDED.max_targets,
    friendly_fire = EXCLUDED.friendly_fire,
    status_effect_id = EXCLUDED.status_effect_id,
    updated_at = NOW();

-- =========================================================
-- PACK HOWL - IMPACT
-- =========================================================

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
    status_effect_chance
)
VALUES (
    'pack_howl',

    'no_reaction',
    'none',

    8.0,
    0.0,

    0.15,
    0.0,

    FALSE,
    FALSE,
    FALSE,

    0,
    0.0,

    0.0,
    0.0,
    0.0,

    FALSE,
    NULL,
    0.0
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
    updated_at = NOW();