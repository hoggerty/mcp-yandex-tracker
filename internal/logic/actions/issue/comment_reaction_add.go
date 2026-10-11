package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/mcp-yandex-tracker/internal/logic/safety"
)

func (s *Service) CommentReactionAddSafety() safety.Level { return safety.Write }

// CommentReactionAdd adds a reaction (e.g. "LIKE", "HEART") to a comment.
func (s *Service) CommentReactionAdd(ctx context.Context, issueKey, commentID, reaction string) (json.RawMessage, error) {
	return safety.Guarded(s.guard, s.CommentReactionAddSafety(), func() (json.RawMessage, error) {
		return s.client.AddCommentReaction(ctx, issueKey, commentID, reaction)
	})
}
