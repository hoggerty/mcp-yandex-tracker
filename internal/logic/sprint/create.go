package sprint

import (
	"context"
	"encoding/json"
)

// Create creates a new sprint on a board.
func (s *Service) Create(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return s.client.CreateSprint(ctx, fields)
}
