package trigger

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) GetSafety() safety.Level { return safety.ReadOnly }

// Get retrieves the parameters of a single trigger.
func (s *Service) Get(ctx context.Context, queueKey, triggerID string) (json.RawMessage, error) {
	if err := s.guard.Check(s.GetSafety()); err != nil {
		return nil, err
	}

	return s.client.GetTrigger(ctx, queueKey, triggerID)
}
