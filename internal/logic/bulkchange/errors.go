package bulkchange

import (
	"context"
	"encoding/json"
)

// Errors retrieves the issues that failed within a bulk-change operation, with error details.
func (s *Service) Errors(ctx context.Context, id string) (json.RawMessage, error) {
	return s.client.GetBulkChangeErrors(ctx, id)
}
