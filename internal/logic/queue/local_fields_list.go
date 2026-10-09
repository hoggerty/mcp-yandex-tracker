package queue

import (
	"context"
	"encoding/json"
)

// LocalFieldsList lists a queue's local fields.
func (s *Service) LocalFieldsList(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return s.client.ListLocalFields(ctx, queueKey)
}
