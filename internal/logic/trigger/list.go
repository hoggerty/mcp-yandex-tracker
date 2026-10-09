package trigger

import (
	"context"
	"encoding/json"
)

// List lists the triggers configured on a queue.
func (s *Service) List(ctx context.Context, queueKey string, perPage int, lastID string) (json.RawMessage, error) {
	return s.client.ListTriggers(ctx, queueKey, perPage, lastID)
}
