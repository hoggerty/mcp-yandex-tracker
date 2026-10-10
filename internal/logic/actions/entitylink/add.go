package entitylink

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) AddSafety() safety.Level { return safety.Write }

// Add links an entity to one or more other entities.
func (s *Service) Add(ctx context.Context, entityType, id string, body any) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.AddSafety(), func() (json.RawMessage, error) {
		return s.client.AddEntityLinks(ctx, entityType, id, body)
	})
}
