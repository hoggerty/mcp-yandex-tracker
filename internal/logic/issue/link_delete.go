package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) LinkDeleteSafety() safety.Level { return safety.Destructive }

// LinkDelete removes a link between two issues.
func (s *Service) LinkDelete(ctx context.Context, issueKey, linkID string) (json.RawMessage, error) {
	if err := s.guard.Check(s.LinkDeleteSafety()); err != nil {
		return nil, err
	}

	return s.client.DeleteIssueLink(ctx, issueKey, linkID)
}
