package board

import (
	"context"
	"encoding/json"
)

// Create creates a new agile board.
func (s *Service) Create(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return s.client.CreateBoard(ctx, fields)
}
