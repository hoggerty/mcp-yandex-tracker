package autoaction

import (
	"context"
	"encoding/json"
)

// LogGet retrieves a single auto action run, including the issues it touched.
func (s *Service) LogGet(ctx context.Context, queueKey, autoActionID, runID string) (json.RawMessage, error) {
	return s.client.GetAutoActionLog(ctx, queueKey, autoActionID, runID)
}
