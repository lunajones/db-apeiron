package handlers

import (
	"context"

	apeironv1 "db-apeiron/gen/apeiron/v1"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ObservabilityHandler struct {
	apeironv1.UnimplementedObservabilityServiceServer

	pool *pgxpool.Pool
}

func NewObservabilityHandler(pool *pgxpool.Pool) *ObservabilityHandler {
	return &ObservabilityHandler{pool: pool}
}

func (h *ObservabilityHandler) GetReadiness(ctx context.Context, _ *apeironv1.Empty) (*apeironv1.ReadinessResponse, error) {
	if h.pool == nil {
		return readiness("NOT_READY", "postgres pool is not initialized"), nil
	}

	if err := h.pool.Ping(ctx); err != nil {
		return readiness("NOT_READY", err.Error()), nil
	}

	return readiness("READY", "postgres reachable"), nil
}

func readiness(status string, message string) *apeironv1.ReadinessResponse {
	return &apeironv1.ReadinessResponse{
		Readiness: &apeironv1.Readiness{
			Status:  status,
			Message: message,
		},
	}
}
