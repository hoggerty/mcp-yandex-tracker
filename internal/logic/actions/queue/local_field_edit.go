package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) LocalFieldEditSafety() safety.Level { return safety.Write }

// LocalFieldEdit edits a queue-local issue field.
func (s *Service) LocalFieldEdit(ctx context.Context, queueKey, fieldKey string, fields map[string]any) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.LocalFieldEditSafety(), func() (json.RawMessage, error) {
		return s.client.EditLocalField(ctx, queueKey, fieldKey, fields)
	})
}
