package board

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) ListSafety() safety.Level { return safety.ReadOnly }

// List lists the agile boards visible to the current user, sorted by ID;
// pass the last board's ID from the previous page as id to continue.
func (s *Service) List(ctx context.Context, perPage, id int) (json.RawMessage, error) {
	if err := s.guard.Check(s.ListSafety()); err != nil {
		return nil, err
	}

	return s.client.ListBoards(ctx, perPage, id)
}
