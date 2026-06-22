db-apeiron/
│
├── cmd/
│   └── db-api/
│       └── main.go
│
├── internal/
│
│   ├── config/
│   │   ├── env.go
│   │   └── config.go
│   │
│   ├── logger/
│   │   └── logger.go
│   │
│   ├── database/
│   │   ├── postgres.go
│   │   ├── pool.go
│   │   ├── migration.go
│   │   └── tx_manager.go
│   │
│   ├── grpc/
│   │   ├── server.go
│   │   │
│   │   ├── interceptors/
│   │   │   ├── logging.go
│   │   │   ├── recovery.go
│   │   │   └── auth.go
│   │   │
│   │   └── handlers/
│   │       ├── cache_handler.go
│   │       ├── creature_data_handler.go
│   │       ├── player_data_handler.go
│   │       ├── world_data_handler.go
│   │       ├── profile_data_handler.go
│   │       ├── skill_data_handler.go
│   │       └── inventory_data_handler.go
│   │
│   ├── repository/
│   │   └── postgres/
│   │       ├── creature_repository.go
│   │       ├── player_repository.go
│   │       ├── world_repository.go
│   │       ├── skill_repository.go
│   │       ├── inventory_repository.go
│   │       └── profile_repository.go
│   │
│   ├── cache/
│   │   ├── template_cache.go
│   │   ├── profile_cache.go
│   │   ├── skill_cache.go
│   │   ├── item_cache.go
│   │   ├── status_effect_cache.go
│   │   └── world_cache.go
│   │
│   ├── bootstrap/
│   │   ├── app.go
│   │   ├── dependencies.go
│   │   └── cache_loader.go
│   │
│   └── shared/
│       ├── errors.go
│       ├── constants.go
│       └── query_utils.go
│
├── proto/
│   └── apeiron/
│       └── v1/
│           ├── common.proto
│           ├── cache_service.proto
│           ├── creature_data_service.proto
│           ├── player_data_service.proto
│           ├── world_data_service.proto
│           ├── profile_data_service.proto
│           ├── skill_data_service.proto
│           └── inventory_data_service.proto
│
├── gen/
│   └── apeiron/
│       └── v1/
│           ├── common.pb.go
│           ├── cache_service.pb.go
│           ├── cache_service_grpc.pb.go
│           ├── creature_data_service.pb.go
│           ├── creature_data_service_grpc.pb.go
│           ├── player_data_service.pb.go
│           ├── player_data_service_grpc.pb.go
│           ├── world_data_service.pb.go
│           ├── world_data_service_grpc.pb.go
│           ├── profile_data_service.pb.go
│           ├── profile_data_service_grpc.pb.go
│           ├── skill_data_service.pb.go
│           ├── skill_data_service_grpc.pb.go
│           ├── inventory_data_service.pb.go
│           └── inventory_data_service_grpc.pb.go
│
├── migrations/
│   ├── 001_extensions.sql
│   ├── 002_creature_template.sql
│   ├── 003_spawn_profile.sql
│   ├── 004_movement_profile.sql
│   ├── 005_combat_core_profile.sql
│   ├── 006_combat_style_profile.sql
│   ├── 007_needs_profile.sql
│   ├── 008_personality_profile.sql
│   ├── 009_ai_decision_profile.sql
│   ├── 010_skill.sql
│   ├── 011_skill_set.sql
│   ├── 012_skill_slot.sql
│   ├── 013_creature_instance_skill_state.sql
│   ├── 014_creature_instance.sql
│   ├── 015_player.sql
│   ├── 016_inventory.sql
│   ├── 017_inventory_item.sql
│   ├── 018_world_region.sql
│   ├── 019_biome.sql
│   ├── 020_spawn_zone.sql
│   ├── 021_skill_projectile_profile.sql
│   ├── 022_skill_hitbox_profile.sql
│   ├── 023_skill_area_effect_profile.sql
│   ├── 024_skill_impact_profile.sql
│   ├── 025_status_effect.sql
│   └── 026_item_template.sql
│
├── seeds/
│   ├── default_profiles.sql
│   ├── default_skills.sql
│   ├── default_items.sql
│   ├── default_status_effects.sql
│   ├── default_creature_templates.sql
│   └── world_defaults.sql
│
├── scripts/
│   ├── generate_proto.bat
│   ├── migrate.bat
│   ├── rollback.bat
│   └── seed.bat
│
├── docs/
│   ├── architecture.md
│   ├── database.md
│   ├── skill-system.md
│   ├── ai-design.md
│   ├── combat-design.md
│   └── world-design.md
│
├── .env
├── .gitignore
├── go.mod
└── go.sum