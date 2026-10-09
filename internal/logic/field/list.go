package field

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const ListSafety = safety.ReadOnly

// List returns all global fields of the organization.
func (s *Service) List(ctx context.Context) (json.RawMessage, error) {
	return s.client.ListFields(ctx)
}
