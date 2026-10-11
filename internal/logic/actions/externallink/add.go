package externallink

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) AddSafety() safety.Level { return safety.Write }

// Add links an issue to an object in an external application.
func (s *Service) Add(ctx context.Context, issueKey, relationship, key, origin string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.AddSafety(), func() (json.RawMessage, error) {
		return s.client.AddExternalLink(ctx, issueKey, relationship, key, origin)
	})
}
