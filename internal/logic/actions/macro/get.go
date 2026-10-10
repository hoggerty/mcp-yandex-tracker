package macro

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) GetSafety() safety.Level { return safety.ReadOnly }

// Get retrieves the parameters of a single macro.
func (s *Service) Get(ctx context.Context, queueKey, macroID string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.GetSafety(), func() (json.RawMessage, error) {
		return s.client.GetMacro(ctx, queueKey, macroID)
	})
}
