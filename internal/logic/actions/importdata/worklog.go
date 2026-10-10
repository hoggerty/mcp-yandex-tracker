package importdata

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) WorklogSafety() safety.Level { return safety.Write }

func (s *Service) Worklog(ctx context.Context, issueID string, fields map[string]any) (json.RawMessage, error) {
	if err := s.guard.Check(s.WorklogSafety()); err != nil {
		return nil, err
	}

	return s.client.ImportWorklog(ctx, issueID, fields)
}
