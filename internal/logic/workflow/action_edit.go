package workflow

import (
	"context"
	"encoding/json"
)

// ActionEdit edits a single action of a workflow step; version must match the workflow's current version.
func (s *Service) ActionEdit(ctx context.Context, id, status, actionID string, version int, fields map[string]any) (json.RawMessage, error) {
	return s.client.PatchWorkflowAction(ctx, id, status, actionID, version, fields)
}
