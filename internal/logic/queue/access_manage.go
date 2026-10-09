package queue

import (
	"context"
	"encoding/json"
)

// AccessManage grants or revokes queue-level permissions (create/write/read/grant) for users, groups, or roles.
func (s *Service) AccessManage(ctx context.Context, queueKey string, permissions map[string]any) (json.RawMessage, error) {
	return s.client.ManageQueueAccess(ctx, queueKey, permissions)
}
