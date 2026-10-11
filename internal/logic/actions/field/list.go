package field

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) ListSafety() safety.Level { return safety.ReadOnly }

// List returns all global fields of the organization.
func (s *Service) List(ctx context.Context) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.ListSafety(), func() (json.RawMessage, error) {
		return s.client.ListFields(ctx)
	})
}
