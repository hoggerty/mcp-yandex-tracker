package trigger

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) LogsListSafety() safety.Level { return safety.ReadOnly }

// LogsList lists the webhook invocation log of a trigger.
func (s *Service) LogsList(ctx context.Context, queueKey, triggerID, issueID string, limit int, from, to string) (json.RawMessage, error) {
	if err := s.guard.Check(s.LogsListSafety()); err != nil {
		return nil, err
	}

	return s.client.ListTriggerLogs(ctx, queueKey, triggerID, issueID, limit, from, to)
}
