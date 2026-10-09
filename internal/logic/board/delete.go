package board

import (
	"context"
	"encoding/json"
)

// Delete deletes an agile board.
func (s *Service) Delete(ctx context.Context, id string) (json.RawMessage, error) {
	return s.client.DeleteBoard(ctx, id)
}
