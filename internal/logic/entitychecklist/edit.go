package entitychecklist

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const EditSafety = safety.Write

// Edit bulk-edits an entity's checklist items.
func (s *Service) Edit(ctx context.Context, entityType, id string, items []map[string]any) (json.RawMessage, error) {
	return s.client.EditEntityChecklist(ctx, entityType, id, items)
}
