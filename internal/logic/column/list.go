package column

import (
	"context"
	"encoding/json"
)

// List lists the columns of an agile board.
func (s *Service) List(ctx context.Context, boardID string) (json.RawMessage, error) {
	return s.client.ListColumns(ctx, boardID)
}
