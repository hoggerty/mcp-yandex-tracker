package autoaction

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) LogGetSafety() safety.Level { return safety.ReadOnly }

// LogGet retrieves a single auto action run, including the issues it touched.
func (s *Service) LogGet(ctx context.Context, queueKey, autoActionID, runID string) (json.RawMessage, error) {
	if err := s.guard.Check(s.LogGetSafety()); err != nil {
		return nil, err
	}

	return s.client.GetAutoActionLog(ctx, queueKey, autoActionID, runID)
}
