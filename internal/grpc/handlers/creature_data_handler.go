package handlers

import (
	"context"

	apeironv1 "db-apeiron/gen/apeiron/v1"
	"db-apeiron/internal/repository/postgres"
)

type CreatureTemplateReader interface {
	GetCreatureTemplate(ctx context.Context, id string) (postgres.CreatureTemplate, error)
}

type CreatureDataHandler struct {
	apeironv1.UnimplementedCreatureDataServiceServer

	templates CreatureTemplateReader
}

func NewCreatureDataHandler(templates CreatureTemplateReader) *CreatureDataHandler {
	return &CreatureDataHandler{templates: templates}
}

func (h *CreatureDataHandler) GetCreatureTemplate(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.CreatureTemplateResponse, error) {
	template, err := h.templates.GetCreatureTemplate(ctx, req.GetId())
	if err != nil {
		return &apeironv1.CreatureTemplateResponse{Found: false}, nil
	}

	return &apeironv1.CreatureTemplateResponse{Found: true, Template: mapCreatureTemplate(template)}, nil
}

func (h *CreatureDataHandler) GetCreatureRuntimeData(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.CreatureRuntimeDataResponse, error) {
	template, err := h.templates.GetCreatureTemplate(ctx, req.GetId())
	if err != nil {
		return &apeironv1.CreatureRuntimeDataResponse{Found: false}, nil
	}

	return &apeironv1.CreatureRuntimeDataResponse{
		Found:                true,
		Template:             mapCreatureTemplate(template),
		SpawnProfileId:       template.SpawnProfileID,
		MovementProfileId:    template.MovementProfileID,
		CombatCoreProfileId:  template.CombatCoreProfileID,
		CombatStyleProfileId: template.CombatStyleProfileID,
		NeedsProfileId:       template.NeedsProfileID,
		PersonalityProfileId: template.PersonalityProfileID,
		AiDecisionProfileId:  template.AIDecisionProfileID,
		SensoryProfileId:     template.SensoryProfileID,
		SkillSetId:           template.SkillSetID,
	}, nil
}

func mapCreatureTemplate(t postgres.CreatureTemplate) *apeironv1.CreatureTemplate {
	return &apeironv1.CreatureTemplate{
		Id:                   t.ID,
		Name:                 t.Name,
		Faction:              t.Faction,
		Tier:                 int32(t.Tier),
		Archetype:            t.Archetype,
		SpawnProfileId:       t.SpawnProfileID,
		MovementProfileId:    t.MovementProfileID,
		CombatCoreProfileId:  t.CombatCoreProfileID,
		CombatStyleProfileId: t.CombatStyleProfileID,
		NeedsProfileId:       t.NeedsProfileID,
		PersonalityProfileId: t.PersonalityProfileID,
		AiDecisionProfileId:  t.AIDecisionProfileID,
		SensoryProfileId:     t.SensoryProfileID,
		SkillSetId:           t.SkillSetID,
	}
}
