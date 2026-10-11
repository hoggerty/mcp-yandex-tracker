package entitycomment

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) DeleteSafety() safety.Level { return safety.Destructive }

// Delete deletes an entity comment.
func (s *Service) Delete(ctx context.Context, entityType, id, commentID string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.DeleteSafety(), func() (json.RawMessage, error) {
		return s.client.DeleteEntityComment(ctx, entityType, id, commentID)
	})
}
