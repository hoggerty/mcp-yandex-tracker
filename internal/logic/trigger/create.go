package trigger

import (
	"context"
	"encoding/json"
)

// Create creates a new trigger on a queue.
func (s *Service) Create(ctx context.Context, queueKey string, fields map[string]any) (json.RawMessage, error) {
	return s.client.CreateTrigger(ctx, queueKey, fields)
}
