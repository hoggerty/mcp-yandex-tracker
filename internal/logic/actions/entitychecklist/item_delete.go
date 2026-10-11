package entitychecklist

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) ItemDeleteSafety() safety.Level { return safety.Destructive }

// ItemDelete deletes a single checklist item of an entity.
func (s *Service) ItemDelete(ctx context.Context, entityType, id, itemID string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.ItemDeleteSafety(), func() (json.RawMessage, error) {
		return s.client.DeleteEntityChecklistItem(ctx, entityType, id, itemID)
	})
}
