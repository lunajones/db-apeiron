package postgres

import (
	"context"
	"time"

	"db-apeiron/internal/database"
)

type MovementReconciliationContract struct {
	ID                     string
	Category               string
	Description            string
	MaxSmoothErrorCM       float64
	HardSnapErrorCM        float64
	SmoothingTimeMS        int
	YawToleranceDeg        float64
	OwnsPosition           bool
	OwnsYaw                bool
	AllowsClientPrediction bool
	InputPolicy            string
	HandoffPolicy          string
	MetadataJSON           string
	CreatedAt              time.Time
	UpdatedAt              time.Time
}

type MovementActionContract struct {
	ID                       string
	ActionType               string
	Description              string
	DurationMS               int
	ActiveMS                 int
	RecoveryMS               int
	DistanceCM               float64
	BaseSpeedCMS             float64
	YawDegrees               float64
	PhaseWindowPolicy        string
	PredictionErrorPolicy    string
	ReconciliationContractID string
	AllowWindupLocomotion    bool
	AllowActiveLocomotion    bool
	AllowRecoveryLocomotion  bool
	AllowYawAdjustment       bool
	RootMotionOwner          string
	ContactPolicy            string
	SpeedCurveJSON           string
	VerticalCurveJSON        string
	MetadataJSON             string
	ReconciliationContract   MovementReconciliationContract
	CreatedAt                time.Time
	UpdatedAt                time.Time
}

func getMovementReconciliationContractByID(ctx context.Context, db database.TxManager, id string) (MovementReconciliationContract, error) {
	var contract MovementReconciliationContract
	err := db.QueryRow(ctx, `
		SELECT
			id,
			category,
			description,
			max_smooth_error_cm,
			hard_snap_error_cm,
			smoothing_time_ms,
			yaw_tolerance_deg,
			owns_position,
			owns_yaw,
			allows_client_prediction,
			input_policy,
			handoff_policy,
			metadata::TEXT,
			created_at,
			updated_at
		FROM apeiron.movement_reconciliation_contract
		WHERE id = $1
	`, id).Scan(
		&contract.ID,
		&contract.Category,
		&contract.Description,
		&contract.MaxSmoothErrorCM,
		&contract.HardSnapErrorCM,
		&contract.SmoothingTimeMS,
		&contract.YawToleranceDeg,
		&contract.OwnsPosition,
		&contract.OwnsYaw,
		&contract.AllowsClientPrediction,
		&contract.InputPolicy,
		&contract.HandoffPolicy,
		&contract.MetadataJSON,
		&contract.CreatedAt,
		&contract.UpdatedAt,
	)
	return contract, err
}

func getMovementActionContractByID(ctx context.Context, db database.TxManager, id string) (MovementActionContract, error) {
	var contract MovementActionContract
	err := db.QueryRow(ctx, `
		SELECT
			id,
			action_type,
			description,
			duration_ms,
			active_ms,
			recovery_ms,
			distance_cm,
			base_speed_cm_s,
			yaw_degrees,
			phase_window_policy,
			prediction_error_policy,
			COALESCE(reconciliation_contract_id, ''),
			allow_windup_locomotion,
			allow_active_locomotion,
			allow_recovery_locomotion,
			allow_yaw_adjustment,
			root_motion_owner,
			contact_policy,
			speed_curve::TEXT,
			vertical_curve::TEXT,
			metadata::TEXT,
			created_at,
			updated_at
		FROM apeiron.movement_action_contract
		WHERE id = $1
	`, id).Scan(
		&contract.ID,
		&contract.ActionType,
		&contract.Description,
		&contract.DurationMS,
		&contract.ActiveMS,
		&contract.RecoveryMS,
		&contract.DistanceCM,
		&contract.BaseSpeedCMS,
		&contract.YawDegrees,
		&contract.PhaseWindowPolicy,
		&contract.PredictionErrorPolicy,
		&contract.ReconciliationContractID,
		&contract.AllowWindupLocomotion,
		&contract.AllowActiveLocomotion,
		&contract.AllowRecoveryLocomotion,
		&contract.AllowYawAdjustment,
		&contract.RootMotionOwner,
		&contract.ContactPolicy,
		&contract.SpeedCurveJSON,
		&contract.VerticalCurveJSON,
		&contract.MetadataJSON,
		&contract.CreatedAt,
		&contract.UpdatedAt,
	)
	if err != nil {
		return contract, err
	}
	if contract.ReconciliationContractID != "" {
		reconciliation, err := getMovementReconciliationContractByID(ctx, db, contract.ReconciliationContractID)
		if err != nil {
			return contract, err
		}
		contract.ReconciliationContract = reconciliation
	}
	return contract, nil
}

func (r *ProfileRepository) GetMovementReconciliationContractByID(ctx context.Context, id string) (MovementReconciliationContract, error) {
	return getMovementReconciliationContractByID(ctx, r.db, id)
}

func (r *ProfileRepository) GetMovementActionContractByID(ctx context.Context, id string) (MovementActionContract, error) {
	return getMovementActionContractByID(ctx, r.db, id)
}
