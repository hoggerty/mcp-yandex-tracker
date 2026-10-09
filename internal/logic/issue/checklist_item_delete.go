package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const ChecklistItemDeleteSafety = safety.Destructive

// ChecklistItemDelete deletes a single item from an issue's checklist.
func (s *Service) ChecklistItemDelete(ctx context.Context, issueKey, itemID string) (json.RawMessage, error) {
	return s.client.DeleteChecklistItem(ctx, issueKey, itemID)
}
