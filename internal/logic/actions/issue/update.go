package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) UpdateSafety() safety.Level { return safety.Write }

func (s *Service) Update(ctx context.Context, issueKey string, fields map[string]any) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.UpdateSafety(), func() (json.RawMessage, error) {
		return s.client.UpdateIssue(ctx, issueKey, fields)
	})
}
