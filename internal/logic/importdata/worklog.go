package importdata

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const WorklogSafety = safety.Write

func (s *Service) Worklog(ctx context.Context, issueID string, fields map[string]any) (json.RawMessage, error) {
	return s.client.ImportWorklog(ctx, issueID, fields)
}
