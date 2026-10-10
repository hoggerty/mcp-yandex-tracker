package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) AccessUserGetSafety() safety.Level { return safety.ReadOnly }

// AccessUserGet retrieves a user's permissions for a queue.
func (s *Service) AccessUserGet(ctx context.Context, queueKey, userID string) (json.RawMessage, error) {
	if err := s.guard.Check(s.AccessUserGetSafety()); err != nil {
		return nil, err
	}

	return s.client.GetQueueUserAccess(ctx, queueKey, userID)
}
