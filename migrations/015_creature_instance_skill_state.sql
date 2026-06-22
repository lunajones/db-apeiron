CREATE TABLE IF NOT EXISTS apeiron.creature_instance_skill_state (
    creature_instance_id TEXT NOT NULL,
    skill_id TEXT NOT NULL,

    cooldown_end_ms BIGINT NOT NULL DEFAULT 0,
    last_used_ms BIGINT NOT NULL DEFAULT 0,

    charges INT NOT NULL DEFAULT 0,
    -- pra skills com múltiplos usos

    is_locked BOOLEAN NOT NULL DEFAULT FALSE,
    -- travado por CC, stun, silence etc

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    PRIMARY KEY (creature_instance_id, skill_id)
);
