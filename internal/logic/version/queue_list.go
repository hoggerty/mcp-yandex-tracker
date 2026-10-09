package version

import (
	"context"
	"encoding/json"
)

// QueueList lists the versions of a queue.
func (s *Service) QueueList(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return s.client.ListQueueVersions(ctx, queueKey)
}
