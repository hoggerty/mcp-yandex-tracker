package issue

import (
	"context"
	"encoding/json"
)

// ChecklistEdit bulk-edits items of an issue's checklist; each item must include its id.
func (s *Service) ChecklistEdit(ctx context.Context, issueKey string, items []map[string]any) (json.RawMessage, error) {
	return s.client.EditChecklistItems(ctx, issueKey, items)
}
