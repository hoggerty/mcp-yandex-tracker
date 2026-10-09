package entity

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const DeleteSafety = safety.Destructive

// Delete deletes an entity. If withBoard is true, its board is deleted too.
func (s *Service) Delete(ctx context.Context, entityType, id string, withBoard bool) (json.RawMessage, error) {
	return s.client.DeleteEntity(ctx, entityType, id, withBoard)
}
