package field

import (
	"context"
	"encoding/json"
)

// EditValues edits a global field's non-name parameters (category, order, description,
// visibility, possible values); version must match the field's current version.
func (s *Service) EditValues(ctx context.Context, id string, version int, fields map[string]any) (json.RawMessage, error) {
	return s.client.PatchField(ctx, id, version, fields)
}
