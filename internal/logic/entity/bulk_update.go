package entity

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) BulkUpdateSafety() safety.Level { return safety.Destructive }

// BulkUpdate updates fields, comment, or links on multiple entities of the same type at once.
func (s *Service) BulkUpdate(ctx context.Context, entityType string, body map[string]any) (json.RawMessage, error) {
	if err := s.guard.Check(s.BulkUpdateSafety()); err != nil {
		return nil, err
	}

	return s.client.BulkUpdateEntities(ctx, entityType, body)
}
