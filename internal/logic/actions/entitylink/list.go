package entitylink

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) ListSafety() safety.Level { return safety.ReadOnly }

// List lists the entities linked to an entity. fields is optional (comma-separated linked entity fields).
func (s *Service) List(ctx context.Context, entityType, id, fields string) (json.RawMessage, error) {
	if err := s.guard.Check(s.ListSafety()); err != nil {
		return nil, err
	}

	return s.client.ListEntityLinks(ctx, entityType, id, fields)
}
