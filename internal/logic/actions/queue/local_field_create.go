package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) LocalFieldCreateSafety() safety.Level { return safety.Write }

// LocalFieldCreate creates a queue-local issue field.
func (s *Service) LocalFieldCreate(ctx context.Context, queueKey string, fields map[string]any) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.LocalFieldCreateSafety(), func() (json.RawMessage, error) {
		return s.client.CreateLocalField(ctx, queueKey, fields)
	})
}
