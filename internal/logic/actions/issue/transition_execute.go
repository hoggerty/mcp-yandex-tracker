package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) ExecuteTransitionSafety() safety.Level { return safety.Write }

func (s *Service) ExecuteTransition(ctx context.Context, issueKey, transitionID string, fields map[string]any) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.ExecuteTransitionSafety(), func() (json.RawMessage, error) {
		return s.client.ExecuteTransition(ctx, issueKey, transitionID, fields)
	})
}
