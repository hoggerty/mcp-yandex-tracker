package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) LocalFieldCreateSafety() safety.Level { return safety.Write }

// LocalFieldCreate creates a queue-local issue field.
func (s *Service) LocalFieldCreate(ctx context.Context, queueKey string, fields map[string]any) (json.RawMessage, error) {
	if err := s.guard.Check(s.LocalFieldCreateSafety()); err != nil {
		return nil, err
	}

	return s.client.CreateLocalField(ctx, queueKey, fields)
}
