package bootstrap

import (
	"context"

	"db-apeiron/internal/cache"
)

type CacheLoader struct {
	Templates     *cache.TemplateCache
	Profiles      *cache.ProfileCache
	Skills        *cache.SkillCache
	Items         *cache.ItemCache
	StatusEffects *cache.StatusEffectCache
}

func NewCacheLoader(
	templates *cache.TemplateCache,
	profiles *cache.ProfileCache,
	skills *cache.SkillCache,
	items *cache.ItemCache,
	statusEffects *cache.StatusEffectCache,
) *CacheLoader {
	return &CacheLoader{
		Templates:     templates,
		Profiles:      profiles,
		Skills:        skills,
		Items:         items,
		StatusEffects: statusEffects,
	}
}

// WarmCreatureTemplate carrega um template de criatura no cache.
// Não carrega perfis automaticamente porque isso exigiria acoplar o loader
// aos campos internos do CreatureTemplate.
func (l *CacheLoader) WarmCreatureTemplate(
	ctx context.Context,
	templateID string,
) error {
	_, err := l.Templates.GetCreatureTemplate(ctx, templateID)
	return err
}

// WarmProfiles carrega perfis usados por criatura/world/template.
// O chamador passa só os IDs existentes.
func (l *CacheLoader) WarmProfiles(
	ctx context.Context,
	ids cache.ProfileLoadIDs,
) error {
	if ids.SpawnProfileID != "" {
		if _, err := l.Profiles.GetSpawnProfile(ctx, ids.SpawnProfileID); err != nil {
			return err
		}
	}

	if ids.MovementProfileID != "" {
		if _, err := l.Profiles.GetMovementProfile(ctx, ids.MovementProfileID); err != nil {
			return err
		}
	}

	if ids.CombatCoreProfileID != "" {
		if _, err := l.Profiles.GetCombatCoreProfile(ctx, ids.CombatCoreProfileID); err != nil {
			return err
		}
	}

	if ids.CombatStyleProfileID != "" {
		if _, err := l.Profiles.GetCombatStyleProfile(ctx, ids.CombatStyleProfileID); err != nil {
			return err
		}
	}

	if ids.NeedsProfileID != "" {
		if _, err := l.Profiles.GetNeedsProfile(ctx, ids.NeedsProfileID); err != nil {
			return err
		}
	}

	if ids.AIPersonalityProfileID != "" {
		if _, err := l.Profiles.GetAIPersonalityProfile(ctx, ids.AIPersonalityProfileID); err != nil {
			return err
		}
	}

	if ids.AIDecisionProfileID != "" {
		if _, err := l.Profiles.GetAIDecisionProfile(ctx, ids.AIDecisionProfileID); err != nil {
			return err
		}
	}

	if ids.SensoryProfileID != "" {
		if _, err := l.Profiles.GetSensoryProfile(ctx, ids.SensoryProfileID); err != nil {
			return err
		}
	}

	return nil
}

// WarmSkill carrega a skill e seus perfis auxiliares.
// Se algum perfil opcional não existir no banco, o repositório vai retornar erro.
// Então use isso só para skill que você sabe que tem esses registros.
func (l *CacheLoader) WarmSkill(
	ctx context.Context,
	skillID string,
) error {
	if _, err := l.Skills.GetSkill(ctx, skillID); err != nil {
		return err
	}

	if _, err := l.Skills.GetProjectileProfile(ctx, skillID); err != nil {
		return err
	}

	if _, err := l.Skills.GetHitboxProfiles(ctx, skillID); err != nil {
		return err
	}

	if _, err := l.Skills.GetAreaEffectProfile(ctx, skillID); err != nil {
		return err
	}

	if _, err := l.Skills.GetImpactProfile(ctx, skillID); err != nil {
		return err
	}

	return nil
}

// WarmSkillBasic carrega só a skill principal.
// Use esse quando projectile/area/hitbox forem opcionais.
func (l *CacheLoader) WarmSkillBasic(
	ctx context.Context,
	skillID string,
) error {
	_, err := l.Skills.GetSkill(ctx, skillID)
	return err
}

// WarmSkillSet carrega skill set, slots e loadout.
func (l *CacheLoader) WarmSkillSet(
	ctx context.Context,
	skillSetID string,
) error {
	if _, err := l.Skills.GetSkillSet(ctx, skillSetID); err != nil {
		return err
	}

	if _, err := l.Skills.GetSlotsBySkillSetID(ctx, skillSetID); err != nil {
		return err
	}

	if _, err := l.Skills.GetSkillSetLoadout(ctx, skillSetID); err != nil {
		return err
	}

	return nil
}

func (l *CacheLoader) WarmItemTemplate(
	ctx context.Context,
	itemTemplateID string,
) error {
	_, err := l.Items.GetItemTemplate(ctx, itemTemplateID)
	return err
}

func (l *CacheLoader) WarmStatusEffect(
	ctx context.Context,
	statusEffectID string,
) error {
	_, err := l.StatusEffects.GetStatusEffect(ctx, statusEffectID)
	return err
}

func (l *CacheLoader) WarmCreatureRuntimeData(
	ctx context.Context,
	req cache.CreatureRuntimeCacheLoadRequest,
) error {
	if req.CreatureTemplateID != "" {
		if err := l.WarmCreatureTemplate(ctx, req.CreatureTemplateID); err != nil {
			return err
		}
	}

	if err := l.WarmProfiles(ctx, req.Profiles); err != nil {
		return err
	}

	if req.SkillSetID != "" {
		if err := l.WarmSkillSet(ctx, req.SkillSetID); err != nil {
			return err
		}
	}

	for _, skillID := range req.SkillIDs {
		if skillID == "" {
			continue
		}

		if err := l.WarmSkillBasic(ctx, skillID); err != nil {
			return err
		}
	}

	for _, itemTemplateID := range req.ItemTemplateIDs {
		if itemTemplateID == "" {
			continue
		}

		if err := l.WarmItemTemplate(ctx, itemTemplateID); err != nil {
			return err
		}
	}

	for _, statusEffectID := range req.StatusEffectIDs {
		if statusEffectID == "" {
			continue
		}

		if err := l.WarmStatusEffect(ctx, statusEffectID); err != nil {
			return err
		}
	}

	return nil
}

func (l *CacheLoader) InvalidateCreatureTemplate(templateID string) {
	l.Templates.InvalidateCreatureTemplate(templateID)
}

func (l *CacheLoader) InvalidateProfile(profileID string) {
	l.Profiles.InvalidateProfile(profileID)
}

func (l *CacheLoader) InvalidateSkill(skillID string) {
	l.Skills.InvalidateSkill(skillID)
}

func (l *CacheLoader) InvalidateSkillSet(skillSetID string) {
	l.Skills.InvalidateSkillSet(skillSetID)
}

func (l *CacheLoader) InvalidateItemTemplate(itemTemplateID string) {
	l.Items.InvalidateItemTemplate(itemTemplateID)
}

func (l *CacheLoader) InvalidateStatusEffect(statusEffectID string) {
	l.StatusEffects.InvalidateStatusEffect(statusEffectID)
}

func (l *CacheLoader) ClearAll() {
	l.Templates.Clear()
	l.Profiles.Clear()
	l.Skills.Clear()
	l.Items.Clear()
	l.StatusEffects.Clear()
}
