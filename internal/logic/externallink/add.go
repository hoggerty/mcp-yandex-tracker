package externallink

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const AddSafety = safety.Write

// Add links an issue to an object in an external application.
func (s *Service) Add(ctx context.Context, issueKey, relationship, key, origin string) (json.RawMessage, error) {
	return s.client.AddExternalLink(ctx, issueKey, relationship, key, origin)
}
