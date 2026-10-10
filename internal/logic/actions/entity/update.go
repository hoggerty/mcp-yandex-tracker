package entity

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) UpdateSafety() safety.Level { return safety.Write }

// Update edits an entity (fields, comment, links).
func (s *Service) Update(ctx context.Context, entityType, id string, body map[string]any) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.UpdateSafety(), func() (json.RawMessage, error) {
		return s.client.UpdateEntity(ctx, entityType, id, body)
	})
}
