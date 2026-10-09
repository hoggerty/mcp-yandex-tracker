package queue

import (
	"context"
	"encoding/json"
)

func (s *Service) Get(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return s.client.GetQueue(ctx, queueKey)
}
