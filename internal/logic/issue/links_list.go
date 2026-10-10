package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) ListLinksSafety() safety.Level { return safety.ReadOnly }

// ListLinks lists the issue links (relates to, blocks, etc.) for an issue.
func (s *Service) ListLinks(ctx context.Context, issueKey string) (json.RawMessage, error) {
	if err := s.guard.Check(s.ListLinksSafety()); err != nil {
		return nil, err
	}

	return s.client.ListIssueLinks(ctx, issueKey)
}
