package bulkchange

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const TransitionSafety = safety.Destructive

// Transition starts a background operation that executes a workflow transition on many issues.
func (s *Service) Transition(ctx context.Context, transition string, issues []string, values map[string]any, notify bool) (json.RawMessage, error) {
	return s.client.BulkTransition(ctx, transition, issues, values, notify)
}
