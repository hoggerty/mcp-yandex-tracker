package worklog

import (
	"context"
	"encoding/json"
)

// Delete removes a worklog record from an issue.
func (s *Service) Delete(ctx context.Context, issueKey, worklogID string) (json.RawMessage, error) {
	return s.client.DeleteWorklog(ctx, issueKey, worklogID)
}
