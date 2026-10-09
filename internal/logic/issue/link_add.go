package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const LinkIssuesSafety = safety.Write

func (s *Service) LinkIssues(ctx context.Context, issueKey, relationship, targetIssueKey string) (json.RawMessage, error) {
	return s.client.LinkIssues(ctx, issueKey, relationship, targetIssueKey)
}
