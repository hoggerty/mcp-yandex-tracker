package project

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) QueuesListSafety() safety.Level { return safety.ReadOnly }

func (s *Service) QueuesList(ctx context.Context, id, expand string) (json.RawMessage, error) {
	if err := s.guard.Check(s.QueuesListSafety()); err != nil {
		return nil, err
	}

	return s.client.ListProjectQueues(ctx, id, expand)
}
