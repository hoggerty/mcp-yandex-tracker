package entityaccess

import (
	"context"
	"encoding/json"
)

// AccessGet retrieves an entity's extended access settings (ACL + permission sources).
func (s *Service) AccessGet(ctx context.Context, entityType, id string) (json.RawMessage, error) {
	return s.client.GetEntityAccess(ctx, entityType, id)
}
