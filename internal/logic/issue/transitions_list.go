package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const ListTransitionsSafety = safety.ReadOnly

func (s *Service) ListTransitions(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return s.client.ListTransitions(ctx, issueKey)
}
