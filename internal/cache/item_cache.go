package cache

import (
	"context"
	"sync"
	"time"

	"db-apeiron/internal/repository/postgres"
)

const defaultItemCacheTTL = 30 * time.Minute

type ItemCache struct {
	repository *postgres.InventoryRepository
	ttl        time.Duration

	itemTemplates map[string]cacheEntry[postgres.ItemTemplate]

	mu sync.RWMutex
}

func NewItemCache(repository *postgres.InventoryRepository) *ItemCache {
	return NewItemCacheWithTTL(repository, defaultItemCacheTTL)
}

func NewItemCacheWithTTL(repository *postgres.InventoryRepository, ttl time.Duration) *ItemCache {
	return &ItemCache{
		repository:    repository,
		ttl:           ttl,
		itemTemplates: make(map[string]cacheEntry[postgres.ItemTemplate]),
	}
}

func (c *ItemCache) GetItemTemplate(ctx context.Context, id string) (postgres.ItemTemplate, error) {
	if value, ok := getCached(&c.mu, c.itemTemplates, id, c.isExpired); ok {
		return value, nil
	}

	item, err := c.repository.GetItemTemplateByID(ctx, id)
	if err != nil {
		return postgres.ItemTemplate{}, err
	}

	setCached(&c.mu, c.itemTemplates, id, item)
	return item, nil
}

func (c *ItemCache) ReloadItemTemplate(ctx context.Context, id string) (postgres.ItemTemplate, error) {
	item, err := c.repository.GetItemTemplateByID(ctx, id)
	if err != nil {
		return postgres.ItemTemplate{}, err
	}

	setCached(&c.mu, c.itemTemplates, id, item)
	return item, nil
}

func (c *ItemCache) InvalidateItemTemplate(id string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.itemTemplates, id)
}

func (c *ItemCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.itemTemplates = make(map[string]cacheEntry[postgres.ItemTemplate])
}

func (c *ItemCache) isExpired(loadedAt time.Time) bool {
	if c.ttl <= 0 {
		return false
	}

	return time.Since(loadedAt) > c.ttl
}
