package queue

import (
	"context"
	"encoding/json"
)

// Delete deletes a queue.
func (s *Service) Delete(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return s.client.DeleteQueue(ctx, queueKey)
}
