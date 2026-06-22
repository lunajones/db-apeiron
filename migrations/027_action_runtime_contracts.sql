-- =========================================================
-- ACTION RUNTIME CONTRACTS
-- Reconstructed from Apeiron recovery artifacts on 2026-06-22.
-- =========================================================

CREATE TABLE IF NOT EXISTS apeiron.movement_reconciliation_contract (
    id TEXT PRIMARY KEY,
    category TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    max_smooth_error_cm FLOAT NOT NULL DEFAULT 35.0,
    hard_snap_error_cm FLOAT NOT NULL DEFAULT 180.0,
    smoothing_time_ms INT NOT NULL DEFAULT 90,
    yaw_tolerance_deg FLOAT NOT NULL DEFAULT 8.0,
    owns_position BOOLEAN NOT NULL DEFAULT FALSE,
    owns_yaw BOOLEAN NOT NULL DEFAULT FALSE,
    allows_client_prediction BOOLEAN NOT NULL DEFAULT TRUE,
    input_policy TEXT NOT NULL DEFAULT 'normal',
    handoff_policy TEXT NOT NULL DEFAULT 'explicit',
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT chk_reconciliation_error CHECK (max_smooth_error_cm >= 0 AND hard_snap_error_cm >= max_smooth_error_cm),
    CONSTRAINT chk_reconciliation_smoothing CHECK (smoothing_time_ms >= 0)
);

ALTER TABLE IF EXISTS apeiron.movement_reconciliation_contract
    ALTER COLUMN id TYPE TEXT USING id::text,
    ADD COLUMN IF NOT EXISTS category TEXT NOT NULL DEFAULT 'generic',
    ADD COLUMN IF NOT EXISTS description TEXT NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS max_smooth_error_cm FLOAT NOT NULL DEFAULT 35.0,
    ADD COLUMN IF NOT EXISTS hard_snap_error_cm FLOAT NOT NULL DEFAULT 180.0,
    ADD COLUMN IF NOT EXISTS smoothing_time_ms INT NOT NULL DEFAULT 90,
    ADD COLUMN IF NOT EXISTS yaw_tolerance_deg FLOAT NOT NULL DEFAULT 8.0,
    ADD COLUMN IF NOT EXISTS owns_position BOOLEAN NOT NULL DEFAULT FALSE,
    ADD COLUMN IF NOT EXISTS owns_yaw BOOLEAN NOT NULL DEFAULT FALSE,
    ADD COLUMN IF NOT EXISTS allows_client_prediction BOOLEAN NOT NULL DEFAULT TRUE,
    ADD COLUMN IF NOT EXISTS input_policy TEXT NOT NULL DEFAULT 'normal',
    ADD COLUMN IF NOT EXISTS handoff_policy TEXT NOT NULL DEFAULT 'explicit',
    ADD COLUMN IF NOT EXISTS metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    ADD COLUMN IF NOT EXISTS created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP NOT NULL DEFAULT NOW();

CREATE TABLE IF NOT EXISTS apeiron.movement_action_contract (
    id TEXT PRIMARY KEY,
    action_type TEXT NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    duration_ms INT NOT NULL,
    active_ms INT NOT NULL,
    recovery_ms INT NOT NULL DEFAULT 0,
    distance_cm FLOAT NOT NULL DEFAULT 0.0,
    base_speed_cm_s FLOAT NOT NULL DEFAULT 0.0,
    yaw_degrees FLOAT NOT NULL DEFAULT 0.0,
    phase_window_policy TEXT NOT NULL DEFAULT 'grounded_action',
    prediction_error_policy TEXT NOT NULL DEFAULT 'bounded_smooth_correction',
    reconciliation_contract_id TEXT,
    allow_windup_locomotion BOOLEAN NOT NULL DEFAULT FALSE,
    allow_active_locomotion BOOLEAN NOT NULL DEFAULT FALSE,
    allow_recovery_locomotion BOOLEAN NOT NULL DEFAULT FALSE,
    allow_yaw_adjustment BOOLEAN NOT NULL DEFAULT TRUE,
    root_motion_owner TEXT NOT NULL DEFAULT 'movement',
    contact_policy TEXT NOT NULL DEFAULT 'none',
    speed_curve JSONB NOT NULL DEFAULT '[]'::jsonb,
    vertical_curve JSONB NOT NULL DEFAULT '[]'::jsonb,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_movement_action_reconciliation
        FOREIGN KEY (reconciliation_contract_id)
        REFERENCES apeiron.movement_reconciliation_contract(id),
    CONSTRAINT chk_movement_action_time CHECK (duration_ms > 0 AND active_ms >= 0 AND recovery_ms >= 0),
    CONSTRAINT chk_movement_action_distance CHECK (distance_cm >= 0)
);

ALTER TABLE IF EXISTS apeiron.movement_action_contract
    ADD COLUMN IF NOT EXISTS reconciliation_contract_id TEXT;

DO $$
DECLARE
    fk RECORD;
BEGIN
    FOR fk IN
        SELECT
            c.conrelid::regclass AS child_table,
            c.conname AS constraint_name,
            a.attname AS child_column
        FROM pg_constraint c
        JOIN LATERAL unnest(c.conkey) AS child_attnum(attnum) ON TRUE
        JOIN pg_attribute a ON a.attrelid = c.conrelid AND a.attnum = child_attnum.attnum
        WHERE c.contype = 'f'
          AND c.confrelid = 'apeiron.movement_action_contract'::regclass
    LOOP
        EXECUTE format(
            'ALTER TABLE %s DROP CONSTRAINT IF EXISTS %I',
            fk.child_table,
            fk.constraint_name
        );
        EXECUTE format(
            'ALTER TABLE %s ALTER COLUMN %I TYPE TEXT USING %I::text',
            fk.child_table,
            fk.child_column,
            fk.child_column
        );
    END LOOP;
END $$;

DO $$
DECLARE
    legacy_table RECORD;
BEGIN
    FOR legacy_table IN
        SELECT table_name
        FROM information_schema.columns
        WHERE table_schema = 'apeiron'
          AND table_name LIKE 'movement_action_%'
          AND column_name = 'contract_id'
    LOOP
        EXECUTE format(
            'ALTER TABLE apeiron.%I DROP CONSTRAINT IF EXISTS %I',
            legacy_table.table_name,
            legacy_table.table_name || '_contract_id_fkey'
        );
        EXECUTE format(
            'ALTER TABLE apeiron.%I ALTER COLUMN contract_id TYPE TEXT USING contract_id::text',
            legacy_table.table_name
        );
    END LOOP;
END $$;

ALTER TABLE IF EXISTS apeiron.movement_action_contract
    ALTER COLUMN id TYPE TEXT USING id::text,
    ALTER COLUMN reconciliation_contract_id TYPE TEXT USING reconciliation_contract_id::text,
    ADD COLUMN IF NOT EXISTS action_type TEXT NOT NULL DEFAULT 'grounded_skill',
    ADD COLUMN IF NOT EXISTS description TEXT NOT NULL DEFAULT '',
    ADD COLUMN IF NOT EXISTS duration_ms INT NOT NULL DEFAULT 1,
    ADD COLUMN IF NOT EXISTS active_ms INT NOT NULL DEFAULT 0,
    ADD COLUMN IF NOT EXISTS recovery_ms INT NOT NULL DEFAULT 0,
    ADD COLUMN IF NOT EXISTS distance_cm FLOAT NOT NULL DEFAULT 0.0,
    ADD COLUMN IF NOT EXISTS base_speed_cm_s FLOAT NOT NULL DEFAULT 0.0,
    ADD COLUMN IF NOT EXISTS yaw_degrees FLOAT NOT NULL DEFAULT 0.0,
    ADD COLUMN IF NOT EXISTS phase_window_policy TEXT NOT NULL DEFAULT 'grounded_action',
    ADD COLUMN IF NOT EXISTS prediction_error_policy TEXT NOT NULL DEFAULT 'bounded_smooth_correction',
    ADD COLUMN IF NOT EXISTS allow_windup_locomotion BOOLEAN NOT NULL DEFAULT FALSE,
    ADD COLUMN IF NOT EXISTS allow_active_locomotion BOOLEAN NOT NULL DEFAULT FALSE,
    ADD COLUMN IF NOT EXISTS allow_recovery_locomotion BOOLEAN NOT NULL DEFAULT FALSE,
    ADD COLUMN IF NOT EXISTS allow_yaw_adjustment BOOLEAN NOT NULL DEFAULT TRUE,
    ADD COLUMN IF NOT EXISTS root_motion_owner TEXT NOT NULL DEFAULT 'movement',
    ADD COLUMN IF NOT EXISTS contact_policy TEXT NOT NULL DEFAULT 'none',
    ADD COLUMN IF NOT EXISTS speed_curve JSONB NOT NULL DEFAULT '[]'::jsonb,
    ADD COLUMN IF NOT EXISTS vertical_curve JSONB NOT NULL DEFAULT '[]'::jsonb,
    ADD COLUMN IF NOT EXISTS metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    ADD COLUMN IF NOT EXISTS created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP NOT NULL DEFAULT NOW();

ALTER TABLE IF EXISTS apeiron.movement_action_contract
    DROP CONSTRAINT IF EXISTS fk_movement_action_reconciliation;

ALTER TABLE IF EXISTS apeiron.movement_action_contract
    ADD CONSTRAINT fk_movement_action_reconciliation
    FOREIGN KEY (reconciliation_contract_id)
    REFERENCES apeiron.movement_reconciliation_contract(id)
    NOT VALID;

DO $$
DECLARE
    legacy_table RECORD;
BEGIN
    FOR legacy_table IN
        SELECT table_name
        FROM information_schema.columns
        WHERE table_schema = 'apeiron'
          AND table_name LIKE 'movement_action_%'
          AND column_name = 'contract_id'
    LOOP
        BEGIN
            EXECUTE format(
                'ALTER TABLE apeiron.%I ADD CONSTRAINT %I FOREIGN KEY (contract_id) REFERENCES apeiron.movement_action_contract(id) NOT VALID',
                legacy_table.table_name,
                legacy_table.table_name || '_contract_id_fkey'
            );
        EXCEPTION
            WHEN duplicate_object THEN NULL;
        END;
    END LOOP;
END $$;

CREATE TABLE IF NOT EXISTS apeiron.skill_action_timing (
    skill_id TEXT PRIMARY KEY,
    windup_ms INT NOT NULL DEFAULT 0,
    active_ms INT NOT NULL DEFAULT 0,
    recovery_ms INT NOT NULL DEFAULT 0,
    cooldown_ms INT NOT NULL DEFAULT 0,
    combo_window_ms INT NOT NULL DEFAULT 0,
    movement_lock_policy TEXT NOT NULL DEFAULT 'contract',
    queue_policy TEXT NOT NULL DEFAULT 'none',
    cancel_policy TEXT NOT NULL DEFAULT 'none',
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_skill_action_timing_skill
        FOREIGN KEY (skill_id)
        REFERENCES apeiron.skill(id)
        ON DELETE CASCADE,
    CONSTRAINT chk_skill_action_timing CHECK (windup_ms >= 0 AND active_ms >= 0 AND recovery_ms >= 0 AND cooldown_ms >= 0)
);

ALTER TABLE IF EXISTS apeiron.skill_action_timing
    ADD COLUMN IF NOT EXISTS windup_ms INT NOT NULL DEFAULT 0,
    ADD COLUMN IF NOT EXISTS active_ms INT NOT NULL DEFAULT 0,
    ADD COLUMN IF NOT EXISTS recovery_ms INT NOT NULL DEFAULT 0,
    ADD COLUMN IF NOT EXISTS cooldown_ms INT NOT NULL DEFAULT 0,
    ADD COLUMN IF NOT EXISTS combo_window_ms INT NOT NULL DEFAULT 0,
    ADD COLUMN IF NOT EXISTS movement_lock_policy TEXT NOT NULL DEFAULT 'contract',
    ADD COLUMN IF NOT EXISTS queue_policy TEXT NOT NULL DEFAULT 'none',
    ADD COLUMN IF NOT EXISTS cancel_policy TEXT NOT NULL DEFAULT 'none',
    ADD COLUMN IF NOT EXISTS metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    ADD COLUMN IF NOT EXISTS created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP NOT NULL DEFAULT NOW();

CREATE TABLE IF NOT EXISTS apeiron.skill_movement_action_binding (
    skill_id TEXT PRIMARY KEY,
    movement_action_contract_id TEXT NOT NULL,
    starts_at_phase TEXT NOT NULL DEFAULT 'windup',
    handoff_policy TEXT NOT NULL DEFAULT 'explicit_recovery_handoff',
    normal_input_policy TEXT NOT NULL DEFAULT 'blocked_during_owned_root',
    target_policy TEXT NOT NULL DEFAULT 'aim_direction',
    contact_policy TEXT NOT NULL DEFAULT 'none',
    is_enabled BOOLEAN NOT NULL DEFAULT TRUE,
    metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_skill_movement_action_skill
        FOREIGN KEY (skill_id)
        REFERENCES apeiron.skill(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_skill_movement_action_contract
        FOREIGN KEY (movement_action_contract_id)
        REFERENCES apeiron.movement_action_contract(id)
);

ALTER TABLE IF EXISTS apeiron.skill_movement_action_binding
    ADD COLUMN IF NOT EXISTS movement_action_contract_id TEXT NOT NULL DEFAULT '';

ALTER TABLE IF EXISTS apeiron.skill_movement_action_binding
    ALTER COLUMN movement_action_contract_id TYPE TEXT USING movement_action_contract_id::text,
    ADD COLUMN IF NOT EXISTS starts_at_phase TEXT NOT NULL DEFAULT 'windup',
    ADD COLUMN IF NOT EXISTS handoff_policy TEXT NOT NULL DEFAULT 'explicit_recovery_handoff',
    ADD COLUMN IF NOT EXISTS normal_input_policy TEXT NOT NULL DEFAULT 'blocked_during_owned_root',
    ADD COLUMN IF NOT EXISTS target_policy TEXT NOT NULL DEFAULT 'aim_direction',
    ADD COLUMN IF NOT EXISTS contact_policy TEXT NOT NULL DEFAULT 'none',
    ADD COLUMN IF NOT EXISTS is_enabled BOOLEAN NOT NULL DEFAULT TRUE,
    ADD COLUMN IF NOT EXISTS metadata JSONB NOT NULL DEFAULT '{}'::jsonb,
    ADD COLUMN IF NOT EXISTS created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP NOT NULL DEFAULT NOW();

ALTER TABLE IF EXISTS apeiron.skill_movement_action_binding
    DROP CONSTRAINT IF EXISTS fk_skill_movement_action_contract;

ALTER TABLE IF EXISTS apeiron.skill_movement_action_binding
    ADD CONSTRAINT fk_skill_movement_action_contract
    FOREIGN KEY (movement_action_contract_id)
    REFERENCES apeiron.movement_action_contract(id)
    NOT VALID;

CREATE INDEX IF NOT EXISTS idx_movement_action_reconciliation
ON apeiron.movement_action_contract(reconciliation_contract_id);

CREATE INDEX IF NOT EXISTS idx_skill_movement_action_contract
ON apeiron.skill_movement_action_binding(movement_action_contract_id);
