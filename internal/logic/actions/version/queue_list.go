package version

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) QueueListSafety() safety.Level { return safety.ReadOnly }

// QueueList lists the versions of a queue.
func (s *Service) QueueList(ctx context.Context, queueKey string) (json.RawMessage, error) {
	if err := s.guard.Check(s.QueueListSafety()); err != nil {
		return nil, err
	}

	return s.client.ListQueueVersions(ctx, queueKey)
}
