package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) FieldsListSafety() safety.Level { return safety.ReadOnly }

// FieldsList retrieves the required fields configured for a queue.
func (s *Service) FieldsList(ctx context.Context, queueKey string) (json.RawMessage, error) {
	if err := s.guard.Check(s.FieldsListSafety()); err != nil {
		return nil, err
	}

	return s.client.GetQueueFields(ctx, queueKey)
}
