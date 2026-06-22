package cache

import (
	"context"
	"sync"
	"time"

	"db-apeiron/internal/repository/postgres"
)

const defaultWorldCacheTTL = 30 * time.Minute

type WorldCache struct {
	repository *postgres.WorldRepository
	ttl        time.Duration

	regions            map[string]cacheEntry[postgres.WorldRegion]
	biomes             map[string]cacheEntry[postgres.Biome]
	spawnZones         map[string]cacheEntry[postgres.SpawnZone]
	biomesByRegion     map[string]cacheEntry[[]postgres.Biome]
	spawnZonesByRegion map[string]cacheEntry[[]postgres.SpawnZone]
	spawnZonesByBiome  map[string]cacheEntry[[]postgres.SpawnZone]

	mu sync.RWMutex
}

func NewWorldCache(repository *postgres.WorldRepository) *WorldCache {
	return NewWorldCacheWithTTL(repository, defaultWorldCacheTTL)
}

func NewWorldCacheWithTTL(repository *postgres.WorldRepository, ttl time.Duration) *WorldCache {
	return &WorldCache{
		repository: repository,
		ttl:        ttl,

		regions:            make(map[string]cacheEntry[postgres.WorldRegion]),
		biomes:             make(map[string]cacheEntry[postgres.Biome]),
		spawnZones:         make(map[string]cacheEntry[postgres.SpawnZone]),
		biomesByRegion:     make(map[string]cacheEntry[[]postgres.Biome]),
		spawnZonesByRegion: make(map[string]cacheEntry[[]postgres.SpawnZone]),
		spawnZonesByBiome:  make(map[string]cacheEntry[[]postgres.SpawnZone]),
	}
}

func (c *WorldCache) GetRegion(ctx context.Context, id string) (postgres.WorldRegion, error) {
	if value, ok := getCached(&c.mu, c.regions, id, c.isExpired); ok {
		return value, nil
	}

	region, err := c.repository.GetRegionByID(ctx, id)
	if err != nil {
		return postgres.WorldRegion{}, err
	}

	setCached(&c.mu, c.regions, id, region)
	return region, nil
}

func (c *WorldCache) GetBiome(ctx context.Context, id string) (postgres.Biome, error) {
	if value, ok := getCached(&c.mu, c.biomes, id, c.isExpired); ok {
		return value, nil
	}

	biome, err := c.repository.GetBiomeByID(ctx, id)
	if err != nil {
		return postgres.Biome{}, err
	}

	setCached(&c.mu, c.biomes, id, biome)
	return biome, nil
}

func (c *WorldCache) GetSpawnZone(ctx context.Context, id string) (postgres.SpawnZone, error) {
	if value, ok := getCached(&c.mu, c.spawnZones, id, c.isExpired); ok {
		return value, nil
	}

	spawnZone, err := c.repository.GetSpawnZoneByID(ctx, id)
	if err != nil {
		return postgres.SpawnZone{}, err
	}

	setCached(&c.mu, c.spawnZones, id, spawnZone)
	return spawnZone, nil
}

func (c *WorldCache) GetBiomesByRegion(ctx context.Context, regionID string) ([]postgres.Biome, error) {
	if value, ok := getCached(&c.mu, c.biomesByRegion, regionID, c.isExpired); ok {
		return value, nil
	}

	biomes, err := c.repository.GetBiomesByRegion(ctx, regionID)
	if err != nil {
		return nil, err
	}

	setCached(&c.mu, c.biomesByRegion, regionID, biomes)
	return biomes, nil
}

func (c *WorldCache) GetSpawnZonesByRegion(ctx context.Context, regionID string) ([]postgres.SpawnZone, error) {
	if value, ok := getCached(&c.mu, c.spawnZonesByRegion, regionID, c.isExpired); ok {
		return value, nil
	}

	spawnZones, err := c.repository.GetSpawnZonesByRegion(ctx, regionID)
	if err != nil {
		return nil, err
	}

	setCached(&c.mu, c.spawnZonesByRegion, regionID, spawnZones)
	return spawnZones, nil
}

func (c *WorldCache) GetSpawnZonesByBiome(ctx context.Context, biomeID string) ([]postgres.SpawnZone, error) {
	if value, ok := getCached(&c.mu, c.spawnZonesByBiome, biomeID, c.isExpired); ok {
		return value, nil
	}

	spawnZones, err := c.repository.GetSpawnZonesByBiome(ctx, biomeID)
	if err != nil {
		return nil, err
	}

	setCached(&c.mu, c.spawnZonesByBiome, biomeID, spawnZones)
	return spawnZones, nil
}

func (c *WorldCache) ReloadRegion(ctx context.Context, id string) (postgres.WorldRegion, error) {
	region, err := c.repository.GetRegionByID(ctx, id)
	if err != nil {
		return postgres.WorldRegion{}, err
	}

	setCached(&c.mu, c.regions, id, region)
	return region, nil
}

func (c *WorldCache) ReloadBiome(ctx context.Context, id string) (postgres.Biome, error) {
	biome, err := c.repository.GetBiomeByID(ctx, id)
	if err != nil {
		return postgres.Biome{}, err
	}

	setCached(&c.mu, c.biomes, id, biome)
	return biome, nil
}

func (c *WorldCache) ReloadSpawnZone(ctx context.Context, id string) (postgres.SpawnZone, error) {
	spawnZone, err := c.repository.GetSpawnZoneByID(ctx, id)
	if err != nil {
		return postgres.SpawnZone{}, err
	}

	setCached(&c.mu, c.spawnZones, id, spawnZone)
	return spawnZone, nil
}

func (c *WorldCache) ReloadBiomesByRegion(ctx context.Context, regionID string) ([]postgres.Biome, error) {
	biomes, err := c.repository.GetBiomesByRegion(ctx, regionID)
	if err != nil {
		return nil, err
	}

	setCached(&c.mu, c.biomesByRegion, regionID, biomes)
	return biomes, nil
}

func (c *WorldCache) ReloadSpawnZonesByRegion(ctx context.Context, regionID string) ([]postgres.SpawnZone, error) {
	spawnZones, err := c.repository.GetSpawnZonesByRegion(ctx, regionID)
	if err != nil {
		return nil, err
	}

	setCached(&c.mu, c.spawnZonesByRegion, regionID, spawnZones)
	return spawnZones, nil
}

func (c *WorldCache) ReloadSpawnZonesByBiome(ctx context.Context, biomeID string) ([]postgres.SpawnZone, error) {
	spawnZones, err := c.repository.GetSpawnZonesByBiome(ctx, biomeID)
	if err != nil {
		return nil, err
	}

	setCached(&c.mu, c.spawnZonesByBiome, biomeID, spawnZones)
	return spawnZones, nil
}

func (c *WorldCache) InvalidateRegion(id string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.regions, id)
	delete(c.biomesByRegion, id)
	delete(c.spawnZonesByRegion, id)
}

func (c *WorldCache) InvalidateBiome(id string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.biomes, id)
	delete(c.spawnZonesByBiome, id)
}

func (c *WorldCache) InvalidateSpawnZone(id string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.spawnZones, id)
}

func (c *WorldCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.regions = make(map[string]cacheEntry[postgres.WorldRegion])
	c.biomes = make(map[string]cacheEntry[postgres.Biome])
	c.spawnZones = make(map[string]cacheEntry[postgres.SpawnZone])
	c.biomesByRegion = make(map[string]cacheEntry[[]postgres.Biome])
	c.spawnZonesByRegion = make(map[string]cacheEntry[[]postgres.SpawnZone])
	c.spawnZonesByBiome = make(map[string]cacheEntry[[]postgres.SpawnZone])
}

func (c *WorldCache) isExpired(loadedAt time.Time) bool {
	if c.ttl <= 0 {
		return false
	}

	return time.Since(loadedAt) > c.ttl
}
