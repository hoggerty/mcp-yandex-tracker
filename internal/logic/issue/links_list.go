package issue

import (
	"context"
	"encoding/json"
)

// ListLinks lists the issue links (relates to, blocks, etc.) for an issue.
func (s *Service) ListLinks(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return s.client.ListIssueLinks(ctx, issueKey)
}
