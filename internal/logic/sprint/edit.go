package sprint

import (
	"context"
	"encoding/json"
)

// Edit edits a sprint; version must match the sprint's current version.
func (s *Service) Edit(ctx context.Context, id string, version int, fields map[string]any) (json.RawMessage, error) {
	return s.client.EditSprint(ctx, id, version, fields)
}
