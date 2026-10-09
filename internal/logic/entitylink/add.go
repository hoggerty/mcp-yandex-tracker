package entitylink

import (
	"context"
	"encoding/json"
)

// Add links an entity to one or more other entities.
func (s *Service) Add(ctx context.Context, entityType, id string, body any) (json.RawMessage, error) {
	return s.client.AddEntityLinks(ctx, entityType, id, body)
}
