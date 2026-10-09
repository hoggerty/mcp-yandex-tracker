package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const GetSafety = safety.ReadOnly

func (s *Service) Get(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return s.client.GetQueue(ctx, queueKey)
}
