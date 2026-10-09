package issue

import (
	"context"
	"encoding/json"
)

func (s *Service) LinkIssues(ctx context.Context, issueKey, relationship, targetIssueKey string) (json.RawMessage, error) {
	return s.client.LinkIssues(ctx, issueKey, relationship, targetIssueKey)
}
