package trigger

import (
	"context"
	"encoding/json"
)

// Get retrieves the parameters of a single trigger.
func (s *Service) Get(ctx context.Context, queueKey, triggerID string) (json.RawMessage, error) {
	return s.client.GetTrigger(ctx, queueKey, triggerID)
}
