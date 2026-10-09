package autoaction

import (
	"context"
	"encoding/json"
)

// Create creates a new auto action on a queue.
func (s *Service) Create(ctx context.Context, queueKey string, fields map[string]any) (json.RawMessage, error) {
	return s.client.CreateAutoAction(ctx, queueKey, fields)
}
