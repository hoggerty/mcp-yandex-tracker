package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) LinkDeleteSafety() safety.Level { return safety.Destructive }

// LinkDelete removes a link between two issues.
func (s *Service) LinkDelete(ctx context.Context, issueKey, linkID string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.LinkDeleteSafety(), func() (json.RawMessage, error) {
		return s.client.DeleteIssueLink(ctx, issueKey, linkID)
	})
}
