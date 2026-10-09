package queue

import (
	"context"
	"encoding/json"
)

// Restore restores a previously deleted queue. Admin-only.
func (s *Service) Restore(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return s.client.RestoreQueue(ctx, queueKey)
}
