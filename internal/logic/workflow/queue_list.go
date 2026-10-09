package workflow

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const QueueListSafety = safety.ReadOnly

// QueueList retrieves the workflow assigned to each issue type in a queue.
func (s *Service) QueueList(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return s.client.GetQueueWorkflows(ctx, queueKey)
}
