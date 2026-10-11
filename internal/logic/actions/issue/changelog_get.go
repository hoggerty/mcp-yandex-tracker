package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) GetChangelogSafety() safety.Level { return safety.ReadOnly }

// GetChangelog retrieves the field change history for an issue.
func (s *Service) GetChangelog(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.GetChangelogSafety(), func() (json.RawMessage, error) {
		return s.client.GetChangelog(ctx, issueKey)
	})
}
