package issue

import (
	"context"
	"encoding/json"

	"github.com/hoggerty/yandex-tracker-mcp/internal/logic/safety"
)

func (s *Service) CommentEditSafety() safety.Level { return safety.Write }

// CommentEdit changes the text (and optionally summonees) of an existing comment.
func (s *Service) CommentEdit(ctx context.Context, issueKey, commentID, text string, summonees []string) (json.RawMessage, error) {
	if err := s.guard.Check(s.CommentEditSafety()); err != nil {
		return nil, err
	}

	return s.client.EditComment(ctx, issueKey, commentID, text, summonees)
}
