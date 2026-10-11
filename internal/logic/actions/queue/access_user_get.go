package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) AccessUserGetSafety() safety.Level { return safety.ReadOnly }

// AccessUserGet retrieves a user's permissions for a queue.
func (s *Service) AccessUserGet(ctx context.Context, queueKey, userID string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.AccessUserGetSafety(), func() (json.RawMessage, error) {
		return s.client.GetQueueUserAccess(ctx, queueKey, userID)
	})
}
