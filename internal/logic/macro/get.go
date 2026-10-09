package macro

import (
	"context"
	"encoding/json"
)

// Get retrieves the parameters of a single macro.
func (s *Service) Get(ctx context.Context, queueKey, macroID string) (json.RawMessage, error) {
	return s.client.GetMacro(ctx, queueKey, macroID)
}
