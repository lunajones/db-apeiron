package cache

type ProfileLoadIDs struct {
	SpawnProfileID             string
	MovementProfileID          string
	CombatCoreProfileID        string
	CombatStyleProfileID       string
	NeedsProfileID             string
	AIPersonalityProfileID     string
	AIDecisionProfileID        string
	SensoryProfileID           string
	CreatureBehaviorContractID string
}

type CreatureRuntimeCacheLoadRequest struct {
	CreatureTemplateID string
	Profiles           ProfileLoadIDs

	SkillSetID      string
	SkillIDs        []string
	ItemTemplateIDs []string
	StatusEffectIDs []string
}
