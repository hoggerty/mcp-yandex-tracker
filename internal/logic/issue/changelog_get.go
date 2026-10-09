package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const GetChangelogSafety = safety.ReadOnly

// GetChangelog retrieves the field change history for an issue.
func (s *Service) GetChangelog(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return s.client.GetChangelog(ctx, issueKey)
}
