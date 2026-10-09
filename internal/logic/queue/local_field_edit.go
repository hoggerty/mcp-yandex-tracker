package queue

import (
	"context"
	"encoding/json"
)

// LocalFieldEdit edits a queue-local issue field.
func (s *Service) LocalFieldEdit(ctx context.Context, queueKey, fieldKey string, fields map[string]any) (json.RawMessage, error) {
	return s.client.EditLocalField(ctx, queueKey, fieldKey, fields)
}
