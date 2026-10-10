package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) ChecklistItemDeleteSafety() safety.Level { return safety.Destructive }

// ChecklistItemDelete deletes a single item from an issue's checklist.
func (s *Service) ChecklistItemDelete(ctx context.Context, issueKey, itemID string) (json.RawMessage, error) {
	if err := s.guard.Check(s.ChecklistItemDeleteSafety()); err != nil {
		return nil, err
	}

	return s.client.DeleteChecklistItem(ctx, issueKey, itemID)
}
