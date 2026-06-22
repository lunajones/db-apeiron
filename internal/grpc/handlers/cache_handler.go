package handlers

import (
	"context"

	apeironv1 "db-apeiron/gen/apeiron/v1"
	"db-apeiron/internal/cache"
)

type CacheLoader interface {
	WarmCreatureRuntimeData(ctx context.Context, req cache.CreatureRuntimeCacheLoadRequest) error

	WarmCreatureTemplate(ctx context.Context, templateID string) error
	WarmSkill(ctx context.Context, skillID string) error
	WarmSkillBasic(ctx context.Context, skillID string) error
	WarmSkillSet(ctx context.Context, skillSetID string) error
	WarmItemTemplate(ctx context.Context, itemTemplateID string) error
	WarmStatusEffect(ctx context.Context, statusEffectID string) error

	InvalidateCreatureTemplate(templateID string)
	InvalidateProfile(profileID string)
	InvalidateSkill(skillID string)
	InvalidateSkillSet(skillSetID string)
	InvalidateItemTemplate(itemTemplateID string)
	InvalidateStatusEffect(statusEffectID string)

	ClearAll()
}

type CacheHandler struct {
	apeironv1.UnimplementedCacheServiceServer

	cacheLoader CacheLoader
}

func NewCacheHandler(cacheLoader CacheLoader) *CacheHandler {
	return &CacheHandler{
		cacheLoader: cacheLoader,
	}
}

func (h *CacheHandler) WarmCreatureRuntimeData(
	ctx context.Context,
	req *apeironv1.WarmCreatureRuntimeDataRequest,
) (*apeironv1.OperationResult, error) {
	profiles := req.GetProfiles()

	err := h.cacheLoader.WarmCreatureRuntimeData(
		ctx,
		cache.CreatureRuntimeCacheLoadRequest{
			CreatureTemplateID: req.GetCreatureTemplateId(),
			Profiles: cache.ProfileLoadIDs{
				SpawnProfileID:         profiles.GetSpawnProfileId(),
				MovementProfileID:      profiles.GetMovementProfileId(),
				CombatCoreProfileID:    profiles.GetCombatCoreProfileId(),
				CombatStyleProfileID:   profiles.GetCombatStyleProfileId(),
				NeedsProfileID:         profiles.GetNeedsProfileId(),
				AIPersonalityProfileID: profiles.GetAiPersonalityProfileId(),
				AIDecisionProfileID:    profiles.GetAiDecisionProfileId(),
				SensoryProfileID:       profiles.GetSensoryProfileId(),
			},
			SkillSetID:      req.GetSkillSetId(),
			SkillIDs:        req.GetSkillIds(),
			ItemTemplateIDs: req.GetItemTemplateIds(),
			StatusEffectIDs: req.GetStatusEffectIds(),
		},
	)
	if err != nil {
		return fail(err), nil
	}

	return ok("creature runtime cache warmed"), nil
}

func (h *CacheHandler) WarmCreatureTemplate(
	ctx context.Context,
	req *apeironv1.IdRequest,
) (*apeironv1.OperationResult, error) {
	if err := h.cacheLoader.WarmCreatureTemplate(ctx, req.GetId()); err != nil {
		return fail(err), nil
	}

	return ok("creature template cache warmed"), nil
}

func (h *CacheHandler) WarmSkill(
	ctx context.Context,
	req *apeironv1.IdRequest,
) (*apeironv1.OperationResult, error) {
	if err := h.cacheLoader.WarmSkill(ctx, req.GetId()); err != nil {
		return fail(err), nil
	}

	return ok("skill cache warmed"), nil
}

func (h *CacheHandler) WarmSkillBasic(
	ctx context.Context,
	req *apeironv1.IdRequest,
) (*apeironv1.OperationResult, error) {
	if err := h.cacheLoader.WarmSkillBasic(ctx, req.GetId()); err != nil {
		return fail(err), nil
	}

	return ok("skill basic cache warmed"), nil
}

func (h *CacheHandler) WarmSkillSet(
	ctx context.Context,
	req *apeironv1.IdRequest,
) (*apeironv1.OperationResult, error) {
	if err := h.cacheLoader.WarmSkillSet(ctx, req.GetId()); err != nil {
		return fail(err), nil
	}

	return ok("skill set cache warmed"), nil
}

func (h *CacheHandler) WarmItemTemplate(
	ctx context.Context,
	req *apeironv1.IdRequest,
) (*apeironv1.OperationResult, error) {
	if err := h.cacheLoader.WarmItemTemplate(ctx, req.GetId()); err != nil {
		return fail(err), nil
	}

	return ok("item template cache warmed"), nil
}

func (h *CacheHandler) WarmStatusEffect(
	ctx context.Context,
	req *apeironv1.IdRequest,
) (*apeironv1.OperationResult, error) {
	if err := h.cacheLoader.WarmStatusEffect(ctx, req.GetId()); err != nil {
		return fail(err), nil
	}

	return ok("status effect cache warmed"), nil
}

func (h *CacheHandler) InvalidateCreatureTemplate(
	ctx context.Context,
	req *apeironv1.IdRequest,
) (*apeironv1.OperationResult, error) {
	_ = ctx

	h.cacheLoader.InvalidateCreatureTemplate(req.GetId())

	return ok("creature template cache invalidated"), nil
}

func (h *CacheHandler) InvalidateProfile(
	ctx context.Context,
	req *apeironv1.IdRequest,
) (*apeironv1.OperationResult, error) {
	_ = ctx

	h.cacheLoader.InvalidateProfile(req.GetId())

	return ok("profile cache invalidated"), nil
}

func (h *CacheHandler) InvalidateSkill(
	ctx context.Context,
	req *apeironv1.IdRequest,
) (*apeironv1.OperationResult, error) {
	_ = ctx

	h.cacheLoader.InvalidateSkill(req.GetId())

	return ok("skill cache invalidated"), nil
}

func (h *CacheHandler) InvalidateSkillSet(
	ctx context.Context,
	req *apeironv1.IdRequest,
) (*apeironv1.OperationResult, error) {
	_ = ctx

	h.cacheLoader.InvalidateSkillSet(req.GetId())

	return ok("skill set cache invalidated"), nil
}

func (h *CacheHandler) InvalidateItemTemplate(
	ctx context.Context,
	req *apeironv1.IdRequest,
) (*apeironv1.OperationResult, error) {
	_ = ctx

	h.cacheLoader.InvalidateItemTemplate(req.GetId())

	return ok("item template cache invalidated"), nil
}

func (h *CacheHandler) InvalidateStatusEffect(
	ctx context.Context,
	req *apeironv1.IdRequest,
) (*apeironv1.OperationResult, error) {
	_ = ctx

	h.cacheLoader.InvalidateStatusEffect(req.GetId())

	return ok("status effect cache invalidated"), nil
}

func (h *CacheHandler) ClearAll(
	ctx context.Context,
	req *apeironv1.Empty,
) (*apeironv1.OperationResult, error) {
	_ = ctx
	_ = req

	h.cacheLoader.ClearAll()

	return ok("all caches cleared"), nil
}

func ok(message string) *apeironv1.OperationResult {
	return &apeironv1.OperationResult{
		Success: true,
		Message: message,
	}
}

func fail(err error) *apeironv1.OperationResult {
	return &apeironv1.OperationResult{
		Success: false,
		Message: err.Error(),
	}
}
