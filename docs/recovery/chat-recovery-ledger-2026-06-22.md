# Chat Recovery Ledger - 2026-06-22

This file records facts recovered from Codex thread history after the db/server data loss.
Use this as source precedence during reconstruction: newer threads override older recovered files
when they clearly describe a later runtime state.

## Source Precedence

1. Current thread, latest runtime fixes and user test reports.
2. Recent Codex thread `DB` (`019e76bb-3b35-7b22-8ffe-b2a12484692e`, updated 2026-06 range).
3. Recent Codex thread `continuar daqui dia 10` (`019e92d3-b2e2-7162-b129-c1c4a681f5a2`).
4. Handoff/reconciliation thread `Separar handoff de movimento` (`019e9ac0-22bf-7bc1-8476-5be9f640c0e4`).
5. WinFR recovered server docs/roadmaps under `B:\ApeironWinFR_Server_Pass2`.
6. VS Code History recovered DB baseline under `B:\AR\db-apeiron`.

Older SQL recovered from VS Code History must not override a later chat fact unless the chat
fact is only speculative.

## Recovered Facts

### Cross-Thread Sweep Status

Pass `2026-06-22 01:55` inspected these Codex threads:

- `DB` (`019e76bb-3b35-7b22-8ffe-b2a12484692e`)
- `continuar daqui dia 10` (`019e92d3-b2e2-7162-b129-c1c4a681f5a2`)
- `Separar handoff de movimento` (`019e9ac0-22bf-7bc1-8476-5be9f640c0e4`)
- `SERVER + unreal combinado` (`019e92d4-9a1f-7c00-b366-d79db34c9e4d`)
- `Sword/shield skill design and VFX` (`019edb84-a5ad-7693-8bab-38ddd3b27363`)

Thread recovery is not just "latest file wins". Use the following merge policy:

1. Prefer the newest completed chat fact when it describes a runtime-tested value.
2. Prefer exact file names and migration numbers from the thread where the change was made.
3. Do not let a recovered old SQL file overwrite a newer chat-described contract.
4. If the reconstructed DB uses compact replacement names/numbers, keep a compatibility map until
   the original numbering can be restored or intentionally replaced.
5. Any seed with multiplicative updates, such as `base_damage = base_damage * 0.5`, must be
   converted to explicit idempotent values before being treated as canonical.

### Original Modern Migration / Seed Numbering From Chat

The recovered disk baseline is older. The chat history proves the live project had a later DB
contract layer with at least these files:

- `migrations/051_movement_action_contract.sql`
- `bootstrap/022_movement_action_contract_seed.sql`
- `migrations/052_creature_behavior_runtime_contract.sql`
- `bootstrap/023_creature_behavior_contract_seed.sql`
- `bootstrap/024_combat_defense_contract_seed.sql`

Current reconstructed DB files in this folder use compacted replacements:

- `migrations/027_action_runtime_contracts.sql` covers part of old `051`.
- `bootstrap/014_action_runtime_contract_seed.sql` covers part of old `022`.
- `migrations/029_creature_behavior_contracts.sql` covers part of old `052`.
- `bootstrap/016_wolf_behavior_contract_seed.sql` covers part of old `023`.
- `migrations/032_combat_defense_contracts.sql` and
  `bootstrap/019_combat_defense_contract_seed.sql` cover part of old `024`.

Reconstruction rule: do not assume `027-032` are the final canonical names. They are a recovery
scaffold until either the original `051/052/023/024` files are recovered, or the project is
deliberately migrated to the compacted numbering with a clean migration ledger.

### Movement Action Contract Magnitude

From thread `DB`, turn `019ea03b-0271-7393-a707-0bb015b5e4dd`:

- A critical gap was found: movement action curves/policies existed, but did not carry absolute
  magnitude.
- The AAA fix was to add absolute fields to `movement_action_contract`:
  - `horizontal_distance_cm`
  - `base_speed_cm_per_sec`
- The curve field remains a scale, not an absolute speed.
- Seeds were updated for:
  - global leap: `420 cm`, `1400 cm/s`
  - global dodge: `400 cm`, `1600 cm/s`
  - light/medium/heavy leap/dodge variants
  - creature-specific magnitudes
- Reconstruction rule: the server must not fall back to old locomotion ability just to invent
  movement action magnitude. The contract needs explicit distance/speed.

### Skill Movement Effect Legacy Compatibility

From thread `DB`, turn `019ebcec-6eb1-7db1-9760-c5268732f84f`:

- `GetSkillMovementEffect(lunge)` originally returned empty because the DB endpoint searched by
  effect `id` only.
- The actual row existed as `id=leap_default`, `skill_id=lunge`.
- After the fix, `GetSkillMovementEffect(lunge)` returned:
  - `found=true`
  - `id=leap_default`
  - `skillId=lunge`
  - `movementType=leap`
  - `distance=420`
  - `speed=1400`
  - `durationMs=300`
- Reconstruction rule: keep a legacy `skill_movement_effect` row for `lunge` until the server
  action-contract path fully replaces this lookup.

### Wolf Behavior Contract

From thread `DB`, turn `019ebcf8-02b9-7ca0-8e9c-ef8bf5899cd4`:

- There was a seed file named `bootstrap/023_creature_behavior_contract_seed.sql`.
- The wolf behavior contract had `commit_angle_max_deg=145`, then it was changed to `180`.
- Reason: the wolf could not commit to attack when the player had their back turned.
- gRPC confirmed contract with:
  - `commitAngleMaxDeg=180`
  - binding `lunge + approach + acquire`
  - distance approximately `180-560 cm`
- Reconstruction rule: `contract_wolf_pack_harasser_v1` must allow back-side commit up to `180`.

From thread `DB`, turns `019ebc8a-b8fe-7b30-957d-6da11eb1a65a` through
`019ebc9d-874d-7d11-8b7e-44b53b5e0226`:

- `migrations/052_creature_behavior_runtime_contract.sql` was extended with:
  - `creature_target_opportunity_policy`
  - `creature_behavior_runtime_contract.target_opportunity_policy_id`
  - `creature_orbit_policy.orbit_locomotion_mode`
- The wolf seed used:
  - `orbit_locomotion_mode = combat_walk`
  - `target_opportunity_policy_id = opportunity_wolf_harasser_v1`
  - `orbit_speed_scale = 0.55`
  - `side_switch_cooldown_ms = 900`
  - `min_orbit_duration_ms = 700`
  - `allow_side_switch_when_target_faces = true`
- Wolf movement tuning was reduced:
  - `max_speed = 340`
  - `acceleration = 1100`
  - `deceleration = 900`
  - `sprint_multiplier = 1.55`
  - `dodge_distance = 330`
- Server-side issue found later: orbit side had been deterministic by `TargetID%2` and then
  reused previous side forever until server logic was changed. DB alone cannot fix that.

Reconstruction rule: wolf combat movement needs both DB policy and server consumption. Do not
restore a DB-only "fix" and assume orbit side/opportunity is complete.

### Wolf Skill Binding / Lunge Runtime

From thread `DB`, turns around `019ebc83` through `019ebcf8`:

- Wolf `lunge` binding needed to be available during `circle/reposition` and
  `approach/acquire`, range about `180-560 cm`.
- The server matcher had to treat `any` as wildcard and `commit_attack` as offensive commitment.
- Hitbox seed bug found: `offset_x` is forward and `offset_y` is lateral. Wolf bite/lunge had
  forward hitboxes accidentally placed sideways by using `offset_y=100`.
- Runtime after DB reset applied:
  - migration `052_creature_behavior_runtime_contract.sql`
  - bootstrap `023_creature_behavior_contract_seed.sql`
- Reconstruction rule: lunge cannot be only a long static hitbox. It needs movement effect/action
  contract and forward hitbox orientation.

### Defense / Stamina Contract

From thread `continuar daqui dia 10`, latest turn:

- Stamina must not be damaged by normal unblocked hits.
- Stamina damage is valid only through block/guard resolution.
- Defense arc should be symmetric around defender facing, not trace aim.
- Defender-facing margin was adjusted visually to roughly `30%` left/right of player cylinder.
- Reconstruction rule: combat defense contract must distinguish:
  - health damage for unblocked hit;
  - stamina/posture pressure only when block/guard is active and frontal.

From thread `continuar daqui dia 10`, turns around `019ec259` through `019ec65a`:

- Block accepted by a valid defense should mitigate `100%` HP damage.
- Differences between shield/sword/dagger belong in stamina, posture, arc, recovery, guard break,
  parry window, reaction, and riposte policy.
- Shield parry window recorded in chat:
  - shield: `60ms -> 220ms`, `160ms` window
  - sword: `70ms -> 190ms`, `120ms` window
- Parry success must:
  - apply zero HP/posture/stamina damage to defender;
  - resolve `source_event = parry_success`, not `guard_break`;
  - interrupt/trap the attacker during the contracted riposte/stagger window;
  - publish enough metadata to distinguish early/late/in-window parry.
- `guard_stability` existed in DB but was not being used by the server until a server fix made it
  reduce block stamina/posture pressure.
- `bootstrap/024_combat_defense_contract_seed.sql` is the later observed seed name for this area.

Reconstruction rule: defense should be reconstructed as a contract family, not as direct damage
branches in combat code.

### Movement / Reconciliation Architecture

From `Separar handoff de movimento` and current thread:

- Existing architecture involved:
  - server authoritative movement;
  - input buffer;
  - pending input replay;
  - DB/server movement contracts;
  - action timeline for dodge/leap;
  - action root history;
  - visual smoothing separated from authority;
  - snapshot timeline / jitter buffer planned;
  - correction debt planned.
- Reconstruction rule: do not collapse normal movement, dodge, leap, turn, and skill movement
  into one generic reconciliation contract when category-specific ownership is needed.

From `SERVER + unreal combinado` and `Separar handoff de movimento`:

- A key production bug was that the bridge dropped `Locomotion` from snapshots, so Unreal received
  empty `locomotion_action` even though the server had computed leap/dodge/landing handoff.
- `ActionMove` with zero direction and stale `client_position` was retained too long and caused
  post-action rubberband.
- Player move intent retention was later tightened:
  - player move direction may remain briefly for continuity;
  - stale client position must expire quickly and cannot remain authoritative.
- Later architecture phases added/preserved:
  - `action_distance_traveled`
  - `action_projected_position`
  - `ActionRootHistory`
  - visual-only correction split from gameplay capsule
  - mesh/camera correction split
  - `CameraFocusHandoff`
  - non-regression ledger for movement.

Reconstruction rule: any restored server/proto/bridge must carry locomotion fields end-to-end.
If `locomotion_action` is empty in Unreal while the server is resolving an action, the build is not
functionally restored.

### Weapon Kit / Combat Modes

From current thread:

- Sword-and-shield has at least two combat modes:
  - `Vanguard`
  - `Bulwark`
- Current active loadout:
  - `Bulwark`: `R = player_shield_bash`, `F = player_shield_rush`, `Q = empty`, `M1 = basic combo`
  - `Vanguard`: no active Q/R/F skills yet; `M1 = basic combo`
- Mode switch duration was reduced to roughly half of the prior value, settling around `250 ms`.

### Player Defense / Spawn Grace / Block UX

From `continuar daqui dia 10`:

- A runtime spawn/session grace was added:
  - newly attached player gets about `8s` authoritative invulnerability;
  - combat pipeline rejects damage during that window;
  - AI suppresses targeting protected players.
- Block visual debug was changed:
  - remove giant local shield sphere;
  - use concise `BLOCK`, `PARRY`, `EVADE` feedback.
- Normal combat movement presentation should support New World style strafe/backpedal relative to
  camera/facing while preserving dodge/leap as action-owned movement.

Reconstruction rule: do not confuse spawn grace with normal combat invulnerability. It belongs to
session/runtime safety, not skill iFrame design.

### Temporal Hitboxes

From current thread and recovered roadmaps:

- Basic attack 1: forward temporal cut from player body to roughly `1.5` player cylinders forward.
- Basic attack 2: right-to-left temporal sword sweep, about `90 degrees`.
- Basic attack 3: shield drive, front contact/carry/push, width later reduced to about `1.5` player cylinders.
- Shield Rush: hitbox follows player front; damage starts around half a cylinder in front, not a full cylinder.
- Shield Bash: wider path than player cylinder, multi-target push.

## Still Missing / Needs Further Thread Sweep

- Exact modern DB migration numbering beyond the recovered baseline.
- Exact `combat_defense_contract` original column set.
- Exact modern proto definitions for weapon kit, movement contracts, behavior contracts, and hitbox motion.
- Exact seed file contents for modern `020+` through `060+` range.
- Continue older pages of:
  - `DB` before cursor `019ea022-7550-7170-91f7-f7d5b0ef02a5`;
  - `server` / `first` threads for earliest schema and world/inventory structure;
  - current latest thread after the DB/server deletion for exact final requested skill movement and
    HUD changes.

## 2026-06-22 - Creature Brain DB Contract Recovery Slice

Recovered and applied the missing wolf behavior contract layer as DB authority instead of leaving
it as server literals or loose JSON only.

Files added/updated:

- `migrations/041_creature_behavior_opportunity_contracts.sql`
- `migrations/042_creature_orbit_legacy_column_finalization.sql`
- `bootstrap/016_wolf_behavior_contract_seed.sql`
- `proto/apeiron/v1/common.proto`
- `proto/apeiron/v1/profile_data_service.proto`
- `proto/apeiron/v1/creature_data_service.proto`
- repository/cache/gRPC handlers for profile and creature runtime data.

Recovered model:

- `creature_target_opportunity_policy`
  - `opportunity_wolf_harasser_v1`
  - `commit_angle_max_deg = 180`
  - `no_ready_skill_memory_policy = observe_only`
  - exposes candidate/cooldown diagnostics in metadata.
- `creature_orbit_policy`
  - `orbit_wolf_harasser_combat_walk_v1`
  - `orbit_locomotion_mode = combat_walk`
  - `orbit_speed_scale = 0.55`
  - `min_orbit_duration_ms = 700`
  - `side_switch_cooldown_ms = 900`
  - keeps side changes from thrashing.
- `creature_skill_behavior_binding`
  - bite: `approach/acquire`, `circle/reposition`
  - lunge: `approach/acquire`, `circle/reposition`
  - maul: `pressure/counter`
  - wolf_dodge: `pressure/evade`

Validation:

- `go test ./...` passes in `db-apeiron`.
- `go build -o bin/db-api.exe ./cmd/db-api` passes.
- Live Postgres verification:
  - migration `041` applied.
  - migration `042` applied.
  - `opportunity_wolf_harasser_v1` exists with `commit_angle_max_deg = 180`.
  - `orbit_wolf_harasser_combat_walk_v1` exists with `orbit_locomotion_mode = combat_walk`.
  - six wolf tactical skill bindings exist.
  - `contract_wolf_pack_harasser_v1` links `steppe_wolf` to the recovered opportunity/orbit policies.

Important design decision:

- Do not add a reverse `creature_template.behavior_contract_id` column yet. The existing
  authoritative relationship is `creature_behavior_runtime_contract.creature_template_id`.
  This avoids a cyclic FK during recovery while still allowing `CreatureRuntimeData` to resolve
  the behavior contract by template id.
