package bulkchange

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) StatusSafety() safety.Level { return safety.ReadOnly }

// Status retrieves the status of a background bulk-change operation.
func (s *Service) Status(ctx context.Context, id string) (json.RawMessage, error) {
	if err := s.guard.Check(s.StatusSafety()); err != nil {
		return nil, err
	}

	return s.client.GetBulkChangeStatus(ctx, id)
}
