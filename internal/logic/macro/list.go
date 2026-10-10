package macro

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) ListSafety() safety.Level { return safety.ReadOnly }

// List lists the macros configured on a queue.
func (s *Service) List(ctx context.Context, queueKey string) (json.RawMessage, error) {
	if err := s.guard.Check(s.ListSafety()); err != nil {
		return nil, err
	}

	return s.client.ListMacros(ctx, queueKey)
}
