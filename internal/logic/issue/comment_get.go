package issue

import (
	"context"
	"encoding/json"
)

// CommentGet retrieves a single issue comment.
func (s *Service) CommentGet(ctx context.Context, issueKey, commentID string) (json.RawMessage, error) {
	return s.client.GetComment(ctx, issueKey, commentID)
}
