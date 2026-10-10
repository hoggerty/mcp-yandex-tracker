package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) ExecuteTransitionSafety() safety.Level { return safety.Write }

func (s *Service) ExecuteTransition(ctx context.Context, issueKey, transitionID string, fields map[string]any) (json.RawMessage, error) {
	if err := s.guard.Check(s.ExecuteTransitionSafety()); err != nil {
		return nil, err
	}

	return s.client.ExecuteTransition(ctx, issueKey, transitionID, fields)
}
