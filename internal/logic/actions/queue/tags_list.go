package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) TagsListSafety() safety.Level { return safety.ReadOnly }

// TagsList retrieves the tags defined on a queue.
func (s *Service) TagsList(ctx context.Context, queueKey string) (json.RawMessage, error) {
	if err := s.guard.Check(s.TagsListSafety()); err != nil {
		return nil, err
	}

	return s.client.GetQueueTags(ctx, queueKey)
}
