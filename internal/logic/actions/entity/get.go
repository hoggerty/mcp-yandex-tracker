package entity

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) GetSafety() safety.Level { return safety.ReadOnly }

// Get retrieves the parameters of a single entity. expand and fields are optional.
func (s *Service) Get(ctx context.Context, entityType, id, expand, fields string) (json.RawMessage, error) {
	if err := s.guard.Check(s.GetSafety()); err != nil {
		return nil, err
	}

	return s.client.GetEntity(ctx, entityType, id, expand, fields)
}
