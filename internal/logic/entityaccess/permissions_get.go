package entityaccess

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const PermissionsGetSafety = safety.ReadOnly

// PermissionsGet retrieves an entity's resolved access permissions.
func (s *Service) PermissionsGet(ctx context.Context, entityType, id string) (json.RawMessage, error) {
	return s.client.GetEntityPermissions(ctx, entityType, id)
}
