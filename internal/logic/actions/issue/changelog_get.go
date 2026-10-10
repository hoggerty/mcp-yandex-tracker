package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) GetChangelogSafety() safety.Level { return safety.ReadOnly }

// GetChangelog retrieves the field change history for an issue.
func (s *Service) GetChangelog(ctx context.Context, issueKey string) (json.RawMessage, error) {
	if err := s.guard.Check(s.GetChangelogSafety()); err != nil {
		return nil, err
	}

	return s.client.GetChangelog(ctx, issueKey)
}
