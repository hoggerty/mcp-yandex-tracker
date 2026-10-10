package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) TagDeleteSafety() safety.Level { return safety.Destructive }

// TagDelete removes a tag from a queue. Admin-only; the tag must be unused by any issue.
func (s *Service) TagDelete(ctx context.Context, queueKey, tag string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.TagDeleteSafety(), func() (json.RawMessage, error) {
		return s.client.DeleteQueueTag(ctx, queueKey, tag)
	})
}
