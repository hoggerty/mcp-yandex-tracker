package project

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) QueuesListSafety() safety.Level { return safety.ReadOnly }

func (s *Service) QueuesList(ctx context.Context, id, expand string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.QueuesListSafety(), func() (json.RawMessage, error) {
		return s.client.ListProjectQueues(ctx, id, expand)
	})
}
