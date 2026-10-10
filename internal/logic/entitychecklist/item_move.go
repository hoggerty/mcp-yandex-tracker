package entitychecklist

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) ItemMoveSafety() safety.Level { return safety.Write }

// ItemMove moves a checklist item so it is positioned before the given item.
func (s *Service) ItemMove(ctx context.Context, entityType, id, itemID, beforeID string) (json.RawMessage, error) {
	if err := s.guard.Check(s.ItemMoveSafety()); err != nil {
		return nil, err
	}

	return s.client.MoveEntityChecklistItem(ctx, entityType, id, itemID, beforeID)
}
