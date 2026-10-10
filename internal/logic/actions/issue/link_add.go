package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) LinkIssuesSafety() safety.Level { return safety.Write }

func (s *Service) LinkIssues(ctx context.Context, issueKey, relationship, targetIssueKey string) (json.RawMessage, error) {
	if err := s.guard.Check(s.LinkIssuesSafety()); err != nil {
		return nil, err
	}

	return s.client.LinkIssues(ctx, issueKey, relationship, targetIssueKey)
}
