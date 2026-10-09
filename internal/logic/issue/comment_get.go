package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

const CommentGetSafety = safety.ReadOnly

// CommentGet retrieves a single issue comment.
func (s *Service) CommentGet(ctx context.Context, issueKey, commentID string) (json.RawMessage, error) {
	return s.client.GetComment(ctx, issueKey, commentID)
}
