package externallink

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) DeleteSafety() safety.Level { return safety.Destructive }

// Delete removes an issue's link to an external application object.
func (s *Service) Delete(ctx context.Context, issueKey, linkID string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.DeleteSafety(), func() (json.RawMessage, error) {
		return s.client.DeleteExternalLink(ctx, issueKey, linkID)
	})
}
