package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) TagsListSafety() safety.Level { return safety.ReadOnly }

// TagsList retrieves the tags defined on a queue.
func (s *Service) TagsList(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.TagsListSafety(), func() (json.RawMessage, error) {
		return s.client.GetQueueTags(ctx, queueKey)
	})
}
