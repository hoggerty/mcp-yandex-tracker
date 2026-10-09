package status

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const ListSafety = safety.ReadOnly

// List lists the issue statuses available in the organization.
func (s *Service) List(ctx context.Context) (json.RawMessage, error) {
	return s.client.ListStatuses(ctx)
}
