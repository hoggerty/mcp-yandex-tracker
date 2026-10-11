package autoaction

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) LogGetSafety() safety.Level { return safety.ReadOnly }

// LogGet retrieves a single auto action run, including the issues it touched.
func (s *Service) LogGet(ctx context.Context, queueKey, autoActionID, runID string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.LogGetSafety(), func() (json.RawMessage, error) {
		return s.client.GetAutoActionLog(ctx, queueKey, autoActionID, runID)
	})
}
