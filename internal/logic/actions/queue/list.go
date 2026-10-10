package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) ListSafety() safety.Level { return safety.ReadOnly }

// List lists the queues visible to the current user.
func (s *Service) List(ctx context.Context, perPage, page int) (json.RawMessage, error) {
	if err := s.guard.Check(s.ListSafety()); err != nil {
		return nil, err
	}

	return s.client.ListQueues(ctx, perPage, page)
}
