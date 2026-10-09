package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const FieldsListSafety = safety.ReadOnly

// FieldsList retrieves the required fields configured for a queue.
func (s *Service) FieldsList(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return s.client.GetQueueFields(ctx, queueKey)
}
