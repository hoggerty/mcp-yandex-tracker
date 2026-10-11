package entitychecklist

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) ItemMoveSafety() safety.Level { return safety.Write }

// ItemMove moves a checklist item so it is positioned before the given item.
func (s *Service) ItemMove(ctx context.Context, entityType, id, itemID, beforeID string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.ItemMoveSafety(), func() (json.RawMessage, error) {
		return s.client.MoveEntityChecklistItem(ctx, entityType, id, itemID, beforeID)
	})
}
