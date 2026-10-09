package issuetype

import (
	"context"
	"encoding/json"
)

// Create creates a new issue type.
func (s *Service) Create(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return s.client.CreateIssueType(ctx, fields)
}
