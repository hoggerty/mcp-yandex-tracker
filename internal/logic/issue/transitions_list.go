package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) ListTransitionsSafety() safety.Level { return safety.ReadOnly }

func (s *Service) ListTransitions(ctx context.Context, issueKey string) (json.RawMessage, error) {
	if err := s.guard.Check(s.ListTransitionsSafety()); err != nil {
		return nil, err
	}

	return s.client.ListTransitions(ctx, issueKey)
}
