package entityaccess

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) AccessEditSafety() safety.Level { return safety.Write }

// AccessEdit changes an entity's extended access settings (ACL grant/revoke, permission sources).
func (s *Service) AccessEdit(ctx context.Context, entityType, id string, body map[string]any) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.AccessEditSafety(), func() (json.RawMessage, error) {
		return s.client.EditEntityAccess(ctx, entityType, id, body)
	})
}
