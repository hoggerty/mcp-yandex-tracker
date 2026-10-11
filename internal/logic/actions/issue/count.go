package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) CountSafety() safety.Level { return safety.ReadOnly }

func (s *Service) Count(ctx context.Context, query string, filter map[string]any) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.CountSafety(), func() (json.RawMessage, error) {
		return s.client.CountIssues(ctx, query, filter)
	})
}
