package entity

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) SearchSafety() safety.Level { return safety.ReadOnly }

// Search lists entities of a given type matching a filter.
func (s *Service) Search(ctx context.Context, entityType string, body map[string]any, fields string, perPage, page int) (json.RawMessage, error) {
	if err := s.guard.Check(s.SearchSafety()); err != nil {
		return nil, err
	}

	return s.client.SearchEntities(ctx, entityType, body, fields, perPage, page)
}
