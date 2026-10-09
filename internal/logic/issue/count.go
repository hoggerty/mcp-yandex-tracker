package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const CountSafety = safety.ReadOnly

func (s *Service) Count(ctx context.Context, query string, filter map[string]any) (json.RawMessage, error) {
	return s.client.CountIssues(ctx, query, filter)
}
