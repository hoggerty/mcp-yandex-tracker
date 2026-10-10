package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) CommentGetSafety() safety.Level { return safety.ReadOnly }

// CommentGet retrieves a single issue comment.
func (s *Service) CommentGet(ctx context.Context, issueKey, commentID string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.CommentGetSafety(), func() (json.RawMessage, error) {
		return s.client.GetComment(ctx, issueKey, commentID)
	})
}
