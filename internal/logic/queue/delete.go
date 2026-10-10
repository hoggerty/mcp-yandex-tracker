package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) DeleteSafety() safety.Level { return safety.Destructive }

// Delete deletes a queue.
func (s *Service) Delete(ctx context.Context, queueKey string) (json.RawMessage, error) {
	if err := s.guard.Check(s.DeleteSafety()); err != nil {
		return nil, err
	}

	return s.client.DeleteQueue(ctx, queueKey)
}
