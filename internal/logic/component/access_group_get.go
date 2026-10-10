package component

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) AccessGroupGetSafety() safety.Level { return safety.ReadOnly }

// AccessGroupGet retrieves a group's permissions for a component.
func (s *Service) AccessGroupGet(ctx context.Context, componentID, groupID string) (json.RawMessage, error) {
	if err := s.guard.Check(s.AccessGroupGetSafety()); err != nil {
		return nil, err
	}

	return s.client.GetComponentGroupAccess(ctx, componentID, groupID)
}
