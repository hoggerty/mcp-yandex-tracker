package bulkchange

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) ErrorsSafety() safety.Level { return safety.ReadOnly }

// Errors retrieves the issues that failed within a bulk-change operation, with error details.
func (s *Service) Errors(ctx context.Context, id string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.ErrorsSafety(), func() (json.RawMessage, error) {
		return s.client.GetBulkChangeErrors(ctx, id)
	})
}
