package worklog

import (
	"context"
	"encoding/json"
)

// List returns all worklog records for an issue.
func (s *Service) List(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return s.client.ListWorklog(ctx, issueKey)
}
