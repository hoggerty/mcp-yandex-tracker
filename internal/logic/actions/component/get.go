package component

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) GetSafety() safety.Level { return safety.ReadOnly }

// Get retrieves the parameters of a single component.
func (s *Service) Get(ctx context.Context, id string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.GetSafety(), func() (json.RawMessage, error) {
		return s.client.GetComponent(ctx, id)
	})
}
