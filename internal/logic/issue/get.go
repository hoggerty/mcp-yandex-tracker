package issue

import (
	"context"
	"encoding/json"
)

// Get retrieves a single issue by its key, e.g. "QUEUE-123".
func (s *Service) Get(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return s.client.GetIssue(ctx, issueKey)
}
