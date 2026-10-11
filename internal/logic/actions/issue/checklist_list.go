package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) ChecklistListSafety() safety.Level { return safety.ReadOnly }

// ChecklistList lists the items of an issue's checklist.
func (s *Service) ChecklistList(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.ChecklistListSafety(), func() (json.RawMessage, error) {
		return s.client.ListChecklistItems(ctx, issueKey)
	})
}
