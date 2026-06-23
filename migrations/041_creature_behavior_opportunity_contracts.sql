-- =========================================================
-- CREATURE BEHAVIOR OPPORTUNITY / ORBIT CONTRACTS
-- Reconstructs the later wolf brain contract layer described in chat recovery:
-- target opportunities, orbit locomotion, and tactical skill bindings are data.
-- =========================================================

CREATE TABLE IF NOT EXISTS apeiron.creature_target_opportunity_policy (
    id TEXT PRIMARY KEY,
    description TEXT NOT NULL DEFAULT '',
    commit_angle_max_deg FLOAT NOT NULL DEFAULT 180.0,
    min_commit_distance_cm FLOAT NOT NULL DEFAULT 0.0,
    max_commit_distance_cm FLOAT NOT NULL DEFAULT 0.0,
    approach_min_distance_cm FLOAT NOT NULL DEFAULT 0.0,
    approach_max_distance_cm FLOAT NOT NULL DEFAULT 0.0,
    bite_range_cm FLOAT NOT NULL DEFAULT 0.0,
    lunge_min_range_cm FLOAT NOT NULL DEFAULT 0.0,
    lunge_max_range_cm FLOAT NOT NULL DEFAULT 0.0,
    maul_pressure_threshold FLOAT NOT NULL DEFAULT 0.0,
    target_memory_ms INT NOT NULL DEFAULT 0,
    no_ready_skill_memory_policy TEXT NOT NULL DEFAULT 'observe_only',
    candidate_cooldown_visibility BOOLEAN NOT NULL DEFAULT TRUE,
    allow_backside_commit BOOLEAN NOT NULL DEFAULT TRUE,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_creature_opportunity_angle CHECK (commit_angle_max_deg >= 0 AND commit_angle_max_deg <= 180),
    CONSTRAINT chk_creature_opportunity_distances CHECK (
        min_commit_distance_cm >= 0
        AND max_commit_distance_cm >= min_commit_distance_cm
        AND approach_min_distance_cm >= 0
        AND approach_max_distance_cm >= approach_min_distance_cm
        AND lunge_min_range_cm >= 0
        AND lunge_max_range_cm >= lunge_min_range_cm
    )
);


CREATE TABLE IF NOT EXISTS apeiron.creature_orbit_policy (
    id TEXT PRIMARY KEY,
    behavior_contract_id TEXT,
    description TEXT NOT NULL DEFAULT '',
    orbit_locomotion_mode TEXT NOT NULL DEFAULT 'combat_walk',
    orbit_speed_scale FLOAT NOT NULL DEFAULT 0.55,
    min_orbit_duration_ms INT NOT NULL DEFAULT 700,
    side_switch_cooldown_ms INT NOT NULL DEFAULT 900,
    allow_side_switch_when_target_faces BOOLEAN NOT NULL DEFAULT TRUE,
    prefer_long_side_commit BOOLEAN NOT NULL DEFAULT TRUE,
    side_flip_chance_multiplier FLOAT NOT NULL DEFAULT 0.35,
    lock_side_during_setup BOOLEAN NOT NULL DEFAULT TRUE,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_creature_orbit_behavior
        FOREIGN KEY (behavior_contract_id)
        REFERENCES apeiron.creature_behavior_runtime_contract(id)
        ON DELETE CASCADE,
    CONSTRAINT chk_creature_orbit_timing CHECK (min_orbit_duration_ms >= 0 AND side_switch_cooldown_ms >= 0),
    CONSTRAINT chk_creature_orbit_scale CHECK (orbit_speed_scale >= 0)
);


CREATE TABLE IF NOT EXISTS apeiron.creature_skill_behavior_binding (
    id TEXT PRIMARY KEY,
    behavior_contract_id TEXT NOT NULL,
    skill_id TEXT NOT NULL,
    tactical_state TEXT NOT NULL,
    decision_phase TEXT NOT NULL,
    setup_policy_id TEXT,
    min_range_cm FLOAT NOT NULL DEFAULT 0.0,
    max_range_cm FLOAT NOT NULL DEFAULT 0.0,
    priority INT NOT NULL DEFAULT 0,
    usage_weight FLOAT NOT NULL DEFAULT 1.0,
    cooldown_group TEXT NOT NULL DEFAULT '',
    requires_line_of_sight BOOLEAN NOT NULL DEFAULT TRUE,
    is_enabled BOOLEAN NOT NULL DEFAULT TRUE,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_creature_skill_binding_behavior
        FOREIGN KEY (behavior_contract_id)
        REFERENCES apeiron.creature_behavior_runtime_contract(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_creature_skill_binding_skill
        FOREIGN KEY (skill_id)
        REFERENCES apeiron.skill(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_creature_skill_binding_setup
        FOREIGN KEY (setup_policy_id)
        REFERENCES apeiron.creature_skill_setup_policy(id)
        ON DELETE SET NULL,
    CONSTRAINT chk_creature_skill_binding_range CHECK (min_range_cm >= 0 AND max_range_cm >= min_range_cm),
    CONSTRAINT chk_creature_skill_binding_weight CHECK (usage_weight >= 0)
);






CREATE INDEX IF NOT EXISTS idx_creature_orbit_behavior
ON apeiron.creature_orbit_policy(behavior_contract_id);

CREATE INDEX IF NOT EXISTS idx_creature_skill_binding_behavior
ON apeiron.creature_skill_behavior_binding(behavior_contract_id);

CREATE INDEX IF NOT EXISTS idx_creature_skill_binding_skill
ON apeiron.creature_skill_behavior_binding(skill_id);

CREATE INDEX IF NOT EXISTS idx_creature_behavior_template
ON apeiron.creature_behavior_runtime_contract(creature_template_id);
