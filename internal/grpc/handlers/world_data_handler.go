package handlers

import (
	"context"

	apeironv1 "db-apeiron/gen/apeiron/v1"
	"db-apeiron/internal/repository/postgres"
)

type WorldReader interface {
	GetRegion(ctx context.Context, id string) (postgres.WorldRegion, error)
	ListRegions(ctx context.Context) ([]postgres.WorldRegion, error)
	GetBiome(ctx context.Context, id string) (postgres.Biome, error)
	GetBiomesByRegion(ctx context.Context, regionID string) ([]postgres.Biome, error)
	GetSpawnZone(ctx context.Context, id string) (postgres.SpawnZone, error)
	GetSpawnZonesByRegion(ctx context.Context, regionID string) ([]postgres.SpawnZone, error)
	GetSpawnZonesByBiome(ctx context.Context, biomeID string) ([]postgres.SpawnZone, error)
}

type WorldDataHandler struct {
	apeironv1.UnimplementedWorldDataServiceServer

	world WorldReader
}

func NewWorldDataHandler(world WorldReader) *WorldDataHandler {
	return &WorldDataHandler{world: world}
}

func (h *WorldDataHandler) GetRegion(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.WorldRegionResponse, error) {
	region, err := h.world.GetRegion(ctx, req.GetId())
	if err != nil {
		return &apeironv1.WorldRegionResponse{Found: false, Error: err.Error()}, nil
	}

	return &apeironv1.WorldRegionResponse{Found: true, Region: mapWorldRegion(region)}, nil
}

func (h *WorldDataHandler) ListRegions(ctx context.Context, _ *apeironv1.Empty) (*apeironv1.WorldRegionsResponse, error) {
	regions, err := h.world.ListRegions(ctx)
	if err != nil {
		return &apeironv1.WorldRegionsResponse{Error: err.Error()}, nil
	}

	out := make([]*apeironv1.WorldRegion, 0, len(regions))
	for _, region := range regions {
		out = append(out, mapWorldRegion(region))
	}

	return &apeironv1.WorldRegionsResponse{Regions: out}, nil
}

func (h *WorldDataHandler) GetBiome(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.BiomeResponse, error) {
	biome, err := h.world.GetBiome(ctx, req.GetId())
	if err != nil {
		return &apeironv1.BiomeResponse{Found: false, Error: err.Error()}, nil
	}

	return &apeironv1.BiomeResponse{Found: true, Biome: mapBiome(biome)}, nil
}

func (h *WorldDataHandler) GetBiomesByRegion(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.BiomesResponse, error) {
	biomes, err := h.world.GetBiomesByRegion(ctx, req.GetId())
	if err != nil {
		return &apeironv1.BiomesResponse{Error: err.Error()}, nil
	}

	out := make([]*apeironv1.Biome, 0, len(biomes))
	for _, biome := range biomes {
		out = append(out, mapBiome(biome))
	}

	return &apeironv1.BiomesResponse{Biomes: out}, nil
}

func (h *WorldDataHandler) GetSpawnZone(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.SpawnZoneResponse, error) {
	spawnZone, err := h.world.GetSpawnZone(ctx, req.GetId())
	if err != nil {
		return &apeironv1.SpawnZoneResponse{Found: false, Error: err.Error()}, nil
	}

	return &apeironv1.SpawnZoneResponse{Found: true, SpawnZone: mapSpawnZone(spawnZone)}, nil
}

func (h *WorldDataHandler) GetSpawnZonesByRegion(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.SpawnZonesResponse, error) {
	return h.spawnZones(ctx, req.GetId(), h.world.GetSpawnZonesByRegion)
}

func (h *WorldDataHandler) GetSpawnZonesByBiome(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.SpawnZonesResponse, error) {
	return h.spawnZones(ctx, req.GetId(), h.world.GetSpawnZonesByBiome)
}

func (h *WorldDataHandler) spawnZones(
	ctx context.Context,
	id string,
	read func(context.Context, string) ([]postgres.SpawnZone, error),
) (*apeironv1.SpawnZonesResponse, error) {
	spawnZones, err := read(ctx, id)
	if err != nil {
		return &apeironv1.SpawnZonesResponse{Error: err.Error()}, nil
	}

	out := make([]*apeironv1.SpawnZone, 0, len(spawnZones))
	for _, spawnZone := range spawnZones {
		out = append(out, mapSpawnZone(spawnZone))
	}

	return &apeironv1.SpawnZonesResponse{SpawnZones: out}, nil
}

func mapWorldRegion(region postgres.WorldRegion) *apeironv1.WorldRegion {
	return &apeironv1.WorldRegion{
		Id:          region.ID,
		Name:        region.Name,
		RegionType:  region.RegionType,
		WorldScale:  int32(region.WorldScale),
		IsInstanced: region.IsInstanced,
		MaxPlayers:  int32(region.MaxPlayers),
		CenterX:     region.CenterX,
		CenterY:     region.CenterY,
		CenterZ:     region.CenterZ,
		SizeX:       region.SizeX,
		SizeY:       region.SizeY,
		SizeZ:       region.SizeZ,
		DangerLevel: region.DangerLevel,
	}
}

func mapBiome(biome postgres.Biome) *apeironv1.Biome {
	return &apeironv1.Biome{
		Id:                 biome.ID,
		Name:               biome.Name,
		RegionId:           biome.RegionID,
		BiomeType:          biome.BiomeType,
		Temperature:        biome.Temperature,
		Humidity:           biome.Humidity,
		VisibilityModifier: biome.VisibilityModifier,
		MovementModifier:   biome.MovementModifier,
		StealthModifier:    biome.StealthModifier,
		AggressionModifier: biome.AggressionModifier,
		ResourceRichness:   biome.ResourceRichness,
		IsSafe:             biome.IsSafe,
	}
}

func mapSpawnZone(spawnZone postgres.SpawnZone) *apeironv1.SpawnZone {
	return &apeironv1.SpawnZone{
		Id:                spawnZone.ID,
		RegionId:          spawnZone.RegionID,
		BiomeId:           spawnZone.BiomeID,
		Name:              spawnZone.Name,
		CenterX:           spawnZone.CenterX,
		CenterY:           spawnZone.CenterY,
		CenterZ:           spawnZone.CenterZ,
		Radius:            spawnZone.Radius,
		MaxEntities:       int32(spawnZone.MaxEntities),
		CurrentEntities:   int32(spawnZone.CurrentEntities),
		RespawnTimeMs:     spawnZone.RespawnTimeMS,
		SpawnDensity:      spawnZone.SpawnDensity,
		AllowedArchetypes: nullString(spawnZone.AllowedArchetypes),
		AggressionLevel:   spawnZone.AggressionLevel,
		LeashEnabled:      spawnZone.LeashEnabled,
	}
}
