package version

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const GetSafety = safety.ReadOnly

// Get retrieves the parameters of a single version.
func (s *Service) Get(ctx context.Context, id string) (json.RawMessage, error) {
	return s.client.GetVersion(ctx, id)
}
