package entityaccess

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) AccessGetSafety() safety.Level { return safety.ReadOnly }

// AccessGet retrieves an entity's extended access settings (ACL + permission sources).
func (s *Service) AccessGet(ctx context.Context, entityType, id string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.AccessGetSafety(), func() (json.RawMessage, error) {
		return s.client.GetEntityAccess(ctx, entityType, id)
	})
}
