package workflow

import (
	"context"
	"encoding/json"
)

// Edit edits a workflow; version must match the workflow's current version.
func (s *Service) Edit(ctx context.Context, id string, version int, fields map[string]any) (json.RawMessage, error) {
	return s.client.PatchWorkflow(ctx, id, version, fields)
}
