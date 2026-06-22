package handlers

import (
	"context"

	apeironv1 "db-apeiron/gen/apeiron/v1"
	"db-apeiron/internal/repository/postgres"
)

type CreatureTemplateReader interface {
	GetCreatureTemplate(ctx context.Context, id string) (postgres.CreatureTemplate, error)
}

type CreatureRuntimeProfileReader interface {
	GetCreatureBehaviorRuntimeContract(ctx context.Context, id string) (postgres.CreatureBehaviorRuntimeContract, error)
	GetCreatureBehaviorRuntimeContractForTemplate(ctx context.Context, templateID string) (postgres.CreatureBehaviorRuntimeContract, error)
	GetCreatureTargetOpportunityPolicy(ctx context.Context, id string) (postgres.CreatureTargetOpportunityPolicy, error)
	GetCreatureOrbitPolicy(ctx context.Context, id string) (postgres.CreatureOrbitPolicy, error)
	GetCreatureEvasionPolicies(ctx context.Context, behaviorContractID string) ([]postgres.CreatureEvasionPolicy, error)
	GetCreatureSkillSetupPolicies(ctx context.Context, behaviorContractID string) ([]postgres.CreatureSkillSetupPolicy, error)
	GetCreatureSkillBehaviorBindings(ctx context.Context, behaviorContractID string) ([]postgres.CreatureSkillBehaviorBinding, error)
}

type CreatureDataHandler struct {
	apeironv1.UnimplementedCreatureDataServiceServer

	templates CreatureTemplateReader
	profiles  CreatureRuntimeProfileReader
}

func NewCreatureDataHandler(templates CreatureTemplateReader, profiles CreatureRuntimeProfileReader) *CreatureDataHandler {
	return &CreatureDataHandler{templates: templates, profiles: profiles}
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

	resp := &apeironv1.CreatureRuntimeDataResponse{
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
	}

	if h.profiles != nil {
		if contract, err := h.profiles.GetCreatureBehaviorRuntimeContractForTemplate(ctx, template.ID); err == nil {
			resp.BehaviorContractId = contract.ID
			resp.BehaviorContract = mapCreatureBehaviorRuntimeContract(contract)
			if contract.TargetOpportunityPolicyID != "" {
				if policy, err := h.profiles.GetCreatureTargetOpportunityPolicy(ctx, contract.TargetOpportunityPolicyID); err == nil {
					resp.TargetOpportunityPolicy = mapCreatureTargetOpportunityPolicy(policy)
				}
			}
			if contract.OrbitPolicyID != "" {
				if policy, err := h.profiles.GetCreatureOrbitPolicy(ctx, contract.OrbitPolicyID); err == nil {
					resp.OrbitPolicy = mapCreatureOrbitPolicy(policy)
				}
			}
		}
		behaviorContractID := resp.GetBehaviorContractId()
		if policies, err := h.profiles.GetCreatureEvasionPolicies(ctx, behaviorContractID); err == nil {
			for _, policy := range policies {
				resp.EvasionPolicies = append(resp.EvasionPolicies, mapCreatureEvasionPolicy(policy))
			}
		}
		if policies, err := h.profiles.GetCreatureSkillSetupPolicies(ctx, behaviorContractID); err == nil {
			for _, policy := range policies {
				resp.SkillSetupPolicies = append(resp.SkillSetupPolicies, mapCreatureSkillSetupPolicy(policy))
			}
		}
		if bindings, err := h.profiles.GetCreatureSkillBehaviorBindings(ctx, behaviorContractID); err == nil {
			for _, binding := range bindings {
				resp.SkillBehaviorBindings = append(resp.SkillBehaviorBindings, mapCreatureSkillBehaviorBinding(binding))
			}
		}
	}

	return resp, nil
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
