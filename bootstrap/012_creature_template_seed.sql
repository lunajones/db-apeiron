-- =========================================================
-- CREATURE DEFAULTS
-- APEIRON MMO
-- STEPPE WOLF - TEMPLATE
-- =========================================================

INSERT INTO apeiron.creature_template (
    id,
    name,
    faction,
    tier,
    archetype,

    movement_profile_id,
    combat_core_profile_id,
    combat_style_profile_id,
    ai_decision_profile_id,
    personality_profile_id,
    sensory_profile_id,
    needs_profile_id,
    skill_set_id,
    spawn_profile_id
)
VALUES (
    'steppe_wolf',
    'Steppe Wolf',
    'wildlife',
    1,
    'beast',

    'movement_steppe_wolf',
    'combat_core_steppe_wolf',
    'combat_style_steppe_wolf',
    'ai_decision_steppe_wolf',
    'personality_steppe_wolf',
    'sensory_steppe_wolf',
    'needs_steppe_wolf',
    'skillset_steppe_wolf',
    'spawn_steppe_wolf'
)
ON CONFLICT (id) DO UPDATE SET
    name = EXCLUDED.name,
    faction = EXCLUDED.faction,
    tier = EXCLUDED.tier,
    archetype = EXCLUDED.archetype,

    movement_profile_id = EXCLUDED.movement_profile_id,
    combat_core_profile_id = EXCLUDED.combat_core_profile_id,
    combat_style_profile_id = EXCLUDED.combat_style_profile_id,
    ai_decision_profile_id = EXCLUDED.ai_decision_profile_id,
    personality_profile_id = EXCLUDED.personality_profile_id,
    sensory_profile_id = EXCLUDED.sensory_profile_id,
    needs_profile_id = EXCLUDED.needs_profile_id,
    skill_set_id = EXCLUDED.skill_set_id,
    spawn_profile_id = EXCLUDED.spawn_profile_id,

    updated_at = NOW();