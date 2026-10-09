package queue

import (
	"context"
	"encoding/json"
)

// TagDelete removes a tag from a queue. Admin-only; the tag must be unused by any issue.
func (s *Service) TagDelete(ctx context.Context, queueKey, tag string) (json.RawMessage, error) {
	return s.client.DeleteQueueTag(ctx, queueKey, tag)
}
