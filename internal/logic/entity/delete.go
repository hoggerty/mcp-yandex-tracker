package entity

import (
	"context"
	"encoding/json"
)

// Delete deletes an entity. If withBoard is true, its board is deleted too.
func (s *Service) Delete(ctx context.Context, entityType, id string, withBoard bool) (json.RawMessage, error) {
	return s.client.DeleteEntity(ctx, entityType, id, withBoard)
}
