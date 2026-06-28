package handlers

import (
	"context"

	apeironv1 "db-apeiron/gen/apeiron/v1"
	"db-apeiron/internal/repository/postgres"
)

type PlayerReader interface {
	GetByID(ctx context.Context, id string) (postgres.Player, error)
	UpdateProgression(ctx context.Context, id string, level int, experience int64, attributePoints int) error
	UpdateAttributes(ctx context.Context, id string, strength, dexterity, intelligence, endurance float64) error
}

type PlayerDataHandler struct {
	apeironv1.UnimplementedPlayerDataServiceServer

	players PlayerReader
}

func NewPlayerDataHandler(players PlayerReader) *PlayerDataHandler {
	return &PlayerDataHandler{players: players}
}

func (h *PlayerDataHandler) GetPlayer(ctx context.Context, req *apeironv1.IdRequest) (*apeironv1.PlayerResponse, error) {
	player, err := h.players.GetByID(ctx, req.GetId())
	if err != nil {
		return &apeironv1.PlayerResponse{Found: false, Error: err.Error()}, nil
	}

	return &apeironv1.PlayerResponse{Found: true, Player: mapPlayer(player)}, nil
}

// UpdatePlayer persists a player's progression (level/experience/attribute_points) and attributes
// (strength/dexterity/intelligence/endurance). Used by the game server to write progression back.
// Coin is not written here (runtime has no coin-set path yet).
func (h *PlayerDataHandler) UpdatePlayer(ctx context.Context, p *apeironv1.Player) (*apeironv1.OperationResult, error) {
	if p.GetId() == "" {
		return &apeironv1.OperationResult{Success: false, Message: "player id required"}, nil
	}
	if err := h.players.UpdateProgression(ctx, p.GetId(), int(p.GetLevel()), p.GetExperience(), int(p.GetAttributePoints())); err != nil {
		return &apeironv1.OperationResult{Success: false, Message: err.Error()}, nil
	}
	if err := h.players.UpdateAttributes(ctx, p.GetId(), p.GetStrength(), p.GetDexterity(), p.GetIntelligence(), p.GetEndurance()); err != nil {
		return &apeironv1.OperationResult{Success: false, Message: err.Error()}, nil
	}
	return &apeironv1.OperationResult{Success: true}, nil
}

func mapPlayer(player postgres.Player) *apeironv1.Player {
	return &apeironv1.Player{
		Id:                 player.ID,
		AccountId:          player.AccountID,
		Name:               player.Name,
		CreatureInstanceId: player.CreatureInstanceID,
		Level:              int32(player.Level),
		Experience:         player.Experience,
		AttributePoints:    int32(player.AttributePoints),
		Strength:           player.Strength,
		Dexterity:          player.Dexterity,
		Intelligence:       player.Intelligence,
		Endurance:          player.Endurance,
		PvpEnabled:         player.PVPEnabled,
		IsInSafeZone:       player.IsInSafeZone,
		GuildId:            nullString(player.GuildID),
		PartyId:            nullString(player.PartyID),
		Reputation:         player.Reputation,
		Coin:               player.Coin,
	}
}
