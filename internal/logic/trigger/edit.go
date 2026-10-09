package trigger

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const EditSafety = safety.Write

// Edit edits a trigger; version must match the trigger's current version.
func (s *Service) Edit(ctx context.Context, queueKey, triggerID string, version int, fields map[string]any) (json.RawMessage, error) {
	return s.client.PatchTrigger(ctx, queueKey, triggerID, version, fields)
}
