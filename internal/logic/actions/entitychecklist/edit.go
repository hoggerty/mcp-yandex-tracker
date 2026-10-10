package entitychecklist

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) EditSafety() safety.Level { return safety.Write }

// Edit bulk-edits an entity's checklist items.
func (s *Service) Edit(ctx context.Context, entityType, id string, items []map[string]any) (json.RawMessage, error) {
	if err := s.guard.Check(s.EditSafety()); err != nil {
		return nil, err
	}

	return s.client.EditEntityChecklist(ctx, entityType, id, items)
}
