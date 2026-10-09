package externallink

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const DeleteSafety = safety.Destructive

// Delete removes an issue's link to an external application object.
func (s *Service) Delete(ctx context.Context, issueKey, linkID string) (json.RawMessage, error) {
	return s.client.DeleteExternalLink(ctx, issueKey, linkID)
}
