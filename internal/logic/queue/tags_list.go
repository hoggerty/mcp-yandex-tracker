package queue

import (
	"context"
	"encoding/json"
)

// TagsList retrieves the tags defined on a queue.
func (s *Service) TagsList(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return s.client.GetQueueTags(ctx, queueKey)
}
