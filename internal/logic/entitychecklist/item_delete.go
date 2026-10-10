package entitychecklist

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) ItemDeleteSafety() safety.Level { return safety.Destructive }

// ItemDelete deletes a single checklist item of an entity.
func (s *Service) ItemDelete(ctx context.Context, entityType, id, itemID string) (json.RawMessage, error) {
	if err := s.guard.Check(s.ItemDeleteSafety()); err != nil {
		return nil, err
	}

	return s.client.DeleteEntityChecklistItem(ctx, entityType, id, itemID)
}
