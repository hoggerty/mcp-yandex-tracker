package entity

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const GetSafety = safety.ReadOnly

// Get retrieves the parameters of a single entity. expand and fields are optional.
func (s *Service) Get(ctx context.Context, entityType, id, expand, fields string) (json.RawMessage, error) {
	return s.client.GetEntity(ctx, entityType, id, expand, fields)
}
