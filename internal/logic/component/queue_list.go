package component

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const QueueListSafety = safety.ReadOnly

// QueueList lists the components of a queue.
func (s *Service) QueueList(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return s.client.ListQueueComponents(ctx, queueKey)
}
