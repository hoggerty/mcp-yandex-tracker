package entitylink

import (
	"context"
	"encoding/json"
)

// List lists the entities linked to an entity. fields is optional (comma-separated linked entity fields).
func (s *Service) List(ctx context.Context, entityType, id, fields string) (json.RawMessage, error) {
	return s.client.ListEntityLinks(ctx, entityType, id, fields)
}
