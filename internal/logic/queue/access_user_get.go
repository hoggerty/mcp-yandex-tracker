package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const AccessUserGetSafety = safety.ReadOnly

// AccessUserGet retrieves a user's permissions for a queue.
func (s *Service) AccessUserGet(ctx context.Context, queueKey, userID string) (json.RawMessage, error) {
	return s.client.GetQueueUserAccess(ctx, queueKey, userID)
}
