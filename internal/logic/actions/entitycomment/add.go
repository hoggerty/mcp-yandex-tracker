package entitycomment

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) AddSafety() safety.Level { return safety.Write }

// Add adds a comment to an entity.
func (s *Service) Add(ctx context.Context, entityType, id string, fields map[string]any) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.AddSafety(), func() (json.RawMessage, error) {
		return s.client.AddEntityComment(ctx, entityType, id, fields)
	})
}
