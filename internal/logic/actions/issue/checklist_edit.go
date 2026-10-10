package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) ChecklistEditSafety() safety.Level { return safety.Write }

// ChecklistEdit bulk-edits items of an issue's checklist; each item must include its id.
func (s *Service) ChecklistEdit(ctx context.Context, issueKey string, items []map[string]any) (json.RawMessage, error) {
	if err := s.guard.Check(s.ChecklistEditSafety()); err != nil {
		return nil, err
	}

	return s.client.EditChecklistItems(ctx, issueKey, items)
}
