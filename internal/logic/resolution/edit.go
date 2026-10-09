package resolution

import (
	"context"
	"encoding/json"
)

// Edit edits an issue resolution; version must match the resolution's current version.
func (s *Service) Edit(ctx context.Context, id string, version int, fields map[string]any) (json.RawMessage, error) {
	return s.client.EditResolution(ctx, id, version, fields)
}
