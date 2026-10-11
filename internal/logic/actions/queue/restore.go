package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) RestoreSafety() safety.Level { return safety.Write }

// Restore restores a previously deleted queue. Admin-only.
func (s *Service) Restore(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.RestoreSafety(), func() (json.RawMessage, error) {
		return s.client.RestoreQueue(ctx, queueKey)
	})
}
