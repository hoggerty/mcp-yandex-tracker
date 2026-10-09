package issue

import (
	"context"
	"encoding/json"
)

// GetChangelog retrieves the field change history for an issue.
func (s *Service) GetChangelog(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return s.client.GetChangelog(ctx, issueKey)
}
