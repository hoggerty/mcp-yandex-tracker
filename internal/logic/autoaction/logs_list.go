package autoaction

import (
	"context"
	"encoding/json"
)

// LogsList lists the run history of an auto action.
func (s *Service) LogsList(ctx context.Context, queueKey, autoActionID string) (json.RawMessage, error) {
	return s.client.ListAutoActionLogs(ctx, queueKey, autoActionID)
}
