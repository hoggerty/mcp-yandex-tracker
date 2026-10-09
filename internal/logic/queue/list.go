package queue

import (
	"context"
	"encoding/json"
)

// List lists the queues visible to the current user.
func (s *Service) List(ctx context.Context, perPage, page int) (json.RawMessage, error) {
	return s.client.ListQueues(ctx, perPage, page)
}
