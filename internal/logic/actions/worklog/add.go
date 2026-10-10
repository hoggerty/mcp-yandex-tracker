package worklog

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) AddSafety() safety.Level { return safety.Write }

// Add records time spent on an issue.
func (s *Service) Add(ctx context.Context, issueKey, start, duration, comment string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.AddSafety(), func() (json.RawMessage, error) {
		return s.client.AddWorklog(ctx, issueKey, start, duration, comment)
	})
}
