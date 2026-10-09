package workflow

import (
	"context"
	"encoding/json"
)

// List lists all workflows in the organization.
func (s *Service) List(ctx context.Context) (json.RawMessage, error) {
	return s.client.ListWorkflows(ctx)
}
