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

### Defense / Stamina Contract

From thread `continuar daqui dia 10`, latest turn:

- Stamina must not be damaged by normal unblocked hits.
- Stamina damage is valid only through block/guard resolution.
- Defense arc should be symmetric around defender facing, not trace aim.
- Defender-facing margin was adjusted visually to roughly `30%` left/right of player cylinder.
- Reconstruction rule: combat defense contract must distinguish:
  - health damage for unblocked hit;
  - stamina/posture pressure only when block/guard is active and frontal.

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

### Weapon Kit / Combat Modes

From current thread:

- Sword-and-shield has at least two combat modes:
  - `Vanguard`
  - `Bulwark`
- Current active loadout:
  - `Bulwark`: `R = player_shield_bash`, `F = player_shield_rush`, `Q = empty`, `M1 = basic combo`
  - `Vanguard`: no active Q/R/F skills yet; `M1 = basic combo`
- Mode switch duration was reduced to roughly half of the prior value, settling around `250 ms`.

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
