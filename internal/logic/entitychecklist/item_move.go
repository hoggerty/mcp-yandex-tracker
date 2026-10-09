package entitychecklist

import (
	"context"
	"encoding/json"
)

// ItemMove moves a checklist item so it is positioned before the given item.
func (s *Service) ItemMove(ctx context.Context, entityType, id, itemID, beforeID string) (json.RawMessage, error) {
	return s.client.MoveEntityChecklistItem(ctx, entityType, id, itemID, beforeID)
}
