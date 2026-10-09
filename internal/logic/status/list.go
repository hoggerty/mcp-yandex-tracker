package status

import (
	"context"
	"encoding/json"
)

// List lists the issue statuses available in the organization.
func (s *Service) List(ctx context.Context) (json.RawMessage, error) {
	return s.client.ListStatuses(ctx)
}
