package workflow

import (
	"context"
	"encoding/json"
)

// Create creates a new workflow.
func (s *Service) Create(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return s.client.CreateWorkflow(ctx, fields)
}
