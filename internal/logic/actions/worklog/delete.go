package worklog

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) DeleteSafety() safety.Level { return safety.Destructive }

// Delete removes a worklog record from an issue.
func (s *Service) Delete(ctx context.Context, issueKey, worklogID string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.DeleteSafety(), func() (json.RawMessage, error) {
		return s.client.DeleteWorklog(ctx, issueKey, worklogID)
	})
}
