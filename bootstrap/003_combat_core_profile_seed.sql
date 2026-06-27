-- =========================================================
-- DEFAULT COMBAT CORE PROFILE
-- APEIRON MMO
-- STEPPE WOLF - FAST LIGHT PREDATOR
-- =========================================================

INSERT INTO apeiron.combat_core_profile (
    id,

    physical_defense,
    magic_defense,

    critical_chance,
    critical_multiplier,

    damage_taken_multiplier,
    damage_dealt_multiplier,

    max_stamina,
    stamina_regen_per_sec,

    dodge_stamina_cost,
    sprint_stamina_cost_per_sec,
    block_stamina_cost_per_sec,
    attack_stamina_cost,

    stamina_exhaustion_threshold,
    stamina_zero_regen_multiplier,

    max_posture,
    posture_recovery_rate,
    posture_damage_multiplier,
    posture_break_duration_ms,

    block_damage_reduction,
    parry_window_ms,
    parry_reward_multiplier,

    can_block,
    can_parry,

    dodge_iframe_ms,
    dodge_cooldown_ms,

    stun_resistance,
    root_resistance,
    knockback_resistance,
    cc_duration_multiplier,

    is_boss,
    is_pvp_immune,

    physical_resistance_rating,
    chemical_resistance_rating,
    biological_resistance_rating,
    resistance_cap
)
VALUES (
    'combat_core_steppe_wolf',

    0.12,
    0.03,

    0.08,
    1.45,

    1.05,
    0.95,

    120.0,
    16.0,

    28.0,
    5.0,
    0.0,
    14.0,

    0.18,
    0.5,

    65.0,
    18.0,
    1.15,
    1800,

    0.0,
    0,
    1.0,

    FALSE,
    FALSE,

    220,
    850,

    0.85,
    0.75,
    0.65,
    1.05,

    FALSE,
    FALSE,

    -- resistance ratings (Physical / Chemical / Biological) + cap
    40.0,
    30.0,
    45.0,
    0.85
)
ON CONFLICT (id) DO UPDATE SET
    physical_defense = EXCLUDED.physical_defense,
    magic_defense = EXCLUDED.magic_defense,
    critical_chance = EXCLUDED.critical_chance,
    critical_multiplier = EXCLUDED.critical_multiplier,
    damage_taken_multiplier = EXCLUDED.damage_taken_multiplier,
    damage_dealt_multiplier = EXCLUDED.damage_dealt_multiplier,
    max_stamina = EXCLUDED.max_stamina,
    stamina_regen_per_sec = EXCLUDED.stamina_regen_per_sec,
    dodge_stamina_cost = EXCLUDED.dodge_stamina_cost,
    sprint_stamina_cost_per_sec = EXCLUDED.sprint_stamina_cost_per_sec,
    block_stamina_cost_per_sec = EXCLUDED.block_stamina_cost_per_sec,
    attack_stamina_cost = EXCLUDED.attack_stamina_cost,
    stamina_exhaustion_threshold = EXCLUDED.stamina_exhaustion_threshold,
    stamina_zero_regen_multiplier = EXCLUDED.stamina_zero_regen_multiplier,
    max_posture = EXCLUDED.max_posture,
    posture_recovery_rate = EXCLUDED.posture_recovery_rate,
    posture_damage_multiplier = EXCLUDED.posture_damage_multiplier,
    posture_break_duration_ms = EXCLUDED.posture_break_duration_ms,
    block_damage_reduction = EXCLUDED.block_damage_reduction,
    parry_window_ms = EXCLUDED.parry_window_ms,
    parry_reward_multiplier = EXCLUDED.parry_reward_multiplier,
    can_block = EXCLUDED.can_block,
    can_parry = EXCLUDED.can_parry,
    dodge_iframe_ms = EXCLUDED.dodge_iframe_ms,
    dodge_cooldown_ms = EXCLUDED.dodge_cooldown_ms,
    stun_resistance = EXCLUDED.stun_resistance,
    root_resistance = EXCLUDED.root_resistance,
    knockback_resistance = EXCLUDED.knockback_resistance,
    cc_duration_multiplier = EXCLUDED.cc_duration_multiplier,
    is_boss = EXCLUDED.is_boss,
    is_pvp_immune = EXCLUDED.is_pvp_immune,
    physical_resistance_rating = EXCLUDED.physical_resistance_rating,
    chemical_resistance_rating = EXCLUDED.chemical_resistance_rating,
    biological_resistance_rating = EXCLUDED.biological_resistance_rating,
    resistance_cap = EXCLUDED.resistance_cap,
    updated_at = NOW();

-- =========================================================
-- PLAYER SWORD + SHIELD CORE PROFILE
-- Canonical vertical-slice player combat profile. The game server
-- must load this through ProfileDataService; it must not invent a
-- local guard/core profile when resolving block, parry, posture, or
-- damage.
-- =========================================================

INSERT INTO apeiron.combat_core_profile (
    id,

    physical_defense,
    magic_defense,

    critical_chance,
    critical_multiplier,

    damage_taken_multiplier,
    damage_dealt_multiplier,

    max_stamina,
    stamina_regen_per_sec,

    dodge_stamina_cost,
    sprint_stamina_cost_per_sec,
    block_stamina_cost_per_sec,
    attack_stamina_cost,

    stamina_exhaustion_threshold,
    stamina_zero_regen_multiplier,

    max_posture,
    posture_recovery_rate,
    posture_damage_multiplier,
    posture_break_duration_ms,

    block_damage_reduction,
    parry_window_ms,
    parry_reward_multiplier,

    can_block,
    can_parry,

    dodge_iframe_ms,
    dodge_cooldown_ms,

    stun_resistance,
    root_resistance,
    knockback_resistance,
    cc_duration_multiplier,

    is_boss,
    is_pvp_immune,

    physical_resistance_rating,
    chemical_resistance_rating,
    biological_resistance_rating,
    resistance_cap
)
VALUES (
    'combat_core_player_sword_shield_v1',

    0.08,
    0.04,

    0.05,
    1.50,

    1.00,
    1.00,

    100.0,
    14.0,

    24.0,
    24.0,
    2.0,
    0.0,

    0.0,
    0.5,

    100.0,
    14.0,
    1.00,
    2200,

    1.0,
    220,
    1.40,

    TRUE,
    TRUE,

    320,
    0,

    1.0,
    1.0,
    1.0,
    1.0,

    FALSE,
    FALSE,

    -- resistance ratings (Physical / Chemical / Biological) + cap
    80.0,
    25.0,
    30.0,
    0.85
)
ON CONFLICT (id) DO UPDATE SET
    physical_defense = EXCLUDED.physical_defense,
    magic_defense = EXCLUDED.magic_defense,
    critical_chance = EXCLUDED.critical_chance,
    critical_multiplier = EXCLUDED.critical_multiplier,
    damage_taken_multiplier = EXCLUDED.damage_taken_multiplier,
    damage_dealt_multiplier = EXCLUDED.damage_dealt_multiplier,
    max_stamina = EXCLUDED.max_stamina,
    stamina_regen_per_sec = EXCLUDED.stamina_regen_per_sec,
    dodge_stamina_cost = EXCLUDED.dodge_stamina_cost,
    sprint_stamina_cost_per_sec = EXCLUDED.sprint_stamina_cost_per_sec,
    block_stamina_cost_per_sec = EXCLUDED.block_stamina_cost_per_sec,
    attack_stamina_cost = EXCLUDED.attack_stamina_cost,
    stamina_exhaustion_threshold = EXCLUDED.stamina_exhaustion_threshold,
    stamina_zero_regen_multiplier = EXCLUDED.stamina_zero_regen_multiplier,
    max_posture = EXCLUDED.max_posture,
    posture_recovery_rate = EXCLUDED.posture_recovery_rate,
    posture_damage_multiplier = EXCLUDED.posture_damage_multiplier,
    posture_break_duration_ms = EXCLUDED.posture_break_duration_ms,
    block_damage_reduction = EXCLUDED.block_damage_reduction,
    parry_window_ms = EXCLUDED.parry_window_ms,
    parry_reward_multiplier = EXCLUDED.parry_reward_multiplier,
    can_block = EXCLUDED.can_block,
    can_parry = EXCLUDED.can_parry,
    dodge_iframe_ms = EXCLUDED.dodge_iframe_ms,
    dodge_cooldown_ms = EXCLUDED.dodge_cooldown_ms,
    stun_resistance = EXCLUDED.stun_resistance,
    root_resistance = EXCLUDED.root_resistance,
    knockback_resistance = EXCLUDED.knockback_resistance,
    cc_duration_multiplier = EXCLUDED.cc_duration_multiplier,
    is_boss = EXCLUDED.is_boss,
    is_pvp_immune = EXCLUDED.is_pvp_immune,
    physical_resistance_rating = EXCLUDED.physical_resistance_rating,
    chemical_resistance_rating = EXCLUDED.chemical_resistance_rating,
    biological_resistance_rating = EXCLUDED.biological_resistance_rating,
    resistance_cap = EXCLUDED.resistance_cap,
    updated_at = NOW();
