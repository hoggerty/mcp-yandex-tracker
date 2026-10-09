package issuetype

import (
	"context"
	"encoding/json"
)

// Edit edits an issue type; version must match the issue type's current version.
func (s *Service) Edit(ctx context.Context, id string, version int, fields map[string]any) (json.RawMessage, error) {
	return s.client.EditIssueType(ctx, id, version, fields)
}
