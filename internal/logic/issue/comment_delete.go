package issue

import (
	"context"
	"encoding/json"
)

// CommentDelete deletes an issue comment.
func (s *Service) CommentDelete(ctx context.Context, issueKey, commentID string) (json.RawMessage, error) {
	return s.client.DeleteComment(ctx, issueKey, commentID)
}
