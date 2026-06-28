-- =========================================================
-- DEV PLAYER — persistent character record for local testing
-- =========================================================
-- A real persistent player row so progression (level/xp/attributes/coin) has somewhere to load from
-- and persist to, replacing the purely in-memory fabricated player. The id matches the game server's
-- default attach id ("local_player"). creature_instance_id is left NULL — the runtime body is spawned
-- at login, not stored here. ON CONFLICT DO NOTHING so re-seeding never resets earned progression.
--
-- NOTE: db-api resets the apeiron schema on every boot (dev), so this row is recreated at level 1 on a
-- db-api restart. Progression persists across client reconnects and game-server restarts while db-api
-- stays up. A proper account/character-creation flow replaces this dev seed later.

-- strength 6 is a DEMO value so Slice 5 (attribute scaling) is visible immediately: 150 max health,
-- +25% physical damage, +10 physical resistance. A real new character starts at strength 1.0; this
-- is replaced once the attribute-point spend command exists.
INSERT INTO apeiron.player (id, account_id, name, strength)
VALUES ('local_player', 'account_dev_local', 'Wanderer', 6.0)
ON CONFLICT (id) DO NOTHING;
