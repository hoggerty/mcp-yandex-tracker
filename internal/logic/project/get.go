package project

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const GetSafety = safety.ReadOnly

func (s *Service) Get(ctx context.Context, id, expand string) (json.RawMessage, error) {
	return s.client.GetProject(ctx, id, expand)
}
