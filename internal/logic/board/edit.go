package board

import (
	"context"
	"encoding/json"
)

// Edit edits an agile board; version must match the board's current version.
func (s *Service) Edit(ctx context.Context, id string, version int, fields map[string]any) (json.RawMessage, error) {
	return s.client.EditBoard(ctx, id, version, fields)
}
