package issue

import (
	"context"
	"encoding/json"
)

// ChecklistList lists the items of an issue's checklist.
func (s *Service) ChecklistList(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return s.client.ListChecklistItems(ctx, issueKey)
}
