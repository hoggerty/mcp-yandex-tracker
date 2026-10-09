package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const TagsListSafety = safety.ReadOnly

// TagsList retrieves the tags defined on a queue.
func (s *Service) TagsList(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return s.client.GetQueueTags(ctx, queueKey)
}
