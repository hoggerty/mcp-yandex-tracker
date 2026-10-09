package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const RestoreSafety = safety.Write

// Restore restores a previously deleted queue. Admin-only.
func (s *Service) Restore(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return s.client.RestoreQueue(ctx, queueKey)
}
