package cache

import (
	"context"
	"sync"
	"time"

	"db-apeiron/internal/repository/postgres"
)

const defaultProfileCacheTTL = 10 * time.Minute

type cacheEntry[T any] struct {
	value    T
	loadedAt time.Time
}

type ProfileCache struct {
	repository *postgres.ProfileRepository
	ttl        time.Duration

	movementProfiles        map[string]cacheEntry[postgres.MovementProfile]
	runtimeMovementProfiles map[string]cacheEntry[postgres.RuntimeMovementReconciliationProfile]
	combatCoreProfiles      map[string]cacheEntry[postgres.CombatCoreProfile]
	combatDefenseContracts  map[string]cacheEntry[postgres.CombatDefenseContract]
	movementActions         map[string]cacheEntry[postgres.MovementActionContract]
	reconciliation          map[string]cacheEntry[postgres.MovementReconciliationContract]
	creatureBehaviors       map[string]cacheEntry[postgres.CreatureBehaviorRuntimeContract]
	creatureEvasion         map[string]cacheEntry[[]postgres.CreatureEvasionPolicy]
	creatureSkillSetups     map[string]cacheEntry[[]postgres.CreatureSkillSetupPolicy]
	creatureOpportunities   map[string]cacheEntry[postgres.CreatureTargetOpportunityPolicy]
	creatureOrbitPolicies   map[string]cacheEntry[postgres.CreatureOrbitPolicy]
	creatureSkillBindings   map[string]cacheEntry[[]postgres.CreatureSkillBehaviorBinding]
	combatStyleProfiles     map[string]cacheEntry[postgres.CombatStyleProfile]
	needsProfiles           map[string]cacheEntry[postgres.NeedsProfile]
	aiPersonalityProfiles   map[string]cacheEntry[postgres.AIPersonalityProfile]
	aiDecisionProfiles      map[string]cacheEntry[postgres.AIDecisionProfile]
	sensoryProfiles         map[string]cacheEntry[postgres.SensoryProfile]
	spawnProfiles           map[string]cacheEntry[postgres.SpawnProfile]

	mu sync.RWMutex
}

func NewProfileCache(repository *postgres.ProfileRepository) *ProfileCache {
	return NewProfileCacheWithTTL(repository, defaultProfileCacheTTL)
}

func NewProfileCacheWithTTL(repository *postgres.ProfileRepository, ttl time.Duration) *ProfileCache {
	return &ProfileCache{
		repository: repository,
		ttl:        ttl,

		movementProfiles:        make(map[string]cacheEntry[postgres.MovementProfile]),
		runtimeMovementProfiles: make(map[string]cacheEntry[postgres.RuntimeMovementReconciliationProfile]),
		combatCoreProfiles:      make(map[string]cacheEntry[postgres.CombatCoreProfile]),
		combatDefenseContracts:  make(map[string]cacheEntry[postgres.CombatDefenseContract]),
		movementActions:         make(map[string]cacheEntry[postgres.MovementActionContract]),
		reconciliation:          make(map[string]cacheEntry[postgres.MovementReconciliationContract]),
		creatureBehaviors:       make(map[string]cacheEntry[postgres.CreatureBehaviorRuntimeContract]),
		creatureEvasion:         make(map[string]cacheEntry[[]postgres.CreatureEvasionPolicy]),
		creatureSkillSetups:     make(map[string]cacheEntry[[]postgres.CreatureSkillSetupPolicy]),
		creatureOpportunities:   make(map[string]cacheEntry[postgres.CreatureTargetOpportunityPolicy]),
		creatureOrbitPolicies:   make(map[string]cacheEntry[postgres.CreatureOrbitPolicy]),
		creatureSkillBindings:   make(map[string]cacheEntry[[]postgres.CreatureSkillBehaviorBinding]),
		combatStyleProfiles:     make(map[string]cacheEntry[postgres.CombatStyleProfile]),
		needsProfiles:           make(map[string]cacheEntry[postgres.NeedsProfile]),
		aiPersonalityProfiles:   make(map[string]cacheEntry[postgres.AIPersonalityProfile]),
		aiDecisionProfiles:      make(map[string]cacheEntry[postgres.AIDecisionProfile]),
		sensoryProfiles:         make(map[string]cacheEntry[postgres.SensoryProfile]),
		spawnProfiles:           make(map[string]cacheEntry[postgres.SpawnProfile]),
	}
}

func (c *ProfileCache) GetSpawnProfile(ctx context.Context, id string) (postgres.SpawnProfile, error) {
	if value, ok := getCached(&c.mu, c.spawnProfiles, id, c.isExpired); ok {
		return value, nil
	}

	profile, err := c.repository.GetSpawnProfileByID(ctx, id)
	if err != nil {
		return postgres.SpawnProfile{}, err
	}

	setCached(&c.mu, c.spawnProfiles, id, profile)
	return profile, nil
}

func (c *ProfileCache) GetMovementProfile(ctx context.Context, id string) (postgres.MovementProfile, error) {
	if value, ok := getCached(&c.mu, c.movementProfiles, id, c.isExpired); ok {
		return value, nil
	}

	profile, err := c.repository.GetMovementProfileByID(ctx, id)
	if err != nil {
		return postgres.MovementProfile{}, err
	}

	setCached(&c.mu, c.movementProfiles, id, profile)
	return profile, nil
}

func (c *ProfileCache) GetRuntimeMovementReconciliationProfile(ctx context.Context, id string) (postgres.RuntimeMovementReconciliationProfile, error) {
	if value, ok := getCached(&c.mu, c.runtimeMovementProfiles, id, c.isExpired); ok {
		return value, nil
	}

	profile, err := c.repository.GetRuntimeMovementReconciliationProfileByID(ctx, id)
	if err != nil {
		return postgres.RuntimeMovementReconciliationProfile{}, err
	}

	setCached(&c.mu, c.runtimeMovementProfiles, id, profile)
	return profile, nil
}

func (c *ProfileCache) GetCombatCoreProfile(ctx context.Context, id string) (postgres.CombatCoreProfile, error) {
	if value, ok := getCached(&c.mu, c.combatCoreProfiles, id, c.isExpired); ok {
		return value, nil
	}

	profile, err := c.repository.GetCombatCoreProfileByID(ctx, id)
	if err != nil {
		return postgres.CombatCoreProfile{}, err
	}

	setCached(&c.mu, c.combatCoreProfiles, id, profile)
	return profile, nil
}

func (c *ProfileCache) GetCombatDefenseContract(ctx context.Context, id string) (postgres.CombatDefenseContract, error) {
	if value, ok := getCached(&c.mu, c.combatDefenseContracts, id, c.isExpired); ok {
		return value, nil
	}

	contract, err := c.repository.GetCombatDefenseContractByID(ctx, id)
	if err != nil {
		return postgres.CombatDefenseContract{}, err
	}

	setCached(&c.mu, c.combatDefenseContracts, id, contract)
	return contract, nil
}

func (c *ProfileCache) GetMovementActionContract(ctx context.Context, id string) (postgres.MovementActionContract, error) {
	if value, ok := getCached(&c.mu, c.movementActions, id, c.isExpired); ok {
		return value, nil
	}

	contract, err := c.repository.GetMovementActionContractByID(ctx, id)
	if err != nil {
		return postgres.MovementActionContract{}, err
	}

	setCached(&c.mu, c.movementActions, id, contract)
	return contract, nil
}

func (c *ProfileCache) GetMovementReconciliationContract(ctx context.Context, id string) (postgres.MovementReconciliationContract, error) {
	if value, ok := getCached(&c.mu, c.reconciliation, id, c.isExpired); ok {
		return value, nil
	}

	contract, err := c.repository.GetMovementReconciliationContractByID(ctx, id)
	if err != nil {
		return postgres.MovementReconciliationContract{}, err
	}

	setCached(&c.mu, c.reconciliation, id, contract)
	return contract, nil
}

func (c *ProfileCache) GetCreatureBehaviorRuntimeContract(ctx context.Context, id string) (postgres.CreatureBehaviorRuntimeContract, error) {
	if value, ok := getCached(&c.mu, c.creatureBehaviors, id, c.isExpired); ok {
		return value, nil
	}

	contract, err := c.repository.GetCreatureBehaviorRuntimeContractByID(ctx, id)
	if err != nil {
		return postgres.CreatureBehaviorRuntimeContract{}, err
	}

	setCached(&c.mu, c.creatureBehaviors, id, contract)
	return contract, nil
}

func (c *ProfileCache) GetCreatureBehaviorRuntimeContractForTemplate(ctx context.Context, templateID string) (postgres.CreatureBehaviorRuntimeContract, error) {
	cacheKey := "template:" + templateID
	if value, ok := getCached(&c.mu, c.creatureBehaviors, cacheKey, c.isExpired); ok {
		return value, nil
	}

	contract, err := c.repository.GetCreatureBehaviorRuntimeContractByTemplateID(ctx, templateID)
	if err != nil {
		return postgres.CreatureBehaviorRuntimeContract{}, err
	}

	setCached(&c.mu, c.creatureBehaviors, cacheKey, contract)
	if contract.ID != "" {
		setCached(&c.mu, c.creatureBehaviors, contract.ID, contract)
	}
	return contract, nil
}

func (c *ProfileCache) GetCreatureEvasionPolicies(ctx context.Context, behaviorContractID string) ([]postgres.CreatureEvasionPolicy, error) {
	if value, ok := getCached(&c.mu, c.creatureEvasion, behaviorContractID, c.isExpired); ok {
		return value, nil
	}

	policies, err := c.repository.GetCreatureEvasionPoliciesByBehaviorContractID(ctx, behaviorContractID)
	if err != nil {
		return nil, err
	}

	setCached(&c.mu, c.creatureEvasion, behaviorContractID, policies)
	return policies, nil
}

func (c *ProfileCache) GetCreatureSkillSetupPolicies(ctx context.Context, behaviorContractID string) ([]postgres.CreatureSkillSetupPolicy, error) {
	if value, ok := getCached(&c.mu, c.creatureSkillSetups, behaviorContractID, c.isExpired); ok {
		return value, nil
	}

	policies, err := c.repository.GetCreatureSkillSetupPoliciesByBehaviorContractID(ctx, behaviorContractID)
	if err != nil {
		return nil, err
	}

	setCached(&c.mu, c.creatureSkillSetups, behaviorContractID, policies)
	return policies, nil
}

func (c *ProfileCache) GetCreatureTargetOpportunityPolicy(ctx context.Context, id string) (postgres.CreatureTargetOpportunityPolicy, error) {
	if value, ok := getCached(&c.mu, c.creatureOpportunities, id, c.isExpired); ok {
		return value, nil
	}

	policy, err := c.repository.GetCreatureTargetOpportunityPolicyByID(ctx, id)
	if err != nil {
		return postgres.CreatureTargetOpportunityPolicy{}, err
	}

	setCached(&c.mu, c.creatureOpportunities, id, policy)
	return policy, nil
}

func (c *ProfileCache) GetCreatureOrbitPolicy(ctx context.Context, id string) (postgres.CreatureOrbitPolicy, error) {
	if value, ok := getCached(&c.mu, c.creatureOrbitPolicies, id, c.isExpired); ok {
		return value, nil
	}

	policy, err := c.repository.GetCreatureOrbitPolicyByID(ctx, id)
	if err != nil {
		return postgres.CreatureOrbitPolicy{}, err
	}

	setCached(&c.mu, c.creatureOrbitPolicies, id, policy)
	return policy, nil
}

func (c *ProfileCache) GetCreatureSkillBehaviorBindings(ctx context.Context, behaviorContractID string) ([]postgres.CreatureSkillBehaviorBinding, error) {
	if value, ok := getCached(&c.mu, c.creatureSkillBindings, behaviorContractID, c.isExpired); ok {
		return value, nil
	}

	bindings, err := c.repository.GetCreatureSkillBehaviorBindingsByBehaviorContractID(ctx, behaviorContractID)
	if err != nil {
		return nil, err
	}

	setCached(&c.mu, c.creatureSkillBindings, behaviorContractID, bindings)
	return bindings, nil
}

func (c *ProfileCache) GetCombatStyleProfile(ctx context.Context, id string) (postgres.CombatStyleProfile, error) {
	if value, ok := getCached(&c.mu, c.combatStyleProfiles, id, c.isExpired); ok {
		return value, nil
	}

	profile, err := c.repository.GetCombatStyleProfileByID(ctx, id)
	if err != nil {
		return postgres.CombatStyleProfile{}, err
	}

	setCached(&c.mu, c.combatStyleProfiles, id, profile)
	return profile, nil
}

func (c *ProfileCache) GetNeedsProfile(ctx context.Context, id string) (postgres.NeedsProfile, error) {
	if value, ok := getCached(&c.mu, c.needsProfiles, id, c.isExpired); ok {
		return value, nil
	}

	profile, err := c.repository.GetNeedsProfileByID(ctx, id)
	if err != nil {
		return postgres.NeedsProfile{}, err
	}

	setCached(&c.mu, c.needsProfiles, id, profile)
	return profile, nil
}

func (c *ProfileCache) GetAIPersonalityProfile(ctx context.Context, id string) (postgres.AIPersonalityProfile, error) {
	if value, ok := getCached(&c.mu, c.aiPersonalityProfiles, id, c.isExpired); ok {
		return value, nil
	}

	profile, err := c.repository.GetAIPersonalityProfileByID(ctx, id)
	if err != nil {
		return postgres.AIPersonalityProfile{}, err
	}

	setCached(&c.mu, c.aiPersonalityProfiles, id, profile)
	return profile, nil
}

func (c *ProfileCache) GetAIDecisionProfile(ctx context.Context, id string) (postgres.AIDecisionProfile, error) {
	if value, ok := getCached(&c.mu, c.aiDecisionProfiles, id, c.isExpired); ok {
		return value, nil
	}

	profile, err := c.repository.GetAIDecisionProfileByID(ctx, id)
	if err != nil {
		return postgres.AIDecisionProfile{}, err
	}

	setCached(&c.mu, c.aiDecisionProfiles, id, profile)
	return profile, nil
}

func (c *ProfileCache) GetSensoryProfile(ctx context.Context, id string) (postgres.SensoryProfile, error) {
	if value, ok := getCached(&c.mu, c.sensoryProfiles, id, c.isExpired); ok {
		return value, nil
	}

	profile, err := c.repository.GetSensoryProfileByID(ctx, id)
	if err != nil {
		return postgres.SensoryProfile{}, err
	}

	setCached(&c.mu, c.sensoryProfiles, id, profile)
	return profile, nil
}

func (c *ProfileCache) ReloadSpawnProfile(ctx context.Context, id string) (postgres.SpawnProfile, error) {
	profile, err := c.repository.GetSpawnProfileByID(ctx, id)
	if err != nil {
		return postgres.SpawnProfile{}, err
	}

	setCached(&c.mu, c.spawnProfiles, id, profile)
	return profile, nil
}

func (c *ProfileCache) ReloadMovementProfile(ctx context.Context, id string) (postgres.MovementProfile, error) {
	profile, err := c.repository.GetMovementProfileByID(ctx, id)
	if err != nil {
		return postgres.MovementProfile{}, err
	}

	setCached(&c.mu, c.movementProfiles, id, profile)
	return profile, nil
}

func (c *ProfileCache) ReloadRuntimeMovementReconciliationProfile(ctx context.Context, id string) (postgres.RuntimeMovementReconciliationProfile, error) {
	profile, err := c.repository.GetRuntimeMovementReconciliationProfileByID(ctx, id)
	if err != nil {
		return postgres.RuntimeMovementReconciliationProfile{}, err
	}

	setCached(&c.mu, c.runtimeMovementProfiles, id, profile)
	return profile, nil
}

func (c *ProfileCache) ReloadCombatCoreProfile(ctx context.Context, id string) (postgres.CombatCoreProfile, error) {
	profile, err := c.repository.GetCombatCoreProfileByID(ctx, id)
	if err != nil {
		return postgres.CombatCoreProfile{}, err
	}

	setCached(&c.mu, c.combatCoreProfiles, id, profile)
	return profile, nil
}

func (c *ProfileCache) ReloadCombatDefenseContract(ctx context.Context, id string) (postgres.CombatDefenseContract, error) {
	contract, err := c.repository.GetCombatDefenseContractByID(ctx, id)
	if err != nil {
		return postgres.CombatDefenseContract{}, err
	}

	setCached(&c.mu, c.combatDefenseContracts, id, contract)
	return contract, nil
}

func (c *ProfileCache) ReloadCombatStyleProfile(ctx context.Context, id string) (postgres.CombatStyleProfile, error) {
	profile, err := c.repository.GetCombatStyleProfileByID(ctx, id)
	if err != nil {
		return postgres.CombatStyleProfile{}, err
	}

	setCached(&c.mu, c.combatStyleProfiles, id, profile)
	return profile, nil
}

func (c *ProfileCache) ReloadNeedsProfile(ctx context.Context, id string) (postgres.NeedsProfile, error) {
	profile, err := c.repository.GetNeedsProfileByID(ctx, id)
	if err != nil {
		return postgres.NeedsProfile{}, err
	}

	setCached(&c.mu, c.needsProfiles, id, profile)
	return profile, nil
}

func (c *ProfileCache) ReloadAIPersonalityProfile(ctx context.Context, id string) (postgres.AIPersonalityProfile, error) {
	profile, err := c.repository.GetAIPersonalityProfileByID(ctx, id)
	if err != nil {
		return postgres.AIPersonalityProfile{}, err
	}

	setCached(&c.mu, c.aiPersonalityProfiles, id, profile)
	return profile, nil
}

func (c *ProfileCache) ReloadAIDecisionProfile(ctx context.Context, id string) (postgres.AIDecisionProfile, error) {
	profile, err := c.repository.GetAIDecisionProfileByID(ctx, id)
	if err != nil {
		return postgres.AIDecisionProfile{}, err
	}

	setCached(&c.mu, c.aiDecisionProfiles, id, profile)
	return profile, nil
}

func (c *ProfileCache) ReloadSensoryProfile(ctx context.Context, id string) (postgres.SensoryProfile, error) {
	profile, err := c.repository.GetSensoryProfileByID(ctx, id)
	if err != nil {
		return postgres.SensoryProfile{}, err
	}

	setCached(&c.mu, c.sensoryProfiles, id, profile)
	return profile, nil
}

func (c *ProfileCache) InvalidateProfile(id string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.spawnProfiles, id)
	delete(c.movementProfiles, id)
	delete(c.runtimeMovementProfiles, id)
	delete(c.combatCoreProfiles, id)
	delete(c.combatDefenseContracts, id)
	delete(c.movementActions, id)
	delete(c.reconciliation, id)
	delete(c.creatureBehaviors, id)
	delete(c.creatureEvasion, id)
	delete(c.creatureSkillSetups, id)
	delete(c.creatureOpportunities, id)
	delete(c.creatureOrbitPolicies, id)
	delete(c.creatureSkillBindings, id)
	delete(c.combatStyleProfiles, id)
	delete(c.needsProfiles, id)
	delete(c.aiPersonalityProfiles, id)
	delete(c.aiDecisionProfiles, id)
	delete(c.sensoryProfiles, id)
}

func (c *ProfileCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.spawnProfiles = make(map[string]cacheEntry[postgres.SpawnProfile])
	c.movementProfiles = make(map[string]cacheEntry[postgres.MovementProfile])
	c.runtimeMovementProfiles = make(map[string]cacheEntry[postgres.RuntimeMovementReconciliationProfile])
	c.combatCoreProfiles = make(map[string]cacheEntry[postgres.CombatCoreProfile])
	c.combatDefenseContracts = make(map[string]cacheEntry[postgres.CombatDefenseContract])
	c.movementActions = make(map[string]cacheEntry[postgres.MovementActionContract])
	c.reconciliation = make(map[string]cacheEntry[postgres.MovementReconciliationContract])
	c.creatureBehaviors = make(map[string]cacheEntry[postgres.CreatureBehaviorRuntimeContract])
	c.creatureEvasion = make(map[string]cacheEntry[[]postgres.CreatureEvasionPolicy])
	c.creatureSkillSetups = make(map[string]cacheEntry[[]postgres.CreatureSkillSetupPolicy])
	c.creatureOpportunities = make(map[string]cacheEntry[postgres.CreatureTargetOpportunityPolicy])
	c.creatureOrbitPolicies = make(map[string]cacheEntry[postgres.CreatureOrbitPolicy])
	c.creatureSkillBindings = make(map[string]cacheEntry[[]postgres.CreatureSkillBehaviorBinding])
	c.combatStyleProfiles = make(map[string]cacheEntry[postgres.CombatStyleProfile])
	c.needsProfiles = make(map[string]cacheEntry[postgres.NeedsProfile])
	c.aiPersonalityProfiles = make(map[string]cacheEntry[postgres.AIPersonalityProfile])
	c.aiDecisionProfiles = make(map[string]cacheEntry[postgres.AIDecisionProfile])
	c.sensoryProfiles = make(map[string]cacheEntry[postgres.SensoryProfile])
}

func (c *ProfileCache) isExpired(loadedAt time.Time) bool {
	if c.ttl <= 0 {
		return false
	}

	return time.Since(loadedAt) > c.ttl
}

func getCached[T any](
	mu *sync.RWMutex,
	items map[string]cacheEntry[T],
	id string,
	isExpired func(time.Time) bool,
) (T, bool) {
	mu.RLock()
	entry, ok := items[id]
	mu.RUnlock()

	if !ok {
		var zero T
		return zero, false
	}

	if isExpired(entry.loadedAt) {
		mu.Lock()
		delete(items, id)
		mu.Unlock()

		var zero T
		return zero, false
	}

	return entry.value, true
}

func setCached[T any](
	mu *sync.RWMutex,
	items map[string]cacheEntry[T],
	id string,
	value T,
) {
	mu.Lock()
	defer mu.Unlock()

	items[id] = cacheEntry[T]{
		value:    value,
		loadedAt: time.Now(),
	}
}
