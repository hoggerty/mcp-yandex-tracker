package entitychecklist

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const ItemMoveSafety = safety.Write

// ItemMove moves a checklist item so it is positioned before the given item.
func (s *Service) ItemMove(ctx context.Context, entityType, id, itemID, beforeID string) (json.RawMessage, error) {
	return s.client.MoveEntityChecklistItem(ctx, entityType, id, itemID, beforeID)
}
