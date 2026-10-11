package version

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) QueueListSafety() safety.Level { return safety.ReadOnly }

// QueueList lists the versions of a queue.
func (s *Service) QueueList(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.QueueListSafety(), func() (json.RawMessage, error) {
		return s.client.ListQueueVersions(ctx, queueKey)
	})
}
