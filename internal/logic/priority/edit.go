package priority

import (
	"context"
	"encoding/json"
)

// Edit edits an issue priority; version must match the priority's current version.
func (s *Service) Edit(ctx context.Context, id string, version int, fields map[string]any) (json.RawMessage, error) {
	return s.client.EditPriority(ctx, id, version, fields)
}
