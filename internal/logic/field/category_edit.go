package field

import (
	"context"
	"encoding/json"
)

// CategoryEdit edits a field category; version must match the category's current version.
func (s *Service) CategoryEdit(ctx context.Context, id string, version int, fields map[string]any) (json.RawMessage, error) {
	return s.client.PatchFieldCategory(ctx, id, version, fields)
}
