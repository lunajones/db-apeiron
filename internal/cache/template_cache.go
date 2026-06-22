package cache

import (
	"context"
	"sync"
	"time"

	"db-apeiron/internal/repository/postgres"
)

const defaultTemplateCacheTTL = 10 * time.Minute

type TemplateCache struct {
	repository *postgres.CreatureRepository
	ttl        time.Duration

	creatureTemplates map[string]cacheEntry[postgres.CreatureTemplate]

	mu sync.RWMutex
}

func NewTemplateCache(repository *postgres.CreatureRepository) *TemplateCache {
	return NewTemplateCacheWithTTL(repository, defaultTemplateCacheTTL)
}

func NewTemplateCacheWithTTL(repository *postgres.CreatureRepository, ttl time.Duration) *TemplateCache {
	return &TemplateCache{
		repository: repository,
		ttl:        ttl,

		creatureTemplates: make(map[string]cacheEntry[postgres.CreatureTemplate]),
	}
}

func (c *TemplateCache) GetCreatureTemplate(
	ctx context.Context,
	id string,
) (postgres.CreatureTemplate, error) {
	if value, ok := getCached(&c.mu, c.creatureTemplates, id, c.isExpired); ok {
		return value, nil
	}

	template, err := c.repository.GetTemplateByID(ctx, id)
	if err != nil {
		return postgres.CreatureTemplate{}, err
	}

	setCached(&c.mu, c.creatureTemplates, id, template)
	return template, nil
}

func (c *TemplateCache) ReloadCreatureTemplate(
	ctx context.Context,
	id string,
) (postgres.CreatureTemplate, error) {
	template, err := c.repository.GetTemplateByID(ctx, id)
	if err != nil {
		return postgres.CreatureTemplate{}, err
	}

	setCached(&c.mu, c.creatureTemplates, id, template)
	return template, nil
}

func (c *TemplateCache) InvalidateCreatureTemplate(id string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.creatureTemplates, id)
}

func (c *TemplateCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.creatureTemplates = make(map[string]cacheEntry[postgres.CreatureTemplate])
}

func (c *TemplateCache) isExpired(loadedAt time.Time) bool {
	if c.ttl <= 0 {
		return false
	}

	return time.Since(loadedAt) > c.ttl
}
