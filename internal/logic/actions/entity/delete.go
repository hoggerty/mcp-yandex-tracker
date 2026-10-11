package entity

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) DeleteSafety() safety.Level { return safety.Destructive }

// Delete deletes an entity. If withBoard is true, its board is deleted too.
func (s *Service) Delete(ctx context.Context, entityType, id string, withBoard bool) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.DeleteSafety(), func() (json.RawMessage, error) {
		return s.client.DeleteEntity(ctx, entityType, id, withBoard)
	})
}
