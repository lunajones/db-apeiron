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

	movementProfiles      map[string]cacheEntry[postgres.MovementProfile]
	combatCoreProfiles    map[string]cacheEntry[postgres.CombatCoreProfile]
	combatStyleProfiles   map[string]cacheEntry[postgres.CombatStyleProfile]
	needsProfiles         map[string]cacheEntry[postgres.NeedsProfile]
	aiPersonalityProfiles map[string]cacheEntry[postgres.AIPersonalityProfile]
	aiDecisionProfiles    map[string]cacheEntry[postgres.AIDecisionProfile]
	sensoryProfiles       map[string]cacheEntry[postgres.SensoryProfile]
	spawnProfiles         map[string]cacheEntry[postgres.SpawnProfile]

	mu sync.RWMutex
}

func NewProfileCache(repository *postgres.ProfileRepository) *ProfileCache {
	return NewProfileCacheWithTTL(repository, defaultProfileCacheTTL)
}

func NewProfileCacheWithTTL(repository *postgres.ProfileRepository, ttl time.Duration) *ProfileCache {
	return &ProfileCache{
		repository: repository,
		ttl:        ttl,

		movementProfiles:      make(map[string]cacheEntry[postgres.MovementProfile]),
		combatCoreProfiles:    make(map[string]cacheEntry[postgres.CombatCoreProfile]),
		combatStyleProfiles:   make(map[string]cacheEntry[postgres.CombatStyleProfile]),
		needsProfiles:         make(map[string]cacheEntry[postgres.NeedsProfile]),
		aiPersonalityProfiles: make(map[string]cacheEntry[postgres.AIPersonalityProfile]),
		aiDecisionProfiles:    make(map[string]cacheEntry[postgres.AIDecisionProfile]),
		sensoryProfiles:       make(map[string]cacheEntry[postgres.SensoryProfile]),
		spawnProfiles:         make(map[string]cacheEntry[postgres.SpawnProfile]),
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

func (c *ProfileCache) ReloadCombatCoreProfile(ctx context.Context, id string) (postgres.CombatCoreProfile, error) {
	profile, err := c.repository.GetCombatCoreProfileByID(ctx, id)
	if err != nil {
		return postgres.CombatCoreProfile{}, err
	}

	setCached(&c.mu, c.combatCoreProfiles, id, profile)
	return profile, nil
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
	delete(c.combatCoreProfiles, id)
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
	c.combatCoreProfiles = make(map[string]cacheEntry[postgres.CombatCoreProfile])
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
