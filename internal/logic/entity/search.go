package entity

import (
	"context"
	"encoding/json"
)

// Search lists entities of a given type matching a filter.
func (s *Service) Search(ctx context.Context, entityType string, body map[string]any, fields string, perPage, page int) (json.RawMessage, error) {
	return s.client.SearchEntities(ctx, entityType, body, fields, perPage, page)
}
