package field

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const EditNameSafety = safety.Write

// EditName renames a global field; version must match the field's current version.
func (s *Service) EditName(ctx context.Context, id string, version int, name map[string]any) (json.RawMessage, error) {
	return s.client.PatchField(ctx, id, version, map[string]any{"name": name})
}
