package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const DeleteSafety = safety.Destructive

// Delete deletes a queue.
func (s *Service) Delete(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return s.client.DeleteQueue(ctx, queueKey)
}
