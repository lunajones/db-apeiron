-- =========================================================
-- DEFAULT SPAWN PROFILE
-- APEIRON MMO
-- STEPPE WOLF - ANCIENT CHINA / NORTHERN FRONTIER
-- =========================================================

INSERT INTO apeiron.spawn_profile (
    id,
    allowed_biomes,
    biome_weights,
    allowed_regions,
    region_weights,
    density_base,
    density_cap,
    respawn_time_seconds,
    spawn_batch_size,
    spawn_variance,
    territorial_bias,
    migration_allowed,
    clustering_strength,
    is_dynamic,
    anti_overcrowd,
    pvp_zone_modifier
)
VALUES
(
    'spawn_steppe_wolf',
    '["steppe", "grassland", "forest_edge"]'::jsonb,
    '[
        {"biome_id": "steppe", "weight": 1.0},
        {"biome_id": "grassland", "weight": 0.9},
        {"biome_id": "forest_edge", "weight": 0.35}
    ]'::jsonb,
    '["region_northern_frontier"]'::jsonb,
    '[
        {"region_id": "region_northern_frontier", "weight": 1.0}
    ]'::jsonb,
    0.75,
    18,
    600,
    2,
    0.35,
    0.65,
    TRUE,
    0.85,
    TRUE,
    TRUE,
    1.15
)
ON CONFLICT (id) DO UPDATE SET
    allowed_biomes = EXCLUDED.allowed_biomes,
    biome_weights = EXCLUDED.biome_weights,
    allowed_regions = EXCLUDED.allowed_regions,
    region_weights = EXCLUDED.region_weights,
    density_base = EXCLUDED.density_base,
    density_cap = EXCLUDED.density_cap,
    respawn_time_seconds = EXCLUDED.respawn_time_seconds,
    spawn_batch_size = EXCLUDED.spawn_batch_size,
    spawn_variance = EXCLUDED.spawn_variance,
    territorial_bias = EXCLUDED.territorial_bias,
    migration_allowed = EXCLUDED.migration_allowed,
    clustering_strength = EXCLUDED.clustering_strength,
    is_dynamic = EXCLUDED.is_dynamic,
    anti_overcrowd = EXCLUDED.anti_overcrowd,
    pvp_zone_modifier = EXCLUDED.pvp_zone_modifier;