package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const LocalFieldGetSafety = safety.ReadOnly

// LocalFieldGet retrieves a single local field of a queue.
func (s *Service) LocalFieldGet(ctx context.Context, queueKey, fieldKey string) (json.RawMessage, error) {
	return s.client.GetLocalField(ctx, queueKey, fieldKey)
}
