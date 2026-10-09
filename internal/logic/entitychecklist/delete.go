package entitychecklist

import (
	"context"
	"encoding/json"
)

// Delete deletes all checklist items of an entity.
func (s *Service) Delete(ctx context.Context, entityType, id string) (json.RawMessage, error) {
	return s.client.DeleteEntityChecklist(ctx, entityType, id)
}
