package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) CommentDeleteSafety() safety.Level { return safety.Destructive }

// CommentDelete deletes an issue comment.
func (s *Service) CommentDelete(ctx context.Context, issueKey, commentID string) (json.RawMessage, error) {
	if err := s.guard.Check(s.CommentDeleteSafety()); err != nil {
		return nil, err
	}

	return s.client.DeleteComment(ctx, issueKey, commentID)
}
