package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) CountSafety() safety.Level { return safety.ReadOnly }

func (s *Service) Count(ctx context.Context, query string, filter map[string]any) (json.RawMessage, error) {
	if err := s.guard.Check(s.CountSafety()); err != nil {
		return nil, err
	}

	return s.client.CountIssues(ctx, query, filter)
}
