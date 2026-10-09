package macro

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const GetSafety = safety.ReadOnly

// Get retrieves the parameters of a single macro.
func (s *Service) Get(ctx context.Context, queueKey, macroID string) (json.RawMessage, error) {
	return s.client.GetMacro(ctx, queueKey, macroID)
}
