package entitycomment

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) ListSafety() safety.Level { return safety.ReadOnly }

// List lists the comments of an entity. expand is optional (e.g. "all", "html", "attachments", "reactions").
func (s *Service) List(ctx context.Context, entityType, id, expand string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.ListSafety(), func() (json.RawMessage, error) {
		return s.client.ListEntityComments(ctx, entityType, id, expand)
	})
}
