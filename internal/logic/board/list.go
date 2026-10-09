package board

import (
	"context"
	"encoding/json"
)

// List lists the agile boards visible to the current user.
func (s *Service) List(ctx context.Context) (json.RawMessage, error) {
	return s.client.ListBoards(ctx)
}
