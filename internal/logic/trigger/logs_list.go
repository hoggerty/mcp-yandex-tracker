package trigger

import (
	"context"
	"encoding/json"
)

// LogsList lists the webhook invocation log of a trigger.
func (s *Service) LogsList(ctx context.Context, queueKey, triggerID, issueID string, limit int, from, to string) (json.RawMessage, error) {
	return s.client.ListTriggerLogs(ctx, queueKey, triggerID, issueID, limit, from, to)
}
