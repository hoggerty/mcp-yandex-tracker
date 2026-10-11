package entity

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) CreateSafety() safety.Level { return safety.Write }

// Create creates a new entity (project, portfolio, or goal).
func (s *Service) Create(ctx context.Context, entityType string, fields map[string]any) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.CreateSafety(), func() (json.RawMessage, error) {
		return s.client.CreateEntity(ctx, entityType, fields)
	})
}
