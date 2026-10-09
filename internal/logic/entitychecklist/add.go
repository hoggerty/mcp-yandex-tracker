package entitychecklist

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const AddSafety = safety.Write

// Add adds an item to an entity's checklist (creating the checklist if needed).
func (s *Service) Add(ctx context.Context, entityType, id string, fields map[string]any) (json.RawMessage, error) {
	return s.client.AddEntityChecklistItem(ctx, entityType, id, fields)
}
