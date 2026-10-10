package externallink

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) DeleteSafety() safety.Level { return safety.Destructive }

// Delete removes an issue's link to an external application object.
func (s *Service) Delete(ctx context.Context, issueKey, linkID string) (json.RawMessage, error) {
	if err := s.guard.Check(s.DeleteSafety()); err != nil {
		return nil, err
	}

	return s.client.DeleteExternalLink(ctx, issueKey, linkID)
}
