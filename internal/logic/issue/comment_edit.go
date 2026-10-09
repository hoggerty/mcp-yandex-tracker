package issue

import (
	"context"
	"encoding/json"
)

// CommentEdit changes the text (and optionally summonees) of an existing comment.
func (s *Service) CommentEdit(ctx context.Context, issueKey, commentID, text string, summonees []string) (json.RawMessage, error) {
	return s.client.EditComment(ctx, issueKey, commentID, text, summonees)
}
