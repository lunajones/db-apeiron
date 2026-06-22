-- =========================================================
-- SWORD AND SHIELD WEAPON KIT / COMBAT MODES
-- =========================================================

INSERT INTO apeiron.weapon_kit (id, name, description, primary_weapon_type, offhand_weapon_type, is_enabled, metadata)
VALUES (
    'weaponkit_sword_shield',
    'Sword and Shield',
    'Sword-and-shield weapon kit. Player swaps between Vanguard and Bulwark combat modes.',
    'sword',
    'shield',
    TRUE,
    '{"source":"reconstructed"}'
)
ON CONFLICT (id) DO UPDATE SET
    name = EXCLUDED.name,
    description = EXCLUDED.description,
    primary_weapon_type = EXCLUDED.primary_weapon_type,
    offhand_weapon_type = EXCLUDED.offhand_weapon_type,
    is_enabled = EXCLUDED.is_enabled,
    metadata = EXCLUDED.metadata,
    updated_at = NOW();

INSERT INTO apeiron.weapon_combat_mode (
    id, weapon_kit_id, name, description, mode_index, switch_duration_ms, is_enabled, metadata
)
VALUES
('mode_sword_shield_vanguard','weaponkit_sword_shield','Vanguard','Sword-forward pressure mode. Future selectable sword skills live here.',0,250,TRUE,'{"source":"reconstructed"}'),
('mode_sword_shield_bulwark','weaponkit_sword_shield','Bulwark','Shield-forward control mode. Current active Q/R/F skills live here.',1,250,TRUE,'{"source":"reconstructed"}')
ON CONFLICT (id) DO UPDATE SET
    weapon_kit_id = EXCLUDED.weapon_kit_id,
    name = EXCLUDED.name,
    description = EXCLUDED.description,
    mode_index = EXCLUDED.mode_index,
    switch_duration_ms = EXCLUDED.switch_duration_ms,
    is_enabled = EXCLUDED.is_enabled,
    metadata = EXCLUDED.metadata,
    updated_at = NOW();

DELETE FROM apeiron.weapon_combat_mode_skill_slot
WHERE combat_mode_id IN ('mode_sword_shield_vanguard','mode_sword_shield_bulwark');

INSERT INTO apeiron.weapon_combat_mode_skill_slot (
    combat_mode_id, input_slot, skill_id, is_basic_attack, is_fatality, is_enabled, metadata
)
VALUES
('mode_sword_shield_vanguard','M1',NULL,FALSE,FALSE,FALSE,'{"emptyUntilSelected":true}'),
('mode_sword_shield_vanguard','Q',NULL,FALSE,FALSE,FALSE,'{"emptyUntilSelected":true}'),
('mode_sword_shield_vanguard','R',NULL,FALSE,FALSE,FALSE,'{"emptyUntilSelected":true}'),
('mode_sword_shield_vanguard','F',NULL,FALSE,FALSE,FALSE,'{"emptyUntilSelected":true}'),
('mode_sword_shield_vanguard','G','player_fatality_placeholder',FALSE,TRUE,FALSE,'{"emptyUntilSelected":true}'),
('mode_sword_shield_bulwark','M1','player_basic_attack_1',TRUE,FALSE,TRUE,'{"comboGroup":"sword_shield_light_combo"}'),
('mode_sword_shield_bulwark','Q',NULL,FALSE,FALSE,FALSE,'{"emptyUntilSelected":true}'),
('mode_sword_shield_bulwark','R','player_shield_bash',FALSE,FALSE,TRUE,'{}'),
('mode_sword_shield_bulwark','F','player_shield_rush',FALSE,FALSE,TRUE,'{}'),
('mode_sword_shield_bulwark','G','player_fatality_placeholder',FALSE,TRUE,FALSE,'{"emptyUntilSelected":true}');
