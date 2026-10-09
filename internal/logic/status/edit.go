package status

import (
	"context"
	"encoding/json"
)

// Edit edits an issue status; version must match the status's current version.
func (s *Service) Edit(ctx context.Context, id string, version int, fields map[string]any) (json.RawMessage, error) {
	return s.client.EditStatus(ctx, id, version, fields)
}
