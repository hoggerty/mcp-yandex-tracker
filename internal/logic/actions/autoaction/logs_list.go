package autoaction

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) LogsListSafety() safety.Level { return safety.ReadOnly }

// LogsList lists the run history of an auto action.
func (s *Service) LogsList(ctx context.Context, queueKey, autoActionID string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.LogsListSafety(), func() (json.RawMessage, error) {
		return s.client.ListAutoActionLogs(ctx, queueKey, autoActionID)
	})
}
