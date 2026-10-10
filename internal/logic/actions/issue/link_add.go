package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) LinkIssuesSafety() safety.Level { return safety.Write }

func (s *Service) LinkIssues(ctx context.Context, issueKey, relationship, targetIssueKey string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.LinkIssuesSafety(), func() (json.RawMessage, error) {
		return s.client.LinkIssues(ctx, issueKey, relationship, targetIssueKey)
	})
}
