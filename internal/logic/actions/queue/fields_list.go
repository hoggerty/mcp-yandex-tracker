package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) FieldsListSafety() safety.Level { return safety.ReadOnly }

// FieldsList retrieves the required fields configured for a queue.
func (s *Service) FieldsList(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.FieldsListSafety(), func() (json.RawMessage, error) {
		return s.client.GetQueueFields(ctx, queueKey)
	})
}
