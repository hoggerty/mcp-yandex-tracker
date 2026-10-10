package priority

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) ListSafety() safety.Level { return safety.ReadOnly }

// List lists the issue priorities available in the organization. If
// localized is false, names are returned in all supported languages.
func (s *Service) List(ctx context.Context, localized bool) (json.RawMessage, error) {
	if err := s.guard.Check(s.ListSafety()); err != nil {
		return nil, err
	}

	return s.client.ListPriorities(ctx, localized)
}
