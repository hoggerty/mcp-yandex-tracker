package trigger

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const ListSafety = safety.ReadOnly

// List lists the triggers configured on a queue.
func (s *Service) List(ctx context.Context, queueKey string, perPage int, lastID string) (json.RawMessage, error) {
	return s.client.ListTriggers(ctx, queueKey, perPage, lastID)
}
