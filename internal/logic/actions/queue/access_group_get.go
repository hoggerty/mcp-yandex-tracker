package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) AccessGroupGetSafety() safety.Level { return safety.ReadOnly }

// AccessGroupGet retrieves a group's permissions for a queue.
func (s *Service) AccessGroupGet(ctx context.Context, queueKey, groupID string) (json.RawMessage, error) {
	if err := s.guard.Check(s.AccessGroupGetSafety()); err != nil {
		return nil, err
	}

	return s.client.GetQueueGroupAccess(ctx, queueKey, groupID)
}
