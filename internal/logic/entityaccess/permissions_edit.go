package entityaccess

import (
	"context"
	"encoding/json"
)

// PermissionsEdit changes an entity's access permissions (grant/revoke users, groups, roles).
func (s *Service) PermissionsEdit(ctx context.Context, entityType, id string, body map[string]any) (json.RawMessage, error) {
	return s.client.EditEntityPermissions(ctx, entityType, id, body)
}
