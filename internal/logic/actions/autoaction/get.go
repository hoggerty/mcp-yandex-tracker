package autoaction

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) GetSafety() safety.Level { return safety.ReadOnly }

// Get retrieves the parameters of an auto action.
func (s *Service) Get(ctx context.Context, queueKey, autoActionID string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.GetSafety(), func() (json.RawMessage, error) {
		return s.client.GetAutoAction(ctx, queueKey, autoActionID)
	})
}
