package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) ChecklistListSafety() safety.Level { return safety.ReadOnly }

// ChecklistList lists the items of an issue's checklist.
func (s *Service) ChecklistList(ctx context.Context, issueKey string) (json.RawMessage, error) {
	if err := s.guard.Check(s.ChecklistListSafety()); err != nil {
		return nil, err
	}

	return s.client.ListChecklistItems(ctx, issueKey)
}
