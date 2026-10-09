package entitychecklist

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const ItemEditSafety = safety.Write

// ItemEdit edits a single checklist item of an entity.
func (s *Service) ItemEdit(ctx context.Context, entityType, id, itemID string, fields map[string]any) (json.RawMessage, error) {
	return s.client.EditEntityChecklistItem(ctx, entityType, id, itemID, fields)
}
