package bootstrap

import (
	"db-apeiron/internal/cache"
	"db-apeiron/internal/database"
	"db-apeiron/internal/repository/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Dependencies struct {
	Pool *pgxpool.Pool
	Tx   database.TxManager

	Repositories Repositories
	Caches       Caches
	CacheLoader  *CacheLoader
}

type Repositories struct {
	Creatures *postgres.CreatureRepository
	Profiles  *postgres.ProfileRepository
	Skills    *postgres.SkillRepository
	Inventory *postgres.InventoryRepository
	Players   *postgres.PlayerRepository
	World     *postgres.WorldRepository
}

type Caches struct {
	Templates     *cache.TemplateCache
	Profiles      *cache.ProfileCache
	Skills        *cache.SkillCache
	Items         *cache.ItemCache
	StatusEffects *cache.StatusEffectCache
}

func NewDependencies(pool *pgxpool.Pool) *Dependencies {
	txManager := database.NewTxManager(pool)

	repositories := Repositories{
		Creatures: postgres.NewCreatureRepository(txManager),
		Profiles:  postgres.NewProfileRepository(txManager),
		Skills:    postgres.NewSkillRepository(txManager),
		Inventory: postgres.NewInventoryRepository(txManager),
		Players:   postgres.NewPlayerRepository(txManager),
		World:     postgres.NewWorldRepository(txManager),
	}

	caches := Caches{
		Templates:     cache.NewTemplateCache(repositories.Creatures),
		Profiles:      cache.NewProfileCache(repositories.Profiles),
		Skills:        cache.NewSkillCache(repositories.Skills),
		Items:         cache.NewItemCache(repositories.Inventory),
		StatusEffects: cache.NewStatusEffectCache(repositories.Skills),
	}

	cacheLoader := NewCacheLoader(
		caches.Templates,
		caches.Profiles,
		caches.Skills,
		caches.Items,
		caches.StatusEffects,
	)

	return &Dependencies{
		Pool: pool,
		Tx:   txManager,

		Repositories: repositories,
		Caches:       caches,
		CacheLoader:  cacheLoader,
	}
}
