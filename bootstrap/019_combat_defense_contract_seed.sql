-- =========================================================
-- COMBAT DEFENSE CONTRACT SEED
-- =========================================================

INSERT INTO apeiron.combat_defense_contract (
    id, name, description, defense_type, frontal_arc_deg,
    defender_margin_left_ratio, defender_margin_right_ratio,
    stamina_damage_only_on_block, health_damage_on_unblocked_hit,
    posture_damage_on_block, perfect_block_window_ms, parry_window_ms,
    guard_damage_multiplier, block_stamina_drain_per_second, metadata
)
VALUES
('player_shield_guard_v1','Player Shield Guard','Symmetric frontal shield guard. Stamina damage applies only on successful block.', 'shield_block',120,0.30,0.30,TRUE,TRUE,TRUE,0,0,1.0,2.0,'{"source":"canonical_bootstrap","frontFacing":"control_rotation_yaw"}'),
('player_perfect_guard_v1','Player Perfect Guard','Short perfect block/parry-capable guard window for future shield skills.', 'perfect_block',130,0.30,0.30,TRUE,TRUE,TRUE,120,90,0.35,2.0,'{"source":"canonical_bootstrap"}'),
('wolf_attack_vs_guard_v1','Wolf Attack Vs Guard','Default creature attack guard interaction. Normal hit damages health; blocked hit pressures stamina/posture.', 'incoming_melee',120,0.30,0.30,TRUE,TRUE,TRUE,0,0,1.0,0.0,'{"source":"canonical_bootstrap"}')
ON CONFLICT (id) DO UPDATE SET
    name = EXCLUDED.name,
    description = EXCLUDED.description,
    defense_type = EXCLUDED.defense_type,
    frontal_arc_deg = EXCLUDED.frontal_arc_deg,
    defender_margin_left_ratio = EXCLUDED.defender_margin_left_ratio,
    defender_margin_right_ratio = EXCLUDED.defender_margin_right_ratio,
    stamina_damage_only_on_block = EXCLUDED.stamina_damage_only_on_block,
    health_damage_on_unblocked_hit = EXCLUDED.health_damage_on_unblocked_hit,
    posture_damage_on_block = EXCLUDED.posture_damage_on_block,
    perfect_block_window_ms = EXCLUDED.perfect_block_window_ms,
    parry_window_ms = EXCLUDED.parry_window_ms,
    guard_damage_multiplier = EXCLUDED.guard_damage_multiplier,
    block_stamina_drain_per_second = EXCLUDED.block_stamina_drain_per_second,
    metadata = EXCLUDED.metadata,
    updated_at = NOW();
