package entityaccess

import (
	"context"
	"encoding/json"
)

// PermissionsGet retrieves an entity's resolved access permissions.
func (s *Service) PermissionsGet(ctx context.Context, entityType, id string) (json.RawMessage, error) {
	return s.client.GetEntityPermissions(ctx, entityType, id)
}
