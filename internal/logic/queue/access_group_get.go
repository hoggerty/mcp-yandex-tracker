package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const AccessGroupGetSafety = safety.ReadOnly

// AccessGroupGet retrieves a group's permissions for a queue.
func (s *Service) AccessGroupGet(ctx context.Context, queueKey, groupID string) (json.RawMessage, error) {
	return s.client.GetQueueGroupAccess(ctx, queueKey, groupID)
}
