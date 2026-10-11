package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) AccessManageSafety() safety.Level { return safety.Destructive }

// AccessManage grants or revokes queue-level permissions (create/write/read/grant) for users, groups, or roles.
func (s *Service) AccessManage(ctx context.Context, queueKey string, permissions map[string]any) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.AccessManageSafety(), func() (json.RawMessage, error) {
		return s.client.ManageQueueAccess(ctx, queueKey, permissions)
	})
}
