package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) ChecklistItemDeleteSafety() safety.Level { return safety.Destructive }

// ChecklistItemDelete deletes a single item from an issue's checklist.
func (s *Service) ChecklistItemDelete(ctx context.Context, issueKey, itemID string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.ChecklistItemDeleteSafety(), func() (json.RawMessage, error) {
		return s.client.DeleteChecklistItem(ctx, issueKey, itemID)
	})
}
