-- =========================================================
-- DEFAULT SKILLS
-- APEIRON MMO
-- GENERIC BEAST SKILLS + STEPPE WOLF SKILL SET
-- =========================================================

-- =========================
-- SKILLS
-- =========================

INSERT INTO apeiron.skill (
    id,
    name,
    description,
    archetype,
    skill_type,

    stamina_cost,
    mana_cost,
    health_cost,

    windup_ms,
    active_frames_ms,
    recovery_ms,
    cast_time_ms,
    cooldown_ms,

    cancel_window_start_ms,
    cancel_window_end_ms,
    iframe_start_ms,
    iframe_end_ms,

    min_range,
    max_range,
    cone_angle,
    max_targets,
    target_type,
    requires_target,

    base_damage,
    damage_type,
    elemental_type,
    posture_damage,
    armor_penetration,
    damage_multiplier,
    critical_bonus_multiplier,

    stun_duration_ms,
    root_duration_ms,
    knockback_force,

    movement_multiplier,
    locks_movement,
    movement_distance,

    combo_group,
    combo_index,
    combo_window_ms,

    is_interruptible,
    is_blockable,
    is_parryable,
    ignores_line_of_sight,
    ignores_collision
)
VALUES
(
    'bite',
    'Bite',
    'Fast close-range bite used by beast-type creatures.',
    'beast',
    'melee_attack',

    14.0,
    0.0,
    0.0,

    180,
    160,
    420,
    0,
    900,

    220,
    360,
    0,
    0,

    0.0,
    1.8,
    55.0,
    1,
    'enemy',
    TRUE,

    18.0,
    'physical',
    NULL,
    22.0,
    0.05,
    1.0,
    1.35,

    0,
    0,
    0.8,

    1.0,
    FALSE,
    0.0,

    'beast_basic_combo',
    1,
    900,

    TRUE,
    TRUE,
    TRUE,
    FALSE,
    FALSE
),
(
    'lunge',
    'Lunge',
    'Short explosive leap used by beast-type creatures to close distance.',
    'beast',
    'gap_closer',

    28.0,
    0.0,
    0.0,

    260,
    180,
    620,
    0,
    9000,

    340,
    520,
    80,
    180,

    1.5,
    11.4,
    35.0,
    1,
    'enemy',
    TRUE,

    26.0,
    'physical',
    NULL,
    34.0,
    0.08,
    1.15,
    1.45,

    0,
    0,
    2.2,

    1.35,
    FALSE,
    4.2,

    'beast_basic_combo',
    2,
    1100,

    TRUE,
    TRUE,
    TRUE,
    FALSE,
    FALSE
),
(
    'pack_howl',
    'Pack Howl',
    'Short howl used by pack creatures to pressure enemies and coordinate aggression.',
    'beast',
    'utility',

    20.0,
    0.0,
    0.0,

    420,
    300,
    700,
    0,
    9000,

    0,
    0,
    0,
    0,

    0.0,
    9.0,
    360.0,
    6,
    'enemy_area',
    FALSE,

    0.0,
    'none',
    NULL,
    10.0,
    0.0,
    1.0,
    1.0,

    0,
    0,
    0.0,

    0.85,
    TRUE,
    0.0,

    NULL,
    NULL,
    0,

    TRUE,
    FALSE,
    FALSE,
    FALSE,
    FALSE
)
ON CONFLICT (id) DO UPDATE SET
    name = EXCLUDED.name,
    description = EXCLUDED.description,
    archetype = EXCLUDED.archetype,
    skill_type = EXCLUDED.skill_type,

    stamina_cost = EXCLUDED.stamina_cost,
    mana_cost = EXCLUDED.mana_cost,
    health_cost = EXCLUDED.health_cost,

    windup_ms = EXCLUDED.windup_ms,
    active_frames_ms = EXCLUDED.active_frames_ms,
    recovery_ms = EXCLUDED.recovery_ms,
    cast_time_ms = EXCLUDED.cast_time_ms,
    cooldown_ms = EXCLUDED.cooldown_ms,

    cancel_window_start_ms = EXCLUDED.cancel_window_start_ms,
    cancel_window_end_ms = EXCLUDED.cancel_window_end_ms,
    iframe_start_ms = EXCLUDED.iframe_start_ms,
    iframe_end_ms = EXCLUDED.iframe_end_ms,

    min_range = EXCLUDED.min_range,
    max_range = EXCLUDED.max_range,
    cone_angle = EXCLUDED.cone_angle,
    max_targets = EXCLUDED.max_targets,
    target_type = EXCLUDED.target_type,
    requires_target = EXCLUDED.requires_target,

    base_damage = EXCLUDED.base_damage,
    damage_type = EXCLUDED.damage_type,
    elemental_type = EXCLUDED.elemental_type,
    posture_damage = EXCLUDED.posture_damage,
    armor_penetration = EXCLUDED.armor_penetration,
    damage_multiplier = EXCLUDED.damage_multiplier,
    critical_bonus_multiplier = EXCLUDED.critical_bonus_multiplier,

    stun_duration_ms = EXCLUDED.stun_duration_ms,
    root_duration_ms = EXCLUDED.root_duration_ms,
    knockback_force = EXCLUDED.knockback_force,

    movement_multiplier = EXCLUDED.movement_multiplier,
    locks_movement = EXCLUDED.locks_movement,
    movement_distance = EXCLUDED.movement_distance,

    combo_group = EXCLUDED.combo_group,
    combo_index = EXCLUDED.combo_index,
    combo_window_ms = EXCLUDED.combo_window_ms,

    is_interruptible = EXCLUDED.is_interruptible,
    is_blockable = EXCLUDED.is_blockable,
    is_parryable = EXCLUDED.is_parryable,
    ignores_line_of_sight = EXCLUDED.ignores_line_of_sight,
    ignores_collision = EXCLUDED.ignores_collision,

    updated_at = NOW();

-- =========================
-- STEPPE WOLF SKILL SET
-- =========================

INSERT INTO apeiron.skill_set (
    id,
    name,
    description,
    is_player_usable,
    is_npc_usable
)
VALUES (
    'skillset_steppe_wolf',
    'Steppe Wolf Skill Set',
    'Default skill set used by steppe wolves.',
    FALSE,
    TRUE
)
ON CONFLICT (id) DO UPDATE SET
    name = EXCLUDED.name,
    description = EXCLUDED.description,
    is_player_usable = EXCLUDED.is_player_usable,
    is_npc_usable = EXCLUDED.is_npc_usable,
    updated_at = NOW();

-- =========================
-- STEPPE WOLF SKILL SLOTS
-- =========================

DELETE FROM apeiron.skill_slot
WHERE skill_set_id = 'skillset_steppe_wolf';

INSERT INTO apeiron.skill_slot (
    skill_set_id,
    skill_id,
    slot_index,
    is_enabled,

    priority,
    usage_weight,
    cooldown_override_ms,

    min_target_hp_percent,
    max_target_hp_percent,
    min_self_hp_percent,
    max_self_hp_percent,

    required_distance_min,
    required_distance_max,

    requires_line_of_sight,

    opener_weight,
    finisher_weight,

    shared_cooldown_group,
    use_only_in_combat
)
VALUES
(
    'skillset_steppe_wolf',
    'bite',
    1,
    TRUE,

    60,
    1.0,
    NULL,

    NULL,
    NULL,
    0.15,
    NULL,

    0.0,
    1.8,

    TRUE,

    0.45,
    0.55,

    'steppe_wolf_melee',
    TRUE
),
(
    'skillset_steppe_wolf',
    'lunge',
    2,
    TRUE,

    84,
    0.48,
    NULL,

    NULL,
    NULL,
    0.25,
    NULL,

    1.8,
    11.4,

    TRUE,

    0.42,
    0.16,

    'steppe_wolf_gap_close',
    TRUE
),
(
    'skillset_steppe_wolf',
    'pack_howl',
    3,
    TRUE,

    40,
    0.35,
    NULL,

    NULL,
    NULL,
    0.30,
    NULL,

    0.0,
    9.0,

    TRUE,

    0.65,
    0.15,

    'steppe_wolf_utility',
    TRUE
);
