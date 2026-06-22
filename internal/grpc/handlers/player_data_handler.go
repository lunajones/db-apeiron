package handlers

import (
	"context"

	apeironv1 "db-apeiron/gen/apeiron/v1"
	"db-apeiron/internal/repository/postgres"
)

type PlayerReader interface {
	GetByID(ctx context.Context, id string) (postgres.Player, error)
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
