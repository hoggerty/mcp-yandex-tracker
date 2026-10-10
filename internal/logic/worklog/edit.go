package worklog

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) EditSafety() safety.Level { return safety.Write }

// Edit changes the duration and/or comment of an existing worklog record.
func (s *Service) Edit(ctx context.Context, issueKey, worklogID, duration, comment string) (json.RawMessage, error) {
	if err := s.guard.Check(s.EditSafety()); err != nil {
		return nil, err
	}

	return s.client.EditWorklog(ctx, issueKey, worklogID, duration, comment)
}
