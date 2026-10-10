package entity

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) UpdateSafety() safety.Level { return safety.Write }

// Update edits an entity (fields, comment, links).
func (s *Service) Update(ctx context.Context, entityType, id string, body map[string]any) (json.RawMessage, error) {
	if err := s.guard.Check(s.UpdateSafety()); err != nil {
		return nil, err
	}

	return s.client.UpdateEntity(ctx, entityType, id, body)
}
