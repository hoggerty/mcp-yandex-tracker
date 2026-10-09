package queue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const AccessManageSafety = safety.Destructive

// AccessManage grants or revokes queue-level permissions (create/write/read/grant) for users, groups, or roles.
func (s *Service) AccessManage(ctx context.Context, queueKey string, permissions map[string]any) (json.RawMessage, error) {
	return s.client.ManageQueueAccess(ctx, queueKey, permissions)
}
