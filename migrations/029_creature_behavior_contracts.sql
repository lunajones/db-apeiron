-- =========================================================
-- CREATURE BEHAVIOR CONTRACTS
-- Reconstructed from wolf behavior/combat AI tuning work.
-- =========================================================

CREATE TABLE IF NOT EXISTS apeiron.creature_behavior_runtime_contract (
    id TEXT PRIMARY KEY,
    creature_template_id TEXT,
    description TEXT NOT NULL DEFAULT '',
    aggression_curve JSONB NOT NULL DEFAULT '{}'::jsonb,
    range_policy JSONB NOT NULL DEFAULT '{}'::jsonb,
    orbit_policy JSONB NOT NULL DEFAULT '{}'::jsonb,
    pressure_policy JSONB NOT NULL DEFAULT '{}'::jsonb,
    stamina_policy JSONB NOT NULL DEFAULT '{}'::jsonb,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_creature_behavior_template
        FOREIGN KEY (creature_template_id)
        REFERENCES apeiron.creature_template(id)
        ON DELETE CASCADE
);

ALTER TABLE IF EXISTS apeiron.creature_behavior_runtime_contract
    ADD COLUMN IF NOT EXISTS creature_template_id TEXT,
    ADD COLUMN IF NOT EXISTS description TEXT NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS aggression_curve JSONB NOT NULL DEFAULT '{}'::jsonb,
    ADD COLUMN IF NOT EXISTS range_policy JSONB NOT NULL DEFAULT '{}'::jsonb,
    ADD COLUMN IF NOT EXISTS orbit_policy JSONB NOT NULL DEFAULT '{}'::jsonb,
    ADD COLUMN IF NOT EXISTS pressure_policy JSONB NOT NULL DEFAULT '{}'::jsonb,
    ADD COLUMN IF NOT EXISTS stamina_policy JSONB NOT NULL DEFAULT '{}'::jsonb,
    ADD COLUMN IF NOT EXISTS metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    ADD COLUMN IF NOT EXISTS created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP NOT NULL DEFAULT NOW();

CREATE TABLE IF NOT EXISTS apeiron.creature_evasion_policy (
    id TEXT PRIMARY KEY,
    behavior_contract_id TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    dodge_skill_id TEXT,
    max_chain_count INT NOT NULL DEFAULT 1,
    stamina_cost_multiplier FLOAT NOT NULL DEFAULT 1.0,
    retreat_chance_multiplier FLOAT NOT NULL DEFAULT 1.0,
    lateral_bias FLOAT NOT NULL DEFAULT 0.5,
    backstep_bias FLOAT NOT NULL DEFAULT 0.5,
    pressure_threshold FLOAT NOT NULL DEFAULT 0.5,
    cooldown_ms INT NOT NULL DEFAULT 0,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_creature_evasion_behavior
        FOREIGN KEY (behavior_contract_id)
        REFERENCES apeiron.creature_behavior_runtime_contract(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_creature_evasion_skill
        FOREIGN KEY (dodge_skill_id)
        REFERENCES apeiron.skill(id)
        ON DELETE SET NULL,
    CONSTRAINT chk_creature_evasion_chain CHECK (max_chain_count >= 0)
);

ALTER TABLE IF EXISTS apeiron.creature_evasion_policy
    ADD COLUMN IF NOT EXISTS behavior_contract_id TEXT NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS description TEXT NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS dodge_skill_id TEXT,
    ADD COLUMN IF NOT EXISTS max_chain_count INT NOT NULL DEFAULT 1,
    ADD COLUMN IF NOT EXISTS stamina_cost_multiplier FLOAT NOT NULL DEFAULT 1.0,
    ADD COLUMN IF NOT EXISTS retreat_chance_multiplier FLOAT NOT NULL DEFAULT 1.0,
    ADD COLUMN IF NOT EXISTS lateral_bias FLOAT NOT NULL DEFAULT 0.5,
    ADD COLUMN IF NOT EXISTS backstep_bias FLOAT NOT NULL DEFAULT 0.5,
    ADD COLUMN IF NOT EXISTS pressure_threshold FLOAT NOT NULL DEFAULT 0.5,
    ADD COLUMN IF NOT EXISTS cooldown_ms INT NOT NULL DEFAULT 0,
    ADD COLUMN IF NOT EXISTS metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    ADD COLUMN IF NOT EXISTS created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP NOT NULL DEFAULT NOW();

CREATE TABLE IF NOT EXISTS apeiron.creature_skill_setup_policy (
    id TEXT PRIMARY KEY,
    behavior_contract_id TEXT NOT NULL,
    skill_id TEXT NOT NULL,
    setup_type TEXT NOT NULL,
    min_setup_ms INT NOT NULL DEFAULT 0,
    max_setup_ms INT NOT NULL DEFAULT 0,
    commit_distance_cm FLOAT NOT NULL DEFAULT 0.0,
    preferred_min_range_cm FLOAT NOT NULL DEFAULT 0.0,
    preferred_max_range_cm FLOAT NOT NULL DEFAULT 0.0,
    movement_tactic TEXT NOT NULL DEFAULT 'none',
    lock_side_during_setup BOOLEAN NOT NULL DEFAULT TRUE,
    is_enabled BOOLEAN NOT NULL DEFAULT TRUE,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_creature_setup_behavior
        FOREIGN KEY (behavior_contract_id)
        REFERENCES apeiron.creature_behavior_runtime_contract(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_creature_setup_skill
        FOREIGN KEY (skill_id)
        REFERENCES apeiron.skill(id)
        ON DELETE CASCADE,
    CONSTRAINT chk_creature_setup_time CHECK (max_setup_ms >= min_setup_ms)
);

ALTER TABLE IF EXISTS apeiron.creature_skill_setup_policy
    ADD COLUMN IF NOT EXISTS behavior_contract_id TEXT NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS skill_id TEXT NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS setup_type TEXT NOT NULL DEFAULT 'none',
    ADD COLUMN IF NOT EXISTS min_setup_ms INT NOT NULL DEFAULT 0,
    ADD COLUMN IF NOT EXISTS max_setup_ms INT NOT NULL DEFAULT 0,
    ADD COLUMN IF NOT EXISTS commit_distance_cm FLOAT NOT NULL DEFAULT 0.0,
    ADD COLUMN IF NOT EXISTS preferred_min_range_cm FLOAT NOT NULL DEFAULT 0.0,
    ADD COLUMN IF NOT EXISTS preferred_max_range_cm FLOAT NOT NULL DEFAULT 0.0,
    ADD COLUMN IF NOT EXISTS movement_tactic TEXT NOT NULL DEFAULT 'none',
    ADD COLUMN IF NOT EXISTS lock_side_during_setup BOOLEAN NOT NULL DEFAULT TRUE,
    ADD COLUMN IF NOT EXISTS is_enabled BOOLEAN NOT NULL DEFAULT TRUE,
    ADD COLUMN IF NOT EXISTS metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    ADD COLUMN IF NOT EXISTS created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP NOT NULL DEFAULT NOW();

CREATE INDEX IF NOT EXISTS idx_creature_setup_behavior
ON apeiron.creature_skill_setup_policy(behavior_contract_id);

CREATE INDEX IF NOT EXISTS idx_creature_setup_skill
ON apeiron.creature_skill_setup_policy(skill_id);
