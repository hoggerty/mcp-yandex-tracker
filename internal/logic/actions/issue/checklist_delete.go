package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) ChecklistDeleteSafety() safety.Level { return safety.Destructive }

// ChecklistDelete deletes an issue's entire checklist.
func (s *Service) ChecklistDelete(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.ChecklistDeleteSafety(), func() (json.RawMessage, error) {
		return s.client.DeleteChecklist(ctx, issueKey)
	})
}
