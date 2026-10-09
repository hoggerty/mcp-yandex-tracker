package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const ListSafety = safety.ReadOnly

// List lists the queues visible to the current user.
func (s *Service) List(ctx context.Context, perPage, page int) (json.RawMessage, error) {
	return s.client.ListQueues(ctx, perPage, page)
}
