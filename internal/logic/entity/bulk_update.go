package entity

import (
	"context"
	"encoding/json"
)

// BulkUpdate updates fields, comment, or links on multiple entities of the same type at once.
func (s *Service) BulkUpdate(ctx context.Context, entityType string, body map[string]any) (json.RawMessage, error) {
	return s.client.BulkUpdateEntities(ctx, entityType, body)
}
