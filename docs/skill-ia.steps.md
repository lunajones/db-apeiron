STATUS

FASE 1 - PROJECT FOUNDATION
[x] 01 - Estrutura de pastas db-api
[x] 02 - main.go composition root
[x] 03 - .env
[x] 04 - config/env.go
[x] 05 - config/config.go
[x] 06 - logger/logger.go
[x] 07 - database/pool.go
[x] 08 - database/postgres.go
[x] 09 - database/migration.go
[x] 10 - grpc/server.go
[x] 11 - grpc/interceptors/logging.go
[x] 12 - grpc/interceptors/recovery.go
[x] 13 - grpc/interceptors/auth.go

FASE 2 - DATABASE CORE
[x] 14 - creature_template
[x] 15 - spawn_profile
[x] 16 - movement_profile
[x] 17 - combat_core_profile
[x] 18 - combat_style_profile
[x] 19 - needs_profile
[x] 20 - personality_profile
[x] 21 - ai_decision_profile
[x] 22 - skill
[x] 23 - skill_set
[x] 24 - skill_slot
[x] 25 - creature_instance_skill_state
[x] 26 - creature_instance
[x] 27 - player
[x] 28 - inventory
[x] 29 - inventory_item
[x] 30 - world_region
[x] 31 - biome
[x] 32 - spawn_zone
[x] 33 - skill_projectile_profile
[x] 34 - skill_hitbox_profile
[x] 35 - skill_area_effect_profile
[x] 36 - skill_impact_profile
[x] 37 - status_effect
[x] 38 - item_template

FASE 3 - DB RULE LAYER
[REMOVED] creature_state_rules
[REMOVED] skill_resolution_rules
[REMOVED] inventory_transaction_rules
[REMOVED] world_spawn_rules
[REMOVED] region_connectivity_rules

Motivo:
DB-APEIRON não executa lógica runtime.
DB-APEIRON guarda dados, contratos, repositories, seeds e cache estático.

FASE 4 - REPOSITORY LAYER
[x] creature_repository
[x] player_repository
[x] world_repository
[x] skill_repository
[x] inventory_repository
[x] profile_repository

FASE 4.1 - SEEDS
[x] seed runner
[x] default_profiles.sql
[x] default_skills.sql
[x] default_items.sql
[x] default_status_effects.sql
[x] default_creature_templates.sql
[x] world_defaults.sql

FASE 5 - CACHE LAYER STATIC ONLY
[x] template_cache
[x] profile_cache
[x] skill_cache
[x] item_cache
[x] status_effect_cache
[x] world_cache
[x] cache_loader
[x] dependencies.go
[x] app.go

FASE 6 - gRPC CONTRACTS
[ ] proto/apeiron/v1/common.proto
[ ] proto/apeiron/v1/cache_service.proto
[ ] proto/apeiron/v1/creature_data_service.proto
[ ] proto/apeiron/v1/player_data_service.proto
[ ] proto/apeiron/v1/world_data_service.proto
[ ] proto/apeiron/v1/profile_data_service.proto
[ ] proto/apeiron/v1/skill_data_service.proto
[ ] proto/apeiron/v1/inventory_data_service.proto
[ ] scripts/generate_proto.bat
[ ] generated Go protobuf files

FASE 7 - gRPC HANDLERS
[ ] cache_handler
[ ] creature_data_handler
[ ] player_data_handler
[ ] world_data_handler
[ ] profile_data_handler
[ ] skill_data_handler
[ ] inventory_data_handler
[ ] register handlers in app.go/server.go

FASE 8 - BOOTSTRAP / DATA LOAD SYSTEM
[x] cache_loader
[x] dependencies.go
[x] app.go
[ ] startup warmup strategy
[ ] optional preload static data
[ ] cache invalidation endpoints
[ ] admin-only cache operations

FASE 9 - TRANSACTION SAFETY
[x] tx_manager.go
[ ] atomic inventory operations
[ ] atomic skill state updates
[ ] safe creature spawn/despawn
[ ] transaction tests

FASE 10 - OBSERVABILITY
[ ] request id interceptor
[ ] structured gRPC logging
[ ] panic recovery interceptor
[ ] metrics
[ ] health check
[x] grpc reflection dev mode

FASE 11 - CONTRACT TESTS
[ ] proto compatibility check
[ ] repository integration tests
[ ] handler tests
[ ] cache tests