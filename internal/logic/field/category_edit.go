package field

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const CategoryEditSafety = safety.Write

// CategoryEdit edits a field category; version must match the category's current version.
func (s *Service) CategoryEdit(ctx context.Context, id string, version int, fields map[string]any) (json.RawMessage, error) {
	return s.client.PatchFieldCategory(ctx, id, version, fields)
}
