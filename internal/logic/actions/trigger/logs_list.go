package trigger

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) LogsListSafety() safety.Level { return safety.ReadOnly }

// LogsList lists the webhook invocation log of a trigger.
func (s *Service) LogsList(ctx context.Context, queueKey, triggerID, issueID string, limit int, from, to string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.LogsListSafety(), func() (json.RawMessage, error) {
		return s.client.ListTriggerLogs(ctx, queueKey, triggerID, issueID, limit, from, to)
	})
}
