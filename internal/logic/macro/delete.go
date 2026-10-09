package macro

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const DeleteSafety = safety.Destructive

// Delete deletes a macro from a queue.
func (s *Service) Delete(ctx context.Context, queueKey, macroID string) (json.RawMessage, error) {
	return s.client.DeleteMacro(ctx, queueKey, macroID)
}
