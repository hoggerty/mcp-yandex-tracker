package entitycomment

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const DeleteSafety = safety.Destructive

// Delete deletes an entity comment.
func (s *Service) Delete(ctx context.Context, entityType, id, commentID string) (json.RawMessage, error) {
	return s.client.DeleteEntityComment(ctx, entityType, id, commentID)
}
