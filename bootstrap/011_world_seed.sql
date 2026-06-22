-- =========================================================
-- WORLD DEFAULTS
-- APEIRON MMO
-- ANCIENT CHINA - NORTHERN FRONTIER
-- =========================================================

-- =========================================================
-- WORLD REGION
-- =========================================================

INSERT INTO apeiron.world_region (
    id,
    name,
    region_type,
    world_scale,
    is_instanced,
    max_players,
    center_x,
    center_y,
    center_z,
    size_x,
    size_y,
    size_z,
    danger_level
)
VALUES (
    'region_northern_frontier',
    'Northern Frontier',
    'wild',
    1,
    FALSE,
    120,
    0.0,
    0.0,
    0.0,
    4000.0,
    800.0,
    4000.0,
    0.55
)
ON CONFLICT (id) DO UPDATE SET
    name = EXCLUDED.name,
    region_type = EXCLUDED.region_type,
    world_scale = EXCLUDED.world_scale,
    is_instanced = EXCLUDED.is_instanced,
    max_players = EXCLUDED.max_players,
    center_x = EXCLUDED.center_x,
    center_y = EXCLUDED.center_y,
    center_z = EXCLUDED.center_z,
    size_x = EXCLUDED.size_x,
    size_y = EXCLUDED.size_y,
    size_z = EXCLUDED.size_z,
    danger_level = EXCLUDED.danger_level,
    updated_at = NOW();

-- =========================================================
-- BIOMES
-- =========================================================

INSERT INTO apeiron.biome (
    id,
    name,
    region_id,
    biome_type,
    temperature,
    humidity,
    visibility_modifier,
    movement_modifier,
    stealth_modifier,
    aggression_modifier,
    resource_richness,
    is_safe
)
VALUES
(
    'steppe',
    'Northern Steppe',
    'region_northern_frontier',
    'grassland',
    0.45,
    0.30,
    1.15,
    1.05,
    0.75,
    1.10,
    0.65,
    FALSE
),
(
    'grassland',
    'Frontier Grassland',
    'region_northern_frontier',
    'grassland',
    0.55,
    0.45,
    1.05,
    1.00,
    0.85,
    1.00,
    0.85,
    FALSE
),
(
    'forest_edge',
    'Sparse Forest Edge',
    'region_northern_frontier',
    'forest',
    0.50,
    0.55,
    0.85,
    0.92,
    1.20,
    0.95,
    1.00,
    FALSE
)
ON CONFLICT (id) DO UPDATE SET
    name = EXCLUDED.name,
    region_id = EXCLUDED.region_id,
    biome_type = EXCLUDED.biome_type,
    temperature = EXCLUDED.temperature,
    humidity = EXCLUDED.humidity,
    visibility_modifier = EXCLUDED.visibility_modifier,
    movement_modifier = EXCLUDED.movement_modifier,
    stealth_modifier = EXCLUDED.stealth_modifier,
    aggression_modifier = EXCLUDED.aggression_modifier,
    resource_richness = EXCLUDED.resource_richness,
    is_safe = EXCLUDED.is_safe,
    updated_at = NOW();

-- =========================================================
-- SPAWN ZONES
-- =========================================================

INSERT INTO apeiron.spawn_zone (
    id,
    region_id,
    biome_id,
    name,
    center_x,
    center_y,
    center_z,
    radius,
    max_entities,
    current_entities,
    respawn_time_ms,
    spawn_density,
    allowed_archetypes,
    aggression_level,
    leash_enabled
)
VALUES
(
    'spawnzone_steppe_wolf_den',
    'region_northern_frontier',
    'steppe',
    'Steppe Wolf Den',
    -620.0,
    0.0,
    340.0,
    180.0,
    10,
    0,
    600000,
    0.65,
    'beast',
    0.70,
    TRUE
),
(
    'spawnzone_open_grassland_hunt',
    'region_northern_frontier',
    'grassland',
    'Open Grassland Hunting Ground',
    220.0,
    0.0,
    -480.0,
    240.0,
    14,
    0,
    480000,
    0.75,
    'beast',
    0.55,
    TRUE
),
(
    'spawnzone_forest_edge_roam',
    'region_northern_frontier',
    'forest_edge',
    'Forest Edge Roaming Path',
    760.0,
    0.0,
    120.0,
    160.0,
    6,
    0,
    720000,
    0.35,
    'beast',
    0.45,
    TRUE
)
ON CONFLICT (id) DO UPDATE SET
    region_id = EXCLUDED.region_id,
    biome_id = EXCLUDED.biome_id,
    name = EXCLUDED.name,
    center_x = EXCLUDED.center_x,
    center_y = EXCLUDED.center_y,
    center_z = EXCLUDED.center_z,
    radius = EXCLUDED.radius,
    max_entities = EXCLUDED.max_entities,
    current_entities = EXCLUDED.current_entities,
    respawn_time_ms = EXCLUDED.respawn_time_ms,
    spawn_density = EXCLUDED.spawn_density,
    allowed_archetypes = EXCLUDED.allowed_archetypes,
    aggression_level = EXCLUDED.aggression_level,
    leash_enabled = EXCLUDED.leash_enabled,
    updated_at = NOW();