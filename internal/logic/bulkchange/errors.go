package bulkchange

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const ErrorsSafety = safety.ReadOnly

// Errors retrieves the issues that failed within a bulk-change operation, with error details.
func (s *Service) Errors(ctx context.Context, id string) (json.RawMessage, error) {
	return s.client.GetBulkChangeErrors(ctx, id)
}
