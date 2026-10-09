package issue

import (
	"context"
	"encoding/json"
)

// ChecklistAdd adds an item to an issue's checklist, creating the checklist if needed.
func (s *Service) ChecklistAdd(ctx context.Context, issueKey string, item map[string]any) (json.RawMessage, error) {
	return s.client.AddChecklistItem(ctx, issueKey, item)
}
