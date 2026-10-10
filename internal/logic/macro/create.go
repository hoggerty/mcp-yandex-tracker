package macro

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) CreateSafety() safety.Level { return safety.Write }

// Create creates a new macro on a queue.
func (s *Service) Create(ctx context.Context, queueKey string, fields map[string]any) (json.RawMessage, error) {
	if err := s.guard.Check(s.CreateSafety()); err != nil {
		return nil, err
	}

	return s.client.CreateMacro(ctx, queueKey, fields)
}
