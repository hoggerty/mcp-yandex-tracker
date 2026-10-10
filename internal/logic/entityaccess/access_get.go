package entityaccess

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) AccessGetSafety() safety.Level { return safety.ReadOnly }

// AccessGet retrieves an entity's extended access settings (ACL + permission sources).
func (s *Service) AccessGet(ctx context.Context, entityType, id string) (json.RawMessage, error) {
	if err := s.guard.Check(s.AccessGetSafety()); err != nil {
		return nil, err
	}

	return s.client.GetEntityAccess(ctx, entityType, id)
}
