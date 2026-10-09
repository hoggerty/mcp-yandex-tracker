package importdata

import (
	"context"
	"encoding/json"
)

func (s *Service) Worklog(ctx context.Context, issueID string, fields map[string]any) (json.RawMessage, error) {
	return s.client.ImportWorklog(ctx, issueID, fields)
}
