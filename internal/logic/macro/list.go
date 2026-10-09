package macro

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const ListSafety = safety.ReadOnly

// List lists the macros configured on a queue.
func (s *Service) List(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return s.client.ListMacros(ctx, queueKey)
}
