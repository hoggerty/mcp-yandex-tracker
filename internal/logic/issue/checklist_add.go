package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) ChecklistAddSafety() safety.Level { return safety.Write }

// ChecklistAdd adds an item to an issue's checklist, creating the checklist if needed.
func (s *Service) ChecklistAdd(ctx context.Context, issueKey string, item map[string]any) (json.RawMessage, error) {
	if err := s.guard.Check(s.ChecklistAddSafety()); err != nil {
		return nil, err
	}

	return s.client.AddChecklistItem(ctx, issueKey, item)
}
