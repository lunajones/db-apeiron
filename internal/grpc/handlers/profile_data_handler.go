package handlers

import (
	"context"
	"encoding/json"

	apeironv1 "db-apeiron/gen/apeiron/v1"
	"db-apeiron/internal/repository/postgres"
)

type ProfileReader interface {
	GetMovementProfile(ctx context.Context, id string) (postgres.MovementProfile, error)
	GetCombatCoreProfile(ctx context.Context, id string) (postgres.CombatCoreProfile, error)
	GetCombatDefenseContract(ctx context.Context, id string) (postgres.CombatDefenseContract, error)
	GetMovementActionContract(ctx context.Context, id string) (postgres.MovementActionContract, error)
	GetMovementReconciliationContract(ctx context.Context, id string) (postgres.MovementReconciliationContract, error)
	GetRuntimeMovementReconciliationProfile(ctx context.Context, id string) (postgres.RuntimeMovementReconciliationProfile, error)
	GetCreatureBehaviorRuntimeContract(ctx context.Context, id string) (postgres.CreatureBehaviorRuntimeContract, error)
	GetCreatureEvasionPolicies(ctx context.Context, behaviorContractID string) ([]postgres.CreatureEvasionPolicy, error)
	GetCreatureSkillSetupPolicies(ctx context.Context, behaviorContractID string) ([]postgres.CreatureSkillSetupPolicy, error)
	GetCreatureTargetOpportunityPolicy(ctx context.Context, id string) (postgres.CreatureTargetOpportunityPolicy, error)
	GetCreatureOrbitPolicy(ctx context.Context, id string) (postgres.CreatureOrbitPolicy, error)
	GetCreatureSkillBehaviorBindings(ctx context.Context, behaviorContractID string) ([]postgres.CreatureSkillBehaviorBinding, error)
}

type ProfileDataHandler struct {
	apeironv1.UnimplementedProfileDataServiceServer

	profiles ProfileReader
}

func NewProfileDataHandler(profiles ProfileReader) *ProfileDataHandler {
	return &ProfileDataHandler{profiles: profiles}
}

func (h *ProfileDataHandler) GetMovementProfile(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.MovementProfileResponse, error) {
	profile, err := h.profiles.GetMovementProfile(ctx, req.GetId())
	if err != nil {
		return &apeironv1.MovementProfileResponse{Found: false, Error: err.Error()}, nil
	}

	return &apeironv1.MovementProfileResponse{Found: true, Profile: mapMovementProfile(profile)}, nil
}

func (h *ProfileDataHandler) GetCombatCoreProfile(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.CombatCoreProfileResponse, error) {
	profile, err := h.profiles.GetCombatCoreProfile(ctx, req.GetId())
	if err != nil {
		return &apeironv1.CombatCoreProfileResponse{Found: false, Error: err.Error()}, nil
	}

	return &apeironv1.CombatCoreProfileResponse{Found: true, Profile: mapCombatCoreProfile(profile)}, nil
}

func (h *ProfileDataHandler) GetCombatDefenseContract(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.CombatDefenseContractResponse, error) {
	contract, err := h.profiles.GetCombatDefenseContract(ctx, req.GetId())
	if err != nil {
		return &apeironv1.CombatDefenseContractResponse{Found: false, Error: err.Error()}, nil
	}

	return &apeironv1.CombatDefenseContractResponse{Found: true, Contract: mapCombatDefenseContract(contract)}, nil
}

func (h *ProfileDataHandler) GetMovementActionContract(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.MovementActionContractResponse, error) {
	contract, err := h.profiles.GetMovementActionContract(ctx, req.GetId())
	if err != nil {
		return &apeironv1.MovementActionContractResponse{Found: false, Error: err.Error()}, nil
	}

	return &apeironv1.MovementActionContractResponse{Found: true, Contract: mapMovementActionContract(contract)}, nil
}

func (h *ProfileDataHandler) GetMovementReconciliationContract(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.MovementReconciliationContractResponse, error) {
	contract, err := h.profiles.GetMovementReconciliationContract(ctx, req.GetId())
	if err != nil {
		return &apeironv1.MovementReconciliationContractResponse{Found: false, Error: err.Error()}, nil
	}

	return &apeironv1.MovementReconciliationContractResponse{Found: true, Contract: mapMovementReconciliationContract(contract)}, nil
}

func (h *ProfileDataHandler) GetRuntimeMovementReconciliationProfile(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.RuntimeMovementReconciliationProfileResponse, error) {
	profile, err := h.profiles.GetRuntimeMovementReconciliationProfile(ctx, req.GetId())
	if err != nil {
		return &apeironv1.RuntimeMovementReconciliationProfileResponse{Found: false, Error: err.Error()}, nil
	}

	return &apeironv1.RuntimeMovementReconciliationProfileResponse{Found: true, Profile: mapRuntimeMovementReconciliationProfile(profile)}, nil
}

func (h *ProfileDataHandler) GetCreatureBehaviorRuntimeContract(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.CreatureBehaviorRuntimeContractResponse, error) {
	contract, err := h.profiles.GetCreatureBehaviorRuntimeContract(ctx, req.GetId())
	if err != nil {
		return &apeironv1.CreatureBehaviorRuntimeContractResponse{Found: false, Error: err.Error()}, nil
	}

	return &apeironv1.CreatureBehaviorRuntimeContractResponse{Found: true, Contract: mapCreatureBehaviorRuntimeContract(contract)}, nil
}

func (h *ProfileDataHandler) GetCreatureEvasionPolicies(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.CreatureEvasionPoliciesResponse, error) {
	policies, err := h.profiles.GetCreatureEvasionPolicies(ctx, req.GetId())
	if err != nil {
		return &apeironv1.CreatureEvasionPoliciesResponse{Found: false, Error: err.Error()}, nil
	}

	out := make([]*apeironv1.CreatureEvasionPolicy, 0, len(policies))
	for _, policy := range policies {
		out = append(out, mapCreatureEvasionPolicy(policy))
	}
	return &apeironv1.CreatureEvasionPoliciesResponse{Found: true, Policies: out}, nil
}

func (h *ProfileDataHandler) GetCreatureSkillSetupPolicies(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.CreatureSkillSetupPoliciesResponse, error) {
	policies, err := h.profiles.GetCreatureSkillSetupPolicies(ctx, req.GetId())
	if err != nil {
		return &apeironv1.CreatureSkillSetupPoliciesResponse{Found: false, Error: err.Error()}, nil
	}

	out := make([]*apeironv1.CreatureSkillSetupPolicy, 0, len(policies))
	for _, policy := range policies {
		out = append(out, mapCreatureSkillSetupPolicy(policy))
	}
	return &apeironv1.CreatureSkillSetupPoliciesResponse{Found: true, Policies: out}, nil
}

func (h *ProfileDataHandler) GetCreatureTargetOpportunityPolicy(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.CreatureTargetOpportunityPolicyResponse, error) {
	policy, err := h.profiles.GetCreatureTargetOpportunityPolicy(ctx, req.GetId())
	if err != nil {
		return &apeironv1.CreatureTargetOpportunityPolicyResponse{Found: false, Error: err.Error()}, nil
	}

	return &apeironv1.CreatureTargetOpportunityPolicyResponse{Found: true, Policy: mapCreatureTargetOpportunityPolicy(policy)}, nil
}

func (h *ProfileDataHandler) GetCreatureOrbitPolicy(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.CreatureOrbitPolicyResponse, error) {
	policy, err := h.profiles.GetCreatureOrbitPolicy(ctx, req.GetId())
	if err != nil {
		return &apeironv1.CreatureOrbitPolicyResponse{Found: false, Error: err.Error()}, nil
	}

	return &apeironv1.CreatureOrbitPolicyResponse{Found: true, Policy: mapCreatureOrbitPolicy(policy)}, nil
}

func (h *ProfileDataHandler) GetCreatureSkillBehaviorBindings(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.CreatureSkillBehaviorBindingsResponse, error) {
	bindings, err := h.profiles.GetCreatureSkillBehaviorBindings(ctx, req.GetId())
	if err != nil {
		return &apeironv1.CreatureSkillBehaviorBindingsResponse{Found: false, Error: err.Error()}, nil
	}

	out := make([]*apeironv1.CreatureSkillBehaviorBinding, 0, len(bindings))
	for _, binding := range bindings {
		out = append(out, mapCreatureSkillBehaviorBinding(binding))
	}
	return &apeironv1.CreatureSkillBehaviorBindingsResponse{Found: true, Bindings: out}, nil
}

func mapMovementProfile(p postgres.MovementProfile) *apeironv1.MovementProfile {
	return &apeironv1.MovementProfile{
		Id:                p.ID,
		MaxSpeed:          p.MaxSpeed,
		Acceleration:      p.Acceleration,
		Deceleration:      p.Deceleration,
		Friction:          p.Friction,
		GravityMultiplier: p.GravityMultiplier,
		Mass:              p.Mass,
		MomentumRetention: p.MomentumRetention,
		TurnRate:          p.TurnRate,
		AirControl:        p.AirControl,
		StrafeEfficiency:  p.StrafeEfficiency,
		BackpedalPenalty:  p.BackpedalPenalty,
		DodgeDistance:     p.DodgeDistance,
		DodgeDurationMs:   int32(p.DodgeDurationMS),
		SprintMultiplier:  p.SprintMultiplier,
		SlopeLimit:        p.SlopeLimit,
		SlideOnSlope:      p.SlideOnSlope,
		IsAirborneEnabled: p.IsAirborneEnabled,
	}
}

func mapCombatCoreProfile(p postgres.CombatCoreProfile) *apeironv1.CombatCoreProfile {
	return &apeironv1.CombatCoreProfile{
		DamageDealtMultiplier:      p.DamageDealtMultiplier,
		CanBlock:                   p.CanBlock,
		BlockDamageReduction:       p.BlockDamageReduction,
		MaxPosture:                 p.MaxPosture,
		PostureDamageMultiplier:    p.PostureDamageMultiplier,
		PostureBreakDurationMs:     int32(p.PostureBreakDurationMS),
		DamageTakenMultiplier:      p.DamageTakenMultiplier,
		CanParry:                   p.CanParry,
		ParryWindowMs:              int32(p.ParryWindowMS),
		ParryRewardMultiplier:      p.ParryRewardMultiplier,
		DodgeIframeMs:              int32(p.DodgeIframeMS),
		MaxStamina:                 p.MaxStamina,
		StaminaRegenPerSec:         p.StaminaRegenPerSec,
		DodgeStaminaCost:           p.DodgeStaminaCost,
		SprintStaminaCostPerSec:    p.SprintStaminaCostPerSec,
		BlockStaminaCostPerSec:     p.BlockStaminaCostPerSec,
		AttackStaminaCost:          p.AttackStaminaCost,
		StaminaExhaustionThreshold: p.StaminaExhaustionThreshold,
		StaminaZeroRegenMultiplier: p.StaminaZeroRegenMultiplier,
	}
}

func mapCombatDefenseContract(c postgres.CombatDefenseContract) *apeironv1.CombatDefenseContract {
	return &apeironv1.CombatDefenseContract{
		Id:                         c.ID,
		Name:                       c.Name,
		Description:                c.Description,
		DefenseType:                c.DefenseType,
		FrontalArcDeg:              c.FrontalArcDeg,
		DefenderMarginLeftRatio:    c.DefenderMarginLeftRatio,
		DefenderMarginRightRatio:   c.DefenderMarginRightRatio,
		StaminaDamageOnlyOnBlock:   c.StaminaDamageOnlyOnBlock,
		HealthDamageOnUnblockedHit: c.HealthDamageOnUnblockedHit,
		PostureDamageOnBlock:       c.PostureDamageOnBlock,
		PerfectBlockWindowMs:       int32(c.PerfectBlockWindowMS),
		ParryWindowMs:              int32(c.ParryWindowMS),
		GuardDamageMultiplier:      c.GuardDamageMultiplier,
		BlockStaminaDrainPerSecond: c.BlockStaminaDrainPerSecond,
		MetadataJson:               c.MetadataJSON,
	}
}

func mapMovementActionContract(c postgres.MovementActionContract) *apeironv1.MovementActionContract {
	return &apeironv1.MovementActionContract{
		Id:                       c.ID,
		ActionType:               c.ActionType,
		Description:              c.Description,
		DurationMs:               int32(c.DurationMS),
		ActiveMs:                 int32(c.ActiveMS),
		RecoveryMs:               int32(c.RecoveryMS),
		DistanceCm:               c.DistanceCM,
		BaseSpeedCmS:             c.BaseSpeedCMS,
		YawDegrees:               c.YawDegrees,
		PhaseWindowPolicy:        c.PhaseWindowPolicy,
		PredictionErrorPolicy:    c.PredictionErrorPolicy,
		ReconciliationContractId: c.ReconciliationContractID,
		AllowWindupLocomotion:    c.AllowWindupLocomotion,
		AllowActiveLocomotion:    c.AllowActiveLocomotion,
		AllowRecoveryLocomotion:  c.AllowRecoveryLocomotion,
		AllowYawAdjustment:       c.AllowYawAdjustment,
		RootMotionOwner:          c.RootMotionOwner,
		ContactPolicy:            c.ContactPolicy,
		SpeedCurve:               mapMovementCurveSamples(c.SpeedCurveJSON),
		VerticalCurve:            mapMovementCurveSamples(c.VerticalCurveJSON),
		MetadataJson:             c.MetadataJSON,
		ReconciliationContract:   mapMovementReconciliationContract(c.ReconciliationContract),
	}
}

func mapMovementReconciliationContract(c postgres.MovementReconciliationContract) *apeironv1.MovementReconciliationContract {
	if c.ID == "" {
		return nil
	}
	return &apeironv1.MovementReconciliationContract{
		Id:                     c.ID,
		Category:               c.Category,
		Description:            c.Description,
		MaxSmoothErrorCm:       c.MaxSmoothErrorCM,
		HardSnapErrorCm:        c.HardSnapErrorCM,
		SmoothingTimeMs:        int32(c.SmoothingTimeMS),
		YawToleranceDeg:        c.YawToleranceDeg,
		OwnsPosition:           c.OwnsPosition,
		OwnsYaw:                c.OwnsYaw,
		AllowsClientPrediction: c.AllowsClientPrediction,
		InputPolicy:            c.InputPolicy,
		HandoffPolicy:          c.HandoffPolicy,
		MetadataJson:           c.MetadataJSON,
	}
}

func mapRuntimeMovementReconciliationProfile(p postgres.RuntimeMovementReconciliationProfile) *apeironv1.RuntimeMovementReconciliationProfile {
	if p.ProfileID == "" {
		return nil
	}
	return &apeironv1.RuntimeMovementReconciliationProfile{
		ProfileId:                         p.ProfileID,
		MaxSpeed:                          p.MaxSpeed,
		SprintSpeedMultiplier:             p.SprintSpeedMultiplier,
		Acceleration:                      p.Acceleration,
		Deceleration:                      p.Deceleration,
		GroundFriction:                    p.GroundFriction,
		AirAcceleration:                   p.AirAcceleration,
		JumpHeight:                        p.JumpHeight,
		JumpDurationMs:                    int32(p.JumpDurationMS),
		RotationRateYaw:                   p.RotationRateYaw,
		GravityScale:                      p.GravityScale,
		BrakingFrictionFactor:             p.BrakingFrictionFactor,
		MaxSlopeDeg:                       p.MaxSlopeDeg,
		StepHeight:                        p.StepHeight,
		BaseDeadzone:                      p.BaseDeadzone,
		GroundedSpeedDeadzoneFactor:       p.GroundedSpeedDeadzoneFactor,
		GroundedSpeedDeadzoneMin:          p.GroundedSpeedDeadzoneMin,
		GroundedSpeedDeadzoneMax:          p.GroundedSpeedDeadzoneMax,
		GroundedTransitionDeadzoneMin:     p.GroundedTransitionDeadzoneMin,
		MoveSustainDeadzone:               p.MoveSustainDeadzone,
		MoveSustainTransitionDeadzone:     p.MoveSustainTransitionDeadzone,
		AirborneDeadzone:                  p.AirborneDeadzone,
		LeapRecentDeadzone:                p.LeapRecentDeadzone,
		LeapAirborneSnapshotDeadzone:      p.LeapAirborneSnapshotDeadzone,
		LeapLandingDeadzoneFactor:         p.LeapLandingDeadzoneFactor,
		LeapLandingDeadzoneMin:            p.LeapLandingDeadzoneMin,
		LeapLandingDeadzoneMax:            p.LeapLandingDeadzoneMax,
		LeapLandingClampIgnoreDeadzone:    p.LeapLandingClampIgnoreDeadzone,
		LeapLandingSoftSnapDeadzone:       p.LeapLandingSoftSnapDeadzone,
		DodgeRecentDeadzone:               p.DodgeRecentDeadzone,
		DodgeActiveDeadzone:               p.DodgeActiveDeadzone,
		DodgeExitDeadzoneFactor:           p.DodgeExitDeadzoneFactor,
		DodgeExitDeadzoneMin:              p.DodgeExitDeadzoneMin,
		DodgeExitDeadzoneMax:              p.DodgeExitDeadzoneMax,
		PostActionGroundedDeadzone:        p.PostActionGroundedDeadzone,
		CorrectionMaxStep:                 p.CorrectionMaxStep,
		HardSnapDistance:                  p.HardSnapDistance,
		SevereDesyncDistance:              p.SevereDesyncDistance,
		VisualSmoothingMs:                 int32(p.VisualSmoothingMS),
		VisualSmoothingMaxDistance:        p.VisualSmoothingMaxDistance,
		RemoteVisualInterpolationMs:       int32(p.RemoteVisualInterpolationMS),
		RemoteVisualMaxExtrapolationMs:    int32(p.RemoteVisualMaxExtrapolationMS),
		RemoteVisualHardSnapDistance:      p.RemoteVisualHardSnapDistance,
		DodgeCarryHandoffMs:               int32(p.DodgeCarryHandoffMS),
		LeapLandingCorrectionGraceMs:      int32(p.LeapLandingCorrectionGraceMS),
		LeapGroundedCarryHandoffMs:        int32(p.LeapGroundedCarryHandoffMS),
		MovementTurnResubmitDotThreshold:  p.MovementTurnResubmitDotThreshold,
		MovementTurnResubmitMinIntervalMs: int32(p.MovementTurnResubmitMinIntervalMS),
		MovementSubmitIntervalMs:          int32(p.MovementSubmitIntervalMS),
		SnapshotPollIntervalMs:            int32(p.SnapshotPollIntervalMS),
		StrafeSpeedMultiplier:             p.StrafeSpeedMultiplier,
		BackpedalSpeedMultiplier:          p.BackpedalSpeedMultiplier,
		StrafeSprintSpeedMultiplier:       p.StrafeSprintSpeedMultiplier,
		BackpedalSprintSpeedMultiplier:    p.BackpedalSprintSpeedMultiplier,
		MetadataJson:                      p.MetadataJSON,
	}
}

func mapCreatureBehaviorRuntimeContract(c postgres.CreatureBehaviorRuntimeContract) *apeironv1.CreatureBehaviorRuntimeContract {
	return &apeironv1.CreatureBehaviorRuntimeContract{
		Id:                        c.ID,
		CreatureTemplateId:        c.CreatureTemplateID,
		Description:               c.Description,
		AggressionCurveJson:       c.AggressionCurveJSON,
		RangePolicyJson:           c.RangePolicyJSON,
		OrbitPolicyJson:           c.OrbitPolicyJSON,
		PressurePolicyJson:        c.PressurePolicyJSON,
		StaminaPolicyJson:         c.StaminaPolicyJSON,
		MetadataJson:              c.MetadataJSON,
		TargetOpportunityPolicyId: c.TargetOpportunityPolicyID,
		OrbitPolicyId:             c.OrbitPolicyID,
	}
}

func mapCreatureEvasionPolicy(p postgres.CreatureEvasionPolicy) *apeironv1.CreatureEvasionPolicy {
	return &apeironv1.CreatureEvasionPolicy{
		Id:                      p.ID,
		BehaviorContractId:      p.BehaviorContractID,
		Description:             p.Description,
		DodgeSkillId:            p.DodgeSkillID,
		MaxChainCount:           int32(p.MaxChainCount),
		StaminaCostMultiplier:   p.StaminaCostMultiplier,
		RetreatChanceMultiplier: p.RetreatChanceMultiplier,
		LateralBias:             p.LateralBias,
		BackstepBias:            p.BackstepBias,
		PressureThreshold:       p.PressureThreshold,
		CooldownMs:              int32(p.CooldownMS),
		MetadataJson:            p.MetadataJSON,
	}
}

func mapCreatureSkillSetupPolicy(p postgres.CreatureSkillSetupPolicy) *apeironv1.CreatureSkillSetupPolicy {
	return &apeironv1.CreatureSkillSetupPolicy{
		Id:                  p.ID,
		BehaviorContractId:  p.BehaviorContractID,
		SkillId:             p.SkillID,
		SetupType:           p.SetupType,
		MinSetupMs:          int32(p.MinSetupMS),
		MaxSetupMs:          int32(p.MaxSetupMS),
		CommitDistanceCm:    p.CommitDistanceCM,
		PreferredMinRangeCm: p.PreferredMinRangeCM,
		PreferredMaxRangeCm: p.PreferredMaxRangeCM,
		MovementTactic:      p.MovementTactic,
		LockSideDuringSetup: p.LockSideDuringSetup,
		IsEnabled:           p.IsEnabled,
		MetadataJson:        p.MetadataJSON,
	}
}

func mapCreatureTargetOpportunityPolicy(p postgres.CreatureTargetOpportunityPolicy) *apeironv1.CreatureTargetOpportunityPolicy {
	return &apeironv1.CreatureTargetOpportunityPolicy{
		Id:                          p.ID,
		Description:                 p.Description,
		CommitAngleMaxDeg:           p.CommitAngleMaxDeg,
		MinCommitDistanceCm:         p.MinCommitDistanceCM,
		MaxCommitDistanceCm:         p.MaxCommitDistanceCM,
		ApproachMinDistanceCm:       p.ApproachMinDistanceCM,
		ApproachMaxDistanceCm:       p.ApproachMaxDistanceCM,
		BiteRangeCm:                 p.BiteRangeCM,
		LungeMinRangeCm:             p.LungeMinRangeCM,
		LungeMaxRangeCm:             p.LungeMaxRangeCM,
		MaulPressureThreshold:       p.MaulPressureThreshold,
		TargetMemoryMs:              int32(p.TargetMemoryMS),
		NoReadySkillMemoryPolicy:    p.NoReadySkillMemoryPolicy,
		CandidateCooldownVisibility: p.CandidateCooldownVisibility,
		AllowBacksideCommit:         p.AllowBacksideCommit,
		MetadataJson:                p.MetadataJSON,
	}
}

func mapCreatureOrbitPolicy(p postgres.CreatureOrbitPolicy) *apeironv1.CreatureOrbitPolicy {
	return &apeironv1.CreatureOrbitPolicy{
		Id:                             p.ID,
		BehaviorContractId:             p.BehaviorContractID,
		Description:                    p.Description,
		OrbitLocomotionMode:            p.OrbitLocomotionMode,
		OrbitSpeedScale:                p.OrbitSpeedScale,
		MinOrbitDurationMs:             int32(p.MinOrbitDurationMS),
		SideSwitchCooldownMs:           int32(p.SideSwitchCooldownMS),
		AllowSideSwitchWhenTargetFaces: p.AllowSideSwitchWhenTargetFaces,
		PreferLongSideCommit:           p.PreferLongSideCommit,
		SideFlipChanceMultiplier:       p.SideFlipChanceMultiplier,
		LockSideDuringSetup:            p.LockSideDuringSetup,
		MetadataJson:                   p.MetadataJSON,
	}
}

func mapCreatureSkillBehaviorBinding(p postgres.CreatureSkillBehaviorBinding) *apeironv1.CreatureSkillBehaviorBinding {
	return &apeironv1.CreatureSkillBehaviorBinding{
		Id:                  p.ID,
		BehaviorContractId:  p.BehaviorContractID,
		SkillId:             p.SkillID,
		TacticalState:       p.TacticalState,
		DecisionPhase:       p.DecisionPhase,
		SetupPolicyId:       p.SetupPolicyID,
		MinRangeCm:          p.MinRangeCM,
		MaxRangeCm:          p.MaxRangeCM,
		Priority:            int32(p.Priority),
		UsageWeight:         p.UsageWeight,
		CooldownGroup:       p.CooldownGroup,
		RequiresLineOfSight: p.RequiresLineOfSight,
		IsEnabled:           p.IsEnabled,
		MetadataJson:        p.MetadataJSON,
	}
}

func mapMovementCurveSamples(raw string) []*apeironv1.MovementCurveSample {
	if raw == "" || raw == "null" {
		return nil
	}

	var samples []struct {
		T     float64 `json:"t"`
		Value float64 `json:"value"`
		V     float64 `json:"v"`
		Z     float64 `json:"z"`
	}
	if err := json.Unmarshal([]byte(raw), &samples); err != nil {
		return nil
	}

	out := make([]*apeironv1.MovementCurveSample, 0, len(samples))
	for _, sample := range samples {
		value := sample.Value
		if value == 0 {
			if sample.V != 0 {
				value = sample.V
			} else if sample.Z != 0 {
				value = sample.Z
			}
		}
		out = append(out, &apeironv1.MovementCurveSample{
			T:     sample.T,
			Value: value,
		})
	}
	return out
}
