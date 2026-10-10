package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) RestoreSafety() safety.Level { return safety.Write }

// Restore restores a previously deleted queue. Admin-only.
func (s *Service) Restore(ctx context.Context, queueKey string) (json.RawMessage, error) {
	if err := s.guard.Check(s.RestoreSafety()); err != nil {
		return nil, err
	}

	return s.client.RestoreQueue(ctx, queueKey)
}
