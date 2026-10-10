package externallink

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) AddSafety() safety.Level { return safety.Write }

// Add links an issue to an object in an external application.
func (s *Service) Add(ctx context.Context, issueKey, relationship, key, origin string) (json.RawMessage, error) {
	if err := s.guard.Check(s.AddSafety()); err != nil {
		return nil, err
	}

	return s.client.AddExternalLink(ctx, issueKey, relationship, key, origin)
}
