package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const ChecklistListSafety = safety.ReadOnly

// ChecklistList lists the items of an issue's checklist.
func (s *Service) ChecklistList(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return s.client.ListChecklistItems(ctx, issueKey)
}
