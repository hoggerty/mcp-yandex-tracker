package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) LocalFieldGetSafety() safety.Level { return safety.ReadOnly }

// LocalFieldGet retrieves a single local field of a queue.
func (s *Service) LocalFieldGet(ctx context.Context, queueKey, fieldKey string) (json.RawMessage, error) {
	if err := s.guard.Check(s.LocalFieldGetSafety()); err != nil {
		return nil, err
	}

	return s.client.GetLocalField(ctx, queueKey, fieldKey)
}
