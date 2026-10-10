package bulkchange

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) UpdateSafety() safety.Level { return safety.Destructive }

// Update starts a background operation that sets field values on many issues.
func (s *Service) Update(ctx context.Context, issues []string, values map[string]any, notify bool) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.UpdateSafety(), func() (json.RawMessage, error) {
		return s.client.BulkUpdate(ctx, issues, values, notify)
	})
}
