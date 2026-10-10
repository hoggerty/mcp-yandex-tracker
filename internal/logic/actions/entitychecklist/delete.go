package entitychecklist

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) DeleteSafety() safety.Level { return safety.Destructive }

// Delete deletes all checklist items of an entity.
func (s *Service) Delete(ctx context.Context, entityType, id string) (json.RawMessage, error) {
	if err := s.guard.Check(s.DeleteSafety()); err != nil {
		return nil, err
	}

	return s.client.DeleteEntityChecklist(ctx, entityType, id)
}
