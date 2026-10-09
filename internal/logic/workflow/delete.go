package workflow

import (
	"context"
	"encoding/json"
)

// Delete deletes a workflow.
func (s *Service) Delete(ctx context.Context, id string) (json.RawMessage, error) {
	return s.client.DeleteWorkflow(ctx, id)
}
