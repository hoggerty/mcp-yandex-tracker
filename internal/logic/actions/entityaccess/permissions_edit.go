package entityaccess

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) PermissionsEditSafety() safety.Level { return safety.Write }

// PermissionsEdit changes an entity's access permissions (grant/revoke users, groups, roles).
func (s *Service) PermissionsEdit(ctx context.Context, entityType, id string, body map[string]any) (json.RawMessage, error) {
	if err := s.guard.Check(s.PermissionsEditSafety()); err != nil {
		return nil, err
	}

	return s.client.EditEntityPermissions(ctx, entityType, id, body)
}
