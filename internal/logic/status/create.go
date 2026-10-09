package status

import (
	"context"
	"encoding/json"
)

// Create creates a new issue status.
func (s *Service) Create(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return s.client.CreateStatus(ctx, fields)
}
