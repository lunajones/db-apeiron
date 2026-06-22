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
	GetMovementActionContract(ctx context.Context, id string) (postgres.MovementActionContract, error)
	GetMovementReconciliationContract(ctx context.Context, id string) (postgres.MovementReconciliationContract, error)
	GetCreatureBehaviorRuntimeContract(ctx context.Context, id string) (postgres.CreatureBehaviorRuntimeContract, error)
	GetCreatureEvasionPolicies(ctx context.Context, behaviorContractID string) ([]postgres.CreatureEvasionPolicy, error)
	GetCreatureSkillSetupPolicies(ctx context.Context, behaviorContractID string) ([]postgres.CreatureSkillSetupPolicy, error)
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
		DamageDealtMultiplier:   p.DamageDealtMultiplier,
		CanBlock:                p.CanBlock,
		BlockDamageReduction:    p.BlockDamageReduction,
		MaxPosture:              p.MaxPosture,
		PostureDamageMultiplier: p.PostureDamageMultiplier,
		PostureBreakDurationMs:  int32(p.PostureBreakDurationMS),
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

func mapCreatureBehaviorRuntimeContract(c postgres.CreatureBehaviorRuntimeContract) *apeironv1.CreatureBehaviorRuntimeContract {
	return &apeironv1.CreatureBehaviorRuntimeContract{
		Id:                  c.ID,
		CreatureTemplateId:  c.CreatureTemplateID,
		Description:         c.Description,
		AggressionCurveJson: c.AggressionCurveJSON,
		RangePolicyJson:     c.RangePolicyJSON,
		OrbitPolicyJson:     c.OrbitPolicyJSON,
		PressurePolicyJson:  c.PressurePolicyJSON,
		StaminaPolicyJson:   c.StaminaPolicyJSON,
		MetadataJson:        c.MetadataJSON,
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
