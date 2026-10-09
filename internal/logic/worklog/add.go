package worklog

import (
	"context"
	"encoding/json"
)

// Add records time spent on an issue.
func (s *Service) Add(ctx context.Context, issueKey, start, duration, comment string) (json.RawMessage, error) {
	return s.client.AddWorklog(ctx, issueKey, start, duration, comment)
}
