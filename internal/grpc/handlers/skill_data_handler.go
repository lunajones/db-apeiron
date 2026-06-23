package handlers

import (
	"context"

	apeironv1 "db-apeiron/gen/apeiron/v1"
	"db-apeiron/internal/repository/postgres"
)

type SkillReader interface {
	GetSkill(ctx context.Context, id string) (postgres.Skill, error)
	GetSkillSet(ctx context.Context, id string) (postgres.SkillSet, error)
	GetSkillSetLoadout(ctx context.Context, skillSetID string) ([]postgres.SkillLoadoutItem, error)
	GetWeaponCombatModeSlots(ctx context.Context, weaponKitID string) ([]postgres.WeaponCombatModeSlot, error)
	GetMovementEffect(ctx context.Context, skillID string) (postgres.SkillMovementEffect, error)
	GetSkillActionTiming(ctx context.Context, skillID string) (postgres.SkillActionTimingContract, error)
	GetSkillMovementActionBinding(ctx context.Context, skillID string) (postgres.SkillMovementActionBinding, error)
	GetHitboxProfiles(ctx context.Context, skillID string) ([]postgres.SkillHitboxProfile, error)
	GetImpactProfile(ctx context.Context, skillID string) (postgres.SkillImpactProfile, error)
}

type SkillDataHandler struct {
	apeironv1.UnimplementedSkillDataServiceServer

	skills SkillReader
}

func NewSkillDataHandler(skills SkillReader) *SkillDataHandler {
	return &SkillDataHandler{skills: skills}
}

func (h *SkillDataHandler) GetSkill(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.SkillResponse, error) {
	skill, err := h.skills.GetSkill(ctx, req.GetId())
	if err != nil {
		return &apeironv1.SkillResponse{Found: false, Error: err.Error()}, nil
	}

	return &apeironv1.SkillResponse{Found: true, Skill: mapSkill(skill)}, nil
}

func (h *SkillDataHandler) GetSkillSet(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.SkillSetResponse, error) {
	skillSet, err := h.skills.GetSkillSet(ctx, req.GetId())
	if err != nil {
		return &apeironv1.SkillSetResponse{Found: false, Error: err.Error()}, nil
	}

	return &apeironv1.SkillSetResponse{Found: true, SkillSet: mapSkillSet(skillSet)}, nil
}

func (h *SkillDataHandler) GetSkillSetLoadout(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.SkillSetLoadoutResponse, error) {
	loadout, err := h.skills.GetSkillSetLoadout(ctx, req.GetId())
	if err != nil {
		return &apeironv1.SkillSetLoadoutResponse{Found: false, Error: err.Error()}, nil
	}

	items := make([]*apeironv1.SkillLoadoutItem, 0, len(loadout))
	for _, item := range loadout {
		items = append(items, &apeironv1.SkillLoadoutItem{
			Slot:  mapSkillSlot(item.Slot),
			Skill: mapSkill(item.Skill),
		})
	}

	return &apeironv1.SkillSetLoadoutResponse{Found: true, Items: items}, nil
}

func (h *SkillDataHandler) GetWeaponCombatModeSlots(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.WeaponCombatModeSlotsResponse, error) {
	slots, err := h.skills.GetWeaponCombatModeSlots(ctx, req.GetId())
	if err != nil {
		return &apeironv1.WeaponCombatModeSlotsResponse{Found: false, Error: err.Error()}, nil
	}

	out := make([]*apeironv1.WeaponCombatModeSlot, 0, len(slots))
	for _, slot := range slots {
		out = append(out, mapWeaponCombatModeSlot(slot))
	}
	return &apeironv1.WeaponCombatModeSlotsResponse{Found: true, Slots: out}, nil
}

func (h *SkillDataHandler) GetSkillMovementEffect(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.SkillMovementEffectResponse, error) {
	effect, err := h.skills.GetMovementEffect(ctx, req.GetId())
	if err != nil {
		return &apeironv1.SkillMovementEffectResponse{Found: false, Error: err.Error()}, nil
	}

	return &apeironv1.SkillMovementEffectResponse{Found: true, Profile: mapSkillMovementEffect(effect)}, nil
}

func (h *SkillDataHandler) GetSkillActionTiming(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.SkillActionTimingResponse, error) {
	contract, err := h.skills.GetSkillActionTiming(ctx, req.GetId())
	if err != nil {
		return &apeironv1.SkillActionTimingResponse{Found: false, Error: err.Error()}, nil
	}

	return &apeironv1.SkillActionTimingResponse{Found: true, Contract: mapSkillActionTimingContract(contract)}, nil
}

func (h *SkillDataHandler) GetSkillMovementActionBinding(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.SkillMovementActionBindingResponse, error) {
	binding, err := h.skills.GetSkillMovementActionBinding(ctx, req.GetId())
	if err != nil {
		return &apeironv1.SkillMovementActionBindingResponse{Found: false, Error: err.Error()}, nil
	}

	return &apeironv1.SkillMovementActionBindingResponse{Found: true, Binding: mapSkillMovementActionBinding(binding)}, nil
}

func (h *SkillDataHandler) GetSkillHitboxProfiles(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.SkillHitboxProfilesResponse, error) {
	profiles, err := h.skills.GetHitboxProfiles(ctx, req.GetId())
	if err != nil {
		return &apeironv1.SkillHitboxProfilesResponse{Found: false, Error: err.Error()}, nil
	}

	out := make([]*apeironv1.SkillHitboxProfile, 0, len(profiles))
	for _, profile := range profiles {
		out = append(out, mapSkillHitboxProfile(profile))
	}

	return &apeironv1.SkillHitboxProfilesResponse{Found: true, Profiles: out}, nil
}

func (h *SkillDataHandler) GetSkillImpactProfile(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.SkillImpactProfileResponse, error) {
	profile, err := h.skills.GetImpactProfile(ctx, req.GetId())
	if err != nil {
		return &apeironv1.SkillImpactProfileResponse{Found: false, Error: err.Error()}, nil
	}

	return &apeironv1.SkillImpactProfileResponse{Found: true, Profile: mapSkillImpactProfile(profile)}, nil
}

func mapSkill(s postgres.Skill) *apeironv1.Skill {
	return &apeironv1.Skill{
		Id:                  s.ID,
		BaseDamage:          s.BaseDamage,
		StaminaCost:         s.StaminaCost,
		ManaCost:            s.ManaCost,
		HealthCost:          s.HealthCost,
		CooldownMs:          int32(s.CooldownMS),
		GlobalCooldownMs:    int32(s.CooldownMS),
		MaxRange:            s.MaxRange,
		RequiresTarget:      s.RequiresTarget,
		RequiresLineOfSight: !s.IgnoresLineOfSight,
		AllowMovement:       !s.LocksMovement,
		MovementLockMs:      int32(s.WindupMS + s.ActiveFramesMS + s.RecoveryMS),
		SkillType:           s.SkillType,
		TargetType:          s.TargetType,
		DamageMultiplier:    s.DamageMultiplier,
		PostureDamage:       s.PostureDamage,
		IsBlockable:         s.IsBlockable,
		IsParryable:         s.IsParryable,
		MaxTargets:          int32(s.MaxTargets),
		MovementDistance:    s.MovementDistance,
		ComboGroup:          nullString(s.ComboGroup),
		ComboStep:           int32(nullInt64(s.ComboIndex)),
		ComboWindowMs:       int32(s.ComboWindowMS),
		ComboResetMs:        int32(s.ComboWindowMS),
		Interruptible:       s.IsInterruptible,
		DamageType:          s.DamageType,
		ElementalType:       nullString(s.ElementalType),
	}
}

func mapSkillSet(s postgres.SkillSet) *apeironv1.SkillSet {
	return &apeironv1.SkillSet{
		Id:             s.ID,
		Name:           s.Name,
		Description:    s.Description,
		IsPlayerUsable: s.IsPlayerUsable,
		IsNpcUsable:    s.IsNPCUsable,
	}
}

func mapSkillSlot(s postgres.SkillSlot) *apeironv1.SkillSlot {
	return &apeironv1.SkillSlot{
		Id:                  s.ID,
		SkillSetId:          s.SkillSetID,
		SkillId:             s.SkillID,
		SlotIndex:           int32(s.SlotIndex),
		IsEnabled:           s.IsEnabled,
		Priority:            int32(s.Priority),
		UsageWeight:         s.UsageWeight,
		CooldownOverrideMs:  nullInt64(s.CooldownOverrideMS),
		MinTargetHpPercent:  nullFloat64(s.MinTargetHPPercent),
		MaxTargetHpPercent:  nullFloat64(s.MaxTargetHPPercent),
		MinSelfHpPercent:    nullFloat64(s.MinSelfHPPercent),
		MaxSelfHpPercent:    nullFloat64(s.MaxSelfHPPercent),
		RequiredDistanceMin: nullFloat64(s.RequiredDistanceMin),
		RequiredDistanceMax: nullFloat64(s.RequiredDistanceMax),
		RequiresLineOfSight: s.RequiresLineOfSight,
		OpenerWeight:        s.OpenerWeight,
		FinisherWeight:      s.FinisherWeight,
		SharedCooldownGroup: nullString(s.SharedCooldownGroup),
		UseOnlyInCombat:     s.UseOnlyInCombat,
	}
}

func mapSkillHitboxProfile(p postgres.SkillHitboxProfile) *apeironv1.SkillHitboxProfile {
	maxTargets := int32(p.MaxTargets)
	if maxTargets <= 0 {
		maxTargets = 1
	}
	targetType := ""
	out := &apeironv1.SkillHitboxProfile{
		Id:                  p.ID,
		SkillId:             p.SkillID,
		HitboxShape:         p.HitboxShape,
		HitboxStartMs:       int32(p.HitboxStartMS),
		HitboxEndMs:         int32(p.HitboxEndMS),
		OffsetX:             p.OffsetX,
		OffsetY:             p.OffsetY,
		OffsetZ:             p.OffsetZ,
		Length:              p.Length,
		Radius:              p.Radius,
		SizeX:               p.SizeX,
		SizeY:               p.SizeY,
		SizeZ:               p.SizeZ,
		HitboxIndex:         int32(p.HitboxIndex),
		Angle:               p.Angle,
		TargetType:          &targetType,
		MaxTargets:          &maxTargets,
		RequiresLineOfSight: true,
		CanHitNeutral:       p.FriendlyFire,
	}
	if p.DamageGroupID.Valid {
		out.DamageGroupId = p.DamageGroupID.String
	}
	if p.MotionProfile != nil {
		out.MotionProfile = mapSkillHitboxMotionProfile(*p.MotionProfile)
	}
	return out
}

func mapSkillHitboxMotionProfile(p postgres.SkillHitboxMotionProfile) *apeironv1.SkillHitboxMotionProfile {
	samples := make([]*apeironv1.SkillHitboxMotionSample, 0, len(p.Samples))
	for _, sample := range p.Samples {
		samples = append(samples, &apeironv1.SkillHitboxMotionSample{
			SampleIndex:   int32(sample.SampleIndex),
			T:             sample.T,
			OffsetX:       sample.OffsetX,
			OffsetY:       sample.OffsetY,
			OffsetZ:       sample.OffsetZ,
			Length:        sample.Length,
			Radius:        sample.Radius,
			SizeX:         sample.SizeX,
			SizeY:         sample.SizeY,
			SizeZ:         sample.SizeZ,
			StartAngleDeg: sample.StartAngleDeg,
			EndAngleDeg:   sample.EndAngleDeg,
		})
	}

	out := &apeironv1.SkillHitboxMotionProfile{
		Id:            p.ID,
		Enabled:       p.Enabled,
		MotionType:    p.MotionType,
		TimeBasis:     p.TimeBasis,
		Interpolation: p.Interpolation,
		SweepShape:    p.SweepShape,
		Samples:       samples,
	}
	if p.DamageGroupID.Valid {
		out.DamageGroupId = p.DamageGroupID.String
	}
	return out
}

func mapWeaponCombatModeSlot(s postgres.WeaponCombatModeSlot) *apeironv1.WeaponCombatModeSlot {
	return &apeironv1.WeaponCombatModeSlot{
		CombatModeId:  s.CombatModeID,
		InputSlot:     s.InputSlot,
		SkillId:       nullString(s.SkillID),
		IsBasicAttack: s.IsBasicAttack,
		IsFatality:    s.IsFatality,
		IsEnabled:     s.IsEnabled,
		MetadataJson:  s.MetadataJSON,
	}
}

func mapSkillMovementEffect(e postgres.SkillMovementEffect) *apeironv1.SkillMovementProfile {
	return &apeironv1.SkillMovementProfile{
		Id:               e.ID,
		MovementType:     e.MovementType,
		Distance:         e.Distance,
		Speed:            e.Speed,
		DurationMs:       int32(e.DurationMS),
		LandingLockMs:    int32(e.RecoveryLockMS),
		SteeringPolicy:   legacyMovementSteeringPolicy(e),
		IgnoresCollision: e.IgnoresCollision,
	}
}

func legacyMovementSteeringPolicy(e postgres.SkillMovementEffect) string {
	if e.CanRotate {
		return "can_rotate"
	}
	return "locked_facing"
}

func mapSkillActionTimingContract(c postgres.SkillActionTimingContract) *apeironv1.SkillActionTimingContract {
	return &apeironv1.SkillActionTimingContract{
		SkillId:            c.SkillID,
		WindupMs:           int32(c.WindupMS),
		ActiveMs:           int32(c.ActiveMS),
		RecoveryMs:         int32(c.RecoveryMS),
		CooldownMs:         int32(c.CooldownMS),
		ComboWindowMs:      int32(c.ComboWindowMS),
		MovementLockPolicy: c.MovementLockPolicy,
		QueuePolicy:        c.QueuePolicy,
		CancelPolicy:       c.CancelPolicy,
		MetadataJson:       c.MetadataJSON,
	}
}

func mapSkillMovementActionBinding(b postgres.SkillMovementActionBinding) *apeironv1.SkillMovementActionBinding {
	return &apeironv1.SkillMovementActionBinding{
		SkillId:                  b.SkillID,
		MovementActionContractId: b.MovementActionContractID,
		StartsAtPhase:            b.StartsAtPhase,
		HandoffPolicy:            b.HandoffPolicy,
		NormalInputPolicy:        b.NormalInputPolicy,
		TargetPolicy:             b.TargetPolicy,
		ContactPolicy:            b.ContactPolicy,
		IsEnabled:                b.IsEnabled,
		MetadataJson:             b.MetadataJSON,
		MovementActionContract:   mapMovementActionContract(b.MovementActionContract),
	}
}

func mapSkillImpactProfile(p postgres.SkillImpactProfile) *apeironv1.SkillImpactProfile {
	return &apeironv1.SkillImpactProfile{
		SkillId:               p.SkillID,
		ImpactType:            p.ImpactType,
		PoiseDamage:           p.PoiseDamage,
		StaggerPower:          p.StaggerPower,
		InterruptPower:        p.InterruptPower,
		HitReaction:           p.HitReaction,
		GuardDamageMultiplier: p.GuardDamageMultiplier,
	}
}
