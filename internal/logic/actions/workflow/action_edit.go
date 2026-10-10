package workflow

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) ActionEditSafety() safety.Level { return safety.Write }

// ActionEdit edits a single action of a workflow step; version must match the workflow's current version.
func (s *Service) ActionEdit(ctx context.Context, id, status, actionID string, version int, fields map[string]any) (json.RawMessage, error) {
	if err := s.guard.Check(s.ActionEditSafety()); err != nil {
		return nil, err
	}

	return s.client.PatchWorkflowAction(ctx, id, status, actionID, version, fields)
}
