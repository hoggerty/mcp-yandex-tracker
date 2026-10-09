package project

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const ListSafety = safety.ReadOnly

func (s *Service) List(ctx context.Context, expand string) (json.RawMessage, error) {
	return s.client.ListProjects(ctx, expand)
}
