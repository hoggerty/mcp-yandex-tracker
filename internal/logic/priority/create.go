package priority

import (
	"context"
	"encoding/json"
)

// Create creates a new issue priority.
func (s *Service) Create(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return s.client.CreatePriority(ctx, fields)
}
