package component

import (
	"context"
	"encoding/json"
)

// Edit edits a component; version must match the component's current version.
func (s *Service) Edit(ctx context.Context, id string, version int, fields map[string]any) (json.RawMessage, error) {
	return s.client.EditComponent(ctx, id, version, fields)
}
