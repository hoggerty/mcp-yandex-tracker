package entity

import (
	"context"
	"encoding/json"
)

// Update edits an entity (fields, comment, links).
func (s *Service) Update(ctx context.Context, entityType, id string, body map[string]any) (json.RawMessage, error) {
	return s.client.UpdateEntity(ctx, entityType, id, body)
}
