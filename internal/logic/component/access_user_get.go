package component

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const AccessUserGetSafety = safety.ReadOnly

// AccessUserGet retrieves a user's permissions for a component.
func (s *Service) AccessUserGet(ctx context.Context, componentID, userID string) (json.RawMessage, error) {
	return s.client.GetComponentUserAccess(ctx, componentID, userID)
}
