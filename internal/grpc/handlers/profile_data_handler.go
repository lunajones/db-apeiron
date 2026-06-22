package handlers

import (
	"context"

	apeironv1 "db-apeiron/gen/apeiron/v1"
	"db-apeiron/internal/repository/postgres"
)

type ProfileReader interface {
	GetMovementProfile(ctx context.Context, id string) (postgres.MovementProfile, error)
	GetCombatCoreProfile(ctx context.Context, id string) (postgres.CombatCoreProfile, error)
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
