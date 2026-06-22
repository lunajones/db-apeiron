package cache

import (
	"context"
	"sync"
	"time"

	"db-apeiron/internal/repository/postgres"
)

const defaultStatusEffectCacheTTL = 30 * time.Minute

type StatusEffectCache struct {
	repository *postgres.SkillRepository
	ttl        time.Duration

	statusEffects map[string]cacheEntry[postgres.StatusEffect]

	mu sync.RWMutex
}

func NewStatusEffectCache(repository *postgres.SkillRepository) *StatusEffectCache {
	return NewStatusEffectCacheWithTTL(repository, defaultStatusEffectCacheTTL)
}

func NewStatusEffectCacheWithTTL(repository *postgres.SkillRepository, ttl time.Duration) *StatusEffectCache {
	return &StatusEffectCache{
		repository:    repository,
		ttl:           ttl,
		statusEffects: make(map[string]cacheEntry[postgres.StatusEffect]),
	}
}

func (c *StatusEffectCache) GetStatusEffect(ctx context.Context, id string) (postgres.StatusEffect, error) {
	if value, ok := getCached(&c.mu, c.statusEffects, id, c.isExpired); ok {
		return value, nil
	}

	statusEffect, err := c.repository.GetStatusEffectByID(ctx, id)
	if err != nil {
		return postgres.StatusEffect{}, err
	}

	setCached(&c.mu, c.statusEffects, id, statusEffect)
	return statusEffect, nil
}

func (c *StatusEffectCache) ReloadStatusEffect(ctx context.Context, id string) (postgres.StatusEffect, error) {
	statusEffect, err := c.repository.GetStatusEffectByID(ctx, id)
	if err != nil {
		return postgres.StatusEffect{}, err
	}

	setCached(&c.mu, c.statusEffects, id, statusEffect)
	return statusEffect, nil
}

func (c *StatusEffectCache) InvalidateStatusEffect(id string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.statusEffects, id)
}

func (c *StatusEffectCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.statusEffects = make(map[string]cacheEntry[postgres.StatusEffect])
}

func (c *StatusEffectCache) isExpired(loadedAt time.Time) bool {
	if c.ttl <= 0 {
		return false
	}

	return time.Since(loadedAt) > c.ttl
}
