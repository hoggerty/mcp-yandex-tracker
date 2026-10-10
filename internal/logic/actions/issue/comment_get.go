package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) CommentGetSafety() safety.Level { return safety.ReadOnly }

// CommentGet retrieves a single issue comment.
func (s *Service) CommentGet(ctx context.Context, issueKey, commentID string) (json.RawMessage, error) {
	if err := s.guard.Check(s.CommentGetSafety()); err != nil {
		return nil, err
	}

	return s.client.GetComment(ctx, issueKey, commentID)
}
