package entityaccess

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) PermissionsGetSafety() safety.Level { return safety.ReadOnly }

// PermissionsGet retrieves an entity's resolved access permissions.
func (s *Service) PermissionsGet(ctx context.Context, entityType, id string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.PermissionsGetSafety(), func() (json.RawMessage, error) {
		return s.client.GetEntityPermissions(ctx, entityType, id)
	})
}
