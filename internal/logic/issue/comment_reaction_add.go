package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const CommentReactionAddSafety = safety.Write

// CommentReactionAdd adds a reaction (e.g. "LIKE", "HEART") to a comment.
func (s *Service) CommentReactionAdd(ctx context.Context, issueKey, commentID, reaction string) (json.RawMessage, error) {
	return s.client.AddCommentReaction(ctx, issueKey, commentID, reaction)
}
