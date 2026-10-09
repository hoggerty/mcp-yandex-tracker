package resolution

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const ListSafety = safety.ReadOnly

// List lists the issue resolutions available in the organization.
func (s *Service) List(ctx context.Context) (json.RawMessage, error) {
	return s.client.ListResolutions(ctx)
}
