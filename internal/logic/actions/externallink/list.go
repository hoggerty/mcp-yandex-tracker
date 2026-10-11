package externallink

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) ListSafety() safety.Level { return safety.ReadOnly }

// List returns an issue's links to objects in external applications.
func (s *Service) List(ctx context.Context, issueKey string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.ListSafety(), func() (json.RawMessage, error) {
		return s.client.ListExternalLinks(ctx, issueKey)
	})
}
