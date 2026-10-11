package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) ChecklistEditSafety() safety.Level { return safety.Write }

// ChecklistEdit bulk-edits items of an issue's checklist; each item must include its id.
func (s *Service) ChecklistEdit(ctx context.Context, issueKey string, items []map[string]any) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.ChecklistEditSafety(), func() (json.RawMessage, error) {
		return s.client.EditChecklistItems(ctx, issueKey, items)
	})
}
