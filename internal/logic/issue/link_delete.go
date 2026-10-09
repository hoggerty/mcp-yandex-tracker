package issue

import (
	"context"
	"encoding/json"
)

// LinkDelete removes a link between two issues.
func (s *Service) LinkDelete(ctx context.Context, issueKey, linkID string) (json.RawMessage, error) {
	return s.client.DeleteIssueLink(ctx, issueKey, linkID)
}
