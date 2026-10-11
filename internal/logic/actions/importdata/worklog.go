package importdata

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) WorklogSafety() safety.Level { return safety.Write }

func (s *Service) Worklog(ctx context.Context, issueID string, fields map[string]any) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.WorklogSafety(), func() (json.RawMessage, error) {
		return s.client.ImportWorklog(ctx, issueID, fields)
	})
}
