package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) ListTransitionsSafety() safety.Level { return safety.ReadOnly }

func (s *Service) ListTransitions(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.ListTransitionsSafety(), func() (json.RawMessage, error) {
		return s.client.ListTransitions(ctx, issueKey)
	})
}
