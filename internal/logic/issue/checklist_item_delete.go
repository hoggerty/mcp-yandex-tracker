package issue

import (
	"context"
	"encoding/json"
)

// ChecklistItemDelete deletes a single item from an issue's checklist.
func (s *Service) ChecklistItemDelete(ctx context.Context, issueKey, itemID string) (json.RawMessage, error) {
	return s.client.DeleteChecklistItem(ctx, issueKey, itemID)
}
