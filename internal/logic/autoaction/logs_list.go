package autoaction

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) LogsListSafety() safety.Level { return safety.ReadOnly }

// LogsList lists the run history of an auto action.
func (s *Service) LogsList(ctx context.Context, queueKey, autoActionID string) (json.RawMessage, error) {
	if err := s.guard.Check(s.LogsListSafety()); err != nil {
		return nil, err
	}

	return s.client.ListAutoActionLogs(ctx, queueKey, autoActionID)
}
