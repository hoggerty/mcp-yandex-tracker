package project

import (
	"context"
	"encoding/json"
)

func (s *Service) QueuesList(ctx context.Context, id, expand string) (json.RawMessage, error) {
	return s.client.ListProjectQueues(ctx, id, expand)
}
