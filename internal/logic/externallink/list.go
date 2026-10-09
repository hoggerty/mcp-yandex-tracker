package externallink

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const ListSafety = safety.ReadOnly

// List returns an issue's links to objects in external applications.
func (s *Service) List(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return s.client.ListExternalLinks(ctx, issueKey)
}
