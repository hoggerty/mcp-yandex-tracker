package entity

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const UpdateSafety = safety.Write

// Update edits an entity (fields, comment, links).
func (s *Service) Update(ctx context.Context, entityType, id string, body map[string]any) (json.RawMessage, error) {
	return s.client.UpdateEntity(ctx, entityType, id, body)
}
