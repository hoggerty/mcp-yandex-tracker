package entitychecklist

import (
	"context"
	"encoding/json"
)

// Add adds an item to an entity's checklist (creating the checklist if needed).
func (s *Service) Add(ctx context.Context, entityType, id string, fields map[string]any) (json.RawMessage, error) {
	return s.client.AddEntityChecklistItem(ctx, entityType, id, fields)
}
