package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) UpdateSafety() safety.Level { return safety.Write }

func (s *Service) Update(ctx context.Context, issueKey string, fields map[string]any) (json.RawMessage, error) {
	if err := s.guard.Check(s.UpdateSafety()); err != nil {
		return nil, err
	}

	return s.client.UpdateIssue(ctx, issueKey, fields)
}
