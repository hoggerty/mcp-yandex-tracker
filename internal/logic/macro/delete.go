package macro

import (
	"context"
	"encoding/json"
)

// Delete deletes a macro from a queue.
func (s *Service) Delete(ctx context.Context, queueKey, macroID string) (json.RawMessage, error) {
	return s.client.DeleteMacro(ctx, queueKey, macroID)
}
