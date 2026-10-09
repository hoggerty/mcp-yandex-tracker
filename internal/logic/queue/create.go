package queue

import (
	"context"
	"encoding/json"
)

// Create creates a new queue.
func (s *Service) Create(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return s.client.CreateQueue(ctx, fields)
}
