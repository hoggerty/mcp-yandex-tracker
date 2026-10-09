package worklog

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const EditSafety = safety.Write

// Edit changes the duration and/or comment of an existing worklog record.
func (s *Service) Edit(ctx context.Context, issueKey, worklogID, duration, comment string) (json.RawMessage, error) {
	return s.client.EditWorklog(ctx, issueKey, worklogID, duration, comment)
}
