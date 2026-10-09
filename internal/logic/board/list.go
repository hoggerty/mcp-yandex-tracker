package board

import (
	"context"
	"encoding/json"
)

// List lists the agile boards visible to the current user, sorted by ID;
// pass the last board's ID from the previous page as id to continue.
func (s *Service) List(ctx context.Context, perPage, id int) (json.RawMessage, error) {
	return s.client.ListBoards(ctx, perPage, id)
}
