package cache

import (
	"context"
	"sync"
	"time"

	"db-apeiron/internal/repository/postgres"
)

const defaultSkillCacheTTL = 10 * time.Minute

type SkillCache struct {
	repository *postgres.SkillRepository
	ttl        time.Duration

	skills             map[string]cacheEntry[postgres.Skill]
	skillSets          map[string]cacheEntry[postgres.SkillSet]
	skillSlots         map[string]cacheEntry[[]postgres.SkillSlot]
	skillLoadouts      map[string]cacheEntry[[]postgres.SkillLoadoutItem]
	projectileProfiles map[string]cacheEntry[postgres.SkillProjectileProfile]
	hitboxProfiles     map[string]cacheEntry[[]postgres.SkillHitboxProfile]
	areaEffectProfiles map[string]cacheEntry[postgres.SkillAreaEffectProfile]
	impactProfiles     map[string]cacheEntry[postgres.SkillImpactProfile]

	mu sync.RWMutex
}

func NewSkillCache(repository *postgres.SkillRepository) *SkillCache {
	return NewSkillCacheWithTTL(repository, defaultSkillCacheTTL)
}

func NewSkillCacheWithTTL(repository *postgres.SkillRepository, ttl time.Duration) *SkillCache {
	return &SkillCache{
		repository: repository,
		ttl:        ttl,

		skills:             make(map[string]cacheEntry[postgres.Skill]),
		skillSets:          make(map[string]cacheEntry[postgres.SkillSet]),
		skillSlots:         make(map[string]cacheEntry[[]postgres.SkillSlot]),
		skillLoadouts:      make(map[string]cacheEntry[[]postgres.SkillLoadoutItem]),
		projectileProfiles: make(map[string]cacheEntry[postgres.SkillProjectileProfile]),
		hitboxProfiles:     make(map[string]cacheEntry[[]postgres.SkillHitboxProfile]),
		areaEffectProfiles: make(map[string]cacheEntry[postgres.SkillAreaEffectProfile]),
		impactProfiles:     make(map[string]cacheEntry[postgres.SkillImpactProfile]),
	}
}

func (c *SkillCache) GetSkill(ctx context.Context, id string) (postgres.Skill, error) {
	if value, ok := getCached(&c.mu, c.skills, id, c.isExpired); ok {
		return value, nil
	}

	skill, err := c.repository.GetSkillByID(ctx, id)
	if err != nil {
		return postgres.Skill{}, err
	}

	setCached(&c.mu, c.skills, id, skill)
	return skill, nil
}

func (c *SkillCache) GetSkillSet(ctx context.Context, id string) (postgres.SkillSet, error) {
	if value, ok := getCached(&c.mu, c.skillSets, id, c.isExpired); ok {
		return value, nil
	}

	skillSet, err := c.repository.GetSkillSetByID(ctx, id)
	if err != nil {
		return postgres.SkillSet{}, err
	}

	setCached(&c.mu, c.skillSets, id, skillSet)
	return skillSet, nil
}

func (c *SkillCache) GetSlotsBySkillSetID(ctx context.Context, skillSetID string) ([]postgres.SkillSlot, error) {
	if value, ok := getCached(&c.mu, c.skillSlots, skillSetID, c.isExpired); ok {
		return value, nil
	}

	slots, err := c.repository.GetSlotsBySkillSetID(ctx, skillSetID)
	if err != nil {
		return nil, err
	}

	setCached(&c.mu, c.skillSlots, skillSetID, slots)
	return slots, nil
}

func (c *SkillCache) GetSkillSetLoadout(ctx context.Context, skillSetID string) ([]postgres.SkillLoadoutItem, error) {
	if value, ok := getCached(&c.mu, c.skillLoadouts, skillSetID, c.isExpired); ok {
		return value, nil
	}

	loadout, err := c.repository.GetSkillSetLoadout(ctx, skillSetID)
	if err != nil {
		return nil, err
	}

	setCached(&c.mu, c.skillLoadouts, skillSetID, loadout)
	return loadout, nil
}

func (c *SkillCache) GetProjectileProfile(ctx context.Context, skillID string) (postgres.SkillProjectileProfile, error) {
	if value, ok := getCached(&c.mu, c.projectileProfiles, skillID, c.isExpired); ok {
		return value, nil
	}

	profile, err := c.repository.GetProjectileProfileBySkillID(ctx, skillID)
	if err != nil {
		return postgres.SkillProjectileProfile{}, err
	}

	setCached(&c.mu, c.projectileProfiles, skillID, profile)
	return profile, nil
}

func (c *SkillCache) GetHitboxProfiles(ctx context.Context, skillID string) ([]postgres.SkillHitboxProfile, error) {
	if value, ok := getCached(&c.mu, c.hitboxProfiles, skillID, c.isExpired); ok {
		return value, nil
	}

	profiles, err := c.repository.GetHitboxProfilesBySkillID(ctx, skillID)
	if err != nil {
		return nil, err
	}

	setCached(&c.mu, c.hitboxProfiles, skillID, profiles)
	return profiles, nil
}

func (c *SkillCache) GetAreaEffectProfile(ctx context.Context, skillID string) (postgres.SkillAreaEffectProfile, error) {
	if value, ok := getCached(&c.mu, c.areaEffectProfiles, skillID, c.isExpired); ok {
		return value, nil
	}

	profile, err := c.repository.GetAreaEffectProfileBySkillID(ctx, skillID)
	if err != nil {
		return postgres.SkillAreaEffectProfile{}, err
	}

	setCached(&c.mu, c.areaEffectProfiles, skillID, profile)
	return profile, nil
}

func (c *SkillCache) GetImpactProfile(ctx context.Context, skillID string) (postgres.SkillImpactProfile, error) {
	if value, ok := getCached(&c.mu, c.impactProfiles, skillID, c.isExpired); ok {
		return value, nil
	}

	profile, err := c.repository.GetImpactProfileBySkillID(ctx, skillID)
	if err != nil {
		return postgres.SkillImpactProfile{}, err
	}

	setCached(&c.mu, c.impactProfiles, skillID, profile)
	return profile, nil
}

func (c *SkillCache) ReloadSkill(ctx context.Context, id string) (postgres.Skill, error) {
	skill, err := c.repository.GetSkillByID(ctx, id)
	if err != nil {
		return postgres.Skill{}, err
	}

	setCached(&c.mu, c.skills, id, skill)
	return skill, nil
}

func (c *SkillCache) ReloadSkillSet(ctx context.Context, id string) (postgres.SkillSet, error) {
	skillSet, err := c.repository.GetSkillSetByID(ctx, id)
	if err != nil {
		return postgres.SkillSet{}, err
	}

	setCached(&c.mu, c.skillSets, id, skillSet)
	return skillSet, nil
}

func (c *SkillCache) ReloadSkillSetLoadout(ctx context.Context, skillSetID string) ([]postgres.SkillLoadoutItem, error) {
	loadout, err := c.repository.GetSkillSetLoadout(ctx, skillSetID)
	if err != nil {
		return nil, err
	}

	setCached(&c.mu, c.skillLoadouts, skillSetID, loadout)
	return loadout, nil
}

func (c *SkillCache) ReloadHitboxProfiles(ctx context.Context, skillID string) ([]postgres.SkillHitboxProfile, error) {
	profiles, err := c.repository.GetHitboxProfilesBySkillID(ctx, skillID)
	if err != nil {
		return nil, err
	}

	setCached(&c.mu, c.hitboxProfiles, skillID, profiles)
	return profiles, nil
}

func (c *SkillCache) ReloadImpactProfile(ctx context.Context, skillID string) (postgres.SkillImpactProfile, error) {
	profile, err := c.repository.GetImpactProfileBySkillID(ctx, skillID)
	if err != nil {
		return postgres.SkillImpactProfile{}, err
	}

	setCached(&c.mu, c.impactProfiles, skillID, profile)
	return profile, nil
}

func (c *SkillCache) InvalidateSkill(id string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.skills, id)
	delete(c.projectileProfiles, id)
	delete(c.hitboxProfiles, id)
	delete(c.areaEffectProfiles, id)
	delete(c.impactProfiles, id)
}

func (c *SkillCache) InvalidateSkillSet(id string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.skillSets, id)
	delete(c.skillSlots, id)
	delete(c.skillLoadouts, id)
}

func (c *SkillCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.skills = make(map[string]cacheEntry[postgres.Skill])
	c.skillSets = make(map[string]cacheEntry[postgres.SkillSet])
	c.skillSlots = make(map[string]cacheEntry[[]postgres.SkillSlot])
	c.skillLoadouts = make(map[string]cacheEntry[[]postgres.SkillLoadoutItem])
	c.projectileProfiles = make(map[string]cacheEntry[postgres.SkillProjectileProfile])
	c.hitboxProfiles = make(map[string]cacheEntry[[]postgres.SkillHitboxProfile])
	c.areaEffectProfiles = make(map[string]cacheEntry[postgres.SkillAreaEffectProfile])
	c.impactProfiles = make(map[string]cacheEntry[postgres.SkillImpactProfile])
}

func (c *SkillCache) isExpired(loadedAt time.Time) bool {
	if c.ttl <= 0 {
		return false
	}

	return time.Since(loadedAt) > c.ttl
}
