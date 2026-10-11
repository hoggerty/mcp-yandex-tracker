package entitycomment

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) EditSafety() safety.Level { return safety.Write }

// Edit edits an entity comment.
func (s *Service) Edit(ctx context.Context, entityType, id, commentID string, fields map[string]any) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.EditSafety(), func() (json.RawMessage, error) {
		return s.client.EditEntityComment(ctx, entityType, id, commentID, fields)
	})
}
