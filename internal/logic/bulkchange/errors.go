package bulkchange

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) ErrorsSafety() safety.Level { return safety.ReadOnly }

// Errors retrieves the issues that failed within a bulk-change operation, with error details.
func (s *Service) Errors(ctx context.Context, id string) (json.RawMessage, error) {
	if err := s.guard.Check(s.ErrorsSafety()); err != nil {
		return nil, err
	}

	return s.client.GetBulkChangeErrors(ctx, id)
}
