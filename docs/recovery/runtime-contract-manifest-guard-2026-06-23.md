# Runtime Contract Manifest Guard - 2026-06-23

This guard exists because Apeiron can compile while the recovered DB is still hollow.
The server boot path has a runtime requirement manifest; the DB bootstrap must mirror
that manifest so missing gameplay contracts fail in tests before Unreal opens with
fallback-like behavior.

## Protected Runtime Surfaces

The DB bootstrap must explicitly seed these canonical surfaces:

- `movement_profile`: `player_default_movement_profile`
- `base_movement_action`: `grounded_move_v1`
- `base_movement_action`: `turn_v1_rate_limited_contextual`
- `base_movement_action`: `dodge_v1_full_iframe`
- `base_movement_action`: `jump_v1_authoritative_grounded_handoff`
- `combat_core_profile`: `combat_core_player_sword_shield_v1`
- `combat_core_profile`: `combat_core_steppe_wolf`
- `defense_contract`: `player_shield_guard_v1`
- `defense_contract`: `wolf_attack_vs_guard_v1`
- `weapon_kit`: `weaponkit_sword_shield`
- `wolf_brain_policy`: `contract_wolf_pack_harasser_v1`

The DB bootstrap must also explicitly bind the active skill runtime:

- `player_basic_attack_1` -> `basic_attack_1_forward_cut_v1`
- `player_basic_attack_2` -> `basic_attack_2_cross_cut_v1`
- `player_basic_attack_3` -> `basic_attack_3_shield_drive_v1`
- `player_shield_bash` -> `shield_bash_front_push_v1`
- `player_shield_rush` -> `shield_rush_front_contact_v1`
- `bite` -> `wolf_bite_melee_commit_v1`
- `lunge` -> `low_fast_lunge_v1`
- `wolf_dodge` -> `wolf_dodge_lateral_leap_v1`
- `maul` -> `wolf_maul_lateral_counter_v1`

## Why This Is Not A Fallback

This is a guard, not a replacement runtime. It does not invent gameplay values in Go.
It checks the SQL bootstrap for the contract rows that the server already treats as
required. If any required row is missing, the DB package test fails and the missing
contract must be restored in migrations/bootstrap/proto/repository code.

## Compatibility Rule

Tables or rows with `legacy`, `compat`, `fallback`, or `recovered` in names or metadata
are evidence to audit. They are not normal gameplay authority unless a runtime consumer
proves it still needs them. Canonical runtime authority must stay in:

- movement action contracts;
- skill action timing;
- skill movement action bindings;
- temporal hitbox motion profiles and damage groups;
- combat core profiles;
- combat defense contracts;
- weapon kit / combat mode slot tables;
- creature behavior runtime contracts and sub-policies.

## Validation

The manifest is enforced by:

- `TestBootstrapSeedsMirrorServerRuntimeRequirementManifest`
- `TestBootstrapSeedsMirrorRequiredSkillActionManifest`
- existing focused bootstrap tests for temporal hitboxes, weapon modes, shield rush
  contact geometry, wolf maul, wolf behavior policy, and movement reconciliation.

Any future reconstruction slice that changes server runtime requirements must update
the DB manifest test in the same commit or deliberately fail until the matching DB
contract is restored.
