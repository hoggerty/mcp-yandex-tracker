package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) LocalFieldsListSafety() safety.Level { return safety.ReadOnly }

// LocalFieldsList lists a queue's local fields.
func (s *Service) LocalFieldsList(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.LocalFieldsListSafety(), func() (json.RawMessage, error) {
		return s.client.ListLocalFields(ctx, queueKey)
	})
}
