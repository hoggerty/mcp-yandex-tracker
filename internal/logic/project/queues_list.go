package project

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const QueuesListSafety = safety.ReadOnly

func (s *Service) QueuesList(ctx context.Context, id, expand string) (json.RawMessage, error) {
	return s.client.ListProjectQueues(ctx, id, expand)
}
