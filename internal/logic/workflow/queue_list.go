package workflow

import (
	"context"
	"encoding/json"
)

// QueueList retrieves the workflow assigned to each issue type in a queue.
func (s *Service) QueueList(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return s.client.GetQueueWorkflows(ctx, queueKey)
}
