package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) LocalFieldsListSafety() safety.Level { return safety.ReadOnly }

// LocalFieldsList lists a queue's local fields.
func (s *Service) LocalFieldsList(ctx context.Context, queueKey string) (json.RawMessage, error) {
	if err := s.guard.Check(s.LocalFieldsListSafety()); err != nil {
		return nil, err
	}

	return s.client.ListLocalFields(ctx, queueKey)
}
