package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) LocalFieldGetSafety() safety.Level { return safety.ReadOnly }

// LocalFieldGet retrieves a single local field of a queue.
func (s *Service) LocalFieldGet(ctx context.Context, queueKey, fieldKey string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.LocalFieldGetSafety(), func() (json.RawMessage, error) {
		return s.client.GetLocalField(ctx, queueKey, fieldKey)
	})
}
