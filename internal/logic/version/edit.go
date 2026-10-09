package version

import (
	"context"
	"encoding/json"
)

// Edit edits a version.
func (s *Service) Edit(ctx context.Context, id string, fields map[string]any) (json.RawMessage, error) {
	return s.client.EditVersion(ctx, id, fields)
}
