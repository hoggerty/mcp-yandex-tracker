package issue

import (
	"context"
	"encoding/json"
)

func (s *Service) Update(ctx context.Context, issueKey string, fields map[string]any) (json.RawMessage, error) {
	return s.client.UpdateIssue(ctx, issueKey, fields)
}
