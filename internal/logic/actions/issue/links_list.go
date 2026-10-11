package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) ListLinksSafety() safety.Level { return safety.ReadOnly }

// ListLinks lists the issue links (relates to, blocks, etc.) for an issue.
func (s *Service) ListLinks(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.ListLinksSafety(), func() (json.RawMessage, error) {
		return s.client.ListIssueLinks(ctx, issueKey)
	})
}
