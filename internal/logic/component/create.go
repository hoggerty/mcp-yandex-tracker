package component

import (
	"context"
	"encoding/json"
)

// Create creates a new component on a queue.
func (s *Service) Create(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return s.client.CreateComponent(ctx, fields)
}
