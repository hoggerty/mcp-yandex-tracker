package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const LinkDeleteSafety = safety.Destructive

// LinkDelete removes a link between two issues.
func (s *Service) LinkDelete(ctx context.Context, issueKey, linkID string) (json.RawMessage, error) {
	return s.client.DeleteIssueLink(ctx, issueKey, linkID)
}
