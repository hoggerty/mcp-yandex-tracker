package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) CommentDeleteSafety() safety.Level { return safety.Destructive }

// CommentDelete deletes an issue comment.
func (s *Service) CommentDelete(ctx context.Context, issueKey, commentID string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.CommentDeleteSafety(), func() (json.RawMessage, error) {
		return s.client.DeleteComment(ctx, issueKey, commentID)
	})
}
