package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const LocalFieldsListSafety = safety.ReadOnly

// LocalFieldsList lists a queue's local fields.
func (s *Service) LocalFieldsList(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return s.client.ListLocalFields(ctx, queueKey)
}
