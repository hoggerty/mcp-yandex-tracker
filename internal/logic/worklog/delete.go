package worklog

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const DeleteSafety = safety.Destructive

// Delete removes a worklog record from an issue.
func (s *Service) Delete(ctx context.Context, issueKey, worklogID string) (json.RawMessage, error) {
	return s.client.DeleteWorklog(ctx, issueKey, worklogID)
}
