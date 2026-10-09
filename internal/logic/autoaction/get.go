package autoaction

import (
	"context"
	"encoding/json"
)

// Get retrieves the parameters of an auto action.
func (s *Service) Get(ctx context.Context, queueKey, autoActionID string) (json.RawMessage, error) {
	return s.client.GetAutoAction(ctx, queueKey, autoActionID)
}
