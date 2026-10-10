package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) CommentReactionAddSafety() safety.Level { return safety.Write }

// CommentReactionAdd adds a reaction (e.g. "LIKE", "HEART") to a comment.
func (s *Service) CommentReactionAdd(ctx context.Context, issueKey, commentID, reaction string) (json.RawMessage, error) {
	if err := s.guard.Check(s.CommentReactionAddSafety()); err != nil {
		return nil, err
	}

	return s.client.AddCommentReaction(ctx, issueKey, commentID, reaction)
}
