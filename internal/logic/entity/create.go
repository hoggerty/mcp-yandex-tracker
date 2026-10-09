package entity

import (
	"context"
	"encoding/json"
)

// Create creates a new entity (project, portfolio, or goal).
func (s *Service) Create(ctx context.Context, entityType string, fields map[string]any) (json.RawMessage, error) {
	return s.client.CreateEntity(ctx, entityType, fields)
}
