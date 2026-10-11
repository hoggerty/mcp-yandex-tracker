package trigger

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) CreateSafety() safety.Level { return safety.Write }

// Create creates a new trigger on a queue.
func (s *Service) Create(ctx context.Context, queueKey string, fields map[string]any) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.CreateSafety(), func() (json.RawMessage, error) {
		return s.client.CreateTrigger(ctx, queueKey, fields)
	})
}
