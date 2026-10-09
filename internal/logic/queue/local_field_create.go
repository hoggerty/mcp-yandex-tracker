package queue

import (
	"context"
	"encoding/json"
)

// LocalFieldCreate creates a queue-local issue field.
func (s *Service) LocalFieldCreate(ctx context.Context, queueKey string, fields map[string]any) (json.RawMessage, error) {
	return s.client.CreateLocalField(ctx, queueKey, fields)
}
