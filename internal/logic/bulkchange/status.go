package bulkchange

import (
	"context"
	"encoding/json"
)

// Status retrieves the status of a background bulk-change operation.
func (s *Service) Status(ctx context.Context, id string) (json.RawMessage, error) {
	return s.client.GetBulkChangeStatus(ctx, id)
}
