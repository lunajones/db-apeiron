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

type RuntimeMovementReconciliationProfile struct {
	ProfileID                         string
	MaxSpeed                          float64
	SprintSpeedMultiplier             float64
	Acceleration                      float64
	Deceleration                      float64
	GroundFriction                    float64
	AirAcceleration                   float64
	JumpHeight                        float64
	JumpDurationMS                    int
	RotationRateYaw                   float64
	GravityScale                      float64
	BrakingFrictionFactor             float64
	MaxSlopeDeg                       float64
	StepHeight                        float64
	BaseDeadzone                      float64
	GroundedSpeedDeadzoneFactor       float64
	GroundedSpeedDeadzoneMin          float64
	GroundedSpeedDeadzoneMax          float64
	GroundedTransitionDeadzoneMin     float64
	MoveSustainDeadzone               float64
	MoveSustainTransitionDeadzone     float64
	AirborneDeadzone                  float64
	LeapRecentDeadzone                float64
	LeapAirborneSnapshotDeadzone      float64
	LeapLandingDeadzoneFactor         float64
	LeapLandingDeadzoneMin            float64
	LeapLandingDeadzoneMax            float64
	LeapLandingClampIgnoreDeadzone    float64
	LeapLandingSoftSnapDeadzone       float64
	DodgeRecentDeadzone               float64
	DodgeActiveDeadzone               float64
	DodgeExitDeadzoneFactor           float64
	DodgeExitDeadzoneMin              float64
	DodgeExitDeadzoneMax              float64
	PostActionGroundedDeadzone        float64
	CorrectionMaxStep                 float64
	HardSnapDistance                  float64
	SevereDesyncDistance              float64
	VisualSmoothingMS                 int
	VisualSmoothingMaxDistance        float64
	RemoteVisualInterpolationMS       int
	RemoteVisualMaxExtrapolationMS    int
	RemoteVisualHardSnapDistance      float64
	DodgeCarryHandoffMS               int
	LeapLandingCorrectionGraceMS      int
	LeapGroundedCarryHandoffMS        int
	MovementTurnResubmitDotThreshold  float64
	MovementTurnResubmitMinIntervalMS int
	MovementSubmitIntervalMS          int
	SnapshotPollIntervalMS            int
	StrafeSpeedMultiplier             float64
	BackpedalSpeedMultiplier          float64
	StrafeSprintSpeedMultiplier       float64
	BackpedalSprintSpeedMultiplier    float64
	MetadataJSON                      string
	CreatedAt                         time.Time
	UpdatedAt                         time.Time
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

type ActionOrientationPolicy struct {
	ID                         string
	OwnerKind                  string
	Description                string
	BodyYawSource              string
	FocusYawSource             string
	AttackYawSource            string
	BodyTurnRateDegS           float64
	FocusTurnRateDegS          float64
	AttackTurnRateDegS         float64
	CommitAlignMS              int
	AttackYawLatchPolicy       string
	AllowHeadLookWhileStrafing bool
	AllowBodySideOnMovement    bool
	MetadataJSON               string
	CreatedAt                  time.Time
	UpdatedAt                  time.Time
}

type ActionEnvelopePolicy struct {
	ID                       string
	OwnerKind                string
	Description              string
	PreCommitMS              int
	AirborneMS               int
	LandingInertiaMS         int
	PreCommitDirectionPolicy string
	AirborneDirectionPolicy  string
	InertiaDirectionPolicy   string
	TacticalReentryPolicy    string
	SpeedCurveJSON           string
	VerticalCurveJSON        string
	MetadataJSON             string
	CreatedAt                time.Time
	UpdatedAt                time.Time
}

type SkillActionPolicyBinding struct {
	SkillID                   string
	ActionOrientationPolicyID string
	ActionEnvelopePolicyID    string
	IsEnabled                 bool
	MetadataJSON              string
	ActionOrientationPolicy   ActionOrientationPolicy
	ActionEnvelopePolicy      ActionEnvelopePolicy
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
}

func getRuntimeMovementReconciliationProfileByID(ctx context.Context, db database.TxManager, id string) (RuntimeMovementReconciliationProfile, error) {
	var profile RuntimeMovementReconciliationProfile
	err := db.QueryRow(ctx, `
		SELECT
			profile_id,
			max_speed,
			sprint_speed_multiplier,
			acceleration,
			deceleration,
			ground_friction,
			air_acceleration,
			jump_height,
			jump_duration_ms,
			rotation_rate_yaw,
			gravity_scale,
			braking_friction_factor,
			max_slope_deg,
			step_height,
			base_deadzone,
			grounded_speed_deadzone_factor,
			grounded_speed_deadzone_min,
			grounded_speed_deadzone_max,
			grounded_transition_deadzone_min,
			move_sustain_deadzone,
			move_sustain_transition_deadzone,
			airborne_deadzone,
			leap_recent_deadzone,
			leap_airborne_snapshot_deadzone,
			leap_landing_deadzone_factor,
			leap_landing_deadzone_min,
			leap_landing_deadzone_max,
			leap_landing_clamp_ignore_deadzone,
			leap_landing_soft_snap_deadzone,
			dodge_recent_deadzone,
			dodge_active_deadzone,
			dodge_exit_deadzone_factor,
			dodge_exit_deadzone_min,
			dodge_exit_deadzone_max,
			post_action_grounded_deadzone,
			correction_max_step,
			hard_snap_distance,
			severe_desync_distance,
			visual_smoothing_ms,
			visual_smoothing_max_distance,
			remote_visual_interpolation_ms,
			remote_visual_max_extrapolation_ms,
			remote_visual_hard_snap_distance,
			dodge_carry_handoff_ms,
			leap_landing_correction_grace_ms,
			leap_grounded_carry_handoff_ms,
			movement_turn_resubmit_dot_threshold,
			movement_turn_resubmit_min_interval_ms,
			movement_submit_interval_ms,
			snapshot_poll_interval_ms,
			strafe_speed_multiplier,
			backpedal_speed_multiplier,
			strafe_sprint_speed_multiplier,
			backpedal_sprint_speed_multiplier,
			metadata::TEXT,
			created_at,
			updated_at
		FROM apeiron.runtime_movement_reconciliation_profile
		WHERE profile_id = $1
	`, id).Scan(
		&profile.ProfileID,
		&profile.MaxSpeed,
		&profile.SprintSpeedMultiplier,
		&profile.Acceleration,
		&profile.Deceleration,
		&profile.GroundFriction,
		&profile.AirAcceleration,
		&profile.JumpHeight,
		&profile.JumpDurationMS,
		&profile.RotationRateYaw,
		&profile.GravityScale,
		&profile.BrakingFrictionFactor,
		&profile.MaxSlopeDeg,
		&profile.StepHeight,
		&profile.BaseDeadzone,
		&profile.GroundedSpeedDeadzoneFactor,
		&profile.GroundedSpeedDeadzoneMin,
		&profile.GroundedSpeedDeadzoneMax,
		&profile.GroundedTransitionDeadzoneMin,
		&profile.MoveSustainDeadzone,
		&profile.MoveSustainTransitionDeadzone,
		&profile.AirborneDeadzone,
		&profile.LeapRecentDeadzone,
		&profile.LeapAirborneSnapshotDeadzone,
		&profile.LeapLandingDeadzoneFactor,
		&profile.LeapLandingDeadzoneMin,
		&profile.LeapLandingDeadzoneMax,
		&profile.LeapLandingClampIgnoreDeadzone,
		&profile.LeapLandingSoftSnapDeadzone,
		&profile.DodgeRecentDeadzone,
		&profile.DodgeActiveDeadzone,
		&profile.DodgeExitDeadzoneFactor,
		&profile.DodgeExitDeadzoneMin,
		&profile.DodgeExitDeadzoneMax,
		&profile.PostActionGroundedDeadzone,
		&profile.CorrectionMaxStep,
		&profile.HardSnapDistance,
		&profile.SevereDesyncDistance,
		&profile.VisualSmoothingMS,
		&profile.VisualSmoothingMaxDistance,
		&profile.RemoteVisualInterpolationMS,
		&profile.RemoteVisualMaxExtrapolationMS,
		&profile.RemoteVisualHardSnapDistance,
		&profile.DodgeCarryHandoffMS,
		&profile.LeapLandingCorrectionGraceMS,
		&profile.LeapGroundedCarryHandoffMS,
		&profile.MovementTurnResubmitDotThreshold,
		&profile.MovementTurnResubmitMinIntervalMS,
		&profile.MovementSubmitIntervalMS,
		&profile.SnapshotPollIntervalMS,
		&profile.StrafeSpeedMultiplier,
		&profile.BackpedalSpeedMultiplier,
		&profile.StrafeSprintSpeedMultiplier,
		&profile.BackpedalSprintSpeedMultiplier,
		&profile.MetadataJSON,
		&profile.CreatedAt,
		&profile.UpdatedAt,
	)
	return profile, err
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

func getActionOrientationPolicyByID(ctx context.Context, db database.TxManager, id string) (ActionOrientationPolicy, error) {
	var policy ActionOrientationPolicy
	err := db.QueryRow(ctx, `
		SELECT
			id,
			owner_kind,
			description,
			body_yaw_source,
			focus_yaw_source,
			attack_yaw_source,
			body_turn_rate_deg_s,
			focus_turn_rate_deg_s,
			attack_turn_rate_deg_s,
			commit_align_ms,
			attack_yaw_latch_policy,
			allow_head_look_while_strafing,
			allow_body_side_on_movement,
			metadata::TEXT,
			created_at,
			updated_at
		FROM apeiron.action_orientation_policy
		WHERE id = $1
	`, id).Scan(
		&policy.ID,
		&policy.OwnerKind,
		&policy.Description,
		&policy.BodyYawSource,
		&policy.FocusYawSource,
		&policy.AttackYawSource,
		&policy.BodyTurnRateDegS,
		&policy.FocusTurnRateDegS,
		&policy.AttackTurnRateDegS,
		&policy.CommitAlignMS,
		&policy.AttackYawLatchPolicy,
		&policy.AllowHeadLookWhileStrafing,
		&policy.AllowBodySideOnMovement,
		&policy.MetadataJSON,
		&policy.CreatedAt,
		&policy.UpdatedAt,
	)
	return policy, err
}

func getActionEnvelopePolicyByID(ctx context.Context, db database.TxManager, id string) (ActionEnvelopePolicy, error) {
	var policy ActionEnvelopePolicy
	err := db.QueryRow(ctx, `
		SELECT
			id,
			owner_kind,
			description,
			pre_commit_ms,
			airborne_ms,
			landing_inertia_ms,
			pre_commit_direction_policy,
			airborne_direction_policy,
			inertia_direction_policy,
			tactical_reentry_policy,
			speed_curve::TEXT,
			vertical_curve::TEXT,
			metadata::TEXT,
			created_at,
			updated_at
		FROM apeiron.action_envelope_policy
		WHERE id = $1
	`, id).Scan(
		&policy.ID,
		&policy.OwnerKind,
		&policy.Description,
		&policy.PreCommitMS,
		&policy.AirborneMS,
		&policy.LandingInertiaMS,
		&policy.PreCommitDirectionPolicy,
		&policy.AirborneDirectionPolicy,
		&policy.InertiaDirectionPolicy,
		&policy.TacticalReentryPolicy,
		&policy.SpeedCurveJSON,
		&policy.VerticalCurveJSON,
		&policy.MetadataJSON,
		&policy.CreatedAt,
		&policy.UpdatedAt,
	)
	return policy, err
}

func getSkillActionPolicyBindingBySkillID(ctx context.Context, db database.TxManager, skillID string) (SkillActionPolicyBinding, error) {
	var binding SkillActionPolicyBinding
	err := db.QueryRow(ctx, `
		SELECT
			skill_id,
			COALESCE(action_orientation_policy_id, ''),
			COALESCE(action_envelope_policy_id, ''),
			is_enabled,
			metadata::TEXT,
			created_at,
			updated_at
		FROM apeiron.skill_action_policy_binding
		WHERE skill_id = $1
		  AND is_enabled = TRUE
	`, skillID).Scan(
		&binding.SkillID,
		&binding.ActionOrientationPolicyID,
		&binding.ActionEnvelopePolicyID,
		&binding.IsEnabled,
		&binding.MetadataJSON,
		&binding.CreatedAt,
		&binding.UpdatedAt,
	)
	if err != nil {
		return binding, err
	}
	if binding.ActionOrientationPolicyID != "" {
		orientation, err := getActionOrientationPolicyByID(ctx, db, binding.ActionOrientationPolicyID)
		if err != nil {
			return binding, err
		}
		binding.ActionOrientationPolicy = orientation
	}
	if binding.ActionEnvelopePolicyID != "" {
		envelope, err := getActionEnvelopePolicyByID(ctx, db, binding.ActionEnvelopePolicyID)
		if err != nil {
			return binding, err
		}
		binding.ActionEnvelopePolicy = envelope
	}
	return binding, nil
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

func (r *ProfileRepository) GetRuntimeMovementReconciliationProfileByID(ctx context.Context, id string) (RuntimeMovementReconciliationProfile, error) {
	return getRuntimeMovementReconciliationProfileByID(ctx, r.db, id)
}

func (r *ProfileRepository) GetMovementReconciliationContractByID(ctx context.Context, id string) (MovementReconciliationContract, error) {
	return getMovementReconciliationContractByID(ctx, r.db, id)
}

func (r *ProfileRepository) GetMovementActionContractByID(ctx context.Context, id string) (MovementActionContract, error) {
	return getMovementActionContractByID(ctx, r.db, id)
}

func (r *ProfileRepository) GetActionOrientationPolicyByID(ctx context.Context, id string) (ActionOrientationPolicy, error) {
	return getActionOrientationPolicyByID(ctx, r.db, id)
}

func (r *ProfileRepository) GetActionEnvelopePolicyByID(ctx context.Context, id string) (ActionEnvelopePolicy, error) {
	return getActionEnvelopePolicyByID(ctx, r.db, id)
}

func (r *ProfileRepository) GetSkillActionPolicyBindingBySkillID(ctx context.Context, skillID string) (SkillActionPolicyBinding, error) {
	return getSkillActionPolicyBindingBySkillID(ctx, r.db, skillID)
}
