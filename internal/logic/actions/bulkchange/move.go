package bulkchange

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) MoveSafety() safety.Level { return safety.Destructive }

// Move starts a background operation that moves many issues to another queue.
func (s *Service) Move(ctx context.Context, targetQueue string, issues []string, values map[string]any, moveAllFields, initialStatus, notify bool) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.MoveSafety(), func() (json.RawMessage, error) {
		return s.client.BulkMove(ctx, targetQueue, issues, values, moveAllFields, initialStatus, notify)
	})
}
