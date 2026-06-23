# DB Legacy Compatibility Audit - 2026-06-22

This audit exists to stop blind cleanup during recovery. A table or column named
`legacy` is not automatically trash. It can be a runtime compatibility surface,
a recovery bridge for older restored databases, or dead data that still needs
proof before removal.

## Classification

- `final_authority`: current canonical schema/data path. Runtime should consume this.
- `compat_runtime_required`: old surface still intentionally exposed to server/client/runtime.
- `recovery_only`: migration bridge for partially recovered or old databases. Not gameplay authority.
- `dead_candidate`: no current runtime/reference evidence found, but removal still needs a test pass.

## Source Precedence

1. Current reconstructed code in DB, server, and Unreal.
2. `docs/recovery/chat-recovery-ledger-2026-06-22.md`.
3. Recent thread facts recorded in recovery docs.
4. Recovered SQL/docs only when they do not conflict with newer runtime facts.

## Audit Table

| Area | Status | Evidence | Decision |
| --- | --- | --- | --- |
| `movement_action_contract` | `final_authority` | `migrations/027_action_runtime_contracts.sql`, `bootstrap/014_action_runtime_contract_seed.sql`, server runtime contract loading, Unreal movement action manifests | Keep as canonical action movement source. It owns distance, speed, phase windows, reconciliation contract id, and metadata. |
| `skill_movement_action_binding` | `final_authority` | `migrations/027_action_runtime_contracts.sql`, `bootstrap/014_action_runtime_contract_seed.sql`, server `movement_action_contract_id` ack/test paths | Keep as canonical skill-to-movement binding. Skill movement must prefer this over legacy movement effect rows. |
| `movement_action_contract.ability_key`, `movement_type`, `contract_version`, `contract_hash` from older schemas | `recovery_only` | `migrations/034_movement_action_legacy_column_compatibility.sql` makes them nullable so restored DBs do not block startup | Do not consume as authority. Keep migration while recovered databases can contain those columns. Later baseline/squash can drop them after fresh DB rebuild tests. |
| `skill_movement_effect` table and `GetSkillMovementEffect(skill_id)` | `compat_runtime_required` | `migrations/031_legacy_skill_movement_effect.sql`, `migrations/040_skill_movement_effect_legacy_column_compatibility.sql`, `bootstrap/018_legacy_skill_movement_effect_seed.sql`, DB handler maps to `SkillMovementProfile`, recovery ledger says keep `lunge` row until action-contract path fully replaces lookup | Keep for now. It is not the preferred authority, but it is still an exposed compatibility endpoint and seed surface. Exit only after server/client/proto prove no migrated skill depends on this lookup. |
| `skill_movement_effect` gameplay values for migrated skills | `dead_candidate` after proof | Current rows carry metadata `prefer: movement_action_contract`; authoritative data exists in movement action contracts | Do not tune gameplay here except to keep compatibility responses coherent. After runtime proof, detach from gameplay selection or remove seed rows. |
| Temporal melee hitbox motion profile model | `final_authority` | `migrations/028_temporal_melee_hitbox.sql`, `bootstrap/015_temporal_hitbox_seed.sql`, server/Unreal temporal hitbox debug paths | Keep as canonical hit detection model. Directional melee should use motion samples/profiles, not static full-arc damage when temporal motion is available. |
| `skill_hitbox_motion_profile.hitbox_profile_id` | `recovery_only` | `migrations/035_temporal_hitbox_legacy_compatibility.sql` preserves old inverse column as nullable compatibility | Do not consume as authority. Current model is profile-to-motion via `skill_hitbox_profile.motion_profile_id` and samples by `motion_profile_id`. |
| `skill_hitbox_motion_sample.id` text fallback | `recovery_only` | `migrations/036_temporal_hitbox_sample_id_finalization.sql` only normalizes recovered TEXT id shape | Keep until recovery is stable. Fresh schema should use the final sample id model from `028`. |
| Creature behavior runtime contract | `final_authority` | `migrations/029_creature_behavior_contracts.sql`, `migrations/041_creature_behavior_opportunity_contracts.sql`, `bootstrap/016_wolf_behavior_contract_seed.sql`, server wolf runtime policy loading | Keep as canonical creature brain contract. It owns runtime behavior links, opportunity policy, orbit policy, skill behavior bindings, and wolf tactical availability. |
| `creature_behavior_runtime_contract.display_name`, `combat_role_id` legacy requirements | `recovery_only` | `migrations/037_creature_behavior_legacy_column_compatibility.sql` drops old NOT NULL requirements | Do not restore UI/catalog ownership here unless a real runtime use is found. Current behavior is keyed by `creature_template_id` and policy ids. |
| Creature setup `setup_type` and `movement_tactic` | `final_authority` | `migrations/041_creature_behavior_opportunity_contracts.sql`, `bootstrap/016_wolf_behavior_contract_seed.sql` use values such as `moving_windup`, `chase_windup`, `pressure_counter` | Keep as canonical setup model for creature skill preparation. |
| Creature setup `setup_tactic` and stale setup checks | `recovery_only` | `migrations/038_creature_setup_legacy_column_compatibility.sql`, `migrations/039_creature_setup_check_finalization.sql` | Keep only as migration tolerance. Do not make AI decisions from `setup_tactic`. |
| Creature orbit modern policy fields | `final_authority` | `migrations/041_creature_behavior_opportunity_contracts.sql`, `bootstrap/016_wolf_behavior_contract_seed.sql` include orbit locomotion mode, side switch and duration rules | Keep as canonical orbit/side-switch tuning. |
| Creature orbit `preferred_radius_cm`, `min_radius_cm`, `max_radius_cm` from older schemas | `recovery_only` | `migrations/042_creature_orbit_legacy_column_finalization.sql` makes old radius columns non-blocking | Keep nullable for recovered DBs. Do not revive them as runtime authority unless code proves a required consumer. |
| `runtime_movement_reconciliation_profile` | `final_authority` | `migrations/043_runtime_movement_reconciliation_profile.sql`, `bootstrap/020_runtime_movement_reconciliation_profile_seed.sql`, server strict load of `player_default_movement_profile`, Unreal `MovementReconciliationProfile` fields | Keep as canonical rich movement/reconciliation profile. Server must fail if this profile is missing instead of inventing values. |
| `033_schema_compatibility.sql` additive movement profile defaults | `recovery_only` | Adds missing columns with old safety defaults for partially recovered DBs | Keep only to unblock old DB shapes. Do not treat its numeric defaults as tuned movement authority. Tuned values must come from seeds/contracts/profile rows. |
| Original modern numbering `051/052/022/023/024` | `dead_candidate` as filenames only | Ledger proves these existed before deletion, but current reconstructed DB uses compact replacement numbering | Do not recreate duplicate migration files solely for numbering. Either recover exact originals or intentionally keep compact numbering with this compatibility map. |

## Immediate Roadmap

### Phase 1 - Classification

Status: done in this document.

The current rule is: do not delete `legacy` paths until each one has a consumer proof and an exit
test. `skill_movement_effect` is the main active compatibility surface. Most `*_legacy_*`
migrations are recovery-only schema tolerance.

### Phase 2 - Runtime Usage Proof

Status: partial, server normal-runtime proof added on 2026-06-23.

Required checks before any removal:

- Search DB/server/Unreal for legacy fields and endpoint usage.
- Server proof added:
  - `server-apeiron/internal/gameapi/legacy_contract_surfaces.go` classifies runtime contract surfaces.
  - `server-apeiron/internal/gameapi/contracts_test.go` proves `ContractSource` does not expose `GetSkillMovementEffect`.
  - `RuntimeStats` now reports `contracts.surface.*` statuses so legacy compatibility cannot hide as normal authority.
- Verify live DB rows for:
  - `skill_movement_effect`
  - `skill_movement_action_binding`
  - `movement_action_contract`
  - `runtime_movement_reconciliation_profile`
  - creature behavior/opportunity/orbit/setup tables.
- Verify gRPC for:
  - `GetSkillMovementEffect(lunge)`
  - skill movement action binding/contract fetch paths
  - `GetRuntimeMovementReconciliationProfile(player_default_movement_profile)`
  - wolf behavior runtime contract and tactical bindings.

Current proof result:

- `skill_movement_effect/GetSkillMovementEffect` remains `compat_runtime_required` because DB still exposes it.
- It is not normal runtime authority for the reconstructed server loader.
- Canonical player/creature skill movement authority is `skill_movement_action_binding` plus `movement_action_contract`.
- Removal/quarantine is not approved until Unreal and any external tools are checked the same way.

### Phase 3 - Isolate Compatibility

Status: pending.

Actions:

- Add comments/status metadata to compatibility migrations/seeds where missing.
- Ensure runtime logs and server loaders identify when a legacy compatibility path is used.
- For migrated skills, assert the server uses movement action contracts as the authority.
- Keep compatibility endpoint responses coherent but never use them to override canonical contracts.

### Phase 4 - Remove Or Quarantine Dead Paths

Status: pending after Phase 2/3.

Rules:

- `recovery_only` stays until a fresh DB baseline and migrated recovered DB both pass.
- `compat_runtime_required` can move to `recovery_only` only after no runtime consumer remains.
- `dead_candidate` can be removed only with a narrow patch, focused tests, and a sentinel audit.

## Done Criteria

- No duplicated live authority for skill movement, creature behavior, temporal hitbox, or movement reconciliation.
- Runtime source of truth is clear for every gameplay value.
- Legacy code/columns either have an explicit compatibility reason or are removed after proof.
- Fresh DB and recovered DB shapes both migrate successfully.
- Server and Unreal consume canonical contracts without fallback values inventing gameplay.
