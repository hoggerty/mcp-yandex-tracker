package component

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) AccessGroupGetSafety() safety.Level { return safety.ReadOnly }

// AccessGroupGet retrieves a group's permissions for a component.
func (s *Service) AccessGroupGet(ctx context.Context, componentID, groupID string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.AccessGroupGetSafety(), func() (json.RawMessage, error) {
		return s.client.GetComponentGroupAccess(ctx, componentID, groupID)
	})
}
