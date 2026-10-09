package macro

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const EditSafety = safety.Write

// Edit edits a macro.
func (s *Service) Edit(ctx context.Context, queueKey, macroID string, fields map[string]any) (json.RawMessage, error) {
	return s.client.PatchMacro(ctx, queueKey, macroID, fields)
}
