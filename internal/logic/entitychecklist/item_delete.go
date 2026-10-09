package entitychecklist

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const ItemDeleteSafety = safety.Destructive

// ItemDelete deletes a single checklist item of an entity.
func (s *Service) ItemDelete(ctx context.Context, entityType, id, itemID string) (json.RawMessage, error) {
	return s.client.DeleteEntityChecklistItem(ctx, entityType, id, itemID)
}
