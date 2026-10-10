package workflow

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) QueueListSafety() safety.Level { return safety.ReadOnly }

// QueueList retrieves the workflow assigned to each issue type in a queue.
func (s *Service) QueueList(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.QueueListSafety(), func() (json.RawMessage, error) {
		return s.client.GetQueueWorkflows(ctx, queueKey)
	})
}
