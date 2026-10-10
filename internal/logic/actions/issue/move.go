package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) MoveSafety() safety.Level { return safety.Write }

// Move moves an issue to another queue.
func (s *Service) Move(ctx context.Context, issueKey, targetQueue string, notify, notifyAuthor, moveAllFields, initialStatus bool, fields map[string]any) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.MoveSafety(), func() (json.RawMessage, error) {
		return s.client.MoveIssue(ctx, issueKey, targetQueue, notify, notifyAuthor, moveAllFields, initialStatus, fields)
	})
}
