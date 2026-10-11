package entitychecklist

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) ItemEditSafety() safety.Level { return safety.Write }

// ItemEdit edits a single checklist item of an entity.
func (s *Service) ItemEdit(ctx context.Context, entityType, id, itemID string, fields map[string]any) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.ItemEditSafety(), func() (json.RawMessage, error) {
		return s.client.EditEntityChecklistItem(ctx, entityType, id, itemID, fields)
	})
}
